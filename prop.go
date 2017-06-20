package goutils

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"
	"io"
	"log"
)

func ParseProperties(properties []byte) (map[string]string, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(properties))
	cfg := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' {
			continue
		} else {
			keyvalue := strings.Split(line, "=")
			if len(keyvalue) >= 2 {
				key := strings.TrimSpace(keyvalue[0])
				value := strings.TrimSpace(line[len(key)+1:])
				var tmp []rune
				for _, r := range value {
					if unicode.IsPrint(r) {
						tmp = append(tmp, r)
					}
				}
				cfg[key] = string(tmp)
			}
		}
	}
	err := scanner.Err()
	if err != nil && err != io.EOF {
		log.Println("mf:error while read input data", err)
		return nil, err
	}
	return cfg, nil
}

