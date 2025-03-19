package main
import "fmt"

// parametro n do tipo inteiro, retorna um inteiro
func fib(n int) int {
    if n <= 1 {
        return n
    }

    return fib(n-1) + fib(n-2)
}

func status() {
    for {
        fmt.Println("calma honey")
    }
}

func main() {
    go status() //roda a função com um fluxo separado de execução - go routine
    result := fib(50)
    fmt.Println(result)
}
