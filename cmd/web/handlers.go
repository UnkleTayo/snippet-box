package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home handler function which write a byte slice
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err)
		return
	}

}

// showSnippet handler FUNCTION
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// convert query string to integar using strconv.Atoi
	// if it cant be converted or value is less than 1 return 4040 page
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		// Changing the response header map after a call to w.WriteHeader() or
		// w.Write() will have no effect on the headers that the user receives.
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed!"))

		// http.Error(w, "Method not Allowed", 405)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create snippet"))
}
