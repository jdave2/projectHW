

package main

import "fmt"

func main(){
    var rad float64
    const PI float64 = 3.14 // Constant
    var areaOfCircle float64
    var ci float64
    fmt.Print("Enter radius of Circle : ")
    fmt.Scanln(&rad)
    areaOfCircle = PI * rad * rad 
    fmt.Println("Area of Circle is : ",areaOfCircle)
    circ = 2 * PI * rad
    fmt.Println("Circumference of Circle is : ",circ)     
}

