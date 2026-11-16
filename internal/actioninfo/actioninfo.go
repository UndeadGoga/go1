package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("Запись %q: %v", data, err)
			continue
		}
		report, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Запись %q: %v", data, err)
			continue
		}
		fmt.Println(strings.TrimSuffix(report, "\n"))
	}
}
