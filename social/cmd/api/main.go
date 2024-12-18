package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		address: os.Getenv("API_ADDRESS"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	
	log.Fatal(app.run(mux))
}
