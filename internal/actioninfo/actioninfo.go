package actioninfo

import (
	"fmt"
	"log"
)

// DataParser разбирает запись об активности и формирует отчёёт
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(info)
	}
}
