package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardValues { // _ dont want to use the variable
		for _, value := range cardSuits {
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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //last parameter is a permisison code
}

func newDeckFromfile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil { //nil = null
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() {

	for i := range d {
		newIndex := rand.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i] //swapping places
	}
}
