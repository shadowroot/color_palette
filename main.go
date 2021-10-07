package main

import (
	"bufio"
	"fmt"
	"os"

	"bitbucket.org/gekky/color_palette/color_palette"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text_input, _ := reader.ReadString('\n')
	cg := color_palette.ParseTextInputString(text_input)
	color, _ := cg.GenerateTematicColorString()
	fmt.Println("Color: ", color)
}
