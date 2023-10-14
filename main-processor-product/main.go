package main

import (
	golangProto "github.com/wahyubagus-ars/golang-proto"
	GenProtoTransaction "github.com/wahyubagus-ars/golang-proto/transaction/generated"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	golangProto.LogInfo("hello")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	_ = GenProtoTransaction.NewProductInfoClient(conn)
}
