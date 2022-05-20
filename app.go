package main

import "fmt"

func main() {
	matrixA := [3][3]int{
		{1, 2, 4},
		{3, 7, 2},
		{5, 8, 4},
	}
	matrixB := [3][3]int{
		{1, 2, 2},
		{2, 2, 1},
		{1, 1, 1},
	}
	matrixC := [3][3]int{
		{11, 14, 12},
		{12, 45, 30},
		{23, 45, 12},
	}

	matrixList := [3][3][3]int{
		matrixA,
		matrixB,
		matrixC,
	}
	for i, m := range matrixList {
		matrixPrint(i+1, m)
	}

	var matrixD [3][9]int
	for x, m := range matrixList {
		for j := 0; j < 3; j++ {
			matrixD[x][j*3] = m[j][0]
			matrixD[x][(j*3)+1] = m[j][1]
			matrixD[x][(j*3)+2] = m[j][2]
		}
	}
	matrixResult(matrixD)
}

func matrixResult(mtx [3][9]int) {
	fmt.Println("Hasil Matriks Gabungan dari ketiga matriks")
	for i := 0; i < 3; i++ {
		fmt.Print("| ")
		for j := 0; j < 9; j++ {
			fmt.Printf("%-2d | ", mtx[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("")
}
func matrixPrint(title int, mtx [3][3]int) {
	fmt.Println("Matriks ", title)
	for i := 0; i < 3; i++ {
		fmt.Print("| ")
		for j := 0; j < 3; j++ {
			fmt.Print(mtx[i][j], " | ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}
