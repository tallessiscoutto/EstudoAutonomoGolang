package main

import "fmt"

//Exercicio 1
/*
func main(){

	var peso float64
	var altura float64

	fmt.Println("Digite seu peso: ")
	fmt.Scanln(&peso)


	fmt.Println("Digite sua altura: ")
	fmt.Scanln(&altura)

	IMC := peso / (altura * altura)

	fmt.Printf("Seu IMC é: %.2f\n", IMC)


}*/
//Exercicio 2
func main(){
	var distancia float64
	var consumo float64
	var precoLitro float64

	fmt.Println("Digite a distancia percorrida(KM): ")
	fmt.Scanln(&distancia)
	fmt.Println("Digite o consumo medio do Carro KM/Litro: ")
	fmt.Scanln(&consumo)
	fmt.Println("Digite o preço por litro: ")
	fmt.Scanln(&precoLitro)

	litros := distancia / consumo

	custo := litros * precoLitro

	fmt.Printf("Você vai gastar: %.2f litros de combutível.\n",litros)
	fmt.Printf("O custo total da viagem é: R$ %.2f\n",custo)


}
