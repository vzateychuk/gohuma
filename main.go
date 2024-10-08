package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type Link struct {
	Related string `json:"related"`
	Items   string `json:"items,omitempty" doc:"A human-readable explanations specific to this occurrence of the problem."`
}

type Error struct {
	Code    string `json:"code,omitempty" example:"400" doc:"HTTP status code"`
	Message string `json:"message,omitempty" doc:"Error message text"`
}

// MyError represents a custom error structure.
type MyError struct {
	Status  int    `json:"status,omitempty" example:"400" doc:"HTTP status code"`
	Message string `json:"message,omitempty" doc:"Error message text"`
	// Errors  []error `json:"errors"`
	Errors []Error `json:"errors,omitempty" doc:"Optional list of individual error details"`
	Links  []Link  `json:"links,omitempty" doc:"Optional list of individual links"`
}

func (e *MyError) Error() string {
	msg := e.Message
	for _, e2 := range e.Errors {
		msg = fmt.Sprintf("%s: %s; %s", e2.Code, e2.Message, msg)
	}
	return msg
}

func (e *MyError) GetStatus() int {
	return e.Status
}

var NewMyError = func(status int, msg string, errs ...error) huma.StatusError {
	var myLinks []Link
	var myErrs []Error
	for i, err := range errs {
		myLinks = append(myLinks, Link{
			Related: fmt.Sprintf("https://example.com/error/%d", i),
			Items:   err.Error(),
		})

		myErrs = append(myErrs, Error{
			Code:    fmt.Sprintf("ERR-%d", i),
			Message: err.Error(),
		})
	}

	return &MyError{
		Status:  status,
		Errors:  myErrs,
		Message: msg,
		Links:   myLinks,
	}
}

func main() {
	// Create a new router & API
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	huma.NewError = NewMyError

	// Register GET /greeting/{name} handler.
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput, error) {
		resp, err := getOutput(input.Name)
		if err != nil {
			// return nil, huma.NewError(403, fmt.Sprintf("Invalid input name: '%s'", input.Name), err)
			return nil, huma.Error401Unauthorized(fmt.Sprintf("Invalid input name: '%s'", input.Name), err)
		}
		return resp, nil
	})

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}

func getOutput(name string) (*GreetingOutput, error) {
	resp := &GreetingOutput{}
	resp.Body.Message = fmt.Sprintf("Hello, %s!", name)
	if strings.HasPrefix(name, "v") {
		return nil, errors.New("name cannot start with 'v'")
	}
	return resp, nil
}
