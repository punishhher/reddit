package main

import (
	"log"
	"reddit123/package/handler"
	"reddit123/package/repository"
	"reddit123/package/service"
)

func main() {
	serverInstance := new(Server)

	postgresConfig := repository.Config{
		Host:     "",
		Port:     "",
		Username: "",
		Password: "",
		DBName:   "",
		SSLMode:  "",
	}

	database, err := repository.NewPostgresDB(postgresConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	repos := repository.NewRepository(database)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	runServer(serverInstance, handlers)

}

func runServer(serverInstance *Server, handlerLayer *handler.Handler) {
	port := "8080"
	router := handlerLayer.InitRoutes()

	if err := serverInstance.Run(port, router); err != nil {
		log.Fatal(err.Error())
	}

	log.Print("server started successfully")
}
