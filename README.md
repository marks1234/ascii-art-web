# Ascii-Art-Web-stylize

## Description

Ascii-Art-Web is an intuitive, responsive, and user-friendly web application designed for crafting ASCII art from text input. Built as a progressive iteration of a prior ASCII art generator, this platform is served via a web server, offering users the option to select various banner styles like shadow, standard, and thinkertoy.

## Authors

- Marco Kuzmina (https://01.kood.tech/git/mkuzmina)

## Usage

### How to Run

1. Clone the repository to your local machine.

   ```bash
   git clone https://github.com/username/Ascii-Art-Web-stylize.git
   ```

2. Access the project directory.

   ```bash
   cd Ascii-Art-Web
   ```

3. Start the server.

   ```bash
   go run server.go
   ```

4. Open your preferred web browser and go to `http://localhost:8080/`.

### User Interface

- Main Page: Allows users to input text and select a banner style. Visually appealing design elements guide the user through the process.
- Feedback System: Real-time previews and error messages improve the user experience.

### Endpoints

- `GET /`: Delivers the primary UI for text input and banner style selection.
- `POST /ascii-art`: Processes the selected banner style and text, and returns the crafted ASCII art.

### HTTP Status Codes

- OK (200) - For successful operations.
- Not Found (404) - When templates or banners are unavailable.
- Bad Request (400) - For erroneous requests.
- Internal Server Error (500) - For unexpected issues.

## Implementation Details

### Algorithm

Written in Go, the application employs the language's standard packages for HTTP servers and HTML templates. ASCII art is generated based on user selections, received via POST requests, and displayed dynamically on the web page.

### Tech Stack

- Language: Go
- Packages: Built solely using Go's standard packages, adhering to best practices.

## Objectives

1. Make the website visually appealing, intuitive, and interactive.
2. Enhance user-friendliness.
3. Implement real-time feedback mechanisms for users.

### Allowed Packages

Only Go's standard packages are permitted for this project.

### Learning Outcomes

- Understand human-computer interface basics.
- Gain foundational knowledge in CSS.
- Learn how to integrate CSS and HTML.
