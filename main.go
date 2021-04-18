package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	// "log"
	// "net/http"
)

type env struct {
	idBot string
	port  string
}

type jsonDB struct {
	Ids         []string
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
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	var result jsonDB
	content, err := ioutil.ReadFile("data.json")
	checkError(err)
	json.Unmarshal([]byte(content), &result)
	return &result
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dotenv := env{idBot: getEnv("telegram_id"), port: getEnv("port")}
	fmt.Println(dotenv.idBot)
	Data = loadConfig()
	fmt.Println(Data.Ids)
}
