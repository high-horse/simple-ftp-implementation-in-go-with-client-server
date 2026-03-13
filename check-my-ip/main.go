package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Get client IP from headers
    ip := r.Header.Get("X-Forwarded-For")
    if ip == "" {
        ip = r.Header.Get("X-Real-IP")
    }
    if ip == "" {
        ip = r.RemoteAddr
    }

    // Serve HTML with embedded JS
    fmt.Fprintf(w, `
    <!DOCTYPE html>
    <html>
    <head><title>User Info</title></head>
    <body>
        <h1>Welcome!</h1>
        <p>Your IP: %s</p>
        <p id="username">Detecting username...</p>

        <script>
            // Simple "fingerprint" using browser info
            const fingerprint = navigator.userAgent + " | " + screen.width + "x" + screen.height;
            document.getElementById("username").innerText = "Client fingerprint: " + fingerprint;
        </script>
    </body>
    </html>
    `, ip)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Serving on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
