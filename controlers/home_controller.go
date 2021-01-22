package controlers

import (
	"net/http"

	"github.com/reyreyhan/golang-try/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
