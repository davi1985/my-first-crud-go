package rest_exceptions

import "net/http"

type RestException struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestException) Error() string {
	return r.Message
}

func NewRestException(message, err string, code int, causes []Cause) *RestException {
	return &RestException{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestException(message string) *RestException {
	return &RestException{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationException(message string, causes []Cause) *RestException {
	return &RestException{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestException {
	return &RestException{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundException(message string) *RestException {
	return &RestException{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenException(message string) *RestException {
	return &RestException{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
