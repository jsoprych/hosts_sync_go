package handler

import (
	"fmt"
	"net/http"
	"strings"
)

// InfoHandler writes client request details with enhanced styling.
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// Determine client IP, checking the X-Forwarded-For header first.
	clientIP := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		parts := strings.Split(forwarded, ",")
		clientIP = strings.TrimSpace(parts[0])
	}

	// Set content type to HTML.
	w.Header().Set("Content-Type", "text/html")

	// Begin the HTML output with inline CSS.
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Client Information</title>
	<style>
		body {
			font-family: Helvetica, Arial, sans-serif;
			background-color: #e0f7fa;
			margin: 0;
			padding: 20px;
		}
		h1, h2 {
			color: #00796b;
		}
		p {
			font-size: 16px;
			color: #004d40;
		}
		.info-box {
			background: #ffffff;
			padding: 15px;
			border-radius: 5px;
			box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
			margin-bottom: 10px;
		}
	</style>
</head>
<body>
	<h1>Client Information</h1>
	<div class="info-box">
		<p><strong>IP Address:</strong> %s</p>
		<p><strong>Request Method:</strong> %s</p>
		<p><strong>Request URL:</strong> %s</p>
		<p><strong>User Agent:</strong> %s</p>
	</div>
	<h2>Headers</h2>`, clientIP, r.Method, r.URL.String(), r.Header.Get("User-Agent"))

	// Print all headers in styled boxes.
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, `<div class="info-box">
			<p><strong>%s:</strong> %s</p>
		</div>`, name, value)
		}
	}

	// End the HTML document.
	fmt.Fprintf(w, `
</body>
</html>`)
}