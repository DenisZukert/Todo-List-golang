package routers

import (
	"Todo/internal/controllers"
	"net/http"
	"strings"
)

func RegisterRoute(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(segments) == 0 || segments[0] == "" {
		http.Error(w, "Endpoint not specified", http.StatusNotFound)
		return
	}

	endpoint := segments[0]
	switch endpoint {
	case "task":
		controllers.TaskController(w, r)
	default:
		http.Error(w, "Endpoint not found", http.StatusNotFound)
	}
}
