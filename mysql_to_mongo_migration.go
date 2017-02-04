package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
  _ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
    "strconv"
	"time"
)

// MessagesMongo stores messages data in mongo
type ReportMongo struct {
	Id               bson.ObjectId        `bson:"_id,omitempty" json:"_id"`
	ReportId 		int                   `bson:"reportId" json:"reportId"` 
	Status   		string      		  `bson:"status" json:"status"` 
	DateOfReport    time.Time             `bson:"DateOfReport" json:"DateOfReport"` 
}

// get mysql messages field.
type ReportMysql struct {
	ReportId	int 	`db:"report_id"`
	Status		int	    `db:"status"`
	DateOfReport string `db:"offence_date_ts"` 
}


// Session initialize mongoDB Session
var Session *mgo.Session

func main() {

	var Merr error
	Session, Merr = mgo.Dial(Mgohostname)
	if Merr != nil {
		panic(Merr.Error())
	}
	defer Session.Close()
	// Optional. Switch the Session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)
	c := Session.DB(Mgodatabasename).C("Report")

	//Conect to database
	db, err := ConnectDB(Type,Connectiondetails)

	//Error
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer db.Close()

	reportMysql := []ReportMysql{}

	err = db.Select(&reportMysql,"SELECT report_id, status,offence_date_ts FROM report where offence_date_ts IS NOT NULL")
	
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var count = 0;

	for key := range reportMysql {

		report := ReportMongo{}

		//var num int;
		var rid int;
		var rstatus int;
		var rstatusMongo string;

	 	rid = reportMysql[key].ReportId;

		//for sender name
		rstatus = reportMysql[key].Status;
		
        //Date sent conversion
		var tempDate time.Time;
	    if reportMysql[key].DateOfReport != "0" {
		    i, err5 := strconv.ParseInt(reportMysql[key].DateOfReport, 10, 64)
		    if err5 != nil {
		        panic(err5)
		    }
		    tempDate = time.Unix(i, 0)
		}

		err = c.Find(bson.M{"reportId": rid}).One(&report)

		if err == nil {

			if rstatus == 2{
				rstatusMongo = "Archived"
			}
			if rstatus == 1{
				rstatusMongo = "Open"
			}

			count++
			fmt.Println("updated : ", count)
			
			// update records if already exist. 
	
			_,err := c.Upsert(bson.M{"reportId": rid}, bson.M{"$set": bson.M{
				"status"	:	rstatusMongo,
				"dateOfReport": tempDate,
				}})
			if err != nil {
				panic(err.Error())
			}
		}
		fmt.Println("Report id :", rid)
	}
	fmt.Println("Total Count:", count)
}

func ArrayStringToInt(r []string) []int{
	var t2 = []int{}
	for _, r := range r {
	    j, err := strconv.Atoi(r)
	    if err != nil {
	        panic(err)
	    }
	    t2 = append(t2, j)
	}
	return t2
}