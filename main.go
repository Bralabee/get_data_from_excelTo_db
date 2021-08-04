package main


import (
"database/sql"
"fmt"
_"github.com/go-sql-driver/mysql"
"github.com/tealeg/xlsx"
)

type Person struct {
	Name       string
	Education  string
	University string
	Industry   string
	Workyear   string
	Position   string
	Salary     string
	Language   string
}

//func init () {
//
//}
func GetExcel() []Person {
	var per1 []Person
	file, err := xlsx.OpenFile("E:\\student_info.xlsx")  //open a file
	if err != nil {
		fmt.Println(err)
	}

	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			var temp1 Person
			var str []string
			for _, cell := range row.Cells {
				str = append(str, cell.String())
			}
			temp1.Name = str[0]
			temp1.Education = str[1]
			temp1.University = str[2]
			temp1.Industry = str[3]
			temp1.Workyear = str[4]
			temp1.Position = str[5]
			temp1.Salary = str[6]
			temp1.Language = str[7]
			if str[1] == "Undergraduate" && str[4] == "1-3 years" {
				per1 = append(per1, temp1)
			}
		}
	}
	for i, v := range per1 {
		fmt.Println(i, v)
	}
	return per1
}

var (
	db *sql.DB
)

//Establish a database connection
func init() {
	var err error
	db, err = sql.Open("mysql", "mysqluser:password@(127.0.0.1:3306)/goDB")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func main() {
	//var db *sql.DB This is the same as the global db. An error is reported! ! ==
	//"User name: password@[connection method](host name: port number)/database name"
	//db, _ := sql.Open("mysql", "mysqluser:password@(127.0.0.1:3306)/goDB")
	defer db.Close()
	//Connect to the database
	err := db.Ping()
	if err != nil {
		fmt.Println("Database connection failed")
		return
	}
	per := GetExcel()
	//db.Query("insert into goDB.person values(?,?,?,?,?,?,?,?)")

	stmt, _ := db.Prepare("insert into goDB.person values(?,?,?,?,?,?,?,?)") //Get prepared statement object
	for _, one := range per {
		fmt.Println(one.Name, one.Education, one.University, one.Industry, one.Workyear, one.Position, one.Salary, one.Language)
		stmt.Exec(one.Name, one.Education, one.University, one.Industry, one.Workyear, one.Position, one.Salary, one.Language) //Call prepared statement
	}

}