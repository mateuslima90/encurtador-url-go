package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mateuslima90/encurtador-url-go/url"
)

var (
	porta int
	urlBase string
)

func init() {
	porta = 8888
	urlBase = fmt.Sprintf("htpp://localhost:%d", porta)
}

type Headers map[string]string

func main() {
	http.HandleFunc("api/encurtar", Encurtador)

	http.HandleFunc("/r/", Redirecionador)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", porta), nil))
}

func Encurtador(w http.ResponseWriter, r *http.Resquest) {

	if r.Method != "POST" {
		responderCom(W, http.StatusMethodNotAllowed, Header{ "Allow": "POST"})
		return
	}

	url, nova, err := url.BuscarOuCriarNovaUrl(extrairUrl(r))

	if err != nil {
		responderCom(w, http.StatusBadRequest, nil)
		return
	}

	var status int
	if nova {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	urlCurta := fmt.Sprintf("%s/r/%s", urlBase, url.Id)
	responderCom(w, status, Headers("Lcation": urlCurta))
}

func responderCom(w http.ResponseWriter, status int, headers Headers) {
	for k,v := range headers {
		w.Header().Set(k, v))
	}
	w.WriteHeader(status)
}

func extrairUrl(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

