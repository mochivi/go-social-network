package main

import (
	"gopher_social/internal/env"
	"gopher_social/internal/store"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create config
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	// Create store
	store := store.NewStorage(nil)

	// Create application
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
