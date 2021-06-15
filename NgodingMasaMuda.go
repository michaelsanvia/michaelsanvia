package main

import "fmt"

func main () {
	var (s1, s2, pi1, pi2 float64
		n,i int
	)
	s1=0
	s2=0
	i=1
	n=1
	pi1=1
	pi2=0
	for pi1-pi2>=0.00001 || pi2-pi1>=0.00001 {
		if n % 2 ==0 {
			s1=s1-(1/float64(i))
			s2=s2-(1/(float64(i)+2))
		}else {
			s1=s1+(1/float64(i))
			s2=s2+(1/(float64(i)+2))
		}
		i=i+2
		n=n+1
		pi1=s1*4
		pi2=s2*4
	}

	fmt.Println("Hasil PI: ", pi1)
	fmt.Println("Hasil PI: ", pi2)
	fmt.Println("Suku ke i", n)

}