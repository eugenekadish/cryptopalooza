package main

import (
	"fmt"

	"github.com/eugenekadish/cryptopalooza/zksnark/qap"
)

// https://eprint.iacr.org/2012/215.pdf
// https://eprint.iacr.org/2013/507.pdf

func main() {

	fmt.Printf(" Example 1 QAP        %t \n", qap.E1QAP())
	fmt.Printf(" Example 1 Strong QAP %t \n", qap.E1SQAP())

	fmt.Printf(" Example 2 QAP        %t \n", qap.E2QAP())
	fmt.Printf(" Example 2 R1CS        %t \n", qap.E2R1CS())

	fmt.Printf(" Example 3 QAP        %t \n", qap.E2QAP())
	fmt.Printf(" Example 3 R1CS        %t \n", qap.E3R1CS())
}
