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
	group1 := &UserService.Group{
		GroupName: "Group A",
		GroupId:   "group_a",
	}
	group2 := &UserService.Group{
		GroupName: "Group B",
		GroupId:   "group_b",
	}
	return &UserService.CreateResponse{
		UserProfile: &UserService.UserProfile{
			UserId:    "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoeyouknow@cool.com",
			Groups:    []*UserService.Group{group1, group2},
		},
	}, nil
}

func (s myUserServiceServer) AddTask(ctx context.Context, req *UserService.CreateTaskRequest) (*UserService.TaskResponse, error) {
	return &UserService.TaskResponse{
		Task: &UserService.Task{
			TaskId:     "This is the task ID",
			TaskName:   "Take out the trash",
			TaskStatus: UserService.TaskStatus_IN_PROGRESS,
			UserId:     "1",
			GroupId:    "group_a",
		},
	}, nil
}

func (s myUserServiceServer) DeleteTask(ctx context.Context, req *UserService.DeleteTaskRequest) (*UserService.TaskResponse, error) {
	return &UserService.TaskResponse{
		Task: &UserService.Task{
			TaskId:     "This is the task ID",
			TaskName:   "Take out the trash",
			TaskStatus: UserService.TaskStatus_IN_PROGRESS,
			UserId:     "1",
			GroupId:    "group_a",
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
