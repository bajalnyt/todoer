package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/bajalnyt/todoer/views/components"
)

//go:embed static/**
var staticFS embed.FS

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		buf := templ.GetBuffer()
		defer templ.ReleaseBuffer(buf)

		accordion := components.AccordionExample()
		err := accordion.Render(context.Background(), buf)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, buf.String())
	})

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	mux.Handle("GET /static/*",
		http.StripPrefix("/static/",
			http.FileServer(
				http.Dir("./static"),
			),
		),
	)

	srv := &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", 8080),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}
