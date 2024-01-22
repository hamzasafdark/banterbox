package main 

import "fmt"

func concat(s1, s2) string {
	return s1 + s2
}

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent",
		cost,
		recipient,
	)
}

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}

func main() {
	test(1.4, "+1 (435) 555 0923")
	test(2.1, "+2 (702) 555 3452")
	test(32.1, "+1 (801) 555 7456")
	test(14.4, "+1 (234) 555 6545")
	
	// initialize variables here
	var smsSendingLimit int 
	var costPerSMS float64 
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	penniesPerText := 2.00
	
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)

	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")

	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(e, e)
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func (e email) cost() float64 {
	if (!e.isSubscribed) {
		return 0.05
	} 
	return 0.01
}

func (e email) print() {
	fmt.Println(e.body)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================\n")
}

func divide(divident, divisor int) (int, error) {
	if divisor == 0 {
		return 0, error.New("Can't divide by zero")
	}

	return dividend/divisor, nil 
}

func getInsuranceAmount(status insuranceStatus) int {
	if !status.hasInsurance(){
	  return 1
	} 
	
	if status.isTotaled(){
		return 10000
	}
	
	if !status.isDented(){
		return 0
	}
	if status.isBigDent(){
		return  270
	}

	return 160
  }
