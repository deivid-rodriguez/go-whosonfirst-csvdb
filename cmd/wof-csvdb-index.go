package main

import (
	"bufio"
	"flag"
	"fmt"
	csvdb "github.com/whosonfirst/go-whosonfirst-csvdb"
	"os"
	"strings"
	"time"
)

func main() {

	var cols = flag.String("columns", "", "Comma-separated list of columns to index")

	flag.Parse()
	args := flag.Args()

	to_index := make([]string, 0)

	for _, c := range strings.Split(*cols, ",") {
		to_index = append(to_index, c)
	}

	t1 := time.Now()

	db, err := csvdb.NewCSVDB()

	if err != nil {
	   panic(err)
	}

	for _, path := range args {

		err := db.IndexCSVFile(path, to_index)

		if err != nil {
			panic(err)
		}
	}

	t2 := time.Since(t1)

	fmt.Printf("> indexes: %d keys: %d rows: %d time to index: %v\n", db.Indexes(), db.Rows(), db.Keys(), t2)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("> query <col>=<id>")
	fmt.Printf("> ")

	for scanner.Scan() {

		input := scanner.Text()
		query := strings.Split(input, "=")

		if len(query) != 2 {
			fmt.Println("invalid query")
			continue
		}

		k := query[0]
		v := query[1]

		fmt.Printf("search for %s=%s\n", k, v)

		t1 := time.Now()

		rows, _ := db.Where(k, v)

		t2 := time.Since(t1)

		fmt.Printf("where %s=%s %d results (%v)\n", k, v, len(rows), t2)

		fmt.Println("")
		fmt.Println("> query <col>=<id>")
		fmt.Printf("> ")
	}
}
