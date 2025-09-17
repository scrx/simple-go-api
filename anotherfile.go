package main
// example another.go file
import (
	"net/http"

	"github.com/scrx/webservice/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)


	
}

