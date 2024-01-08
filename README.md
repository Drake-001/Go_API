TITLE: EV Battery Data API Server

DESCRIPTION:
This Golang-based API server offers access to Electric Vehicle (EV) battery data, including battery health, charge cycles, temperature readings, and performance metrics. It's designed to serve as a backend server providing JSON data through HTTP requests.

DEPENDENCIES:
Go (Golang), A JSON file "battery_data.json" with EV battery data.

INSTALLING AND RUNNING ON THE SERVER:
Clone the repository or download the source code to your local machine.
Ensure Go is installed on your system. If not, download and install it from Go's official website.
Place the battery_data.json file in the same directory as your main Go script. The JSON file should contain the EV battery data in the appropriate format.

RUN THE SERVER:
go run main.go
This will start the server on http://localhost:8080.

API ENDPOINTS:
GET /batterydata: Returns a list of EV battery data in JSON format.
GET /: Returns a simple string message indicating that the EV Battery Data API Server is running.

USAGE:
Once the server is running, you can access the API endpoints using any HTTP client (web browser, curl, or Postman) Example:
curl http://localhost:8080/batterydata

OUTPUT:
The server returns JSON-formatted data representing various attributes of EV batteries, such as health percentage, charge cycles, temperature, and other metrics.

Author
Drake Davis
