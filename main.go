package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Bienvenidx a go guess ahre!")
	fmt.Println("elegi un número en secreto entre 1 y 100, y la IA va a intentar adivinarlo.")

	low := 1
	high := 100
	attempts := 0
	reader := bufio.NewReader(os.Stdin)

	for {
		guess := (low + high) / 2
		attempts++

		fmt.Printf("¿Es %d? (responde: mayor, menor, correcto): ", guess)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "correcto" {
			fmt.Printf("La IA adivinó tu número (%d) en %d intentos porque es una ia chill de cojones\n", guess, attempts)
			break
		} else if response == "mayor" {
			low = guess + 1
		} else if response == "menor" {
			high = guess - 1
		} else {
			fmt.Println("Respuesta no válida. Por favor responde con: mayor, menor o correcto.")
		}

		if low > high {
			fmt.Println("Parece que hubo un error. ¿Seguiste las reglas? Reinicia el juego.")
			break
		}
	}
}
