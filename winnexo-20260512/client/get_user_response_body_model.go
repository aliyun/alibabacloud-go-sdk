// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetUserResponseBody
	GetAccountId() *string
	SetCode(v string) *GetUserResponseBody
	GetCode() *string
	SetDisplayName(v string) *GetUserResponseBody
	GetDisplayName() *string
	SetGmtCreate(v string) *GetUserResponseBody
	GetGmtCreate() *string
	SetIsActive(v bool) *GetUserResponseBody
	GetIsActive() *bool
	SetLastLoginTime(v string) *GetUserResponseBody
	GetLastLoginTime() *string
	SetMessage(v string) *GetUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
	SetRoleCodes(v []*string) *GetUserResponseBody
	GetRoleCodes() []*string
	SetUserGroupIds(v []*string) *GetUserResponseBody
	GetUserGroupIds() []*string
	SetWnUserId(v string) *GetUserResponseBody
	GetWnUserId() *string
}

type GetUserResponseBody struct {
	// WINNEXO 登录账号
	//
	// example:
	//
	// exampleAccountId
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 用户显示名称
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 加入租户时间
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 启用/停用状态
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// 最后登录时间
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	LastLoginTime *string `json:"lastLoginTime,omitempty" xml:"lastLoginTime,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 用户拥有的系统角色 code 列表
	//
	// example:
	//
	// string_value
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
	// 用户所属用户组ID列表
	//
	// example:
	//
	// string_value
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// WINNEXO 平台用户ID
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetUserResponseBody) GetIsActive() *bool {
	return s.IsActive
}

func (s *GetUserResponseBody) GetLastLoginTime() *string {
	return s.LastLoginTime
}

func (s *GetUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *GetUserResponseBody) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *GetUserResponseBody) GetWnUserId() *string {
	return s.WnUserId
}

func (s *GetUserResponseBody) SetAccountId(v string) *GetUserResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetDisplayName(v string) *GetUserResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBody) SetGmtCreate(v string) *GetUserResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetUserResponseBody) SetIsActive(v bool) *GetUserResponseBody {
	s.IsActive = &v
	return s
}

func (s *GetUserResponseBody) SetLastLoginTime(v string) *GetUserResponseBody {
	s.LastLoginTime = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetRoleCodes(v []*string) *GetUserResponseBody {
	s.RoleCodes = v
	return s
}

func (s *GetUserResponseBody) SetUserGroupIds(v []*string) *GetUserResponseBody {
	s.UserGroupIds = v
	return s
}

func (s *GetUserResponseBody) SetWnUserId(v string) *GetUserResponseBody {
	s.WnUserId = &v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	return dara.Validate(s)
}
