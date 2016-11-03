package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// pushingRecorder wraps httptest.NewRecorder and implements http.Pusher
type (
	pushingRecorder struct {
		*httptest.ResponseRecorder                // wraps an httptest.ResponseRecorder
		Pushes                     []RecordedPush // records all observed server pushes
	}

	RecordedPush struct {
		Path    string
		Options *http.PushOptions
	}
)

func (pw *pushingRecorder) Push(target string, opts *http.PushOptions) error {
	pw.Pushes = append(pw.Pushes, RecordedPush{
		Path:    target,
		Options: opts,
	})
	return nil
}

func TestServerPush(t *testing.T) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// server push is available if w implements http.Pusher
		if p, ok := w.(http.Pusher); ok {
			p.Push("/static/gopher.png", nil)
		} else {
			fmt.Println("not a pusher")
		}

		// load the main page
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<img src="/static/gopher.png" />`))

	})

	rw := &pushingRecorder{ResponseRecorder: httptest.NewRecorder()}

	req, _ := http.NewRequest("GET", "/index.html", nil)

	handler(rw, req)

	// t.Logf("Pushes: %#v\n", rw.Pushes)
	expected := []RecordedPush{RecordedPush{Path: "/static/gopher.png", Options: nil}}

	if !reflect.DeepEqual(expected, rw.Pushes) {
		t.Errorf("unexpected Server Pushes occurred: expected '%#v' but got '%#v'\n", expected, rw.Pushes)
	}

}
