package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kuzin57/FormalPractic/second_practic/adapter"
)

func main() {
	grammarAdapter := adapter.BuildAdapter(os.Args[1], os.Args[2])
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		if grammarAdapter.Read(word) {
			fmt.Println("YES")
			grammarAdapter.Flush()
			continue
		}
		fmt.Println("NO")
		grammarAdapter.Flush()
	}
}
