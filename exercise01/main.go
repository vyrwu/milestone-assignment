package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title        string   `json:"title"`
	ISBN         string   `json:"isbn"`
	Downloadable []string `json:"downloadable"`
	Author       string   `json:"author,omitempty"`
	Pages        int      `json:"pages,omitempty"`
}

type Disk struct {
	Format       string   `json:"format"`
	Title        string   `json:"title"`
	ISBN         string   `json:"isbn"`
	Downloadable []string `json:"downloadable"`
	Tracks       []Track  `json:"tracks,omitempty"`
}

type Track struct {
	Typ    string `json:"type"`
	Title  string `json:"title"`
	Artist string `json:"artist,omitempty"`
	Length int    `json:"length, omitempty"`
}

func main() {
	fmt.Println("---------------------------")
	defer fmt.Println("---------------------------")
	vector := &Disk{
		Format: "CD",
		Title:  "Vector",
		ISBN:   "978-3-16-148410-0",
		Downloadable: []string{
			"audio",
		},
		Tracks: []Track{
			{
				Typ:    "audio",
				Title:  "Clear",
				Artist: "Haken",
				Length: 107,
			},
			{
				Typ:    "audio",
				Title:  "The Good Doctor",
				Artist: "Haken",
				Length: 238,
			},
		},
	}

	bytes, _ := json.Marshal(vector)
	fmt.Println(string(bytes))
}
