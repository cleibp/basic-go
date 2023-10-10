package main

import (
	"fmt"
)

func main() {
	var nome string = "Cleiton"
	var idade int = 20
	var sexo rune = 'M'
	var peso float64 = 70.5
	var ativo bool = true

	var val1, val2, adicao, subtracao, multiplicacao, modulo int
	var divisao float64

	var idadeTernario int
	var idadeTer string

	var idadeIF int

	var dia int

	var a int

	var b int

	var m int
	var n int

	const PI float64 = 3.14159265

	// Comentário de uma linha

	/*
	 * Comentário
	 * de varias linhas
	 */

	// ESCREVER NA TELA
	fmt.Println("#### ESCREVER NA TELA ####")
	fmt.Println("Olá Mundo")
	fmt.Println()

	// VARIÁVEIS
	fmt.Println("### VARIÁVEIS E TIPOS BÁSICOS ###")
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Sexo:", string(sexo))
	fmt.Println("Peso:", peso)
	fmt.Println("Ativo:", ativo)
	fmt.Println()

	// CONSTANTE
	fmt.Println("### CONSTANTE ###")
	fmt.Println("PI:", PI)
	fmt.Println()

	// OPERACOES
	fmt.Println("#### OPERACOES ####")
	fmt.Print("Informe o valor 1: ")
	fmt.Scan(&val1)

	fmt.Print("Informe o valor 2: ")
	fmt.Scan(&val2)

	adicao = val1 + val2 // Pode usar: (+, -, *, /, %)
	subtracao = val1 - val2
	multiplicacao = val1 * val2
	divisao = float64(val1) / float64(val2)
	modulo = val1 % val2
	fmt.Println("Soma:", adicao)
	fmt.Println("Subtracao:", subtracao)
	fmt.Println("Multiplicacao:", multiplicacao)
	fmt.Println("Divisao:", divisao)
	fmt.Println("Modulo:", modulo)
	fmt.Println()

	// TERNARIO
	fmt.Println("### TERNARIO ###")
	fmt.Print("Digite um número: ")
	fmt.Scan(&idadeTernario)
	if idadeTernario >= 18 {
		idadeTer = "Maior de idade"
	} else {
		idadeTer = "Menor de idade"
	}
	fmt.Println(idadeTer)
	fmt.Println()

	// IF ELSE IF ELSE
	fmt.Println("### IF ELSE IF ELSE ###")
	fmt.Print("Informe a idade: ")
	fmt.Scan(&idadeIF)
	if idadeIF < 12 {
		fmt.Println("CRIANCA")
	} else if idadeIF >= 12 && idadeIF < 18 {
		fmt.Println("ADOLESCENTE")
	} else {
		fmt.Println("ADULTO")
	}
	fmt.Println()

	// CASE
	fmt.Println("### CASE ###")
	fmt.Print("Informe um numero 1 - 7 para semana: ")
	fmt.Scan(&dia)

	switch dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sabado")
	default:
		fmt.Println("Valor nao existente")
	}
	fmt.Println()

	// REPEAT
	fmt.Println("### REPEAT ###")
	fmt.Println("Não tem REPEAT")
	fmt.Println()

	// DO WHILE
	fmt.Println("### DO WHILE ###")
	for a < 10 {
		fmt.Println(a)
		a = a + 1
	}
	fmt.Println()

	// WHILE
	fmt.Println("### WHILE ###")
	b = 0
	for b < 10 {
		fmt.Println(b)
		b = b + 1
	}
	fmt.Println()

	// FOR
	fmt.Println("### FOR ###")
	for c := 0; c <= 10; c++ {
		fmt.Println(c)
	}
	fmt.Println()

  // ARRAY
	fmt.Println("### ARRAY ###")
	numbers := []int{10, 20, 30, 40}

	for _, number := range numbers {
		fmt.Println(number)
	}
	fmt.Println()

	// MATRIZ
	fmt.Println("### MATRIZ ###")
	var matriz [3][3]int

	// Inicialização da matriz
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = i*3 + j + 1
		}
	}

	// Acesso aos elementos da matriz
	fmt.Println("Elementos da matriz:")
	for i := 0; i < 3; i++ {
		row := ""
		for j := 0; j < 3; j++ {
			row += fmt.Sprintf("%d ", matriz[i][j])
		}
		fmt.Println(row)
	}
	fmt.Println()

	// FUNCAO
	fmt.Println("### FUNCAO ###")
	fmt.Print("Digite o valor 1: ")
	fmt.Scan(&m)

	fmt.Print("Digite o valor 2: ")
	fmt.Scan(&n)

	resultado := soma(m, n)
	fmt.Printf("A soma de %d e %d é igual a %d\n", m, n, resultado)
	fmt.Println()

	// PROCEDURE
	fmt.Println("### PROCEDURE ###")
	fmt.Println("Não tem PROCEDURE")
	fmt.Println()

  // PONTEIRO
  fmt.Println("### PONTEIRO ###")
  fmt.Println("Não tem PONTEIRO")
  fmt.Println("Não é necessário liberar memória manualmente, como em C ou C++. A variável será automaticamente coletada pelo coletor de lixo quando não estiver mais em uso.")
  fmt.Println()
}

func soma(a int, b int) int {
	return a + b
}
