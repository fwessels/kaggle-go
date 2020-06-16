package main

import (
	"fmt"
	"log"
	"os/exec"
)

// $ kaggle datasets list --help
// usage: kaggle datasets list [-h] [--sort-by SORT_BY] [--size SIZE]
//                             [--file-type FILE_TYPE] [--license LICENSE_NAME]
//                             [--tags TAG_IDS] [-s SEARCH] [-m] [--user USER]
//                             [-p PAGE] [-v] [--max-size MAX_SIZE]
//                             [--min-size MIN_SIZE]
//
// optional arguments:
//   -h, --help            show this help message and exit
//   --sort-by SORT_BY     Sort list results. Default is 'hottest'. Valid options are 'hottest', 'votes', 'updated', and 'active'
//   --size SIZE           DEPRECATED. Please use --max-size and --min-size to filter dataset sizes.
//   --file-type FILE_TYPE
//                         Search for datasets with a specific file type. Default is 'all'. Valid options are 'all', 'csv', 'sqlite', 'json', and 'bigQuery'. Please note that bigQuery datasets cannot be downloaded
//   --license LICENSE_NAME
//                         Search for datasets with a specific license. Default is 'all'. Valid options are 'all', 'cc', 'gpl', 'odb', and 'other'
//   --tags TAG_IDS        Search for datasets that have specific tags. Tag list should be comma separated
//   -s SEARCH, --search SEARCH
//                         Term(s) to search for
//   -m, --mine            Display only my items
//   --user USER           Find public datasets owned by a specific user or organization
//   -p PAGE, --page PAGE  Page number for results paging. Page size is 20 by default
//   -v, --csv             Print results in CSV format (if not set print in table format)
//   --max-size MAX_SIZE   Specify the maximum size of the dataset to return (bytes)
//   --min-size MIN_SIZE   Specify the minimum size of the dataset to return (bytes)
//
func list() (res []string) {

	// kaggle datasets list --csv --file-type csv --min-size 1024000000 --max-size 1096000000 -p 1

	args := make([]string, 0, 16)
	args = append(args, "datasets")
	args = append(args, "list")
	if true {
		args = append(args, "--csv")
	}
	args = append(args, "--sort-by")
	args = append(args, "votes")
	args = append(args, "--file-type")
	args = append(args, "csv")
	args = append(args, "--min-size")
	args = append(args, "102400000")
	args = append(args, "--max-size")
	args = append(args, "1024000000")
	args = append(args, "-p")

	for page := 1; page <= 3; page++ {

		args = append(args, fmt.Sprintf("%d", page))

		cmd := exec.Command("kaggle", args...)
		cmb, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Println(string(cmb))

		args = args[:len(args)-1]
	}

	return
}

func main() {
	list()
}