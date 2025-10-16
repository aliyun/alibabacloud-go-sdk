// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDirectoryUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListDirectoryUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDirectoryUsersResponseBody
	GetRequestId() *string
	SetUsers(v []*ListDirectoryUsersResponseBodyUsers) *ListDirectoryUsersResponseBody
	GetUsers() []*ListDirectoryUsersResponseBodyUsers
}

type ListDirectoryUsersResponseBody struct {
	// The token used to start the next query. If the value of this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The usernames corresponding to the AD directory. If the AD directory contains only the Administrator and Guest accounts, the Users array will be empty.
	Users []*ListDirectoryUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListDirectoryUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoryUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDirectoryUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDirectoryUsersResponseBody) GetUsers() []*ListDirectoryUsersResponseBodyUsers {
	return s.Users
}

func (s *ListDirectoryUsersResponseBody) SetNextToken(v string) *ListDirectoryUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersResponseBody) SetRequestId(v string) *ListDirectoryUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDirectoryUsersResponseBody) SetUsers(v []*ListDirectoryUsersResponseBodyUsers) *ListDirectoryUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListDirectoryUsersResponseBody) Validate() error {
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

type ListDirectoryUsersResponseBodyUsers struct {
	// The number of assigned cloud computers.
	//
	// example:
	//
	// 2
	AssignedDesktopNumber *int32 `json:"AssignedDesktopNumber,omitempty" xml:"AssignedDesktopNumber,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// Alice
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	DisplayNameNew *string `json:"DisplayNameNew,omitempty" xml:"DisplayNameNew,omitempty"`
	// The email address.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the user.
	//
	// example:
	//
	// Alice
	EndUser *string `json:"EndUser,omitempty" xml:"EndUser,omitempty"`
	// The mobile number.
	//
	// example:
	//
	// 130********
	Phone             *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListDirectoryUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoryUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponseBodyUsers) GetAssignedDesktopNumber() *int32 {
	return s.AssignedDesktopNumber
}

func (s *ListDirectoryUsersResponseBodyUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListDirectoryUsersResponseBodyUsers) GetDisplayNameNew() *string {
	return s.DisplayNameNew
}

func (s *ListDirectoryUsersResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *ListDirectoryUsersResponseBodyUsers) GetEndUser() *string {
	return s.EndUser
}

func (s *ListDirectoryUsersResponseBodyUsers) GetPhone() *string {
	return s.Phone
}

func (s *ListDirectoryUsersResponseBodyUsers) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListDirectoryUsersResponseBodyUsers) SetAssignedDesktopNumber(v int32) *ListDirectoryUsersResponseBodyUsers {
	s.AssignedDesktopNumber = &v
	return s
}

func (s *ListDirectoryUsersResponseBodyUsers) SetDisplayName(v string) *ListDirectoryUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListDirectoryUsersResponseBodyUsers) SetDisplayNameNew(v string) *ListDirectoryUsersResponseBodyUsers {
	s.DisplayNameNew = &v
	return s
}

func (s *ListDirectoryUsersResponseBodyUsers) SetEmail(v string) *ListDirectoryUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListDirectoryUsersResponseBodyUsers) SetEndUser(v string) *ListDirectoryUsersResponseBodyUsers {
	s.EndUser = &v
	return s
}

func (s *ListDirectoryUsersResponseBodyUsers) SetPhone(v string) *ListDirectoryUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *ListDirectoryUsersResponseBodyUsers) SetUserPrincipalName(v string) *ListDirectoryUsersResponseBodyUsers {
	s.UserPrincipalName = &v
	return s
}

func (s *ListDirectoryUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
