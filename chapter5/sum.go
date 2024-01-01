package main

// 参数数量可变的函数称为可变参数函数。
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	temp := vals[0]
	for _, val := range vals {
		if val > temp {
			temp = val
		}
	}
	return temp
}

//fmt.Println(sum())           // "0"
//fmt.Println(sum(3))          // "3"
//fmt.Println(sum(1, 2, 3, 4)) // "10"
