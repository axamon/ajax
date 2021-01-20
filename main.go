package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var INDEX_HTML []byte

type A struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func init() {
	INDEX_HTML, _ = ioutil.ReadFile("./html/index.html")
}

var database = "db.json"

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/q", q)
	mux.HandleFunc("/index", IndexHandler)
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	mux.HandleFunc("/post", PostHandler)
	http.ListenAndServe(":8888", mux)
}

// localhost:8080/?id=1
func q(w http.ResponseWriter, r *http.Request) {

	uidEnd := strings.Split(uuid.New().String(), "-")[4]

	uid := time.Now().Format("20060102") + "-" + uidEnd

	fmt.Println(uid)

	rID, ok := r.URL.Query()["id"]

	if !ok || len(rID[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	var a []A

	data, err := ioutil.ReadFile(database)
	if err != nil {
		log.Println(err)
	}

	// fmt.Println(string(data))

	json.Unmarshal(data, &a)

	var b A

	b.ID = uid
	b.Value = strconv.Itoa(rand.Intn(100))

	a = append(a, b)

	bites, err := json.MarshalIndent(a, "", "")
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(database, bites, os.FileMode(666))
	if err != nil {
		log.Println(err)
	}

	for _, e := range a {
		if strings.HasPrefix(e.ID, rID[0]) {
			fmt.Fprintf(w, "%s,%s\n", e.ID, e.Value)
		}
	}

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /")
	w.Write(INDEX_HTML)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("in posthandler", r.Form)
	var value = r.FormValue("textfield")
	w.Write([]byte(value))
}
