package main

import (
	"fmt"

	"github.com/Bruno-Bianchi/go-api-rest/database"
	"github.com/Bruno-Bianchi/go-api-rest/models"
	"github.com/Bruno-Bianchi/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
