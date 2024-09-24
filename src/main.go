package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Aplicacao exemplo - HARDCLOUD - V4.4.0")
}




package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Obtém o nome do pod da variável de ambiente
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "Nome do Pod não encontrado"
	}

	// Exibe o nome do pod na resposta HTTP
	fmt.Fprintf(w, "<h1>Nome do Pod: %s</h1>", podName)
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Printf("Servidor rodando na porta %s\n", port)
	http.ListenAndServe(port, nil)
}






func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}
