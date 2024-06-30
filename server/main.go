package main

import (
	"context"
	"log"
	"net"

	"gihub.com/wesleywillians/grpc-hello-go/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (*server) Hello(ctx context.Context, request *pb.HelloResquest) (*pb.HelloResponse, error) {
	result := "hello " + request.GetName()

	response := &pb.HelloResponse{
		Msg: result,
	}

	return response, nil
}

func main() {
	listener, error := net.Listen("tcp", "0.0.0.0:8080")
	if error != nil {
		log.Fatalf("failed to List %v", error)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &server{})

	if error := grpcServer.Serve(listener); error != nil {
		log.Fatalf("failed to List %v", error)
	}
}
