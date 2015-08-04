package csrf

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/zenazn/goji/web"
)

// Tests that options functions are applied to the middleware.
func TestOptions(t *testing.T) {
	var h http.Handler

	age := 86400
	domain := "goji.io"
	path := "/forms/"
	header := "X-AUTH-TOKEN"
	field := "authenticity_token"
	errorHandler := unauthorizedHandler

	testOpts := []func(*csrf) error{
		MaxAge(age),
		Domain(domain),
		Path(path),
		HttpOnly(false),
		Secure(false),
		RequestHeader(header),
		FieldName(field),
		ErrorHandler(web.HandlerFunc(errorHandler)),
	}

	// Parse our test options and check that they set the related struct fields.
	cs := parseOptions(h, testOpts...)

	if cs.opts.MaxAge != age {
		t.Errorf("MaxAge not set correctly: got %v want %v", cs.opts.MaxAge, age)
	}

	if cs.opts.Domain != domain {
		t.Errorf("Domain not set correctly: got %v want %v", cs.opts.Domain, domain)
	}

	if cs.opts.Path != path {
		t.Errorf("Path not set correctly: got %v want %v", cs.opts.Path, path)
	}

	if cs.opts.HttpOnly != false {
		t.Errorf("HttpOnly not set correctly: got %v want %v", cs.opts.HttpOnly, false)
	}

	if cs.opts.Secure != false {
		t.Errorf("Secure not set correctly: got %v want %v", cs.opts.Secure, false)
	}

	if cs.opts.RequestHeader != header {
		t.Errorf("RequestHeader not set correctly: got %v want %v", cs.opts.RequestHeader, header)
	}

	if cs.opts.FieldName != field {
		t.Errorf("FieldName not set correctly: got %v want %v", cs.opts.FieldName, field)
	}

	if !reflect.ValueOf(cs.opts.ErrorHandler).IsValid() {
		t.Errorf("ErrorHandler not set correctly: got %v want %v",
			reflect.ValueOf(cs.opts.ErrorHandler).IsValid(), reflect.ValueOf(errorHandler).IsValid())
	}
}