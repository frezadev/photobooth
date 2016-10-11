package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	wd = func() string {
		d, _ := os.Getwd()
		return d + "/"
	}()
)

type DateInfo struct {
	DateId    time.Time
	MonthId   int
	MonthDesc string
	QtrId     int
	QtrDesc   string
	Year      int
}

func GetDateInfo(t time.Time) DateInfo {
	di := DateInfo{}

	year := t.Year()
	month := int(t.Month())

	monthid := strconv.Itoa(year) + LeftPad2Len(strconv.Itoa(month), "0", 2)
	monthdesc := t.Month().String() + " " + strconv.Itoa(year)

	qtr := 0
	if month%3 > 0 {
		qtr = int(math.Ceil(float64(month / 3)))
		qtr = qtr + 1
	} else {
		qtr = month / 3
	}

	qtrid := strconv.Itoa(year) + LeftPad2Len(strconv.Itoa(qtr), "0", 2)
	qtrdesc := "Q" + strconv.Itoa(qtr) + " " + strconv.Itoa(year)

	di.DateId, _ = time.Parse("2006-01-02 15:04:05", t.UTC().Format("2006-01-02")+" 00:00:00")
	di.Year = year
	di.MonthDesc = monthdesc
	di.MonthId, _ = strconv.Atoi(monthid)
	di.QtrDesc = qtrdesc
	di.QtrId, _ = strconv.Atoi(qtrid)

	return di
}

func MonthIDToDateInfo(mid int) (dateInfo DateInfo) {
	monthid := strconv.Itoa(mid)
	year := monthid[0:4]
	month := monthid[4:6]
	day := "01"

	iMonth, _ := strconv.Atoi(string(month))
	iMonth = iMonth - 1

	dtStr := year + "-" + month + "-" + day
	date, _ := time.Parse("2006-01-02", dtStr)

	dateInfo = GetDateInfo(date)

	return
}

func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func ErrorHandler(e error, position string) {
	if e != nil {
		fmt.Printf("ERROR on %v: %v \n", position, e.Error())
	}
}

func ErrorLog(e error, position string, errorList []error) []error {
	if e != nil {
		errorList = append(errorList, e)
		// fmt.Printf("ERROR on %v: %v \n", position, e.Error())
	}
	return errorList
}

/*func WriteErrors(errorList fmt.M, fileName string) (e error) {
	config := ReadConfig()
	source := config["datasource"]
	dataSourceFolder := "errors"
	fileName = fileName + "_" + fmt.GenerateRandomString("", 5) + ".txt"
	fmt.Printf("Saving Errors... %v\n", fileName)

	errors := ""

	for x, err := range errorList {
		errors = errors + "" + fmt.Sprintf("#%v: %#v \n", x, err)
	}

	e = ioutil.WriteFile(source+"\\"+dataSourceFolder+"\\"+fileName, []byte(errors), 0644)
	return
}*/

func ReadConfig() map[string]string {
	ret := make(map[string]string)
	file, err := os.Open(wd + "/conf/app.conf")
	if err == nil {
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, _, e := reader.ReadLine()
			if e != nil {
				break
			}

			sval := strings.Split(string(line), "=")
			ret[sval[0]] = sval[1]
		}
	} else {
		fmt.Println(err.Error())
	}

	return ret
}

func ReadJson(source string, result interface{}) {
	file, err := os.Open(wd + source)
	if err == nil {
		defer file.Close()

		jsonParser := json.NewDecoder(file)
		err = jsonParser.Decode(&result)

		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}
