package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

var (
	errExist = 1
	sucExist = 0
	META_COMMAND_SUCCESS = 0
	META_COMMAND_UNRECOGINZED_COMMAND = 1
)



type InputBuffer struct{
	buffer string
	bufferLen int
	inputLen int
}


func main(){
	inputBuffer := new(InputBuffer)
	for {
		printPrompt()
		readInput(inputBuffer)
		if inputBuffer.buffer[0] == '.' {
			switch commandRet := doMetaCommand(inputBuffer); commandRet {
			case META_COMMAND_SUCCESS:
				continue
			case META_COMMAND_UNRECOGINZED_COMMAND:
				fmt.Printf("Unrecoginzed command %s. \n", inputBuffer.buffer)
			}
		}
	}
}


func printPrompt() {
	fmt.Printf("db > ")
}


func doMetaCommand(inputBuffer *InputBuffer) int {
	if inputBuffer.buffer == ".exit" {
		os.Exit(sucExist)
	} else {
		return META_COMMAND_UNRECOGINZED_COMMAND
	}
	return META_COMMAND_SUCCESS
}

func readInput(inputBuffer *InputBuffer) {
        reader := bufio.NewReader(os.Stdin)
        text, err := reader.ReadString('\n')
        inputBuffer.inputLen = len(inputBuffer.buffer)
        if err != nil {
                fmt.Println(err)
                os.Exit(errExist)
        }
        inputBuffer.buffer = strings.Trim(text, "\n")
}