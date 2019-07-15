package switResponseCode

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	NO_PERMISSION:        "no permission",
	EXPIRED_TOKEN:        "expired token",
	INVALID_TOKEN:        "invalid token",
	NOT_EXIST_TOKEN:      "not exist token",
	VALID_TOKEN_NO_FOUND: "valid token , not found data",

	INVALID_PARAMETER:          "parameter is invalid",
	DUPLICATE_WORKSPACE_DOMAIN: "Sorry, that is already in use.",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
