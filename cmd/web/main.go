package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/treyarte/bookings/internal/config"
	"github.com/treyarte/bookings/internal/handlers"
	"github.com/treyarte/bookings/internal/render"
)

const port = ":8080"


var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", port)

	srv := &http.Server{
		Addr:port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
	// _ = http.ListenAndServe(port, nil)
}