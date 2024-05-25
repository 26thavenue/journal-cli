package journal

import (
	_ "fmt"
	 "os"
	"strings"
	"time"
)

type entry struct {
	Date time.Time
	Task string
}

type Journal []entry

func (j *Journal) AddEntry(task string) {
	je := entry{
		Date: time.Now(),
		Task: task,
	}

	*j = append(*j, je)
}




func (j *Journal) Save(filename string) error {
	var builder strings.Builder

	for _, e := range *j{
		builder.WriteString( "-" + " " +e.Date.Format(time.RFC850) + " " + e.Task  + "\n"  )
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(builder.String()); err != nil {
		return err
	}

	return nil
}



func (j *Journal) DeleteAll(filename string) error {
	*j = Journal{}
	var builder strings.Builder

	for _, e := range *j{
		builder.WriteString( "-" + " " +e.Date.Format(time.RFC850) + " " + e.Task  + "\n"  )
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(builder.String()); err != nil {
		return err
	}

	return nil
}

// func(j *Journal) DeleteEntry(filename string, task string) error{
// 	var builder strings.Builder

// }