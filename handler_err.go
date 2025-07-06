package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	resondWithError(w, 400, "Something went wrong...")
}