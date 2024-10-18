package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ProxyHandler(w, r)
}
