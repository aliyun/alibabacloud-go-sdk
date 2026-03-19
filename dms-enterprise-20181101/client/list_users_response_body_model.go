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
	TotalCount *int64                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserList   *ListUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
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
	if s.UserList != nil {
		if err := s.UserList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.User != nil {
		for _, item := range s.User {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersResponseBodyUserListUser struct {
	CurExecuteCount  *int64                                         `json:"CurExecuteCount,omitempty" xml:"CurExecuteCount,omitempty"`
	CurResultCount   *int64                                         `json:"CurResultCount,omitempty" xml:"CurResultCount,omitempty"`
	DingRobot        *string                                        `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	Email            *string                                        `json:"Email,omitempty" xml:"Email,omitempty"`
	LastLoginTime    *string                                        `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	MaxExecuteCount  *int64                                         `json:"MaxExecuteCount,omitempty" xml:"MaxExecuteCount,omitempty"`
	MaxResultCount   *int64                                         `json:"MaxResultCount,omitempty" xml:"MaxResultCount,omitempty"`
	Mobile           *string                                        `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	NickName         *string                                        `json:"NickName,omitempty" xml:"NickName,omitempty"`
	NotificationMode *string                                        `json:"NotificationMode,omitempty" xml:"NotificationMode,omitempty"`
	ParentUid        *string                                        `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	RoleIdList       *ListUsersResponseBodyUserListUserRoleIdList   `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Struct"`
	RoleNameList     *ListUsersResponseBodyUserListUserRoleNameList `json:"RoleNameList,omitempty" xml:"RoleNameList,omitempty" type:"Struct"`
	SignatureMethod  *string                                        `json:"SignatureMethod,omitempty" xml:"SignatureMethod,omitempty"`
	State            *string                                        `json:"State,omitempty" xml:"State,omitempty"`
	Uid              *string                                        `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId           *string                                        `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Webhook          *string                                        `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
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
	if s.RoleIdList != nil {
		if err := s.RoleIdList.Validate(); err != nil {
			return err
		}
	}
	if s.RoleNameList != nil {
		if err := s.RoleNameList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
