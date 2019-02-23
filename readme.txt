Simple GO Lang REST API
Simple RESTful API to create, read and update users. 

Quick Start
# Install mux router
go get -u github.com/gorilla/mux
go build
./go_restapi
Endpoints
Get All Users
GET api/users
Get Single User
GET api/users/{id}
Create User
POST api/users
