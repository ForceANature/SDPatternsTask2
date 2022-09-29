package main

import ("fmt")

type Coffee interface {
	getDescription() string
	cost() float64
}
type Espresso struct{}

func (i Espresso) cost() float64 {
	return 10.99
}
func (i Espresso) getDescription() string {
	return "Espresso: 10.99$\n"
}

type Cappuccino struct{}

func (ip Cappuccino) cost() float64 {
	return 14.99
}
func (ip Cappuccino) getDescription() string {
	return "Cappuccino : 14.99$\n"
}

type Latte struct{}

func (ipm Latte) cost() float64 {
	return 8.99
}
func (ipm Latte) getDescription() string {
	return "Latte : 8.99$\n"
}

type Sugar struct {
	coffee Coffee
}

func (ap Sugar) cost() float64 {
	p := ap.coffee.cost()
	return p + 0.99
}
func (ap Sugar) getDescription() string {
	return ap.coffee.getDescription() + "\t+ Sugar: 0.99$\n"
}

type Whip struct {
	coffee Coffee
}

func (c Whip) cost() float64 {
	p := c.coffee.cost()
	return p + 0.19
}

func (c Whip) getDescription() string {
	return c.coffee.getDescription() + "\t+ Whip: 0.19$\n"
}

type Honey struct {
	coffee Coffee
}

func (aw Honey) cost() float64 {
	p := aw.coffee.cost()
	return p + 1.99
}

func (aw Honey) getDescription() string {
	return aw.coffee.getDescription() + "\t+ Honey: 1.99$\n"
}

type Mint struct {
	coffee Coffee
}

func (c Mint) cost() float64 {
	p := c.coffee.cost()
	return p + 49.99
}
func (c Mint) getDescription() string {
	return c.coffee.getDescription() + "\t+ Mint: 0.19$\n"
}
func Menu() {
	fmt.Println("===Action===\n1 - order\n2 - leave\n============")
	var frsw int
	fmt.Scanf("%d\n", &frsw)
	switch frsw {
	case 1:
		chooseBase()
	default:
		fmt.Println("Goodbye!")
	}
}
func condimenList(b Coffee) {
	fmt.Println("===anything=else?===\n1 - add condiment\n2 - back to menu\n3 - end transaction\n============")
	var frsw int
	fmt.Scanf("%d\n", &frsw)
	switch frsw {
	case 1:
		Condiments(b)
	case 2:
		Menu()
	default:
		fmt.Println("Goodbye!")
	}
}
func chooseBase() {
	fmt.Println("===Menu===\n1 - Espresso\n2 - Cappuccino\n3 - Latte\n============")
	var frsw int
	fmt.Scanf("%d\n", &frsw)
	var base Coffee
	switch frsw {
	case 1:
		base = Espresso{}
	case 2:
		base = Cappuccino{}
	case 3:
		base = Latte{}
	}
	check(base)
	condimenList(base)
}
func Condiments(base Coffee) {
	fmt.Println("===List of Condiments===\n1 - Sugar\n2 - Whip\n3 - Honey\n4 - Mint\n============")
	var frsw int
	fmt.Scanf("%d\n", &frsw)
	var baseWithComponent Coffee
	switch frsw {
	case 1:
		baseWithComponent = Sugar{base}
	case 2:
		baseWithComponent = Whip{base}
	case 3:
		baseWithComponent = Honey{base}
	case 4:
		baseWithComponent = Mint{base}
	}
	check(baseWithComponent)
	condimenList(baseWithComponent)
}
func check(b Coffee) {
	fmt.Println("\n===Your order===")
	fmt.Println(b.getDescription())
	fmt.Printf("Total cost: %v$\n", fmt.Sprintf("%.2f", b.cost()))
	fmt.Println("============")
}
func main() {
	fmt.Println("Welcome to MugShot!")
	fmt.Println("The place to mugged or shot")
	Menu()
}
