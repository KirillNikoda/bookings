package main

import (
	config2 "github.com/KirillNikoda/bookings/pkg/config"
	handlers2 "github.com/KirillNikoda/bookings/pkg/handlers"
	render2 "github.com/KirillNikoda/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

var (
	portNumber string = ":3000"
	app        config2.AppConfig
	session    *scs.SessionManager
)


// main is the main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render2.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache:", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers2.NewRepo(&app)
	handlers2.NewHandlers(repo)

	render2.NewTemplates(&app)



	//log.Fatal(http.ListenAndServe(portNumber, nil))

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	log.Fatal(srv.ListenAndServe())
}
