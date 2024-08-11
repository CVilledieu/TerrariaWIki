package server

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.Static("/static", "public")
	e.Renderer = newTemplate()
	e.GET("/", nil)

	e.Logger.Fatal(e.Start(":2222"))
}

type Template struct {
	templates *template.Template
}

func newTemplate() *Template {
	return &Template{templates: template.Must(template.ParseGlob("public/views/*.html"))}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
