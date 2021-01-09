package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	in := `first_name,last_name,username
"Rob,","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))
	s, _ := r.ReadAll()
	var x string
	var count int16
	var timerFlag int
	flag.IntVar(&timerFlag, "timer", 5, "set the duration of quiz")
	flag.Parse()
	timer := time.AfterFunc(time.Duration(timerFlag)*time.Second, func() {
		fmt.Printf("\ntime is up. you got %d right in %d seconds", count, timerFlag)
		os.Exit(0)
	})
	defer timer.Stop()
	for i := range s {
		fmt.Print(s[i][0])
		fmt.Scan(&x)
		if s[i][1] == x {
			count++
		}
	}
	fmt.Println(count)
}
