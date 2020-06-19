package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Gorbitsa struct {
	memory [256]uint8
	regX   uint8
	regPc  uint32
}

type Instruction struct {
	op    byte
	param int
}

func (gorbitsa *Gorbitsa) hasNext() bool {
	return gorbitsa.regPc < 256
}

func (gorbitsa *Gorbitsa) executeNext(instructions [256]Instruction) {
	switch instructions[gorbitsa.regPc].op {
	case 'G':
	case 'O':
	case 'R':
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input: ")
		gorbitsa.regX, _ = reader.ReadByte()
	case 'B':
	case 'I':
	case 'T':
		fmt.Println(gorbitsa.regX - '0')
	case 'S':
	case 'A':
	}
	gorbitsa.regPc++
}

func newInstruction(instr string) Instruction {
	var instruction = Instruction{}
	instruction.op = instr[0]
	i, err := strconv.Atoi(instr[1:len(instr)])
	if err != nil {
		instruction.param = i
	}
	return instruction
}

func compileProgram(program string) [256]Instruction {
	var instructions [256]Instruction
	for i, s := range strings.Split(program, " ") {
		instructions[i] = newInstruction(s)
	}
	return instructions
}

func main() {
	const program = "R T"
	var instructions = compileProgram(program)

	var gorbitsa = Gorbitsa{}
	for gorbitsa.hasNext() {
		gorbitsa.executeNext(instructions)
	}
}
