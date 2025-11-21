package main

import "fmt"

/*func main(){
	//Variaveis que utilizam := identificam o tipo de dado automaticamente
	var nome string = "Talles"
	idade := 22
	altura := 1.85
	ativo := true

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Ativo:", ativo)
}*/

//Exercicio
func main(){
	var nome string 
	var idade int
	altura := 1.85
	ativo:= true

	fmt.Print("Digite o seu nome: ")
	fmt.Scanln(&nome)
	fmt.Print("Digite a sua idade: ")
	fmt.Scanln(&idade)
	fmt.Printf("Nome: %s Idade: %d Altura: %.2f\n Ativo: %t", nome, idade, altura, ativo)

}