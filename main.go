package main

import (
	"net/http"

	"github.com/scrx/webservice/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)


	// here is a comment that I have added in very important one!
}

