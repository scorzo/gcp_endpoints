package main

import (
	pb "GCP_endpoints/myproto"
	"context"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		// handle error
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewMyServiceClient(conn)

	// Create a request
	req := &pb.MyRequest{
		A: 1,
		B: "some_string",
	}

	// Send the request
	response, err := client.MyMethod(context.Background(), req)
	if err != nil {
		// handle error
	}

	// Print the response
	println(response.Success)
	println(response.Message)
}
