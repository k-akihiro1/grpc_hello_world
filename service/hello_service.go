package service

import (
	"context"
	"grpc_hello_world/hello"
)

type HelloService struct {
	hello.UnimplementedHelloServiceServer
}

func (s *HelloService) Hello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Message: "Hello " + in.Name + "!"}, nil
}
