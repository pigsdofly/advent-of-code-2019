package advent

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func convert(val int) int {
	result := int(math.Floor(float64(val)/3) - 2)
	if result <= 0 {
		return 0
	} else {
		return result + convert(result)
	}
}

func main() {
	filename := "input"
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	s := string(dat)
	s_a := strings.Split(s, "\n")
	var sum int
	for i := 0; i < len(s_a); i++ {
		j, _ := strconv.Atoi(s_a[i])
		sum += convert(j)
	}
	fmt.Println(sum)
}
