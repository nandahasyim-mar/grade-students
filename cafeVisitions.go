package main

import (
	"fmt"
	// "strconv"
)

func cafeVisition(name string, age int, money int) string {
	pesan := ""
	if name == "" {
		pesan = "Anda tidak boleh masuk!"
	} else if age < 17 {
		if money < 50000 {
			pesan = "Uang anda tidak cukup. Anda harus pulang!"
		} else {
			pesan = "Anda bisa pesan minum. Sisa uang anda"
		}
	} else if age >= 17 {
		if money < 300000 {
			pesan = "Uang anda tidak cukup. Anda harus pulang!"
		} else {
			pesan = "Anda bisa pesan anggur. Sisa uang anda"
		}
	}

	return pesan
}

func main() {
	fmt.Println(cafeVisition("Nanda Hasyim Marfianshar", 24, 50000000000))
}
