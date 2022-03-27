
package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/pkg/handlers"
	"text/template"
	"bufio"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Book struct{
	ID     []int    `json:"id"`
	Title  []string `json:"title"`
	Author []string `json:"author"`
	Desc   []string `json:"desc"`
}

//func getStrings()[]string{
//	
//	var st Book
//	path := "github.com/pkg/mocks/book.go"
//	byteValue, err := ioutil.ReadFile(path)
//	if err != nil {
//		panic(err)
//	  }
//	err = json.Unmarshal(byteValue, &st)
//	if err != nil {
//		panic(err)
//	  }
//	return st.Title
//}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	title := getStrings("github.com/pkg/mocks/book.go")
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Book{
		Title: title,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}


func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("bill.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}



func main(){
    
	


	router:=mux.NewRouter()

    router.HandleFunc("/books", viewHandler)
	//-----------------можно удалить
	router.HandleFunc("/books/add", newHandler)
    //-----------------можно удалить


	router.HandleFunc("/books/1",handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}",handlers.GetBook).Methods(http.MethodGet)
	// "/books"
	router.HandleFunc("/books/look/add",handlers.AddBooks).Methods(http.MethodGet)
	// "/books/{id}"
	router.HandleFunc("/books/u/{id}",handlers.UpdateBook).Methods(http.MethodGet)
	// "/books/{id}"
	router.HandleFunc("/books/dated/{id}",handlers.DeleteBook).Methods(http.MethodGet)
	

	log.Println("API is running!")
	http.ListenAndServe(":4000",router)
}