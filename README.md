# Go Huma Rest API
Practicing with HUMA rest API framework with the Golang.
That is a simple API that demonstrate how to use the Huma Rest API for documentation.

## Run
```shell
$ go run .
```

## Calling the API
```shell
$ curl http://localhost:8888/greeting/world

HTTP/1.1 200 OK                                                         
Content-Length: 86                                                      
Content-Type: application/cbor                                          
Date: Sun, 05 Nov 2023 16:57:00 GMT                                     
Link: </schemas/GreetingOutputBody.json>; rel="describedBy"             
                                                                        
{                                                                       
  $schema: "http://localhost:8888/schemas/GreetingOutputBody.json"      
  message: "Hello, world!"                                              
}
```

## API Documentation
Go to http://localhost:8888/docs to see the interactive generated documentation for the API.

These docs are generated from the OpenAPI specification. You can use this file to generate documentation, client libraries, commandline clients, mock servers, and more. Two versions are provided by Huma. It is recommended to use OpenAPI 3.1, but OpenAPI 3.0.3 is also available for compatibility with older tools:
- OpenAPI 3.1 JSON: http://localhost:8888/openapi.json
- OpenAPI 3.1 YAML: http://localhost:8888/openapi.yaml
- OpenAPI 3.0.3 JSON: http://localhost:8888/openapi-3.0.json
- OpenAPI 3.0.3 YAML: http://localhost:8888/openapi-3.0.yaml

## References

['Your first API' tutorial](https://huma.rocks/tutorial/your-first-api/)


