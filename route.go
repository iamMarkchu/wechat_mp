package main

import (
	"net/http"
	"regexp"
	"io"
)

type WebController struct {
	Function    func(w http.ResponseWriter, r *http.Request)
	Method  string
	Pattern string
}

var mux []WebController

func init() {
	mux = append(mux, WebController{homeGet, "GET", "^/$"})
	mux = append(mux, WebController{homePost, "POST", "^/$"})
	mux = append(mux, WebController{homePost, "GET", "^/home$"})
}

type httpHandler struct{}

func (this *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, WebController := range mux {
		if m, _ := regexp.MatchString(WebController.Pattern, r.URL.Path); m {
			if r.Method == WebController.Method {
				WebController.Function(w, r)
				return
			}
		}
	}
	io.WriteString(w, "")
	return
}
