package main

import (
	"context"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "GCP_endpoints/myproto"
	"google.golang.org/grpc"
)

type server struct {
	pb.MyServiceServer
}

func (s *server) MyMethod(ctx context.Context, in *pb.MyRequest) (*pb.MyResponse, error) {
	log.Printf("Received request: a = %d, b = %s", in.A, in.B)
	return &pb.MyResponse{Success: true, Message: "Success"}, nil
}

func (s *server) mustEmbedUnimplementedMyServiceServer() {}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterMyServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
