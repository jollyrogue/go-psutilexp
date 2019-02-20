package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	sampleInterval := time.Duration(1000 * time.Millisecond)

	fmt.Printf("Getting CPU usage percent...\n")
	for sample := 0; sample <= 10; sample++ {
		cpuPercentTotal, err := cpu.Percent(sampleInterval, false)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			break
		}

		cpuPercentPerCPU, err := cpu.Percent(sampleInterval, true)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			break
		}

		fmt.Printf("Length: %d %v\n", len(cpuPercentTotal), cpuPercentTotal)
		for index, value := range cpuPercentTotal {
			fmt.Printf("CPU %d Percent: %f\n", index, value)
		}

		fmt.Printf("Length: %d %v\n", len(cpuPercentPerCPU), cpuPercentPerCPU)
		for index, value := range cpuPercentPerCPU {
			fmt.Printf("CPU %d Percent: %f\n", index, value)
		}
		fmt.Println()
	}

	fmt.Println("Sleeping...\n")
	time.Sleep(0 * time.Second)

}
