package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}

	pi()

}

var n int64 = 10000000000
var h float64 = 1.0 / float64(n)

func f(a float64) float64 {
	return 4.0 / (1.0 + a*a)
}

func chunk(start, end int64, c chan float64) {
	var sum float64 = 0.0
	for i := start; i < end; i++ {
		x := h * (float64(i) + 0.5)
		sum += f(x)
	}
	c <- sum * h
}
func pi() {

	//记录开始时间
	start := time.Now()

	var pi float64
	np := runtime.NumCPU()
	runtime.GOMAXPROCS(np)
	c := make(chan float64, np)
	fmt.Println("np: ", np)

	for i := 0; i < np; i++ {
		//利用多处理器，并发处理
		go chunk(int64(i)*n/int64(np), (int64(i)+1)*n/int64(np), c)
	}

	for i := 0; i < np; i++ {
		tmp := <-c
		fmt.Println("c->: ", tmp)

		pi += tmp
		fmt.Println("pai: ", pi)

	}

	fmt.Println("Pi: ", pi)

	//记录结束时间
	end := time.Now()

	//输出执行时间，单位为毫秒。
	fmt.Printf("spend time: %vs\n", end.Sub(start).Seconds())

}

// 如何 profile： https://juejin.cn/post/7090137193292759077
// pprof -http=":8080" 'execing-processes' 'cpu.pprof'
