package main

import (
	"bufio"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
)

// função para arrendondar caracteres do tipo number
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// função para fixer um numero de caracteres para um number
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func main() {
	fmt.Println("Abrindo servidor...")
	//declara a porta
	args := os.Args
	fmt.Print(args)
	//retorna erro caso não tenha sido declarada nenhuma porta
	if len(args) == 1 {
		fmt.Println("Por favor digite a porta que deseja abrir o servidor")
		return
	}
	// escuta na porta passada via terminal
	ln, err := net.Listen("tcp", args[1])
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("\nTCP Server inicializado na porta", args[1])
	// aceita conexão na porta seleciona
	conn, err := ln.Accept()
	if err != nil {
		fmt.Print(err)
	}
	// se tudo ocorreu corretamente servidor é aberto
	fmt.Print("Cliente conectado\n")

	//// ao final fecha a conexão
	defer ln.Close()

	// loop eterno
	for {
		mensagem, _ := bufio.NewReader(conn).ReadString('\n')
		if strings.ToUpper(strings.TrimSpace(string(mensagem))) == "SAIR" {
			fmt.Println("servidor encerrado.")
			return
		}
		if strings.ToUpper(strings.TrimSpace(string(mensagem))) == "REGISTRAR" {

			quantidade, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			quantidadeNumber, _ := strconv.Atoi(strings.TrimSpace(quantidade))

			for i := 0; i < quantidadeNumber; i++ {

				nomeAluno, _ := bufio.NewReader(conn).ReadString('\n')
				nomeString := (strings.TrimSpace(nomeAluno))

				nota1, _ := bufio.NewReader(conn).ReadString('\n')
				nota1string, _ := strconv.ParseFloat(strings.TrimSpace(nota1), 64)

				nota2, _ := bufio.NewReader(conn).ReadString('\n')
				nota2string, _ := strconv.ParseFloat(strings.TrimSpace(nota2), 64)

				nota3, _ := bufio.NewReader(conn).ReadString('\n')
				nota3string, _ := strconv.ParseFloat(strings.TrimSpace(nota3), 64)
				//imprime a mesangem no servidor
				var media = toFixed((nota1string+nota2string+nota3string)/3, 2)
				fmt.Println("----------------------------\n\nNome: ", nomeString)
				fmt.Println("\nNotas: ", nota1string, "|", nota2string, "|", nota3string)
				fmt.Println("\nMédia: ", media, "\n\n----------------------------")
				//imprime a mensagem no client
			}
		}
	}
}
