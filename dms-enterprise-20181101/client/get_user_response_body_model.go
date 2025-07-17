// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetUserResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetUserResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserResponseBody
	GetSuccess() *bool
	SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody
	GetUser() *GetUserResponseBodyUser
}

type GetUserResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Unknown server error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 804BB128-CAFA-5DD0-BA1E-43DDE488****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about the user.
	User *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserResponseBody) GetUser() *GetUserResponseBodyUser {
	return s.User
}

func (s *GetUserResponseBody) SetErrorCode(v string) *GetUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserResponseBody) SetErrorMessage(v string) *GetUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetSuccess(v bool) *GetUserResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyUser struct {
	// The number of queries that are performed on the current day.
	//
	// example:
	//
	// 3
	CurExecuteCount *int64 `json:"CurExecuteCount,omitempty" xml:"CurExecuteCount,omitempty"`
	// The number of rows that are queried on the current day.
	//
	// example:
	//
	// 28
	CurResultCount *int64 `json:"CurResultCount,omitempty" xml:"CurResultCount,omitempty"`
	// The DingTalk chatbot URL that is used to receive notifications.
	//
	// >
	//
	// 	- The system returns this parameter if the user has set a DingTalk chatbot URL in the console. To set a DingTalk chatbot URL in the console, move the pointer over the profile picture in the upper-right corner and click the Edit icon next to **Notice**.
	//
	// 	- The system does not return this parameter if the user has not set a DingTalk chatbot URL.
	//
	// example:
	//
	// https://XXX.dingtalk.com/robot/send?access_token=***
	DingRobot *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	// The email address that is used to receive notifications.
	//
	// >
	//
	// 	- The system returns this parameter if the user has set an email address in the console. To set an email address in the console, move the pointer over the profile picture in the upper-right corner and click the Edit icon next to **Notice**.
	//
	// 	- The system does not return this parameter if the user has not set an email address.
	//
	// example:
	//
	// Uesr_email
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The last point in time when the user logged on to the console.
	//
	// example:
	//
	// 2021-11-08 11:26:21
	LastLoginTime *string `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	// The maximum number of queries that can be performed on the current day.
	//
	// example:
	//
	// 2000
	MaxExecuteCount *int64 `json:"MaxExecuteCount,omitempty" xml:"MaxExecuteCount,omitempty"`
	// The maximum number of rows that can be queried on the current day.
	//
	// example:
	//
	// 10000
	MaxResultCount *int64 `json:"MaxResultCount,omitempty" xml:"MaxResultCount,omitempty"`
	// The mobile number of the user.
	//
	// >
	//
	// 	- The system returns this parameter if the user has set a mobile phone number in the console. To set a mobile phone number in the console, move the pointer over the profile picture in the upper-right corner and click the Edit icon next to **Notice**.
	//
	// 	- The system does not return this parameter if the user has not set a mobile phone number.
	//
	// example:
	//
	// 1389223****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// User_NickName
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The notification method. The system returns one or more values. Valid values:
	//
	// 	- **SMS**: text message
	//
	// 	- **EMAIL**: email.
	//
	// 	- **DINGDING**: DingTalk.
	//
	// 	- **DINGROBOT**: DingTalk chatbot.
	//
	// 	- **WEBHOOK**: webhook.
	//
	// example:
	//
	// EMAIL
	NotificationMode *string `json:"NotificationMode,omitempty" xml:"NotificationMode,omitempty"`
	// The UID of the Alibaba Cloud account of the user.
	//
	// > An Alibaba Cloud account can contain one or more RAM users.
	//
	// example:
	//
	// 140692647406****
	ParentUid *int64 `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	// The list of role IDs.
	RoleIdList *GetUserResponseBodyUserRoleIdList `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Struct"`
	// The list of role names.
	RoleNameList *GetUserResponseBodyUserRoleNameList `json:"RoleNameList,omitempty" xml:"RoleNameList,omitempty" type:"Struct"`
	// The signature method that is used to secure connections when a webhook URL is used. Valid values:
	//
	// 	- **NONE**: no signature.
	//
	// 	- **HMAC_SHA1**: HMAC_SHA1.
	//
	// example:
	//
	// NONE
	SignatureMethod *string `json:"SignatureMethod,omitempty" xml:"SignatureMethod,omitempty"`
	// The status of the user. Valid values:
	//
	// 	- **NORMAL**: The user is normal.
	//
	// 	- **DISABLE**: The user is disabled.
	//
	// 	- **DELETE**: The user is deleted.
	//
	// example:
	//
	// NORMAL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The UID of the user.
	//
	// example:
	//
	// 22275482072787****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The webhook URL that is used to receive notifications.
	//
	// >
	//
	// 	- If the user has set a webhook URL, DMS sends notifications to the specified URL.
	//
	// 	- The system does not return this parameter if the user has not set a webhook URL.
	//
	// example:
	//
	// http://dms-XXX.aliyun.com:8***
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) GetCurExecuteCount() *int64 {
	return s.CurExecuteCount
}

func (s *GetUserResponseBodyUser) GetCurResultCount() *int64 {
	return s.CurResultCount
}

func (s *GetUserResponseBodyUser) GetDingRobot() *string {
	return s.DingRobot
}

func (s *GetUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBodyUser) GetLastLoginTime() *string {
	return s.LastLoginTime
}

func (s *GetUserResponseBodyUser) GetMaxExecuteCount() *int64 {
	return s.MaxExecuteCount
}

func (s *GetUserResponseBodyUser) GetMaxResultCount() *int64 {
	return s.MaxResultCount
}

func (s *GetUserResponseBodyUser) GetMobile() *string {
	return s.Mobile
}

func (s *GetUserResponseBodyUser) GetNickName() *string {
	return s.NickName
}

func (s *GetUserResponseBodyUser) GetNotificationMode() *string {
	return s.NotificationMode
}

func (s *GetUserResponseBodyUser) GetParentUid() *int64 {
	return s.ParentUid
}

func (s *GetUserResponseBodyUser) GetRoleIdList() *GetUserResponseBodyUserRoleIdList {
	return s.RoleIdList
}

func (s *GetUserResponseBodyUser) GetRoleNameList() *GetUserResponseBodyUserRoleNameList {
	return s.RoleNameList
}

func (s *GetUserResponseBodyUser) GetSignatureMethod() *string {
	return s.SignatureMethod
}

func (s *GetUserResponseBodyUser) GetState() *string {
	return s.State
}

func (s *GetUserResponseBodyUser) GetUid() *string {
	return s.Uid
}

func (s *GetUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *GetUserResponseBodyUser) GetWebhook() *string {
	return s.Webhook
}

func (s *GetUserResponseBodyUser) SetCurExecuteCount(v int64) *GetUserResponseBodyUser {
	s.CurExecuteCount = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCurResultCount(v int64) *GetUserResponseBodyUser {
	s.CurResultCount = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDingRobot(v string) *GetUserResponseBodyUser {
	s.DingRobot = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmail(v string) *GetUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLastLoginTime(v string) *GetUserResponseBodyUser {
	s.LastLoginTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMaxExecuteCount(v int64) *GetUserResponseBodyUser {
	s.MaxExecuteCount = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMaxResultCount(v int64) *GetUserResponseBodyUser {
	s.MaxResultCount = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMobile(v string) *GetUserResponseBodyUser {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBodyUser) SetNickName(v string) *GetUserResponseBodyUser {
	s.NickName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetNotificationMode(v string) *GetUserResponseBodyUser {
	s.NotificationMode = &v
	return s
}

func (s *GetUserResponseBodyUser) SetParentUid(v int64) *GetUserResponseBodyUser {
	s.ParentUid = &v
	return s
}

func (s *GetUserResponseBodyUser) SetRoleIdList(v *GetUserResponseBodyUserRoleIdList) *GetUserResponseBodyUser {
	s.RoleIdList = v
	return s
}

func (s *GetUserResponseBodyUser) SetRoleNameList(v *GetUserResponseBodyUserRoleNameList) *GetUserResponseBodyUser {
	s.RoleNameList = v
	return s
}

func (s *GetUserResponseBodyUser) SetSignatureMethod(v string) *GetUserResponseBodyUser {
	s.SignatureMethod = &v
	return s
}

func (s *GetUserResponseBodyUser) SetState(v string) *GetUserResponseBodyUser {
	s.State = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUid(v string) *GetUserResponseBodyUser {
	s.Uid = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetWebhook(v string) *GetUserResponseBodyUser {
	s.Webhook = &v
	return s
}

func (s *GetUserResponseBodyUser) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyUserRoleIdList struct {
	RoleIds []*int32 `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUserRoleIdList) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUserRoleIdList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserRoleIdList) GetRoleIds() []*int32 {
	return s.RoleIds
}

func (s *GetUserResponseBodyUserRoleIdList) SetRoleIds(v []*int32) *GetUserResponseBodyUserRoleIdList {
	s.RoleIds = v
	return s
}

func (s *GetUserResponseBodyUserRoleIdList) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyUserRoleNameList struct {
	RoleNames []*string `json:"RoleNames,omitempty" xml:"RoleNames,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUserRoleNameList) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUserRoleNameList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserRoleNameList) GetRoleNames() []*string {
	return s.RoleNames
}

func (s *GetUserResponseBodyUserRoleNameList) SetRoleNames(v []*string) *GetUserResponseBodyUserRoleNameList {
	s.RoleNames = v
	return s
}

func (s *GetUserResponseBodyUserRoleNameList) Validate() error {
	return dara.Validate(s)
}
