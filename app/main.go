package main

import "log"
import "github.com/singhsanket143/Go-Server-Base/types"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	log.Fatal(app.run())
}