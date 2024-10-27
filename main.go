package main

import (
	"embed"
	"net/http"

	"github.com/bajalnyt/todoer/views/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

//go:embed static/**
var staticFS embed.FS

func main() {
	e := echo.New()

	e.Static("/", "static")

	e.GET("/", func(c echo.Context) error {
		buf := templ.GetBuffer()
		defer templ.ReleaseBuffer(buf)

		accordion := components.AccordionExample()
		if err := accordion.Render(c.Request().Context(), buf); err != nil {
			return err
		}

		return c.HTML(http.StatusOK, buf.String())
	})

	e.Start(":8080")
}
