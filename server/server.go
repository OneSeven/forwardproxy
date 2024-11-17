package server

import (
	"context"
	"encoding/base64"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"slices"
	"sync"
	"time"
)

// EncodeAuthCredentials base64-encode credentials
func EncodeAuthCredentials(user, pass string) (result []byte) {
	raw := []byte(user + ":" + pass)
	result = make([]byte, base64.StdEncoding.EncodedLen(len(raw)))
	base64.StdEncoding.Encode(result, raw)
	return
}

type Server struct {
	UnimplementedServerServer
	Users        sync.Map
	UsersTraffic sync.Map
	Connect      atomic.Bool
	GrpcServer   *grpc.Server
}

func NewServer() *Server {
	server := &Server{}
	server.Connect.Store(false)
	return server
}

type UserTraffic struct {
	Username  string `json:"username"`
	Traffic   atomic.Uint64
	Ip        atomic.String
	CreatedAt int64
	UpdatedAt atomic.Int64
}

func (s *Server) AddUser(ctx context.Context, user *User) (*emptypb.Empty, error) {
	auth := string(EncodeAuthCredentials(user.Username, user.Password))
	s.Users.Store(auth, user)
	timeNow := time.Now().Unix()

	if _, ok := s.UsersTraffic.Load(user.Username); !ok {
		userTraffic := &UserTraffic{
			Username:  user.Username,
			CreatedAt: timeNow,
		}
		userTraffic.Traffic.Store(0)
		userTraffic.Ip.Store("")
		userTraffic.UpdatedAt.Store(timeNow)
		s.UsersTraffic.Store(user.Username, userTraffic)
	}
	return nil, nil
}

func (s *Server) UpdateUser(ctx context.Context, user *User) (*emptypb.Empty, error) {
	auth := string(EncodeAuthCredentials(user.Username, user.Password))
	s.Users.Store(auth, user)
	return nil, nil
}

func (s *Server) DeleteUser(ctx context.Context, user *User) (*emptypb.Empty, error) {
	auth := string(EncodeAuthCredentials(user.Username, user.Password))
	s.Users.Delete(auth)
	return nil, nil
}

func (s *Server) SyncUser(ctx context.Context, list *UserList) (*emptypb.Empty, error) {
	for _, user := range list.UserList {
		auth := string(EncodeAuthCredentials(user.Username, user.Password))
		if _, ok := s.Users.Load(auth); !ok {
			_, _ = s.AddUser(ctx, user)
		}
	}
	s.Users.Range(func(key, value interface{}) bool {
		existUser := value.(*User)
		if slices.IndexFunc(list.UserList, func(user *User) bool {
			if existUser.Username == user.Username {
				return true
			}
			return false
		}) < 0 {
			auth := string(EncodeAuthCredentials(existUser.Username, existUser.Password))
			s.Users.Delete(auth)
		}
		return true
	})
	return nil, nil
}

func (s *Server) TrafficStats(interval *Interval, stream grpc.ServerStreamingServer[UserList]) error {
	s.Connect.Store(true)
	defer func() {
		s.Connect.Store(false)
	}()
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			var userList []*User
			s.UsersTraffic.Range(func(key, value interface{}) bool {
				userTraffic := value.(*UserTraffic)
				userList = append(userList, &User{
					Username:  userTraffic.Username,
					Traffic:   userTraffic.Traffic.Load(),
					Ip:        userTraffic.Ip.Load(),
					CreatedAt: userTraffic.CreatedAt,
					UpdatedAt: userTraffic.UpdatedAt.Load(),
				})
				return true
			})
			err := stream.Send(&UserList{
				UserList: userList,
			})
			if err != nil {
				return err
			}
		}
		time.Sleep(time.Millisecond * time.Duration(interval.Interval))
	}
}
