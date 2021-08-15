package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type error interface {
	Error() string
}

func GetNumberInput(message string) (value int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(message)
	scanner.Scan()

	text := scanner.Text()
	number, err := strconv.ParseInt(text, 10, 0)

	if err != nil {
		fmt.Println("You must enter a valid number!")
		return 0, err
	}

	return int(number), nil
}
