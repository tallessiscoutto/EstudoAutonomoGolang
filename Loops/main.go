package main

/*
func main1(){

	for{{
	fmt.Println("1 - Calcular IMC")
	fmt.Println("2 - Dizer Olá")
	fmt.Println("0 - Sair ")

	var opcao int
	fmt.Print("Escolha: ")
	fmt.Scanln(&opcao)

	switch opcao {
	case 1:
		var peso float64
		var altura float64
		fmt.Println("Digite seu peso: ")
		fmt.Scanln(&peso)


		fmt.Println("Digite sua altura: ")
		fmt.Scanln(&altura)

		IMC := peso / (altura * altura)

		fmt.Printf("Seu IMC é: %f\n", IMC)

	case 2:
		fmt.Println("Olá!")
	case 0:
		fmt.Println("Saíndo...")
		return
	default:
		fmt.Println("Opção inválida")
}}
}}
*/
import "fmt"

func main(){
	var produtostring []string 
	var produtoint []float64
	var total float64
	var produto string
	var preco float64
		
	for{
		fmt.Println("1 - Adicionar Produto")
		fmt.Println("2 - Mostrar total ")
		fmt.Println("3 - Mostrar Lista de produtos")
		fmt.Println("0 - sair")

		var opcao int
		fmt.Print("Escolha: ")
		fmt.Scanln(&opcao)

		switch opcao{
		case 1:
			for{
				fmt.Println("Digite o nome do produto ou  'sair' pra sair")
				fmt.Scanln(&produto)
				if produto == "sair" {
					preco = 0
					break
				}
				produtostring = append(produtostring, produto)
				fmt.Println("Digite o preço do produto ")
				fmt.Scanln(&preco)
				produtoint = append(produtoint, preco)

				
			}
			
		case 2:
			total = 0
			for i := 0; i< len(produtoint); i++{
				total = total + produtoint[i]
			}
			 
			fmt.Printf("Preço total da compra %f\n",total)

		case 3:
			fmt.Printf("Produtos:\n")
			for i:= 0; i< len(produtostring); i++{
				fmt.Printf("- %s\n",produtostring[i])
			}

		case 0:
			fmt.Printf("Saindo...")
			return
		}
	}
}