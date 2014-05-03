package main

import "fmt"
// import "time"
// import "sync"

func multiply( name string, in chan uint64, out chan uint64, start bool) {
	if start {
		out <- 1
	}
	for {
		// select {
		// case num := <- in:
			num := <- in
			fmt.Println("channel ", name, " received ", num)
			sum := num * 2
			fmt.Println("channel ", name, " sending ", sum)
			out <- sum
			if sum == 0 {
				return
			}
		// default:
		// 	fmt.Println( "no activity ", name )
		// 	return
		// }
		// if sum > 1000 {
		// 	close( in )
		// 	fmt.Println( "Channel closed for ", name )
		// 	return
		// }
	}
}

func sharedChannelMultiply( name string, channel chan uint64, start bool) {

	if start {
		channel <- 1
	}
	for {
		// select {
		// case num := <- in:
		num := <- channel
		fmt.Println("channel ", name, " received ", num)
		sum := num + 2
		fmt.Println("channel ", name, " sending ", sum)
		channel <- sum
		if sum == 0 {
			return
		}
	}

}

func main() {
	channel1 := make( chan uint64 )
	// channel2 := make( chan uint64 )


	// go multiply( "one", channel1, channel2, true)
	// go multiply( "two", channel2, channel1, false)

	go sharedChannelMultiply( "one", channel1, true)
	go sharedChannelMultiply( "two", channel1, false)

    var input string
    fmt.Scanln(&input)
    fmt.Println("done")

}