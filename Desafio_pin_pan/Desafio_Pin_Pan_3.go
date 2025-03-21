package main

func main() {

	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			println("Pin Pan", i)

		} else if (i%3 == 0) && (!(i%5 == 0)) {
			println("Pin", i)

		} else if (!(i%3 == 0)) && (i%5 == 0) {
			println("Pan", i)
		} else {
			println(i)
		}

	}

}
