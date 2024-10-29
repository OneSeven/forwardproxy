// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: server.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Server_AddUser_FullMethodName      = "/server.Server/AddUser"
	Server_UpdateUser_FullMethodName   = "/server.Server/UpdateUser"
	Server_DeleteUser_FullMethodName   = "/server.Server/DeleteUser"
	Server_SyncUser_FullMethodName     = "/server.Server/SyncUser"
	Server_TrafficStats_FullMethodName = "/server.Server/TrafficStats"
)

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SyncUser(ctx context.Context, in *UserList, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TrafficStats(ctx context.Context, in *Interval, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UserList], error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Server_AddUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Server_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Server_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) SyncUser(ctx context.Context, in *UserList, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Server_SyncUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) TrafficStats(ctx context.Context, in *Interval, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UserList], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Server_ServiceDesc.Streams[0], Server_TrafficStats_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Interval, UserList]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Server_TrafficStatsClient = grpc.ServerStreamingClient[UserList]

// ServerServer is the server API for Server service.
// All implementations must embed UnimplementedServerServer
// for forward compatibility.
type ServerServer interface {
	AddUser(context.Context, *User) (*emptypb.Empty, error)
	UpdateUser(context.Context, *User) (*emptypb.Empty, error)
	DeleteUser(context.Context, *User) (*emptypb.Empty, error)
	SyncUser(context.Context, *UserList) (*emptypb.Empty, error)
	TrafficStats(*Interval, grpc.ServerStreamingServer[UserList]) error
	mustEmbedUnimplementedServerServer()
}

// UnimplementedServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServerServer struct{}

func (UnimplementedServerServer) AddUser(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedServerServer) UpdateUser(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedServerServer) DeleteUser(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedServerServer) SyncUser(context.Context, *UserList) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncUser not implemented")
}
func (UnimplementedServerServer) TrafficStats(*Interval, grpc.ServerStreamingServer[UserList]) error {
	return status.Errorf(codes.Unimplemented, "method TrafficStats not implemented")
}
func (UnimplementedServerServer) mustEmbedUnimplementedServerServer() {}
func (UnimplementedServerServer) testEmbeddedByValue()                {}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	// If the following call pancis, it indicates UnimplementedServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_SyncUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).SyncUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_SyncUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).SyncUser(ctx, req.(*UserList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_TrafficStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Interval)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerServer).TrafficStats(m, &grpc.GenericServerStream[Interval, UserList]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Server_TrafficStatsServer = grpc.ServerStreamingServer[UserList]

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Server_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Server_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Server_DeleteUser_Handler,
		},
		{
			MethodName: "SyncUser",
			Handler:    _Server_SyncUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TrafficStats",
			Handler:       _Server_TrafficStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server.proto",
}
