package router

// Gerar vai retornar um Router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)

import "github.com/gorilla/mux"

//Gerar vai retornar um Router com as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
