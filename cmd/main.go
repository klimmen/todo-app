package main

import (
	"fmt"
	"log"

	// "github.com/go-delve/delve/service"
	"github.com/klimmen/todo-app"
	"github.com/klimmen/todo-app/pkg/handler"
	"github.com/klimmen/todo-app/pkg/repository"
	"github.com/klimmen/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Start")
	if err := InitConfig(); err != nil {
		log.Fatalf("error iniializing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
