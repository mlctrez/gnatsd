package server

import (
	"net/url"
)

// ConnectToRoute connects the embedded server to the provided cluster url.
func (s *Server) ConnectToRoute(url *url.URL) {
	s.startGoRoutine(func() { s.connectToRoute(url, false) })
}
