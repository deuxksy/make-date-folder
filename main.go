package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Printf("%d, %s\n", len(os.Args), os.Args)
	if len(os.Args) < 3 {
		panic("생성할 폴더를 위치 와 실행할 횟수를 입력해 주세요 ex) make-date-folder d:\\ 365")
	}
	path := os.Args[1]
	count, _ := strconv.Atoi(os.Args[2])
	if strings.TrimSpace(path) != "" {
		now := time.Now()
		for i := 1; i <= count; i++ {
			now := now.AddDate(0, 0, i)
			year := now.Format("2006")
			month := now.Format("01")
			date := now.Format("02")
			mf(path, year, month, date)
		}
	}
}

func mf(path string, year string, month string, date string) {
	err := os.MkdirAll(fmt.Sprintf(
		"%s\\%s\\%s\\%s",
		path,
		year,
		month,
		date,
	), 0755)
	if err != nil {
		return
	}
}
