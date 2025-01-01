package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(texts []string) ([]float64, error) {
	var floats []float64
	for _, text := range texts {
		floatPrice, err := strconv.ParseFloat(text, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
