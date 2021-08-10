package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tzey/twittorcgs/bd"
	"github.com/tzey/twittorcgs/models"
)

// para crear en la BD el reg de usuario
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contrasena ha de ser de 6 caracteres o mas", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "El email ya esta en la BD", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrion un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
