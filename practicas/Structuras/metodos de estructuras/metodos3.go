package main
import(
	"fmt"
	"math"
)
type Myfloat float64
func (f Myfloat) abs()float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}
func main(){
	f:= Myfloat(-math.Sqrt2)
	fmt.Println(f.abs())
}