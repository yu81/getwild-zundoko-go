package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	counter := 0
	for {
		counter++
		getWild := getWildAndTough()
		fmt.Println(getWild)
		if isGetWildAndTough(getWild) {
			printWithDuration(hugeGet, 200*time.Millisecond)
			printWithDuration(hugeWild, 400*time.Millisecond)
			printWithDuration(hugeAnd, 300*time.Millisecond)
			printWithDuration(hugeTough, 1000*time.Millisecond)
			fmt.Printf("%d 回目のチャレンジでGet Wild And Tough\n", counter)
			os.Exit(0)
		}
		if isGetChanceAndLuck(getWild) {
			printWithDuration(hugeGet, 200*time.Millisecond)
			printWithDuration(hugeChance, 400*time.Millisecond)
			printWithDuration(hugeAnd, 300*time.Millisecond)
			printWithDuration(hugeLuck, 1000*time.Millisecond)
			fmt.Printf("%d 回目のチャレンジでGet Chance And Luck\n", counter)
			os.Exit(0)
		}
	}
}

func getRandomIndex() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(int64(len(words)))
}

func getWordFromIndex(w []string, n int) string {
	if n < 0 || n > len(w) {
		return ""
	}
	return w[n]
}

func getWildAndTough() string {
	const l = 4
	result := make([]string, l)
	for i := 0; i < l; i++ {
		result[i] += getWordFromIndex(words, int(getRandomIndex()))
	}
	return strings.Join(result, " ")
}

func printWithDuration(s string, d time.Duration) {
	fmt.Println(s)
	time.Sleep(d)
}

func isGetWildAndTough(s string) bool {
	return s == "Get Wild And Tough"
}

func isGetChanceAndLuck(s string) bool {
	return s == "Get Chance And Luck"
}

var words = []string{"Get", "Wild", "And", "Tough", "Chance", "Luck"}

// generated with figlet
// https://github.com/cmatsuoka/figlet
const (
	hugeGet = `  ____      _
 / ___| ___| |_
| |  _ / _ \ __|
| |_| |  __/ |_
 \____|\___|\__|
`

	hugeWild = `__        ___ _     _
\ \      / (_) | __| |
 \ \ /\ / /| | |/ _` + "`" + "|" + `
  \ V  V / | | | (_| |
   \_/\_/  |_|_|\__,_|
`

	hugeAnd = `    _              _
   / \   _ __   __| |
  / _ \ | '_ \ / _` + "`" + "|" + `
 / ___ \| | | | (_| |
/_/   \_\_| |_|\__,_|
`

	hugeTough = ` _                    _
| |_ ___  _   _  __ _| |__
| __/ _ \| | | |/ _` + "`" + ` | '_ \
| || (_) | |_| | (_| | | | |
\__\___/ \__,_|\__, |_| |_|
                |___/
`

	hugeChance = `  ____ _
 / ___| |__   __ _ _ __   ___ ___
| |   | '_ \ / _` + "`" + ` | '_ \ / __/ _ \
| |___| | | | (_| | | | | (_|  __/
\____|_| |_|\__,_|_| |_|\___\___|
`

	hugeLuck = ` _               _
| |   _   _  ___| | __
| |  | | | |/ __| |/ /
| |__| |_| | (__|   <
|_____\__,_|\___|_|\_\
`
)
