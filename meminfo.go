package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type meminfo struct {
	name       string
	byteValue  int
	humanValue float32
	sizing     string
}

func isError(e error) {
	if e != nil {
		panic(e)
	}
}

func findLine(v string) string {
	file, err := os.Open("meminfo.txt")
	isError(err)

	line := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), v) {
			return scanner.Text()
		}
		line++
	}
	panic("error parsing /proc/meminfo for $v")
}

func humanReadable(i int) float32 {
	var f float32
	s := strconv.Itoa(i)
	switch len(s) {
	case 0:
		f = 0
	case 1:
		f = float32(i)
	case 2:
		f = float32(i)
	case 3:
		f = float32(i)
	case 4:
		f = float32(i) / 1000
	case 5:
		f = float32(i) / 1000
	case 6:
		f = float32(i) / 1000
	case 7:
		f = float32(i) / 1000000
	case 8:
		f = float32(i) / 1000000
	case 9:
		f = float32(i) / 1000000
	case 10:
		f = float32(i) / 1000000000
	case 11:
		f = float32(i) / 1000000000
	case 12:
		f = float32(i) / 1000000000
	default:
		panic("Memory out of bound")
	}
	return f
}

func sizing(i string) string {
	var s string
	switch len(i) {
	case 0:
		s = "KB"
	case 1:
		s = "KB"
	case 2:
		s = "KB"
	case 3:
		s = "KB"
	case 4:
		s = "MB"
	case 5:
		s = "MB"
	case 6:
		s = "MB"
	case 7:
		s = "GB"
	case 8:
		s = "GB"
	case 9:
		s = "GB"
	case 10:
		s = "TB"
	case 11:
		s = "TB"
	case 12:
		s = "TB"
	default:
		s = ""
	}
	return s
}

func ratio(total string, free string) int {
	switch {
	case total == "KB" && free == "KB":
		return 1
	case total == "MB" && free == "KB":
		return 1000
	case total == "GB" && free == "KB":
		return 1000000
	case total == "TB" && free == "KB":
		return 1000000000
	case total == "MB" && free == "MB":
		return 1
	case total == "GB" && free == "MB":
		return 1000
	case total == "TB" && free == "MB":
		return 1000000
	case total == "GB" && free == "GB":
		return 1
	case total == "TB" && free == "GB":
		return 1000
	case total == "TB" && free == "TB":
		return 1
	}
	return 0
}
func main() {
	// Total Memory
	t := strings.Fields(findLine("MemTotal"))
	bValue, _ := strconv.Atoi(t[1])
	memTotal := meminfo{"Total memory", bValue, humanReadable(bValue), sizing(t[1])}
	fmt.Println(memTotal)

	// Total free memory
	t = strings.Fields(findLine("MemFree"))
	bValue, _ = strconv.Atoi(t[1])
	memFree := meminfo{"Free Memory", bValue, humanReadable(bValue), sizing(t[1])}
	fmt.Println(memFree)

	// Total Used
	memUsed := meminfo{"Used memory", memTotal.byteValue - memFree.byteValue, humanReadable(memTotal.byteValue - memFree.byteValue), sizing(strconv.Itoa(memTotal.byteValue - memFree.byteValue))}
	fmt.Println(memUsed)
}
