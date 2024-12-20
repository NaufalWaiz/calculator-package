package calculator

import "errors"

func Penjumlahan(a, b float64) float64 {
	return a + b
}

func Pengurangan(a, b float64) float64 {
	return a - b
}

func Perkalian(a, b float64) float64 {
	return a * b
}

func Pembagian(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa membagi dengan 0")
	}
	return a / b, nil
}
