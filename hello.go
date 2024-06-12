

package main

import(fmt)

func main() {
	fmt.Println("Hello World!")

    // Define a list of integers
    numbers := []int{1, 2, 3, 4, 5}

    // Initialize sum variable
    sum := 0

    // Loop through the list and add each number to the sum
    for _, num := range numbers {
        sum += num
    }

    // Print out the sum
    fmt.Println("The sum of the numbers is:", sum)
}
