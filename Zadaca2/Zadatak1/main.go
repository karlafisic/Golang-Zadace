package main

import "fmt"

type Pravokutnik struct {
	sirina float32
	visina float32
}
type Adresa struct {
	grad  string
	ulica string
}

func main() {

	pravokutnik1 := Pravokutnik{
		visina: 5.7,
		sirina: 2.1,
	}

	pravokutnik1.Povrsina()
	pravokutnik1.Opseg()

}

func (p Pravokutnik) Povrsina() {
	var pov float32
	pov = p.sirina * p.visina
	fmt.Printf("Površina pravokutnika je: %f", pov)
}

func (p Pravokutnik) Opseg() {
	var ops float32
	ops = 2 * (p.sirina + p.visina)
	fmt.Printf("\nOpseg pravokutnika je: %f", ops)
}

