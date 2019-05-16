package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	data := [][]string{
		[]string{"one", "15", "10/20", "(1,1,1)"},
		[]string{"two", "20", "20/26", "(1,1,1)"},
		[]string{"three", "30", "15/20", "(1,1,1)"},
		[]string{"four", "50", "50/80", "(1,1,1)"}}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"NPC", "Speed", "Power", "Location"})

	table.AppendBulk(data)

	table.Render()

}
