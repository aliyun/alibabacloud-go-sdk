// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListUsersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListUsersResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUsersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListUsersResponseBody
	GetTotalCount() *int64
	SetUserList(v *ListUsersResponseBodyUserList) *ListUsersResponseBody
	GetUserList() *ListUsersResponseBodyUserList
}

type ListUsersResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69161997-6030-53BA-A333-DBEC83B732FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details of users.
	UserList *ListUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUsersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUsersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersResponseBody) GetUserList() *ListUsersResponseBodyUserList {
	return s.UserList
}

func (s *ListUsersResponseBody) SetErrorCode(v string) *ListUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUsersResponseBody) SetErrorMessage(v string) *ListUsersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUserList(v *ListUsersResponseBodyUserList) *ListUsersResponseBody {
	s.UserList = v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyUserList struct {
	User []*ListUsersResponseBodyUserListUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUserList) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserList) GetUser() []*ListUsersResponseBodyUserListUser {
	return s.User
}

func (s *ListUsersResponseBodyUserList) SetUser(v []*ListUsersResponseBodyUserListUser) *ListUsersResponseBodyUserList {
	s.User = v
	return s
}

func (s *ListUsersResponseBodyUserList) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyUserListUser struct {
	// The number of queries that were performed on the current day.
	//
	// example:
	//
	// 0
	CurExecuteCount *int64 `json:"CurExecuteCount,omitempty" xml:"CurExecuteCount,omitempty"`
	// The number of rows that were queried on the current day.
	//
	// example:
	//
	// 0
	CurResultCount *int64 `json:"CurResultCount,omitempty" xml:"CurResultCount,omitempty"`
	// The DingTalk chatbot URL that is used to receive notifications.
	//
	// >
	//
	// 	- The system returns this parameter if you have set a DingTalk chatbot URL in the console. To set a DingTalk chatbot URL in the console, move the pointer over the profile picture in the upper-right corner and click the Edit icon next to **Notice**.
	//
	// 	- The system does not return this parameter if you have not set a DingTalk chatbot URL.
	//
	// example:
	//
	// https://XXX.dingtalk.com/robot/send?access_token=9b7a4a562cbe7fcdea9962afac7b9d7b4504d564948083419750f9cafa78e4ef
	DingRobot *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	// The email address that is used to receive notifications.
	//
	// >
	//
	// 	- The system returns this parameter if you have set an email address in the console. To set an email address in the console, move the pointer over the profile picture in the upper-right corner and click the Edit icon next to **Notice**.
	//
	// 	- The system does not return this parameter if you have not set an email address.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The time when the user last logged on to the console.
	//
	// example:
	//
	// 2021-11-08 13:43:43
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
	// 50000
	MaxResultCount *int64 `json:"MaxResultCount,omitempty" xml:"MaxResultCount,omitempty"`
	// The mobile phone number of the user.
	//
	// >
	//
	// 	- The system returns this parameter if you have set a mobile phone number in the console. To set a mobile phone number in the console, move the pointer over the profile picture in the upper-right corner and click the Edit icon next to **Notice**.
	//
	// 	- The system does not return this parameter if you have not set a mobile phone number.
	//
	// example:
	//
	// 1389999****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// test_NickName
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
	// DINGROBOT
	NotificationMode *string `json:"NotificationMode,omitempty" xml:"NotificationMode,omitempty"`
	// The ID of the Alibaba Cloud account of the user.
	//
	// example:
	//
	// 140692647406****
	ParentUid *string `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	// The IDs of the roles.
	RoleIdList *ListUsersResponseBodyUserListUserRoleIdList `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Struct"`
	// The names of roles.
	RoleNameList *ListUsersResponseBodyUserListUserRoleNameList `json:"RoleNameList,omitempty" xml:"RoleNameList,omitempty" type:"Struct"`
	// The signature method that is used to secure connections when a webhook URL is used. Valid values:
	//
	// 	- **NONE**: no signature.
	//
	// 	- **HMAC_SHA1**: HMAC_SHA1.
	//
	// example:
	//
	// HMAC_SHA1
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
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 167382665015****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 8****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The webhook URL that is used to receive notifications.
	//
	// >
	//
	// 	- If you have set a webhook URL, DMS sends notifications to the specified URL.
	//
	// 	- The system does not return this parameter if you have not set a webhook URL.
	//
	// example:
	//
	// http://dms-XXX.aliyun.com:8***
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s ListUsersResponseBodyUserListUser) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUserListUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserListUser) GetCurExecuteCount() *int64 {
	return s.CurExecuteCount
}

func (s *ListUsersResponseBodyUserListUser) GetCurResultCount() *int64 {
	return s.CurResultCount
}

func (s *ListUsersResponseBodyUserListUser) GetDingRobot() *string {
	return s.DingRobot
}

func (s *ListUsersResponseBodyUserListUser) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyUserListUser) GetLastLoginTime() *string {
	return s.LastLoginTime
}

func (s *ListUsersResponseBodyUserListUser) GetMaxExecuteCount() *int64 {
	return s.MaxExecuteCount
}

func (s *ListUsersResponseBodyUserListUser) GetMaxResultCount() *int64 {
	return s.MaxResultCount
}

func (s *ListUsersResponseBodyUserListUser) GetMobile() *string {
	return s.Mobile
}

func (s *ListUsersResponseBodyUserListUser) GetNickName() *string {
	return s.NickName
}

func (s *ListUsersResponseBodyUserListUser) GetNotificationMode() *string {
	return s.NotificationMode
}

func (s *ListUsersResponseBodyUserListUser) GetParentUid() *string {
	return s.ParentUid
}

func (s *ListUsersResponseBodyUserListUser) GetRoleIdList() *ListUsersResponseBodyUserListUserRoleIdList {
	return s.RoleIdList
}

func (s *ListUsersResponseBodyUserListUser) GetRoleNameList() *ListUsersResponseBodyUserListUserRoleNameList {
	return s.RoleNameList
}

func (s *ListUsersResponseBodyUserListUser) GetSignatureMethod() *string {
	return s.SignatureMethod
}

func (s *ListUsersResponseBodyUserListUser) GetState() *string {
	return s.State
}

func (s *ListUsersResponseBodyUserListUser) GetUid() *string {
	return s.Uid
}

func (s *ListUsersResponseBodyUserListUser) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersResponseBodyUserListUser) GetWebhook() *string {
	return s.Webhook
}

func (s *ListUsersResponseBodyUserListUser) SetCurExecuteCount(v int64) *ListUsersResponseBodyUserListUser {
	s.CurExecuteCount = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetCurResultCount(v int64) *ListUsersResponseBodyUserListUser {
	s.CurResultCount = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetDingRobot(v string) *ListUsersResponseBodyUserListUser {
	s.DingRobot = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetEmail(v string) *ListUsersResponseBodyUserListUser {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetLastLoginTime(v string) *ListUsersResponseBodyUserListUser {
	s.LastLoginTime = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetMaxExecuteCount(v int64) *ListUsersResponseBodyUserListUser {
	s.MaxExecuteCount = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetMaxResultCount(v int64) *ListUsersResponseBodyUserListUser {
	s.MaxResultCount = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetMobile(v string) *ListUsersResponseBodyUserListUser {
	s.Mobile = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetNickName(v string) *ListUsersResponseBodyUserListUser {
	s.NickName = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetNotificationMode(v string) *ListUsersResponseBodyUserListUser {
	s.NotificationMode = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetParentUid(v string) *ListUsersResponseBodyUserListUser {
	s.ParentUid = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetRoleIdList(v *ListUsersResponseBodyUserListUserRoleIdList) *ListUsersResponseBodyUserListUser {
	s.RoleIdList = v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetRoleNameList(v *ListUsersResponseBodyUserListUserRoleNameList) *ListUsersResponseBodyUserListUser {
	s.RoleNameList = v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetSignatureMethod(v string) *ListUsersResponseBodyUserListUser {
	s.SignatureMethod = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetState(v string) *ListUsersResponseBodyUserListUser {
	s.State = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetUid(v string) *ListUsersResponseBodyUserListUser {
	s.Uid = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetUserId(v string) *ListUsersResponseBodyUserListUser {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetWebhook(v string) *ListUsersResponseBodyUserListUser {
	s.Webhook = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyUserListUserRoleIdList struct {
	RoleIds []*int32 `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUserListUserRoleIdList) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUserListUserRoleIdList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserListUserRoleIdList) GetRoleIds() []*int32 {
	return s.RoleIds
}

func (s *ListUsersResponseBodyUserListUserRoleIdList) SetRoleIds(v []*int32) *ListUsersResponseBodyUserListUserRoleIdList {
	s.RoleIds = v
	return s
}

func (s *ListUsersResponseBodyUserListUserRoleIdList) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyUserListUserRoleNameList struct {
	RoleNames []*string `json:"RoleNames,omitempty" xml:"RoleNames,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUserListUserRoleNameList) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUserListUserRoleNameList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserListUserRoleNameList) GetRoleNames() []*string {
	return s.RoleNames
}

func (s *ListUsersResponseBodyUserListUserRoleNameList) SetRoleNames(v []*string) *ListUsersResponseBodyUserListUserRoleNameList {
	s.RoleNames = v
	return s
}

func (s *ListUsersResponseBodyUserListUserRoleNameList) Validate() error {
	return dara.Validate(s)
}
