// Value Recievers
// Recives values and does calculations
// cannot modify a stuct value

package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934 // 1 mph  = 1.60934 kmph

type car struct {
	gas_pedal      uint16 //min 0 max 65535
	brake_pedal    uint16
	steering_wheel int16 // -32k to +32 k
	top_speed_kmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

// Value reciever makes a copy
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

// Just a function instead of a method
func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

func main() {
	a_car := car{gas_pedal: 22341, // can use semicolon as we are inside a function
		brake_pedal:    0,
		steering_wheel: 12561,
		top_speed_kmh:  225.0}
	// b_car := car{21341, 0, 12561, 225.0} // shorten format
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	a_car.new_top_speed(400)
	fmt.Println(a_car.kmh())
}
