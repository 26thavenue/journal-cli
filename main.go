package main

import (
	"flag"
	"fmt"
	"log"
	_ "os"

	"github.com/26thavenue/journal-cli/journal"
)

func main() {
	
	add := flag.String("add", "", "Add a new journal entry")

	delAll := flag.Bool("delAll", false, "Delete all journal entries")


	var filename = "journal.md"
	flag.Parse()

	j := &journal.Journal{}




	switch{
		case *delAll:
			err := j.DeleteAll(filename)

			

			if err != nil{
				log.Fatalf("could not save the journal: %v", err)
			}

			fmt.Println("All entries succesfully deleted")


		case *add != "" :
			j.AddEntry(*add)

			err:= j.Save(filename)

			if err != nil{
				log.Fatalf("could not save the journal: %v", err)
			}

			fmt.Println("Entry succesfully saved")

		default:
			fmt.Println("Invalid action")

	}
}