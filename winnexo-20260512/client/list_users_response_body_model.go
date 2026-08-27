// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUsersResponseBody
	GetCode() *string
	SetItems(v []*ListUsersResponseBodyItems) *ListUsersResponseBody
	GetItems() []*ListUsersResponseBodyItems
	SetMessage(v string) *ListUsersResponseBody
	GetMessage() *string
	SetPage(v int64) *ListUsersResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListUsersResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListUsersResponseBody
	GetTotal() *int64
}

type ListUsersResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The user information.
	Items []*ListUsersResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of data entries in the project.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUsersResponseBody) GetItems() []*ListUsersResponseBodyItems {
	return s.Items
}

func (s *ListUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUsersResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListUsersResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetItems(v []*ListUsersResponseBodyItems) *ListUsersResponseBody {
	s.Items = v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetPage(v int64) *ListUsersResponseBody {
	s.Page = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int64) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotal(v int64) *ListUsersResponseBody {
	s.Total = &v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersResponseBodyItems struct {
	// The account ID.
	//
	// example:
	//
	// exampleAccountId
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The display name of the tool.
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
	// Indicates whether the account is activated:
	//
	// - 1: Activated.
	//
	// - 0: Not activated.
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
	// The list of new system role codes (full replacement, at least one role must be included). Valid values: SUPER_ADMIN / SYSTEM_ADMIN / SEMANTIC_ADMIN / SKILL_ADMIN / KB_ADMIN / AGENT_ADMIN / APPLICATION_USER.
	//
	// example:
	//
	// string_value
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 1
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListUsersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyItems) GetAccountId() *string {
	return s.AccountId
}

func (s *ListUsersResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListUsersResponseBodyItems) GetIsActive() *bool {
	return s.IsActive
}

func (s *ListUsersResponseBodyItems) GetLastLoginTime() *string {
	return s.LastLoginTime
}

func (s *ListUsersResponseBodyItems) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *ListUsersResponseBodyItems) GetUserId() *int64 {
	return s.UserId
}

func (s *ListUsersResponseBodyItems) SetAccountId(v string) *ListUsersResponseBodyItems {
	s.AccountId = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetDisplayName(v string) *ListUsersResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetGmtCreate(v string) *ListUsersResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetIsActive(v bool) *ListUsersResponseBodyItems {
	s.IsActive = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetLastLoginTime(v string) *ListUsersResponseBodyItems {
	s.LastLoginTime = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetRoleCodes(v []*string) *ListUsersResponseBodyItems {
	s.RoleCodes = v
	return s
}

func (s *ListUsersResponseBodyItems) SetUserId(v int64) *ListUsersResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
