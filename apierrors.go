package apierrors

import "encoding/json"

// Error represents a JSON:API error
type Error struct {
	ID     string                 `json:"id,omitempty"`
	Links  *Links                 `json:"links,omitempty"`
	Status string                 `json:"status,omitempty"`
	Code   int                    `json:"code,omitempty"`
	Title  string                 `json:"title,omitempty"`
	Detail string                 `json:"detail,omitempty"`
	Source *ErrorSource           `json:"source,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
}

// Links represents a set of links related to an error
type Links struct {
	About string `json:"about,omitempty"`
}

// ErrorSource represents the source of an error
type ErrorSource struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

// Errors represents a list of JSON:API errors
type Errors struct {
	Errors []Error `json:"errors"`
}

// NewError creates a new JSON:API error
func NewError(status string, code int, title, detail string) *Error {
	return &Error{
		Status: status,
		Code:   code,
		Title:  title,
		Detail: detail,
		Meta:   make(map[string]interface{}),
	}
}

// ToJSON marshals the Errors object to JSON
func (e *Errors) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// AddError adds an error to the Errors list
func (e *Errors) AddError(err Error) {
	e.Errors = append(e.Errors, err)
}

// AddMeta adds a key-value pair to the error's meta field
func (err *Error) AddMeta(key string, value interface{}) {
	if err.Meta == nil {
		err.Meta = make(map[string]interface{})
	}
	err.Meta[key] = value
}
