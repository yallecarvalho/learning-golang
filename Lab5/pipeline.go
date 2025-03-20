package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

//below random string functions are based on Jon Calhoun code
const charset = "abcdefghijklmnopqrstuvwxyz1234567890"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return StringWithCharset(length, charset)
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

//gera 100 strings aleatorias de 5 caracteres usando o RandString e envia cada uma para o canal out, depois, fecha o canal sinalizando que nao haverá mais dados
func produce(out chan string){ //envia strings pra um canal
    for i := 0; i < 100; i++ {
            out <- RandString(5) //guarda os valores gerados por RandString em out
    }
	close(out) //fecha o canal pra ter um livelock
}

func filtrarAlphas(in chan string, out chan string) { //filtra de um canal de input e envia por um canal de output
	for word := range in { // le strings do canal de entrada (in)
		if isLetter(word){ //verifica se a string contem apenas letras
			out <- word //envia pro canal de saida
		}
	}
	close(out)  // IMPORTANTE fechar o canal depois que todos os dados forem processados
}


func main() {
    rawContent := make(chan string)  //cria canal pra strings nao filtradas
	filtrarAlphas := make(chan string)  //cria canal pra strings filtradas (somente letras)

	go produce(rawContent)  //inicia a go routine para produzir strings
	go filtrarAlphas(rawContent, filteredAlphas)   //inicia a goroutine para filtrar strings

	// Consumir e imprimir os valores filtrados
	for alpha := range filetredAlphas {   //le do canal filtrado e imprime
		fmt.Printf("alpha:  <%s>\n", <- alpha)
	}

}
