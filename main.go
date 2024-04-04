package main

import "fmt"
import "github.com/Agusotto1/goDesde0/variables"
 func main(){
	 fmt.Println("Hola Mundo")
	 variables.MuestroEnteros()
	 //variables.RestoVariables()
	estado, texto := variables.ConvertirTexto(1500)
	fmt.Println(texto, estado)
 }