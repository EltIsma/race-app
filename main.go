package main

import (
	"log"
	"os"
	"runnersApp/config"
	"runnersApp/server"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting Runners App")
	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())
	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println("Initializing HTTP server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}


func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "runners-" + env
	}

	return "runners"
}