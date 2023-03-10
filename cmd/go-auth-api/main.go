package main

import (
	"log"

	v1 "github.com/ishanshre/go-auth-api/api/v1"
)

func main() {
	// initiates the new connection to the postgresql database
	store, err := v1.NewPostgresStore()
	if err != nil {
		log.Fatalf("error in connecting to database: %s", err)
	}
	// runs the init method and creates table or if not exists
	if err := store.Init(); err != nil {
		log.Fatalf("Error in creating table: %s", err)
	}
	server := v1.NewApiServer(":8000", store)
	server.Run()
}
