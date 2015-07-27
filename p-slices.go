package main

  import (
  "fmt"
  )

  type Person struct {
    Name string
    Rank int
  }


  func main () {

    // defining and initializing
    eric := &Person{
      Name: "Eric",
      Rank: 1,
    }

    candice := &Person{
      Name: "Candice",
      Rank: 2,
    }

    finley := &Person{
      Name: "Finley",
      Rank: 3,
    }

    // collect values
    family := []*Person{finley, candice, eric}

    // print original momory addresses
    whatAmI("Orginal", family)

    // re-order slice based on Rank
    sortedFamily := sortByRank(family)

    // print memory addresses after sort
    whatAmI("New address", sortedFamily)
  }



  func whatAmI(str string, arr []*Person) {
    for _, person := range arr {
      fmt.Printf("Name is %s and Rank is %d\n", person.Name, person.Rank)
      fmt.Printf("%s Address is %p\n", str, person)
    }
  }

  func sortByRank(arr []*Person) ([]*Person){
    // sort coming soon
    temp := arr[0]
    arr[0] = arr[2]
    arr[2] = temp
    return arr
  }

