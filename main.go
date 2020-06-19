package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Gorbitsa struct {
	memory       [256]uint8
	instructions [256]uint8
	regX         uint8
	regPc        uint32
}

type Instruction struct {
	op    byte
	param int
}

func makeInstruction(instr string) Instruction {
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
		instructions[i] = makeInstruction(s)
	}
	return instructions
}

func main() {
	const program = "R T"
	var instructions = compileProgram(program)

	fmt.Println(instructions)
}
