package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server Error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "The server encountered an internal error")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Request Error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest, "The server encountered a bad request")
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not Found Error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusNotFound, "The server encountered a not found error")
}
