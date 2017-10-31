package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

const (
	errExist = iota
	sucExist 
)

type metaCommandCode int
const (
	_ metaCommandCode = iota
	metaCommandSuccess
	metaCommandUnrecoginzedCommand
)

type prepareResultCode int
const (
	_ prepareResultCode = iota
	prepareSuccess
	prepareUnrecognizedStatement
)


type InputBuffer struct {
	buffer string
	bufferLen int
	inputLen int
}

type statementType int
const (
	_ statementType = iota
	insertType
	selectType
	deleteType
)

type Statement struct {
	sType statementType
}

func prepareStatement(inputBuffer *InputBuffer, statement *Statement) prepareResultCode {
	if len(inputBuffer.buffer) < 6 {
		return prepareUnrecognizedStatement
	}
	if inputBuffer.buffer[:6] == "insert" {
		statement.sType = insertType
		return prepareSuccess
	}
	if inputBuffer.buffer[:6] == "select" {
		statement.sType = selectType
		return prepareSuccess
	}
	if inputBuffer.buffer[:6] == "delete" {
		statement.sType = deleteType
		return prepareSuccess
	}
	return prepareUnrecognizedStatement
}

func executeStatement(statement *Statement) {
	switch statement.sType {
	case selectType:
		fmt.Println("do select")
	case deleteType:
		fmt.Println("do delete")
	case insertType:
		fmt.Println("do insert")
	}
}


func main(){
	inputBuffer := new(InputBuffer)
	for {
		printPrompt()
		readInput(inputBuffer)
		if inputBuffer.buffer[0] == '.' {
			switch doMetaCommand(inputBuffer) {
			case metaCommandSuccess:
			case metaCommandUnrecoginzedCommand:
				fmt.Printf("Unrecoginzed command '%s'. \n", inputBuffer.buffer)
			}
		}

		// prepare statement
		statement := new(Statement)
		switch prepareStatement(inputBuffer, statement) {
		case prepareSuccess:
		case prepareUnrecognizedStatement:
			fmt.Printf("Unrecoginzed keyword at start of '%s'. \n", inputBuffer.buffer)
			continue
		}

		// execute statement
		executeStatement(statement)
		fmt.Println("Executed.")
	}
}


func printPrompt() {
	fmt.Printf("db > ")
}


func doMetaCommand(inputBuffer *InputBuffer) metaCommandCode {
	if inputBuffer.buffer == ".exit" {
		os.Exit(sucExist)
	} else {
		return metaCommandUnrecoginzedCommand
	}
	return metaCommandSuccess
}

func readInput(inputBuffer *InputBuffer) {
        reader := bufio.NewReader(os.Stdin)
        text, err := reader.ReadString('\n')
        inputBuffer.inputLen = len(inputBuffer.buffer)
        if err != nil {
                fmt.Println(err)
                os.Exit(errExist)
        }
        inputBuffer.buffer = strings.TrimSpace(text)
}