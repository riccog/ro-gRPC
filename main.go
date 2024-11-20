package main

import (
	"context"
	"github.com/riccog/ro-gRPC/UserService"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myUserServiceServer struct {
	UserService.UnimplementedUserServiceServer // Embed this line
}

func (s myUserServiceServer) GetUser(ctx context.Context, req *UserService.CreateRequest) (*UserService.CreateResponse, error) {
	return &UserService.CreateResponse{
		UserProfile: &UserService.UserProfile{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoeyouknow@cool.com",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	UserService.RegisterUserServiceServer(server, &myUserServiceServer{})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
