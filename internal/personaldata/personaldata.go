package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64 // кг
	Height float64 // метры
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
