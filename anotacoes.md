Go

- todo fluxo de execução é uma goroutine. Em cada programa há pelo menos UMA goroutine. goroutine seria similar a thread, porém mais leve.
Novas goroutines podem ser criadas com "go". Você coloca go na frente da chamada da função.


For infinito em go:
```
package main
import "fmt"

func main(){
    for {
        fmt.Println(i)
        i+=1
    }
}
```

For normal:

```
package main
import "fmt"

func main() {
    banda :=[]string{"Axel", "Slash"} //criacao de array com dados
    for _, integrante := range banda {
        fmt.Println("Integrante:", integrante)
    }
}

Canais - permitem receber e enviar dados entre goroutines.
```
ch := make(chan int) //criacao de um canal ch de inteiros
ch <- x // envia o valor x pro canal ch
y = <- ch // recebe o valor e atribui o valor recebido a y
<- ch // recebe o valor e descarta
```