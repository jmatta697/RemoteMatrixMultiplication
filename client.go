package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"shared" //Path to the package contains shared struct
	"strconv"
	"strings"
)

type MatrixMultRPC struct {
	client *rpc.Client
}

func (t *MatrixMultRPC) MultiplyMatrix(matrix1, matrix2 [][]int) [][]int {
	args := &shared.MatrixArgs{M1: matrix1, M2: matrix2}
	var reply [][]int
	err := t.client.Call("MatrixMultiply.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func getMatrixSizeFromUser() (int64, error) {

	reader := bufio.NewReader(os.Stdin)
	var intStr string
	fmt.Println("What are the size of the matrices? \n " +
		"(For example, if two 6 X 6 matrices are desired, enter '6'.)")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		intStr = strings.Replace(text, "\n", "", -1)

		//check if input is an int
		i, err := strconv.ParseInt(intStr, 10, 32)
		if err != nil {
			fmt.Println("Enter a valid INTEGER!")
		} else {
			//check if number is greater than zero
			if i <= 0 {
				fmt.Println("Enter an integer GREATER THAN ZERO.")
			} else {
				break
			}
		}
	}
	return strconv.ParseInt(intStr, 10, 32)

}

func getIntegerFromUser() (int64, error)  {
	reader := bufio.NewReader(os.Stdin)
	var userInt string

	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		userInt = strings.Replace(text, "\n", "", -1)
		//check if input is an int
		_, err := strconv.ParseInt(userInt, 10, 32)
		if err != nil {
			fmt.Print("ERROR: Enter a valid INTEGER! > ")
		} else {
			//check if number is greater than zero
			break
		}
	}
	return strconv.ParseInt(userInt, 10, 32)
}

func buildMatrixFromUserInput(matricesSize int) [][]int {
	var matrix [][]int

	for i:=0; i<int(matricesSize); i++ {
		fmt.Printf("*** ROW %d ***\n", i+1)
		var tempRow []int
		for j:=0; j<int(matricesSize); j++{
			fmt.Print("Enter an integer: ")
			inputInt, _ := getIntegerFromUser()
			tempRow = append(tempRow, int(inputInt))
		}
		matrix = append(matrix, tempRow)
	}
	return matrix
}

func main() {

	matricesSize, _ := getMatrixSizeFromUser()
	fmt.Println(matricesSize)
	fmt.Println(fmt.Sprintf("Matrices Size: %d", matricesSize))

	//declare int 2d matrix
	fmt.Println("----- MATRIX 1 -----")
	firstMatrix := buildMatrixFromUserInput(int(matricesSize))
	fmt.Println("----- MATRIX 2 -----")
	secondMatrix := buildMatrixFromUserInput(int(matricesSize))

	fmt.Println(firstMatrix)
	fmt.Println(secondMatrix)

	// Tries to connect to localhost:1234 (The port on which rpc server is listening)
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connecting:", err)
	}
	// Create a struct, that mimics all methods provided by interface.
	// It is not compulsory, we are doing it here, just to simulate a traditional method call.
	matrixMultiply := &MatrixMultRPC{client: rpc.NewClient(conn)}

	fmt.Println(matrixMultiply.MultiplyMatrix(firstMatrix, secondMatrix))
}
