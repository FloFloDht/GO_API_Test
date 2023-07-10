package albums_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/mod/albums"
	"github.com/gin-gonic/gin"
)

func TestGetAlbums(t *testing.T) {
	// Crée un contexte de test avec une requête GET factice
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/albums", nil)
	c.Request = req

	// Exécute la fonction de test
	albums.GetAlbums(c)

	// Vérifie le code de statut de la réponse
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}
}

func TestAddAlbums(t *testing.T) {
	// Crée un contexte de test avec une requête POST factice et un corps de requête JSON
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"id":"4","title":"New Album","artist":"Artist","price":10.99}`))
	c.Request = req

	// Exécute la fonction de test
	albums.AddAlbums(c)

	// Vérifie le code de statut de la réponse
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, w.Code)
	}
}

func TestGetAlbumByID(t *testing.T) {

	// Crée un contexte de test avec une requête GET factice contenant un paramètre d'ID
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/albums/2", nil)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "2"})

	// Exécute la fonction de test
	albums.GetAlbumByID(c)

	// Vérifie le code de statut de la réponse
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}
}
