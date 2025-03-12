package handler

import (
	"fmt"
	"net/http"
	"strings"
)

// InfoHandler writes client request details to the response.
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// Get client IP from RemoteAddr or X-Forwarded-For if behind a proxy.
	clientIP := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		// In case of multiple proxies, the first IP is the client.
		parts := strings.Split(forwarded, ",")
		clientIP = strings.TrimSpace(parts[0])
	}

	// Write out the basic client info.
	fmt.Fprintf(w, "<h1>Client Information</h1>")
	fmt.Fprintf(w, "<p><strong>IP Address:</strong> %s</p>", clientIP)
	fmt.Fprintf(w, "<p><strong>Request Method:</strong> %s</p>", r.Method)
	fmt.Fprintf(w, "<p><strong>Request URL:</strong> %s</p>", r.URL.String())
	fmt.Fprintf(w, "<p><strong>User Agent:</strong> %s</p>", r.Header.Get("User-Agent"))

	// Optionally, print all headers.
	fmt.Fprintf(w, "<h2>Headers</h2>")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "<p><strong>%s:</strong> %s</p>", name, value)
		}
	}
}