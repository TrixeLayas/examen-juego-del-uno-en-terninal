//marco antonio vizcarra valle #25760062

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Color string
	Value string
}

type Player struct {
	Name string
	Hand []Card
}

type Game struct {
	Deck        []Card
	Players     []Player
	DiscardPile []Card
	Turn        int
}

func (g *Game) InitDeck() {
	colors := []string{"Rojo", "Verde", "Azul", "Amarillo"}
	values := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, c := range colors {
		for _, v := range values {
			g.Deck = append(g.Deck, Card{Color: c, Value: v})
		}
	}

	rand.Shuffle(len(g.Deck), func(i, j int) {
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	})
}

func (g *Game) DealCards() {
	for i := 0; i < len(g.Players); i++ {
		for j := 0; j < 5; j++ {
			g.Players[i].Hand = append(g.Players[i].Hand, g.DrawCard())
		}
	}
	// Inicializar pila de descarte con una carta del mazo
	g.DiscardPile = append(g.DiscardPile, g.DrawCard())
}

func (g *Game) DrawCard() Card {
	if len(g.Deck) == 0 {
		// Si el mazo se agota, barajamos la pila de descarte (excepto la última carta)
		if len(g.DiscardPile) <= 1 {
			fmt.Println("¡No quedan más cartas en el juego!")
			return Card{Color: "N/A", Value: "N/A"}
		}
		topCard := g.DiscardPile[len(g.DiscardPile)-1]
		g.Deck = g.DiscardPile[:len(g.DiscardPile)-1]
		g.DiscardPile = []Card{topCard}

		rand.Shuffle(len(g.Deck), func(i, j int) {
			g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
		})
		fmt.Println("--- Mazo agotado, barajando pila de descarte ---")
	}

	card := g.Deck[len(g.Deck)-1]
	g.Deck = g.Deck[:len(g.Deck)-1]
	return card
}

func PrintCard(c Card) string {
	colorCode := ""
	switch c.Color {
	case "Rojo":
		colorCode = "\033[31m"
	case "Verde":
		colorCode = "\033[32m"
	case "Amarillo":
		colorCode = "\033[33m"
	case "Azul":
		colorCode = "\033[34m"
	}
	reset := "\033[0m"
	return fmt.Sprintf("%s[%s %s]%s", colorCode, c.Color, c.Value, reset)
}

func (g *Game) Wait() {
	fmt.Print("\nPresiona Enter para continuar...")
	fmt.Scanln()
	// "Limpiar" pantalla con saltos de línea
	for i := 0; i < 30; i++ {
		fmt.Println()
	}
}

func (g *Game) IsValidMove(c Card) bool {
	topCard := g.DiscardPile[len(g.DiscardPile)-1]
	return c.Color == topCard.Color || c.Value == topCard.Value
}

func (g *Game) PlayTurn() bool {
	player := &g.Players[g.Turn]
	topCard := g.DiscardPile[len(g.DiscardPile)-1]

	fmt.Printf("\n--- %s ---\n", player.Name)
	fmt.Printf("Carta en mesa: %s\n", PrintCard(topCard))
	fmt.Println("Tu mano:")
	for i, card := range player.Hand {
		fmt.Printf("%d: %s\n", i+1, PrintCard(card))
	}

	hasValidMove := false
	for _, card := range player.Hand {
		if g.IsValidMove(card) {
			hasValidMove = true
			break
		}
	}

	if !hasValidMove {
		fmt.Println("No tienes movimientos válidos. Robando carta...")
		newCard := g.DrawCard()
		fmt.Printf("Has robado: %s\n", PrintCard(newCard))
		if g.IsValidMove(newCard) {
			fmt.Println("¡La carta robada es jugable! Jugándola automáticamente...")
			g.DiscardPile = append(g.DiscardPile, newCard)
		} else {
			player.Hand = append(player.Hand, newCard)
		}
		return false
	}

	var choice int
	for {
		fmt.Print("Selecciona el número de la carta a jugar: ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > len(player.Hand) {
			fmt.Println("Selección inválida. Intenta de nuevo.")
			continue
		}

		selectedCard := player.Hand[choice-1]
		if g.IsValidMove(selectedCard) {
			// Jugar la carta
			g.DiscardPile = append(g.DiscardPile, selectedCard)
			// Eliminar de la mano
			player.Hand = append(player.Hand[:choice-1], player.Hand[choice:]...)
			fmt.Printf("Has jugado: %s\n", PrintCard(selectedCard))
			break
		} else {
			fmt.Println("Esa carta no se puede jugar. Debe coincidir color o valor.")
		}
	}

	if len(player.Hand) == 0 {
		fmt.Printf("\n¡Felicidades %s! ¡Has ganado!\n", player.Name)
		return true
	}

	return false
}

func (g *Game) NextTurn() {
	g.Turn = (g.Turn + 1) % len(g.Players)
}

func main() {
	// Inicializar generador de números aleatorios
	rand.Seed(time.Now().UnixNano())

	game := Game{
		Players: []Player{
			{Name: "Jugador 1"},
			{Name: "Jugador 2"},
		},
	}

	fmt.Println("¡Bienvenido al Juego UNO!")
	game.InitDeck()
	game.DealCards()

	for {
		gameOver := game.PlayTurn()
		if gameOver {
			break
		}
		game.Wait()
		game.NextTurn()
	}
	fmt.Println("¡Gracias por jugar!")
}
