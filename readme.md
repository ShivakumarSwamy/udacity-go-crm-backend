# Udacity CRM Backend
A Go solution based Customer Relationship Management(CRM) Backend service to handle the lifecycle of a customer.  

## Pre-requisites
- Install Go 1.21 via its [website](https://go.dev/learn/) or via [goenv](https://github.com/go-nv/goenv)

## Run Locally
Run the below command to run the application server
```shell
go run main.go
```

Note: While running if required packages are missing it will download automatically and install under $GOPATH

## Run Tests
Run the below command to run the test
```shell
go test
OR 
# verbose mode
go test -v
```

## API Usage
After the application server is up, you can use the below api's to interact with customers data

### Get All Customers
```
REQUEST
curl --location 'http://localhost:3000/customers'
```
```
RESPONSE 200 OK
[
    {
        "id": "78ee45a9-7626-4c83-9a97-b572c9c318eb",
        "name": "Bob",
        "role": "USER",
        "email": "bob@company.com",
        "phone": 273598645,
        "contacted": false
    },...
]
```

### Get Single Customer

#### When Customer Found
```
REQUEST
curl --location 'http://localhost:3000/customers/431baecf-6535-452f-884e-1da18ff0d5a2'
```
```
RESPONSE 200 OK
{
    "id": "431baecf-6535-452f-884e-1da18ff0d5a2",
    "name": "Foo",
    "role": "ADMIN",
    "email": "foo@company.com",
    "phone": 56792834,
    "contacted": false
}
```

#### When Customer Not Found
```
REQUEST
curl --location 'http://localhost:3000/customers/93i493'
```
```
RESPONSE 404 NOT FOUND
"Unable to find the customer"
```


### Add Customer
```
REQUEST
curl --location 'http://localhost:3000/customers' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Name",
    "role": "USER",
    "email": "Email",
    "phone": 674245823,
    "contacted": false
}'
```
```
RESPONSE 201 CREATED
{
    "id": "67624a0e-c9f0-45ab-ae60-bd2765b18b20",
    "name": "Name",
    "role": "USER",
    "email": "Email",
    "phone": 674245823,
    "contacted": false
}
```

### Update Customer

#### Existing Customer
```
REQUEST
curl --location --request PUT 'http://localhost:3000/customers/431baecf-6535-452f-884e-1da18ff0d5a2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "431baecf-6535-452f-884e-1da18ff0d5a2",
    "name": "Foo",
    "role": "ADMIN",
    "email": "foo@company.com",
    "phone": 7698656785,
    "contacted": true
}'
```
```
RESPONSE 200 OK
{
    "id": "431baecf-6535-452f-884e-1da18ff0d5a2",
    "name": "Foo",
    "role": "ADMIN",
    "email": "foo@company.com",
    "phone": 7698656785,
    "contacted": true
}
```

#### When Customer Not Found
```
REQUEST
curl --location --request PUT 'http://localhost:3000/customers/hbdgjd' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "431baecf-6535-452f-884e-1da18ff0d5a2",
    "name": "Foo",
    "role": "ADMIN",
    "email": "foo@company.com",
    "phone": 7698656785,
    "contacted": false
}'
```
```
RESPONSE 404 NOT FOUND
"Unable to find the customer"
```

### Delete Customer

#### Existing Customer
```
REQUEST
curl --location --request DELETE 'http://localhost:3000/customers/4eb4af25-b104-4aa2-a321-9ab671c4fa36'
```
```
RESPONSE 204 NO CONTENT
```

#### When Customer Not Found
```
REQUEST
curl --location --request DELETE 'http://localhost:3000/customers/ndgjhrngj'
```
```
RESPONSE 404 NOT FOUND
"Unable to find the customer"
```