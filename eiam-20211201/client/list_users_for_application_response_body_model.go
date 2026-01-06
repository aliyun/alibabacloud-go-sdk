// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUsersForApplicationResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUsersForApplicationResponseBody
	GetTotalCount() *int64
	SetUsers(v []*ListUsersForApplicationResponseBodyUsers) *ListUsersForApplicationResponseBody
	GetUsers() []*ListUsersForApplicationResponseBodyUsers
}

type ListUsersForApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The IDs of the accounts.
	Users []*ListUsersForApplicationResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersForApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersForApplicationResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersForApplicationResponseBody) GetUsers() []*ListUsersForApplicationResponseBodyUsers {
	return s.Users
}

func (s *ListUsersForApplicationResponseBody) SetRequestId(v string) *ListUsersForApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersForApplicationResponseBody) SetTotalCount(v int64) *ListUsersForApplicationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersForApplicationResponseBody) SetUsers(v []*ListUsersForApplicationResponseBodyUsers) *ListUsersForApplicationResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersForApplicationResponseBody) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersForApplicationResponseBodyUsers struct {
	// 应用角色列表。
	ApplicationRoles []*ListUsersForApplicationResponseBodyUsersApplicationRoles `json:"ApplicationRoles,omitempty" xml:"ApplicationRoles,omitempty" type:"Repeated"`
	// The ID of the account.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUsersForApplicationResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForApplicationResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForApplicationResponseBodyUsers) GetApplicationRoles() []*ListUsersForApplicationResponseBodyUsersApplicationRoles {
	return s.ApplicationRoles
}

func (s *ListUsersForApplicationResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersForApplicationResponseBodyUsers) SetApplicationRoles(v []*ListUsersForApplicationResponseBodyUsersApplicationRoles) *ListUsersForApplicationResponseBodyUsers {
	s.ApplicationRoles = v
	return s
}

func (s *ListUsersForApplicationResponseBodyUsers) SetUserId(v string) *ListUsersForApplicationResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersForApplicationResponseBodyUsers) Validate() error {
	if s.ApplicationRoles != nil {
		for _, item := range s.ApplicationRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersForApplicationResponseBodyUsersApplicationRoles struct {
	// 应用角色标识。
	//
	// example:
	//
	// app_role_mkv7rgt4ds8d8v0qtzev2mxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
}

func (s ListUsersForApplicationResponseBodyUsersApplicationRoles) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForApplicationResponseBodyUsersApplicationRoles) GoString() string {
	return s.String()
}

func (s *ListUsersForApplicationResponseBodyUsersApplicationRoles) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *ListUsersForApplicationResponseBodyUsersApplicationRoles) SetApplicationRoleId(v string) *ListUsersForApplicationResponseBodyUsersApplicationRoles {
	s.ApplicationRoleId = &v
	return s
}

func (s *ListUsersForApplicationResponseBodyUsersApplicationRoles) Validate() error {
	return dara.Validate(s)
}
