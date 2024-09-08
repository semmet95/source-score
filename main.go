package main

import (
	"log"
	pong "oapi-gin-petstore/api"
	"oapi-gin-petstore/pkg/api"

	"github.com/gin-gonic/gin"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=api/config.yaml api/minimal-api.yaml
func main() {
	server := gin.Default()
	api.RegisterHandlers(server, pong.NewPong())

	err := server.Run()
	if err != nil {
		log.Fatalf("failed to start the server: %s\n", err)
	} else {
		log.Println("Server is up and running!")
	}
}
