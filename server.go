package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	ascii_art "ascii_art/packages" // Import the custom package for ASCII art transformation
)

// asciiHandler handles the ASCII art transformation based on user input.
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Process the submitted form

	// HTML template for displaying the form
	tmpl := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Form Example</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f5f5f5; margin: 20px;">
			<h1 style="color: #333366; text-align: center;">Hello, what ASCII would you like?</h1>
			<form action="/" method="post" style="width: 300px; margin: auto;">
				<textarea name="text_for_ascii" placeholder="Enter Text to get the ASCII representation" rows="4" cols="80" style="width: 100%; padding: 10px; margin: 5px 0; box-sizing: border-box;"></textarea><br>
				<div style="margin: 5px 0;">
					<input type="radio" id="standard" name="type" value="standard.txt" checked="checked" style="margin-right: 5px;">
					<label for="standard" style="font-size: 14px;">Standard</label>
				</div>
				<div style="margin: 5px 0;">
					<input type="radio" id="shadow" name="type" value="shadow.txt" style="margin-right: 5px;">
					<label for="shadow" style="font-size: 14px;">Shadow</label>
				</div>
				<div style="margin: 5px 0;">
					<input type="radio" id="thinkertoy" name="type" value="thinkertoy.txt" style="margin-right: 5px;">
					<label for="thinkertoy" style="font-size: 14px;">Thinkertoy</label>
				</div>
				<button type="submit" style="background-color: #4CAF50; color: white; padding: 14px 20px; border: none; cursor: pointer;">Submit</button>
			</form>
		</body>
	</html>
	`

	// Parse and execute the HTML template
	t := template.Must(template.New("tmpl").Parse(tmpl))
	t.Execute(w, "")

	r.ParseForm() // Parsing again; consider removing this redundant step
	// Retrieve form values
	text_for_ascii := r.FormValue("text_for_ascii")
	type_of_ascii := r.FormValue("type")

	if r.Method == "POST" {

		// Transform text to ASCII art using custom package
		lines_of_ascii := ascii_art.AsciiArtTransform(text_for_ascii, type_of_ascii)

		// Prepare the ASCII art for display
		store := "<br>"
		for i := 0; i < 8*len(lines_of_ascii); i++ {
			store += lines_of_ascii[i]
		}
		store = `<pre style="color: #333366; text-align: left;">` + store + "</pre>"

		// Parse and execute the ASCII art template
		t = template.Must(template.New("tmpl").Parse(store))
		t.Execute(w, "")

		return
	}

	// fmt.Fprintf(w, "POST request successful\n") // Display a message for successful POST request
}

// main initializes the HTTP server.
func main() {
	http.HandleFunc("/", asciiHandler) // Register the asciiHandler function for "/" path

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err) // Log fatal errors and exit
	}
}
