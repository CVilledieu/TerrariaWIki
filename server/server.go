package server

import (
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.Static("/static", "public")
	e.Renderer = newTemplate()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/GameStageOpen", func(c echo.Context) error {
		return c.Render(http.StatusOK, "GameStage", nil)
	})

	e.GET("/GameStageClose", func(c echo.Context) error {
		return c.Render(http.StatusOK, "Core_Content", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

type Template struct {
	templates *template.Template
}

func newTemplate() *Template {
	return &Template{templates: template.Must(template.ParseGlob(`public/views/html/*/*.html`))}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("public/views/html/error/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
}
