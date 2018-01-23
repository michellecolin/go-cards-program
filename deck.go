package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // _ dont want to use the variable
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) { //can return multiple values
	return d[:handSize], d[handSize:]
}

func (d deck) print() { //receiver  //this or self the receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}
