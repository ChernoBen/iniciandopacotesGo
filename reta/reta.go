package main

import "math"

//Metodos, variaveis, funcoes e etc com letra maiuscula são de visibilidade publica(dentro e fora do pacote)
// logo letras minusculas indicam que sao privados dentro do pacote
// ponto no inicio do nome = publico/ ponto e pontoUnderline privado

// por ser publico deve possuir comentario
type Ponto struct {
	x float64
	y float64
}

//calcular catetos
func catetos(a, b Ponto) (cx, cy float64) {
	//obetendo valores absolutos com pacote math/note que funcao recebe 2 instancias da struct Ponto
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return //return limpo/retorno nomeado
}

// metodo publico p/ calcular distancia
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	//retornando a raiz quadrada da soma dos quadrados dos catetos
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
