package main

import "fmt"

func main(){

	var peso float64
	var altura float64

	fmt.Println("Digite seu peso: ")
	fmt.Scanln(&peso)


	fmt.Println("Digite sua altura: ")
	fmt.Scanln(&altura)

	IMC := peso / (altura * altura)

	fmt.Printf("Seu IMC é: %.2f \n", IMC)

	if IMC < 18.5{
		fmt.Printf("Classificação: abaixo do peso ")
	}else if IMC < 24.9{
		fmt.Printf("Classificação: peso normal ")
	}else if IMC < 29.9{
		fmt.Printf("Classificação: Sobrepeso ")
	}else if IMC < 34.9 {
		fmt.Printf("Classificação: Obesidade 1 ")
	}else if IMC < 39.9{
		fmt.Printf("Classificação: Obesidade 2 ")
	}else{
		fmt.Printf("Classificação: Obesidade 3 ")
		
}
}