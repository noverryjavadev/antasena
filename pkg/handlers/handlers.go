package handlers

import (
	"github.com/noverryjavadev/antasena/pkg/config"
	"github.com/noverryjavadev/antasena/pkg/models"
	"github.com/noverryjavadev/antasena/pkg/render"
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

// NewRepo create new repository
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	data := &models.TemplateData{
		StringMap: stringMap,
	}

	render.RenderTemplate(w, "about.page.tmpl", data)
}
