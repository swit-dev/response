package response

var MsgFlags = map[int32]string{
	SUCCESS:                       "success",
	ERROR:                         "fail",
	AUTH_FAIL:                     "no permission",
	EXPIRED_TOKEN:                 "expired token",
	INVALID_TOKEN:                 "invalid token",
	NOT_EXIST_TOKEN:               "not exist token",
	VALID_TOKEN_NO_FOUND:          "valid token , not found data",
	DUPLICATE_EMAIL:               "This email is already linked with a Swit account.",
	AUTH_CODE_INVALID:             "wrong auth code",
	EMAIL_PASSWORD_NOT_MATCH:      "Please recheck your email and password and try again.",
	FAILED_SEND_MAIL:              "Sorry, Failed to send mail. Please try again.",
	INVALID_EMAIL:                 "email is not valid",
	INVALID_PASSWORD:              "password is not valid",
	NOT_FOUND_EMAIL:               "email not found",
	NOT_FOUND_REDIS_KEY:           "not found token",
	INVALID_CHAR:                  "letters, numbers, and dashes only between 4 and 20 characters with dash(-) not in the first or the last position.",
	DUPLICATE_WORKSPACE_DOMAIN:    "Sorry, that is already in use.",
	INVALID_LOGIN:                 "please login",
	DUPLICATE_WORKSPACE:           "The person is already in your workspace.",
	NO_MEMBERS_WORKSPACE:          "no members in the workspace",
	NOT_FOUND_WORKSPACE:           "workspace or channel is not valid",
	DUPLICATE_CHANNEL:             "duplicate channel name",
	NOT_INVITED_EMAIL:             "this is not an invited email",
	INVALID_CHANNEL:               "channel is invalid",
	INVALID_FILE_SIZE:             "file size too big",
	TOKEN_INVALID:                 "token is invalid",
	NOT_LEAVE_CHANNEL:             "can not leave default channel",
	CHANNEL_HOST_AUTH_FAIL:        "The channel host can not leave that channel. Please assign the Channel Host privilege and try again.",
	MASTER_AUTH_FAIL:              "The Master can not leave that Workspace. Please assign the Master privilege and try again.",
	DUPLICATE_URL:                 "Sorry, that URL address is already in use.",
	NOT_FIND_PROFILE:              "can't not find user profile",
	PASSWORD_NOT_MATCH:            "Passwords do not match.",
	INVALID_IMAGE:                 "image is not valid",
	INVALID_PROJECT:               "project is invalid",
	DUPLICATE_PROJECT:             "duplicate project name",
	PROJECT_HOST_AUTH_FAIL:        "The project host can not leave that project. Please assign the Project Host privilege and try again.",
	INVALID_WORKSPACE_ACCESS:      "You can’t access this workspace. Please check your Swit account again.",
	INVALID_PROJECT_JOIN:          "Not a user joined this project",
	DUPLICATE_TAG:                 "duplicate tag",
	DUPLICATE_MEMBER:              "The member is already participating.",
	INVALID_PRIVATE_PROJECT:       "This task is for a private project and only members who are participating in the project can be viewed.",
	JOIN_PROJECT:                  "You are not in the this project. Would you like to join this project?",
	NOT_ALLOWED_CHAR:              "Special characters are not allowed.",
	PARAMETER_INVALID:             "parameter is invalid",
	SWIT_ERROR:                    "unintended error occurred(router)",
	NOT_FOUND_DATA:                "data not found",
	INVALID_ACCESS_TOKEN:          "access token is invalid",
	VALIDATION_ERROR:              "validation error",
	SWIT_ERROR_MODEL:              "unintended error occurred(model)",
	FAILED_SAVE:                   "failed to save",
	FAILED_UPDATE:                 "failed to update",
	FAILED_DELETE:                 "failed to delete",
	DUPLICATE_DATA:                "duplicate data",
	DELETED_CHANNEL:               "This channel has been deleted.",
	INVALID_WORDS:                 "You can only have 50 words.",
	INVAILD_ARCHIVED_CHANNEL:      "Archived Channel",
	INVAILD_ARCHIVED_PROJECT:      "Archived Project",
	INVALID_PROJECT_COUNT:         "Up to 5 projects",
	INVALID_TASK_COUNT:            "Up to 300 tasks",
	INVALID_TRIAL_PERIOD:          "two-week trial period is over",
	EXPIRED_FILE:                  "expired file",
	CAPACITY_EXCEEDED:             "capacity exceeded",
	NO_FOUND_INFORMATION:          "The information cannot be found. Please try again after refreshing",
	FAILED_DELETE_CATEGORY:        "can not delete category",
	FAIL_ANALYSIS_BATCH:           "fail analysis batch",
	NOT_UPLOAD_SIZE:               "Files bigger than 2GB cannot be uploaded.",
	IMPORT_FILE_ERROR:             "Unable to unzip exported files from Slack. Please try again.",
	EXIST_IN_PROGRESS_IMPORT:      "exist in progress import.",
	INVALID_DOWNLOAD_LINK:         "Invalid download link.",
	INVALID_IMPORT:                "Invalid import.",
	EXIST_IN_ACTION_IMPORT:        "Exist in action import.",
	INVALID_DOMAIN:                "Please recheck your domain and try again.",
	INVALID_DOMAIN_COUNT:          "Up to 5 domain",
	INVALID_DOMAIN_PUBLIC:         "",
	INVALID_PASSWORD_CHAR:         "Use at least 8 characters.",
	INVALID_PASSWORD_SPECIAL_CHAR: "Please use only letters (a~z and A~z), number(0~9) and allowed special characters.",
}

func GetMsg(code int32) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
