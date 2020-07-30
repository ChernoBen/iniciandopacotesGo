package main

/*
Criando pacotes em src/github/ChernoBen
para poder ser utilizando quando e onde quiser
*/
import (
	"fmt"

	"github.com/ChernoBen/area"
)

func main() {
	result := area.Circ(6.0)
	fmt.Println(area.Rect(5.0, 6.0))
	fmt.Println(result)
}
