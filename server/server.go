package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fukata/golang-stats-api-handler"
)

type Server struct {
	AppVersion string

	mux *http.ServeMux
}

func New(appVersion string) *Server {
	s := &Server{
		AppVersion: appVersion,
	}

	s.mux = http.NewServeMux()

	s.mux.HandleFunc("/api/stats", stats_api.Handler)
	s.mux.HandleFunc("/api/sleep", s.versionHandler)

	return s
}

func (s *Server) versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	b, _ := json.Marshal(struct {
		Version string `json:"version"`
	}{Version: s.AppVersion})

	w.WriteHeader(http.StatusOK)
	w.Write(b)

	time.Sleep(5 * time.Second)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func outputErrorMsg(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	b, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{Error: msg})

	w.WriteHeader(status)
	w.Write(b)
}
