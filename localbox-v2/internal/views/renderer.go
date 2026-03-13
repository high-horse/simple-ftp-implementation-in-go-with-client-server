package views

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

var templates = map[string]*template.Template{}

func LoadTemplates() {
    files, err := filepath.Glob("web/templates/*.tmpl")
    if err != nil {
        log.Fatalf("Failed to read templates: %v", err)
    }

    for _, file := range files {
        name := filepath.Base(file)
        tmpl, err := template.ParseFiles("web/templates/layout.tmpl", file)
        if err != nil {
            log.Fatalf("Failed to parse template %s: %v", name, err)
        }
        templates[name] = tmpl
    }

    log.Printf("Loaded %d templates", len(templates))
}

func isHTMX(r *http.Request) bool {
    return r.Header.Get("HX-Request") == "true"
}


func Render(w http.ResponseWriter, page string, data any, r *http.Request) {
    tmpl, ok := templates[page]
    if !ok {
        http.Error(w, "Template not found", http.StatusInternalServerError)
        return
    }

    if isHTMX(r) {
        // Render only the "body" block for HTMX
        if err := tmpl.ExecuteTemplate(w, "body", data); err != nil {
            log.Printf("HTMX template execution error: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
        return
    }

    // Normal page request: render full layout
    if err := tmpl.ExecuteTemplate(w, "layout.tmpl", data); err != nil {
        log.Printf("Template execution error: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
