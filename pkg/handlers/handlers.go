package handlers

import (
	config2 "github.com/KirillNikoda/bookings/pkg/config"
	models2 "github.com/KirillNikoda/bookings/pkg/models"
	render2 "github.com/KirillNikoda/bookings/pkg/render"
	"net/http"
)


// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config2.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config2.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render2.RenderTemplate(w, "home.page.tmpl", &models2.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render2.RenderTemplate(w, "about.page.tmpl", &models2.TemplateData{
		StringMap: stringMap,
	})
}
