package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"

	"nick/packages/dataset/environdec"
	"nick/packages/maps"
	"nick/packages/system"
	//Git Packages
)

func main(){

  categoriesMap, categoriesIntMap :=  maps.GetEnvirondecCategories()


  var selectedCategory int

  // Sort keys to print them always in the same order
  keys := make([]int, 0)
  for k := range categoriesIntMap {
      keys = append(keys, k)
  }
  sort.Ints(keys)
  fmt.Println("Hellow, please select a category:")


  for {
     
    for _, k := range keys {
      fmt.Println("\t", k+1, "-", categoriesIntMap[k])
    }
    fmt.Println("------------------------------------------------")

    // Get user input
    _, err := fmt.Scanln(&selectedCategory)
    if err != nil {
      log.Fatalln(err.Error())
    }
    if selectedCategory > 0 && selectedCategory <= len(keys){
      break
    }
    
    system.CallClear() 
    fmt.Println("------------------------------------------------")
    fmt.Println("Not in range, please select a valid category.")
    fmt.Println("------------------------------------------------")
  }

  fmt.Println("------------------------------------------------")

  fmt.Println("How many PDFs do you want to download?:")
  fmt.Println("1 - all")

  var limit int
  fmt.Scanln(&limit)

  if limit == 1 {

    system.CallClear()
    fmt.Println("Downloading", limit, "PDFs from", categoriesIntMap[selectedCategory-1], "category...")  
    // Create category folder if doesn't exist yet
    path := categoriesIntMap[selectedCategory-1]
  	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
  		err := os.Mkdir(path, os.ModePerm)
      if err != nil {
        log.Println(err)
      }
    }
    environdec.GetEnvirondecPDFs(categoriesMap[categoriesIntMap[selectedCategory-1]],"500", categoriesIntMap[selectedCategory-1])  
  }

  //geoScopesMap := maps.GetEnvirondecGeoScopes()
}

