package main

import "log"
import "GoTwitter/types"

func main() {
	cfg := types.Config{
		Addr: ":8080",
	}

	app := &types.Application{
		Config: cfg,
	}

	log.Fatal(app.Run())
}