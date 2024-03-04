package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	Response struct {
		Endpoints      []Endpoint     `json:"endpoints"`
		VersionWeights map[string]int `json:"version_weights"`
		ServiceConfig  string         `json:"service_config"`
	}

	Endpoint struct {
		Address string `json:"address"`
		Version string `json:"version"`
		Weight  int    `json:"weight"`
	}
)

var HelloWorldServerResponse = Response{
	Endpoints: []Endpoint{
		{
			Address: "127.0.0.1:8080",
			Version: "v1",
		},
		{
			Address: "127.0.0.1:8090",
			Version: "v2",
			Weight:  34,
		},
		{
			Address: "127.0.0.1:8095",
			Version: "v2",
			Weight:  67,
		},
	},
	VersionWeights: map[string]int{
		"v1": 10,
		"v2": 90,
	},
	ServiceConfig: "{\"loadBalancingPolicy\": \"habr_balancer\"}",
}

func main() {
	log.Println("DNS SERVER UP")
	http.HandleFunc("/endpoints", HandleEndpointRequest)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleEndpointRequest(w http.ResponseWriter, req *http.Request) {
	log.Println("new request with queryy ", req.URL.RawQuery)

	target := req.URL.Query().Get("target")
	switch target {
	case "hello_world_server":
		marshalledResponse, err := json.Marshal(HelloWorldServerResponse)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(marshalledResponse)
		if err != nil {
			log.Println(err)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
