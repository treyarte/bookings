package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//Form creates a custom form struct, embeds a url.values object
type Form struct {
	url.Values
	Errors errors
}

//valid return true if there are no errors otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		val := f.Get(field)
		if strings.TrimSpace(val) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field);

	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false;
	}

	return true;
}

//MinLength checks for string min length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field be at least %d characters long", length))
		return false
	}

	return true
}