package main

import (
	"time"

	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	userService "github.com/AlexRojasB/go-mongoAtlas-connection.git/repositories/user.repository"
)

func main() {
	user := m.User{
		Name:      "Alexander",
		Email:     "alexrrojas.b@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userService.Create(user)

	if err != nil {

	} else {

	}
}
