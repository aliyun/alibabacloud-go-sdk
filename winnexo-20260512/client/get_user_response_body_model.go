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
	// The account ID.
	//
	// example:
	//
	// exampleAccountId
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Indicates whether the account is activated. Valid values:
	//
	//  - **true**: Activated.
	//
	// - **false**: Not activated.
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// The last logon time.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	LastLoginTime *string `json:"lastLoginTime,omitempty" xml:"lastLoginTime,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of system role codes (full replacement, must contain at least one role). Valid values: SUPER_ADMIN, SYSTEM_ADMIN, SEMANTIC_ADMIN, SKILL_ADMIN, KB_ADMIN, AGENT_ADMIN, and APPLICATION_USER.
	//
	// example:
	//
	// string_value
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
	// The list of user group IDs to which the user belongs.
	//
	// example:
	//
	// string_value
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The WINNEXO platform user ID.
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
