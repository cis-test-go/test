// Package config has the configuration of the MySql and MongoDB databases
package main

import (
	"os"
	"github.com/jmoiron/sqlx"
)

// User stores  the database user
var	User              string = SetConfigValue("MYSQLUSER", "root")
// Password stores the database user password
var	Password          string = SetConfigValue("MYSQLPSSWD", "root")
// Host stores the database host
var	Host              string = SetConfigValue("MYSQLHOST", "127.0.0.1:3306")
// Database stores the name of the database
var	Database          string = SetConfigValue("MYSQLDB", "mydb")
// Type stores the type of database
var	Type              string = SetConfigValue("TYPE", "mysql")
// ConnectionDetails stores all combination of the other variables
var	Connectiondetails string = "" + User + ":" + Password + "@tcp(" + Host + ")/" + Database + ""

//MongoDB configuration
// Mgohostname stores the mongodb host name
var Mgohostname      string = SetConfigValue("MGOHOSTNAME", "127.0.0.1:27017")
// Mgodatabasename stores the database name
var	Mgodatabasename  string = SetConfigValue("MGODATABASENAME", "mydbmongo")

// SetConfigValue set environment by value and string
func SetConfigValue(e string, d string) string {
	if v := os.Getenv(e); v != "" {
		return string(v)
	}
	return d
}

// ConnectDB creates a connection and ping to the specific database
func ConnectDB(Type string, ConnectionDetails string) (*sqlx.DB, error) {
	return sqlx.Connect(Type, ConnectionDetails)
}
