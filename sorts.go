package main

import (
	"net/http"

	"github.com/arlinrus/serv.git/service"
)

func main() {
	mux := http.NewServeMux()

	scrv := service.New()

	mux.Handle("/poll", scrv.Poll)
	mux.Handle("/serv", scrv.Vote)

	http.ListenAndServe(":8000", mux)

}
