package main

import "time"

type Note struct {
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type Notes []Note
