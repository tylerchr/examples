package main

import "net/http"

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.Handle("/index.html", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// server push is available if w implements http.Pusher
		if p, ok := w.(http.Pusher); ok {
			p.Push("/static/gopher.png", nil)
		}

		// load the main page
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<img src="/static/gopher.png" />`))

	}))

	http.ListenAndServeTLS(":4430", "cert.pem", "key.pem", nil)

}
