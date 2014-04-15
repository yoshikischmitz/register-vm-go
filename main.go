package main

import (
	"fmt"
)

var (
	// Program Registers
	regs = []int{0, 0, 0, 0}
	// The Actual Program
	prog = []int{0x1064, 0x11C8, 0x2201, 0x0000}

	// Program counter
	pc = -1

	// Instruction Fields
	instrNum = 0
	reg1     = 0
	reg2     = 0
	reg3     = 0
	imm      = 0

	running = true
)

func fetch() int {
	pc++
	return prog[pc]
}

func decode(instr int) {
	instrNum = (instr & 0xF000) >> 12
	reg1 = (instr & 0xF00) >> 8
	reg2 = (instr & 0xF0) >> 4
	reg3 = (instr & 0xF)
	imm = (instr & 0xFF)
}

func eval() {
	switch instrNum {
	case 0:
		// Halt
		fmt.Printf("halt\n")
		running = false
	case 1:
		// Loadi
		fmt.Printf("Loading r%d, #%d\n", reg1, imm)
		regs[reg1] = imm
	case 2:
		// Add
		fmt.Printf("add r%d r%d r%d\n", reg1, reg2, reg3)
		regs[reg1] = regs[reg2] + regs[reg3]
	}
}

func showRegs() {
	fmt.Printf("Regs = ")
	for i := 0; i < len(regs); i++ {
		fmt.Printf("%04X  ", regs[i])
	}
	fmt.Printf("\n")
}

func main() {
	for running {
		showRegs()
		instr := fetch()
		decode(instr)
		eval()
	}
}
