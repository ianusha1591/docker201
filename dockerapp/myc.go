package main
import (
    "fmt"
    "html/template"
    "log"
    "net/http"
"gopkg.in/mgo.v2"
"encoding/json"
)
type Data struct {
    Name string
}
func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("myc.tmpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
name:= r.FormValue("username")
data  := &Data{name}
 b, err :=json.Marshal(data)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
}
 fmt.Println(string(b))
session, err := mgo.Dial("mongodb://35.243.216.194:27017")
      if err != nil {
            panic(err)
  }
        defer session.Close()
         session.SetMode(mgo.Monotonic, true)
        c := session.DB("test6").C("people")
index := mgo.Index{
               Key:        []string{"username"},
                Unique:     true,
               DropDups:   true,
                Background: true,
 Sparse:     true,
       }
        err = c.EnsureIndex(index)
        if err != nil  && mgo.IsDup(err) {
panic(err)
        }
err = c.Insert(data)
        if err != nil && mgo.IsDup(err) {
panic(err)
}
}
}
func getdetails(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
session, err := mgo.Dial("mongodb://35.243.216.194:27017")
        if err != nil {
                fmt.Println("Failed to establish connection to Mongo server:", err)
        }
defer session.Close()
        
        c := session.DB("test6").C("people")
var games []Data
        if err := c.Find(nil).All(&games); err != nil {
                fmt.Println("Failed to write results:", err)
        }
        
fmt.Println("Results All: ", games)
userJson, err:= json.Marshal(games)
if err!=nil {
panic(err)
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
w.Write(userJson)
}
func main() {
 http.HandleFunc("/myc", login)
http.HandleFunc("/get", getdetails)
    err := http.ListenAndServe(":80", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
}
}
   
