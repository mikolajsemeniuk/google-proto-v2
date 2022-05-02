package main

import (
	"context"
	"log"
	"net"
	"os"
	"project/greet"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
)

type server struct{}

func (*server) Greet(c context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	count := req.GetCount()

	if count == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Received a zero value for Count")
	}

	res := &greet.GreetResponse{
		Count: count + 1,
	}

	return res, nil
}

func main() {
	connection, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	greet.RegisterGreetServiceServer(s, &server{})
	log.Println("Server listen on 50051")

	if err := s.Serve(connection); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
