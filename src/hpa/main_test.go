package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoop(t *testing.T) {
	t.Run("Loop", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		loop(response, request)

		got := response.Body.String()
		want := "Code.education Rocks!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
