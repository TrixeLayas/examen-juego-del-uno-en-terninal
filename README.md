# Juego de UNO en Go - Explicación Técnica

Este documento explica detalladamente la implementación del juego UNO contenida en el archivo `examen.go`. El código ha sido estructurado para cumplir con los requerimientos de lógica, control de flujo y estructuras de datos.

---

## 1. Estructuras de Datos 

El juego se basa en tres estructuras principales que organizan la información:

- **`Card` **: Representa una carta individual con su `Color` (Rojo, Verde, Azul, Amarillo) y su `Value` (números del 0 al 9).
- **`Player` **: Almacena el `Name` del jugador y su `Hand` (mano), que es un slice de cartas.
- **`Game` **: Controla el estado global, incluyendo el `Deck` (mazo principal), los `Players`, la `DiscardPile` (pila de descarte) y el `Turn` actual.

---

## 2. Inicialización del Mazo (Líneas 26-39)

La función `InitDeck` es la encargada de preparar el juego:
- : Utiliza dos bucles anidados para crear una carta de cada valor (0-9) para cada uno de los cuatro colores.
- : Mezcla el mazo de forma aleatoria usando `rand.Shuffle`.

---

## 3. Repartición de Cartas (Líneas 41-50)

La función `DealCards`:
- : Reparte 5 cartas a cada jugador llamando repetidamente a `DrawCard`.
- : Coloca la primera carta en la pila de descarte para iniciar el juego.

---

## 4. Lógica de Robo y Reciclaje 

La función `DrawCard` es crítica para el flujo continuo:
- : Si el mazo se agota, toma todas las cartas de la pila de descarte (excepto la última), las pone de nuevo en el mazo y las baraja. Esto evita que el juego se detenga.
- : Extrae la carta superior del mazo y la devuelve.

---

## 5. Interfaz Visual y Colores 

- **`PrintCard` **: Utiliza códigos ANSI para que las cartas se impriman con su color real en la terminal, mejorando la experiencia del usuario.
- **`Wait` **: Implementa una pausa entre turnos y "limpia" la pantalla con saltos de línea para que un jugador no pueda ver las cartas del otro.

---

## 6. Lógica de Jugada y Validación (Líneas 100-164)

- **`IsValidMove` **: Compara la carta que el jugador quiere usar con la que está en la mesa. Es válida si coinciden en color **o** en valor.
- **`PlayTurn`**: Es el corazón del juego:
    - Muestra el estado actual .
    - **Auto-robo **: Si el jugador no tiene cartas válidas, roba una automáticamente. Si la robada sirve, se juega de inmediato.
    - **Selección **: Pide al usuario el número de carta. Valida que la selección sea correcta y que la carta cumpla las reglas de UNO.
    - **Condición de Victoria **: Si al jugador le quedan 0 cartas, la función retorna `true`, señalando el fin del juego.

---

## 7. Ejecución Principal (Líneas 170-193)

El `main`:
- Inicializa la semilla aleatoria para que cada partida sea diferente .
- Crea los jugadores y arranca el ciclo de juego  que continúa hasta que alguien gana.

---

## Cómo ejecutar el juego

Asegúrate de tener Go instalado y ejecuta:

```bash
go run examen.go
```
