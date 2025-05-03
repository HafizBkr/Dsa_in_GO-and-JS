package main

import (
	"fmt"
)

func main(){

	//simple array
      var sArray [5]int
	  sArray[0]=1
	  sArray[1]=2
	  sArray[2]=3
	  sArray[3]=4
	  sArray[4]=5
	  fmt.Println(sArray)
	  //array with values
	  var sArray2 [5]int=[5]int{1,2,3,4,5}
	  fmt.Println(sArray2)

	//slice array
	var array=[...]int{4,4,43,3,2,1,2,4}
	var tab []string
	tab=[]string{"hello", "world"}

	
    //2 dimension arrays
	const v int =4
	const h int =4
	var array2 = [v][h]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	// parcourir les element du tableau 
	for i := 0; i < v; i++ {
		for j := 0; j < h; j++ {
			array2[i][j] = i + j
		}
	}
	fmt.Println(array2)

	//array with values
	fmt.Println(array, tab)

}