package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type env struct {
	idBot string
	port  string
}

type jsonDB struct {
	Ids []string
}

type messageDB struct {
	Ids     []string `json:"ids"`
	Message string   `json:"message"`
}

type responseMessage struct {
	Success bool
	Message string
}

var (
	Data = &jsonDB{}
)

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	envKey := os.Getenv(key)
	if envKey == "" {
		log.Fatalf("Error Variable not found: %s", key)
	}
	return envKey
}

func loadConfig() *jsonDB {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()

	var result jsonDB
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	json.Unmarshal([]byte(content), &result)
	return &result
}

func handleSend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message messageDB
	_ = json.NewDecoder(r.Body).Decode(&message)
	response := getPaylod(message)
	if !response.Success {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(response)
}

func getPaylod(message messageDB) responseMessage {
	var response responseMessage
	response.Success = true
	response.Message = "Message sended"
	if message.Ids == nil {
		message.Ids = Data.Ids
	}
	if message.Message == "" {
		response.Message = "No required field - message"
		response.Success = false
	}
	return response
}

func main() {
	r := mux.NewRouter()
	dotenv := env{idBot: getEnv("telegram_id"), port: getEnv("port")}
	fmt.Println(dotenv.idBot)
	Data = loadConfig()
	fmt.Println(Data.Ids)
	r.HandleFunc("/notify/send", handleSend).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+dotenv.port, r))
}
