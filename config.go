package goutils

import (
	"log"
	"strconv"
)

type Config map[string]string

func (c Config) GetIntParameter(key string, def int) int {
	value := c[key]
	if value == "" {
		return def
	} else {
		valueInInt, err := strconv.Atoi(value)
		if err != nil {
			log.Println("error parse ", key, ". use default ", def)
			return def
		} else {
			return valueInInt
		}
	}
}