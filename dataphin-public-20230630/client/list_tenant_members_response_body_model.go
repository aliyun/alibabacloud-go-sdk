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
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListTenantMembersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 110
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserList   []*ListTenantMembersResponseBodyPageResultUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
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
	// example:
	//
	// zhangsan
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// dd123123
	DingNumber *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// zhangsan
	DisplayNameWithoutStatus *string `json:"DisplayNameWithoutStatus,omitempty" xml:"DisplayNameWithoutStatus,omitempty"`
	// example:
	//
	// true
	EnableWhiteIp *string `json:"EnableWhiteIp,omitempty" xml:"EnableWhiteIp,omitempty"`
	// example:
	//
	// 1730000000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1730000000000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 132321
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123@aliyun.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// example:
	//
	// 13888888888
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// susan
	NickName *string   `json:"NickName,omitempty" xml:"NickName,omitempty"`
	RealName *string   `json:"RealName,omitempty" xml:"RealName,omitempty"`
	RoleList []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// example:
	//
	// 213213232422222
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// aliyun
	SourceType    *string                                                         `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	UserGroupList []*ListTenantMembersResponseBodyPageResultUserListUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
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
	// example:
	//
	// true
	Active      *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 121313
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
