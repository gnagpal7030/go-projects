package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gnagpal7030/go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewGreetingServiceClient(cc)
	request := &pb.GreetingServiceRequest{Name: "Gophers"}

	resp, err := client.Greeting(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive response => %s ", resp.Message)
}
