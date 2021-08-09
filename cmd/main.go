package main

import (
	"fmt"
	"log"

	"github.com/klimmen/todo-app"
	"github.com/klimmen/todo-app/pkg/handler"
)

func main() {
	fmt.Println("Hello world!")
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
