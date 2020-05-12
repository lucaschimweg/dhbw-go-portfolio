package fibonacci

func fibChan(n int64, res chan<- int64) {
	if n == 0 {
		res <- 0
		return
	}

	if n == 1 {
		res <- 1
		return
	}

	ch := make(chan int64)
	go fibChan(n-1, ch)
	go fibChan(n-2, ch)

	res <- (<-ch) + (<-ch)
}

func Fib(n int64) int64 {
	ch := make(chan int64)
	go fibChan(n, ch)
	return <- ch
}
