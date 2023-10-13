package main

import (
	"context"
	"fmt"
	"grpc_hello_world/hello"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hello.UnimplementedHelloServiceServer
}

func (s *server) Hello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Message: "Hello " + in.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})

	fmt.Println("Server started...")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
