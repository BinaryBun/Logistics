package main

import (
  "log"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

func start_db() (*sql.DB) {
  login, passwd, name :=  "root", "binarybun", "pis"
  db, err := sql.Open("mysql", login+":"+passwd+"@tcp(127.0.0.1:3306)/"+name)

  if err != nil {
    log.Println("Error: ", err.Error())
  } else {
    return db
  }
  return nil
}

func get_logs(id string) ([][4]string) {
  end_data := [][4]string {}
  data := [4]string {}

  db := start_db()
  defer db.Close()

  //log.Println("select package_id, place, `data` from `logs` WHERE id <=> " + id)
  query := "select track, place, weight, `data` from `logs` "
  query += "inner join package on package_id <=> package.id WHERE track <=> " + id
  incert, err_inc := db.Query(query)
  defer incert.Close()

  if err_inc != nil {
    log.Println("Error incert: ", err_inc.Error())
  } else {
    for incert.Next() {
      nil := incert.Scan(&data[0], &data[1], &data[2], &data[3])
      if nil == nil {}

      data[0] = "Track: " + data[0]
      data[1] = "Place: " + data[1]
      data[2] = "Weight: " + data[2]
      data[3] = "Date: " + data[3]
      end_data = append(end_data, data)
    }
  }
  if len(end_data) == 0 {
    end_data = append(end_data, [4]string {"", "", "", ""})
  }
  return end_data
}
