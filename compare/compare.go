/*
compare.go 

This package analyses the SPIN verification speed for two models,
defined in PROMELA (hand translated from NuSMV models),
outputing to a .csv file for printing

LTL properties generated automatically by the Spot library.

*/

package main

import (
	"fmt"
	"os"
	"os/exec"

	//"github.com/user/stringutil"
)

func main() {
	
	fmt.Println("Hello Go!")

    	cmd := exec.Command("./src/github.com/jsshafer/spot-2.6.3/bin/randltl")
    	err := cmd.Run()
    	if err != nil {
    	    os.Exit(1)
    	}

	// For {single.pml, multi.pml}
	
		// For n = {1, 20, 40, 60, 80, 100}
		
			// For 1 .. numTrials

				// Generate n properties using `randltl`
				// run SPIN, record into results.csv

}
