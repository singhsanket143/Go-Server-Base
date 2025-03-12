package types

import "net/http"
import "time"
import "log"

type Application struct {
	Config Config
}

type Config struct {
	Addr string
}

func (app *Application) pingHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	return
}

func (app *Application) Run() error {

	router := http.NewServeMux()

	router.HandleFunc("/v1/ping", app.pingHandler)
	

	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: router,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Server has started at %s", app.Config.Addr)

	return server.ListenAndServe()
}