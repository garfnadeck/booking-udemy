package handlers

import (
	"github.com/garfnadeck/bookings-udemy/pkg/config"
	"github.com/garfnadeck/bookings-udemy/pkg/models"
	"github.com/garfnadeck/bookings-udemy/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["Test"] = "Hello, again"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(
		w, "about.page.tmpl", &models.TemplateData{
			StringMap: stringMap,
		},
	)
}
