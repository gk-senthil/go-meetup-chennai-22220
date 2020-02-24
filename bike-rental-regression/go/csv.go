package main
import "fmt"
import "encoding/csv"
import "os"
import "strconv"

func main(){
	// Open the CSV.
	f, err := os.Open("day.csv")
	// Read in the CSV records.
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil{
		fmt.Println(err)
	}
	var i = 0
	var month_counts[13] int
	for _, record := range records {
		i++
		if (i<=5){
			fmt.Println(record)
		}
		month, err := strconv.Atoi (record[4])
		if (err != nil){

		}
		month_counts[month] += 1
	}
	for i := 1; i < 13; i++ {
		fmt.Println("Month: "+strconv.Itoa(i)+" "+strconv.Itoa(month_counts[i]))
	}
}
