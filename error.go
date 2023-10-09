package errorcode

import (
	"errors"
)

var (
	SUCCESS                   = errors.New("SUCCESS")
	DATABASE_TIMEOUT          = errors.New("DATABASE TIMEOUT")
	UNAUTHORIZED              = errors.New("UNAUTHORIZED")
	BAD_REQUEST               = errors.New("BAD REQUEST")
	NOT_FOUND                 = errors.New("NOT FOUND")
	NOT_ACCEPTABLE            = errors.New("NOT ACCEPTABLE")
	METHOD_NOT_ALLOWED        = errors.New("METHOD NOT ALLOWED")
	NOT_PERMITTED             = errors.New("NOT PERMITTED")
	NOT_PERMISSIONS           = errors.New("NOT_PERMISSIONS")
	UNKNOWN                   = errors.New("UNKNOWN")
	DATA_EXIST                = errors.New("DATA EXIST")
	MODEL_NULL                = errors.New("MODEL NULL")
	DATA_CORRUPT              = errors.New("DATA_CORRUPT")
	DATA_EMPTY                = errors.New("DATA_EMPTY")
	StatusInternalServerError = errors.New("StatusInternalServerError")
	StatusTooManyRequests     = errors.New("StatusTooManyRequests")
	StatusForbidden           = errors.New("StatusForbidden")
	StatusProcessing          = errors.New("StatusProcessing")
	StatusNoContent           = errors.New("StatusNoContent")
	StatusNotModified         = errors.New("StatusNotModified")
	StatusConflict            = errors.New("StatusConflict")
	StatusNotImplemented      = errors.New("StatusNotImplemented")
)

func GetCode(e error) int {
	switch e {
	case nil:
		return 200
	case SUCCESS:
		return 200
	case DATABASE_TIMEOUT:
		return 408
	case UNAUTHORIZED:
		return 401
	case BAD_REQUEST:
		return 400
	case NOT_FOUND:
		return 404
	case NOT_ACCEPTABLE:
		return 406
	case METHOD_NOT_ALLOWED:
		return 405
	case UNKNOWN:
		return 407
	case MODEL_NULL:
		return 303
	case DATA_CORRUPT:
		return 309
	case StatusInternalServerError:
		return 500
	case StatusTooManyRequests:
		return 429
	case StatusForbidden:
		return 403
	case StatusProcessing:
		return 102
	case StatusNoContent:
		return 204
	case StatusNotModified:
		return 304
	case StatusConflict:
		return 409
	case StatusNotImplemented:
		return 501
	case NOT_PERMISSIONS:
		return 507
	case NOT_PERMITTED:
		return 508
	default:
		return 407
	}
}

func GetError(code int) error {
	switch code {
	case 200:
		return SUCCESS
	case 408:
		return DATABASE_TIMEOUT
	case 401:
		return UNAUTHORIZED
	case 400:
		return BAD_REQUEST
	case 404:
		return NOT_FOUND
	case 406:
		return NOT_ACCEPTABLE
	case 405:
		return METHOD_NOT_ALLOWED
	case 407:
		return UNKNOWN
	case 409:
		return DATA_EXIST
	case 303:
		return MODEL_NULL
	case 309:
		return DATA_CORRUPT
	default:
		return UNKNOWN
	}
}

func GetString(e error) string {
	if e == nil {
		return SUCCESS.Error()
	}
	return e.Error()
}
