package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string	`json:"author"`
	Year int `json:"publication_year,omitempty"`
	ISBN string `json:"isbn,omitempty"`
	ReaderID *int `json:"reader_id,omitempty"`
}

func main() {
	jsonString := `{
	"id": 101,
	"title": "Что такое короткоя перемена",
	"author": "Стдуент отчисленный",
	"publication_year": 2026,
	"isbn": "",
	"reader_id": null
	}`

	jsonDataBytes := []byte(jsonString)
	var loadedBook Book
	err := json.Unmarshal(jsonDataBytes, &loadedBook)
	if err != nil {
		log.Fatalf("Ошибка десериализации: %v", err)
	}

	fmt.Printf("Загруженная книга: %+v\n", loadedBook)
}
