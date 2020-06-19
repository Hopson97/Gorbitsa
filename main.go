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
	param uint8
}

func (gorbitsa *Gorbitsa) hasNext() bool {
	return gorbitsa.regPc < 255
}

func (gorbitsa *Gorbitsa) executeNext(instructions [256]Instruction) {
	var instruction = &instructions[gorbitsa.regPc]
	switch instruction.op {
	case 'G':
		gorbitsa.regX = gorbitsa.memory[instruction.param]
	case 'O':
		gorbitsa.memory[instruction.param] = gorbitsa.regX
	case 'R':
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input: ")
		gorbitsa.regX, _ = reader.ReadByte()
	case 'B':
		if gorbitsa.regX == 0 {
			gorbitsa.regPc = uint32(instruction.param)
		}
	case 'I':
		gorbitsa.regX += instruction.param
	case 'T':
		fmt.Print(gorbitsa.regX)
	case 'S':
		gorbitsa.regX = instruction.param
	case 'A':
		gorbitsa.regX += gorbitsa.memory[instruction.param]

	}
	gorbitsa.regPc++
}

func newInstruction(instr string) Instruction {
	var instruction = Instruction{}
	instruction.op = instr[0]
	i, err := strconv.Atoi(instr[1:len(instr)])
	if err == nil {
		instruction.param = uint8(i)
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
	const program = "R O0 O1 R I255 O2 G0 A1 O0 G2 I255 O2 B15 S0 B6 G0 T"
	var instructions = compileProgram(program)

	var gorbitsa = Gorbitsa{}
	for gorbitsa.hasNext() {
		gorbitsa.executeNext(instructions)
	}
}
