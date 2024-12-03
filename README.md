# Juego de Adivinanza con IA en Go

Este proyecto es un simple pero divertido juego de adivinanza desarrollado en **Go** donde una inteligencia artificial (IA) intenta adivinar el número que el jugador tiene en mente.

## Objetivo

La IA intentará adivinar un número que el jugador elige en secreto. El jugador dará pistas ("Mayor", "Menor", "Correcto") y la IA ajustará sus predicciones hasta adivinar correctamente.

## Funcionalidades principales

- El jugador elige un número dentro de un rango (por defecto, entre 1 y 100).
- La IA intenta adivinar el número usando una búsqueda optimizada (búsqueda binaria).
- El jugador proporciona retroalimentación sobre las conjeturas de la IA.
- Se calcula cuántos intentos necesita la IA para adivinar correctamente.

## Requisitos

- Go (1.20 o superior) instalado en tu sistema.

## Ejecución

1. Clona este repositorio o copia el código fuente.
2. Ejecuta el programa:
   ```bash
   go run main.go
