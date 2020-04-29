package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	logrus.Info("Hello world")

	port := os.Getenv("PORT")

	if len(port) == 0 {
		logrus.Fatal("PORT is not set")
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":"+port, nil)
}
