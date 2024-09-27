package utils

import (
	"errors"
	"os"
)

func ReadFile(fileName string) ([]byte, error) {
	if IsEmty(fileName) {
		return []byte(nil), errors.New("Dosya Adı Boş Olamaz")
	}
	byte, err := os.ReadFile(fileName)
	CheckError(err)
	return byte, nil
}
