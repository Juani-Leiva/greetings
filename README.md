# Saludos en Go

Este paquete proporciona una forma simple de saludos personalizados

## Instalacion

Ejecuta el siguiente comando para instalar el paquete

```bash
go get -u github.com/Juani-Leiva/greetings.git
```

## Uso

Aqui tienes un ejemplo de como utilizar el paquete en tu codigo:

```go
package main
import (
    "fmt"
    "github.com/Juani-Leiva/greetings.git"
)
func main(){
    message, err := greetings.Hello("Alex")

    if err != nil {
        fmt.Println ("Ocurrio un error" , err)
        return
    }
    fmt.Println(message)
}
```
