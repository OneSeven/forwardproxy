package server

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     0,                // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      0,                // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 10 * time.Second, // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  20 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               5 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

func GRPCServer(server *Server) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp),
	)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", 9887))
	if err != nil {
		log.Fatalf("grpc failed to listen: %v", err)
	}
	RegisterServerServer(grpcServer, server)
	fmt.Printf("node server listening at %v\n", lis.Addr())

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()
	server.GrpcServer = grpcServer
	return grpcServer
}
