package user_service_test

import (
	"testing"
	"time"

	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	userService "github.com/AlexRojasB/go-mongoAtlas-connection.git/services/user.service"
)

func TestCreate(t *testing.T) {
	user := m.User{
		Name:      "Alexander",
		Email:     "alexrrojas.b@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userService.Create(user)

	if err != nil {
		t.Error("Error en la prueba de persistencia de datos")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}

func TestRead(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
