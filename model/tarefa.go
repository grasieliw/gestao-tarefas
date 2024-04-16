package model

type Tarefa struct {
	Nome        string
	Descricao   string
	Categoria   string
	IdTarefa    string
	IdUsuario   []string
	DataEntrega string
}
