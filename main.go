package reta

import "math"

//inicializando com letra maiuscula é publico (visivel dentro e fora do pacote)
//inicializando com letra maiuscula é privado (visivel dentro do pacote)
//todos arquivos viram pacotes
//Ponto - publico
//ponto ou _Ponto - privado

// Ponto representa uma coordenada no ponto cartesiano
type Ponto struct {
	X float64
	Y float64
}

func Catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.X - a.X)
	cy = math.Abs(b.Y - a.Y)
	return
}

// Distancia calcula distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := Catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
