package types

import "net/http"
import "time"
import "log"

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) pingHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	return
}

func (app *application) run() error {

	router := http.NewServeMux()

	router.HandleFunc("GET /v1/ping", app.pingHandler)
	

	server := &http.Server{
		Addr: app.config.addr,
		Handler: router,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Server has starter at %s", app.config.addr)

	return server.ListenAndServe()
}