package main

import (
	"fmt"
	"net/http"
)

const (
	authUser = "user"
	authPass = "pass"
)

func main() {
	http.HandleFunc("/auth", auth)
	http.ListenAndServe(":9094", nil)
}

func auth(w http.ResponseWriter, r *http.Request) {
	if user, pass, ok := r.BasicAuth(); !ok || user != authUser || pass != authPass {
		w.Header().Add("WWW-Authenticate", `Basic realm="my private area"`)
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Not Authorized", 402)
		return
	}
	fmt.Fprintln(w, "Authed")
}
