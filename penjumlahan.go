package main 
import (
	"fmt"
)

func main() {
	fmt.Println("================================")
	fmt.Println("       Cek Predikat Nilai Anda")
	fmt.Println("================================")

	
	var a,b int
	var hasil int

	fmt.Scan(&a,&b)

	hasil = a + b
	fmt.Println("Nilai Anda: ", hasil)
	
	
	if hasil >= 80 {
		fmt.Println("Predikat anda: A")
	} else if hasil >= 70 {
		fmt.Println("Predikat anda: B")
	} else if hasil >= 50 {
		fmt.Println("Predikat anda: C")
	} else if hasil >= 40 {
		fmt.Println("Predikat anda: D")
	} else {
		fmt.Println("Predikat anda: E")
	}
}