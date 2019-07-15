package switResponseCode

import (
	"reflect"
)

const (
	SUCCESS = 200
	ERROR   = 500

	NO_PERMISSION        = 401
	EXPIRED_TOKEN        = 402
	INVALID_TOKEN        = 403
	NOT_EXIST_TOKEN      = 405
	VALID_TOKEN_NO_FOUND = 406

	DUPLICATE_WORKSPACE_DOMAIN = 509

	INVALID_PARAMETER = 600

	SWIT_ERROR = 601
)

//JSONResponse JSONResponse
func NvlArray(data interface{}) interface{} {
	if reflect.ValueOf(data).Len() == 0 {
		return []string{}
	}
	return data
}
