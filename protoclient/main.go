package main

import (
	pb "GCP_endpoints/myproto"
	"context"
	"fmt"
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
	fmt.Println("Regular Client Request")
	fmt.Println(response.Success)
	fmt.Println(response.Message)

	/*	// Use the reflection service to query the server for services and methods
		refClient := reflection.NewServerReflectionClient(conn)
		reqRef := &reflection.ServerReflectionRequest{
			Host: "localhost:50051",
		}
		resStream, err := refClient.ServerReflectionInfo(context.Background(), reqRef)
		if err != nil {
			// handle error
		}

		fmt.Println("Reflection Request")
		for {
			res, err := resStream.Recv()
			if err != nil {
				break
			}
			fmt.Println(res)
		}*/
}
