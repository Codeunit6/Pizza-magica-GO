package main

import (
	"fmt"
	"math/rand"
	"time"
)

func comer(pizza chan int, done chan bool, ingredientes []string) {
	var numero int
	for {
		numero = rand.Intn(3)
		j, more := <-pizza
		if more {
			fmt.Println("Comiendo rebanada numero: ", j, " es pizza ", ingredientes[numero])
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("Pizza terminada :(")
			time.Sleep(5 * time.Second)
			done <- true
			return
		}
	}
}
func crearpizza(pizza chan int) {
	for j := 1; j <= 50; j++ {
		pizza <- j
	}
	return
}
func main() {
	//arreglo de ingrediente de pizza
	ingredientes := []string{"peperoni", "champiñones", "mexicana"}
	// Canales, parametros [int, almacenamiento buffer]
	pizza := make(chan int, 50)
	done := make(chan bool, 2)

	//Parametro [channel pizza 'int']
	crearpizza(pizza)
	//Go rutina parametros [channel pizza 'int', channel done 'bool=true|false']
	go comer(pizza, done, ingredientes)

	// Cerrar canal parametro [channel pizza 'int']
	close(pizza)
	// Extraer del canal done
	<-done
	// Recrear funcion
	defer main()
}
