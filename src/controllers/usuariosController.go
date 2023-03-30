package controllers

import "net/http"

//Criar usuário insere um usuário no BD
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!!"))
}

//BuscarUsuarios busca todos os usuários salvos BD
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

//BuscarUsuarioId busca um usuário só no BD
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!!"))
}

//AtualizandoUsuario atualiza as informações de um usuário no BD
func AtualizandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!!"))
}

//DeletarUsuario deleta um usuário no BD
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!!"))
}
