package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
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
	Data   = &jsonDB{}
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
	fmt.Println("new req with response: ", response)
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
	if response.Success {
		sendMessage(message.Message, message.Ids)
	}
	return response
}

func sendMessage(message string, users []string) {
	for i := range users {
		bot, err := tgbotapi.NewBotAPI(getEnv("telegram_id"))
		if err != nil {
			log.Panic(err)
		}
		idChat, _ := strconv.ParseInt(users[i], 10, 64)
		msg := tgbotapi.NewMessage(idChat, "Message from Notitication API:\n" +message)
		bot.Send(msg)
	}
}

func main() {
	r := mux.NewRouter()
	dotenv := env{idBot: getEnv("telegram_id"), port: getEnv("port")}
	Data = loadConfig()
	r.HandleFunc("/notify/send", handleSend).Methods("POST")
	fmt.Println("Server listen", dotenv.port)
	log.Fatal(http.ListenAndServe(":"+dotenv.port, r))
}
