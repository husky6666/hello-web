package main

import (
	"context"
	"google.golang.org/grpc"
	"hello-web/pb_service"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	service := pb_service.NewHelloServiceClient(conn)
	response, err := service.SayHello(context.Background(), &pb_service.HelloRequest{Name: "husky"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Println(response.GetMessage())
	defer conn.Close()
}
