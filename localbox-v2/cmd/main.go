package main

import (
    "log"
    "net/http"
    "tiny-drop/internal/views"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    views.LoadTemplates()

    r := chi.NewRouter()
    r.Use(middleware.Logger)

    // Static files if needed
    r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

    // Routes
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        views.Render(w, "home.tmpl", map[string]string{
            "Title": "Home",
            "User":  "Camel",
        }, r)
    })

    r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
        views.Render(w, "about.tmpl", map[string]string{
            "Title": "About",
            "Description": "about the page from server",
        }, r)
    })

    log.Println("Server running at http://localhost:8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
