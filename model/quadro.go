package model

type Quadro struct {
	IdQuadro   string
	IdUsuario  []string
	InfoTarefa []InfoTarefa
	NomeQuadro string
}

type InfoTarefa struct {
	IdTarefa  string
	Categoria string //[A Fazer/Em Andamento/Concluido]
}
