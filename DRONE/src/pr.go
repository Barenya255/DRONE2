package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"repo/tools"
	"repo/worker"
)

func main() {
	fmt.Println("start")
	//fmt.Printf("%v-----\n", os.Args[0])
	//fmt.Printf("%v-----\n", os.Args[1]) //workerID
	//fmt.Printf("%v-----\n", os.Args[2]) //partitionNum
	//fmt.Printf("%v-----\n", os.Args[3]) //graph
	//fmt.Printf("%v-----\n", os.Args[4]) //partitionstrategy
	for i := 0; i < len(os.Args); i++ {
		log.Printf("args[%d]: %v\n", i, os.Args[i])
	}
	workerID, err := strconv.Atoi(os.Args[1])
	PartitionNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("conv fail!")
	}
	tools.SetDataPath(os.Args[3])

	is_rep := false
	if len(os.Args) == 5 && os.Args[4] == "rep" {
		is_rep = true
	}
	worker.RunPRWorker(workerID, PartitionNum, is_rep)
	fmt.Println("stop")
}
