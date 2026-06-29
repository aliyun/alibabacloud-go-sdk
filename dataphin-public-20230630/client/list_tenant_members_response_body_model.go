// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTenantMembersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTenantMembersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTenantMembersResponseBody
	GetMessage() *string
	SetPageResult(v *ListTenantMembersResponseBodyPageResult) *ListTenantMembersResponseBody
	GetPageResult() *ListTenantMembersResponseBodyPageResult
	SetRequestId(v string) *ListTenantMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTenantMembersResponseBody
	GetSuccess() *bool
}

type ListTenantMembersResponseBody struct {
	// The error code. OK indicates a successful request.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paginated query result.
	PageResult *ListTenantMembersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTenantMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTenantMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTenantMembersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTenantMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTenantMembersResponseBody) GetPageResult() *ListTenantMembersResponseBodyPageResult {
	return s.PageResult
}

func (s *ListTenantMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTenantMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTenantMembersResponseBody) SetCode(v string) *ListTenantMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListTenantMembersResponseBody) SetHttpStatusCode(v int32) *ListTenantMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTenantMembersResponseBody) SetMessage(v string) *ListTenantMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListTenantMembersResponseBody) SetPageResult(v *ListTenantMembersResponseBodyPageResult) *ListTenantMembersResponseBody {
	s.PageResult = v
	return s
}

func (s *ListTenantMembersResponseBody) SetRequestId(v string) *ListTenantMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantMembersResponseBody) SetSuccess(v bool) *ListTenantMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListTenantMembersResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTenantMembersResponseBodyPageResult struct {
	// The total number of entries.
	//
	// example:
	//
	// 110
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of users.
	UserList []*ListTenantMembersResponseBodyPageResultUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListTenantMembersResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListTenantMembersResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTenantMembersResponseBodyPageResult) GetUserList() []*ListTenantMembersResponseBodyPageResultUserList {
	return s.UserList
}

func (s *ListTenantMembersResponseBodyPageResult) SetTotalCount(v int32) *ListTenantMembersResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResult) SetUserList(v []*ListTenantMembersResponseBodyPageResultUserList) *ListTenantMembersResponseBodyPageResult {
	s.UserList = v
	return s
}

func (s *ListTenantMembersResponseBodyPageResult) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTenantMembersResponseBodyPageResultUserList struct {
	// The account name.
	//
	// example:
	//
	// zhangsan
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The DingTalk number.
	//
	// example:
	//
	// dd123123
	DingNumber *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The display name of the user without status.
	//
	// example:
	//
	// zhangsan
	DisplayNameWithoutStatus *string `json:"DisplayNameWithoutStatus,omitempty" xml:"DisplayNameWithoutStatus,omitempty"`
	// Indicates whether the IP address whitelist is enabled.
	//
	// example:
	//
	// true
	EnableWhiteIp *string `json:"EnableWhiteIp,omitempty" xml:"EnableWhiteIp,omitempty"`
	// The time when the user was created.
	//
	// example:
	//
	// 1730000000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the user was last modified.
	//
	// example:
	//
	// 1730000000000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 132321
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The email address.
	//
	// example:
	//
	// 123@aliyun.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 13888888888
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The username.
	//
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// susan
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The real name of the user.
	//
	// example:
	//
	// 张三
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// The list of member roles.
	RoleList []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// The user source ID.
	//
	// example:
	//
	// 213213232422222
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The user source.
	//
	// example:
	//
	// aliyun
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The list of user groups to which the user belongs.
	UserGroupList []*ListTenantMembersResponseBodyPageResultUserListUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
	// The IP address whitelist.
	//
	// example:
	//
	// 0.0.0.0/0
	WhiteIp *string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty"`
}

func (s ListTenantMembersResponseBodyPageResultUserList) String() string {
	return dara.Prettify(s)
}

func (s ListTenantMembersResponseBodyPageResultUserList) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetAccountName() *string {
	return s.AccountName
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetDingNumber() *string {
	return s.DingNumber
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetDisplayNameWithoutStatus() *string {
	return s.DisplayNameWithoutStatus
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetEnableWhiteIp() *string {
	return s.EnableWhiteIp
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetId() *string {
	return s.Id
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetMail() *string {
	return s.Mail
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetName() *string {
	return s.Name
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetNickName() *string {
	return s.NickName
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetRealName() *string {
	return s.RealName
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetRoleList() []*string {
	return s.RoleList
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetSourceId() *string {
	return s.SourceId
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetUserGroupList() []*ListTenantMembersResponseBodyPageResultUserListUserGroupList {
	return s.UserGroupList
}

func (s *ListTenantMembersResponseBodyPageResultUserList) GetWhiteIp() *string {
	return s.WhiteIp
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetAccountName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.AccountName = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetDingNumber(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.DingNumber = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetDisplayName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.DisplayName = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetDisplayNameWithoutStatus(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.DisplayNameWithoutStatus = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetEnableWhiteIp(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.EnableWhiteIp = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetGmtCreate(v int64) *ListTenantMembersResponseBodyPageResultUserList {
	s.GmtCreate = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetGmtModified(v int64) *ListTenantMembersResponseBodyPageResultUserList {
	s.GmtModified = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetId(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.Id = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetMail(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.Mail = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetMobilePhone(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.MobilePhone = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.Name = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetNickName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.NickName = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetRealName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.RealName = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetRoleList(v []*string) *ListTenantMembersResponseBodyPageResultUserList {
	s.RoleList = v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetSourceId(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.SourceId = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetSourceType(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.SourceType = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetUserGroupList(v []*ListTenantMembersResponseBodyPageResultUserListUserGroupList) *ListTenantMembersResponseBodyPageResultUserList {
	s.UserGroupList = v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetWhiteIp(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.WhiteIp = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) Validate() error {
	if s.UserGroupList != nil {
		for _, item := range s.UserGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTenantMembersResponseBodyPageResultUserListUserGroupList struct {
	// Indicates whether the user group is enabled.
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description.
	//
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// 121313
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListTenantMembersResponseBodyPageResultUserListUserGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListTenantMembersResponseBodyPageResultUserListUserGroupList) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) GetActive() *bool {
	return s.Active
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) GetDescription() *string {
	return s.Description
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) GetId() *string {
	return s.Id
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) GetName() *string {
	return s.Name
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) SetActive(v bool) *ListTenantMembersResponseBodyPageResultUserListUserGroupList {
	s.Active = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) SetDescription(v string) *ListTenantMembersResponseBodyPageResultUserListUserGroupList {
	s.Description = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) SetId(v string) *ListTenantMembersResponseBodyPageResultUserListUserGroupList {
	s.Id = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) SetName(v string) *ListTenantMembersResponseBodyPageResultUserListUserGroupList {
	s.Name = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) Validate() error {
	return dara.Validate(s)
}
