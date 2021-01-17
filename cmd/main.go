package main

import (
	"fmt"
	"time"
)

const (
	//formato universal de hora: Mon Jan 2 15:04:05 2006 MST
	//para recordar facilmente:  day  1  2  3  4  5   6   -7

	//se puede utilizar segun lo que se requiera, y en el orden que quiera
	layout    = "Mon Jan 2 15:04:05 2006"
	layout2   = "15:04 2006"
	layout3   = "2006-Jan-02 15:04 Mon"
	shortForm = "2006-Jan-02"
)

func main() {
	//la cadena recibida debe tener el formato especifico al que se debe convertir, en este caso layout2
	time1 := "12:35 2070"
	//convertir la cadena a un objeto de tipo time
	t, err := time.Parse(layout2, time1)
	if err != nil {
		panic(err)
	}
	// t := time.Now().Format("15:04")
	//imprimir formatiando porque si no imprime de manera incorrecta con otros valores por defecto
	fmt.Println(t.Format(layout2))
	// fmt.Printf("type: %T, value:%v ", t, t)
	// stringmode := string(t)

}
