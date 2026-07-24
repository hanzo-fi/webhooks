package server

import (
	"net/http"

	"github.com/hanzo-fi/go-libs/v2/api"
)

type ServiceInfo struct {
	Version string `json:"version"`
}

func (h *serverHandler) getInfo(info ServiceInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api.RawOk(w, info)
	}
}
