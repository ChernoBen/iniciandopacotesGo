package main

import "fmt"

func main() {
	//inicializando chamando funcoes dentro do pacote
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}
	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}

//rodando em windows : go run main.go reta.go
//linux go run *.go
