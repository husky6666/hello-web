package main

import (
	"context"
	"google.golang.org/grpc"
	"hello-web/pb_service"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	service := pb_service.NewHelloServiceClient(conn)
	for i := 0; i < 100; i++ {
		beginTime := time.Now()
		response, err := service.SayHello(context.Background(), &pb_service.HelloRequest{Name: "husky"})
		if err != nil {
			log.Fatalf("%v", err)
		}
		endTime := time.Now()
		log.Printf("请求时长:%v/微妙, 响应消息:%v", endTime.Sub(beginTime).Microseconds(), response.GetMessage())
	}
	defer conn.Close()
}
