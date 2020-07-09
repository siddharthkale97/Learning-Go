package main

import (
	"fmt"
)

type car struct {
	gas_pedal      uint16 //min 0 max 65535
	brake_pedal    uint16
	steering_wheel int16 // -32k to +32 k
	top_speed_kmh  float64
}

func main() {
	a_car := car{gas_pedal: 22341, // can use semicolon as we are inside a function
		brake_pedal:    0,
		steering_wheel: 12561,
		top_speed_kmh:  225.0}
	// b_car := car{21341, 0, 12561, 225.0} // shorten format
	fmt.Println(a_car.gas_pedal)
}
