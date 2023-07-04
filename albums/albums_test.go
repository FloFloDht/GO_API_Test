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

	// Vérifie le corps de la réponse
	expectedBody := `[{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":"2","title":"Jeru","artist":"Gerry Mulligan","price":17.99},{"id":"3","title":"Sarah Vaughan and Clifford Brown","artist":"Sarah Vaughan","price":39.99}]`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s but got %s", expectedBody, w.Body.String())
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

	// Vérifie le corps de la réponse
	expectedBody := `{"id":"4","title":"New Album","artist":"Artist","price":10.99}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s but got %s", expectedBody, w.Body.String())
	}
}

func TestGetAlbumByID(t *testing.T) {
	// Crée un contexte de test avec une requête GET factice contenant un paramètre d'ID
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/albums/2", nil)
	c.Request = req

	// Exécute la fonction de test
	albums.GetAlbumByID(c)

	// Vérifie le code de statut de la réponse
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	// Vérifie le corps de la réponse
	expectedBody := `{"id":"2","title":"Jeru","artist":"Gerry Mulligan","price":17.99}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s but got %s", expectedBody, w.Body.String())
	}
}
