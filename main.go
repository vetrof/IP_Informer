package main

import (
	"IP_Informer/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Get("/ip-info/{ip}", handleApiInfo)
	http.ListenAndServe(":8000", router)
}

func handleApiInfo(writer http.ResponseWriter, request *http.Request) {
	ip := chi.URLParam(request, "ip")

	if ip == "" {
		http.Error(writer, "IP not provided", http.StatusBadRequest)
		return
	}

	answer, err := service.IpInformer(ip)
	if err != nil {
		http.Error(writer, "Error fetching IP info", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(answer)
}
