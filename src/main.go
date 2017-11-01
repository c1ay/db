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
    prepareSyntaxError
)

type executeResult int
const (
    _ executeResult = iota
    executeSuccess
    executeFail
)


type InputBuffer struct {
    buffer string
    bufferLen int
    inputLen int
}


type Row struct {
    id uint32
    username string
    email string
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
    rowToInsert Row
}

type Table struct {
    pages interface{}
    numRows uint32
}


func serializeRow(row *Row, destination interface{}) {
}


func prepareStatement(inputBuffer *InputBuffer, statement *Statement) prepareResultCode {
    if len(inputBuffer.buffer) < 6 {
        return prepareUnrecognizedStatement
    }
    if inputBuffer.buffer[:6] == "insert" {
        statement.sType = insertType
        argsAssigned, err := fmt.Sscanf(inputBuffer.buffer, "insert %d %s %s", &statement.rowToInsert.id,
             &statement.rowToInsert.username, &statement.rowToInsert.email)
        if err != nil {
            return prepareSyntaxError
        }
        if argsAssigned < 3 {
            return prepareSyntaxError
        }
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

func executeStatement(statement *Statement) executeResult {
    switch statement.sType {
    case selectType:
        // fmt.Println("do select")
        return executeSelect(statement)
    case deleteType:
        // fmt.Println("do delete")
        return executeDelete(statement)
    case insertType:
        // fmt.Println("do insert")
        return executeInsert(statement)
    }
    return executeFail
}

func executeSelect(statement *Statement) executeResult {
    return executeSuccess
}

func executeInsert(statement *Statement) executeResult {
    return executeSuccess
}

func executeDelete(statement *Statement) executeResult {
    return executeSuccess
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
                continue
            }
        }

        // prepare statement
        statement := new(Statement)
        switch prepareStatement(inputBuffer, statement) {
        case prepareSuccess:
        case prepareSyntaxError:
            fmt.Printf("Syntax error. Could not parse statement. \n")
            continue
        case prepareUnrecognizedStatement:
            fmt.Printf("Unrecoginzed keyword at start of '%s'. \n", inputBuffer.buffer)
            continue
        }

        // execute statement
        executeStatement(statement)
        fmt.Println("Executed.")
    }
}


func printPrompt() {fmt.Printf("db > ")}


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