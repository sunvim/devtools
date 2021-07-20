package main

import (
	"context"
	"log"

	"../grace"
)

// pls run: GO111MODULE=off go run ./grace_main.go
func main() {
	log.Println("start grace test")

	_, mainServ := grace.New(context.Background())

	mainServ.Wait()

}
