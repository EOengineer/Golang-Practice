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
    fmt.Printf("Original address %p\n", family[0])
    fmt.Printf("Original address %p\n", family[1])
    fmt.Printf("Original address %p\n", family[2])

    // sorting not yet implemented.
    // currently serves to prove in place function handoff
    sortedFamily := sortByRank(family)

    fmt.Printf("New address %p\n", sortedFamily[0])
    fmt.Printf("New address %p\n", sortedFamily[1])
    fmt.Printf("New address %p\n", sortedFamily[2])
  }

  func sortByRank(arr []*Person) ([]*Person){
    for _, person := range arr {
      // these memory addresses will match those above, no duplicated objects created.
      fmt.Printf("Name is %s and Rank is %d\n", person.Name, person.Rank)
      fmt.Printf("Inner Loop Address is %p\n", person)
    }
    test := []*Person{arr[2], arr[1], arr[0]}
    return test
  }
