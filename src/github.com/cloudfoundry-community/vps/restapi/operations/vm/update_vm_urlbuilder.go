package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
)

// UpdateVMURL generates an URL for the update Vm operation
type UpdateVMURL struct {
}

// Build a url path and query string
func (o *UpdateVMURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/vms"

	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UpdateVMURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UpdateVMURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UpdateVMURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UpdateVMURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UpdateVMURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *UpdateVMURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}