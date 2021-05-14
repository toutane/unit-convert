package length

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%gft", f)
}
