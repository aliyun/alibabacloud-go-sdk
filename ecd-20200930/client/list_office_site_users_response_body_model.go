// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfficeSiteUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListOfficeSiteUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListOfficeSiteUsersResponseBody
	GetRequestId() *string
	SetUsers(v []*ListOfficeSiteUsersResponseBodyUsers) *ListOfficeSiteUsersResponseBody
	GetUsers() []*ListOfficeSiteUsersResponseBodyUsers
}

type ListOfficeSiteUsersResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC21DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The usernames of the AD accounts.\\
	//
	// If the only Administrator and Guest users exist in the enterprise AD directory, an empty User array is returned.
	Users []*ListOfficeSiteUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListOfficeSiteUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOfficeSiteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOfficeSiteUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOfficeSiteUsersResponseBody) GetUsers() []*ListOfficeSiteUsersResponseBodyUsers {
	return s.Users
}

func (s *ListOfficeSiteUsersResponseBody) SetNextToken(v string) *ListOfficeSiteUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBody) SetRequestId(v string) *ListOfficeSiteUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBody) SetUsers(v []*ListOfficeSiteUsersResponseBodyUsers) *ListOfficeSiteUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListOfficeSiteUsersResponseBody) Validate() error {
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

type ListOfficeSiteUsersResponseBodyUsers struct {
	AssignedDesktopNumber *int32 `json:"AssignedDesktopNumber,omitempty" xml:"AssignedDesktopNumber,omitempty"`
	// The display name of the AD account.
	//
	// example:
	//
	// Alice
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	DisplayNameNew *string `json:"DisplayNameNew,omitempty" xml:"DisplayNameNew,omitempty"`
	Email          *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username of the AD account.
	//
	// example:
	//
	// Alice
	EndUser           *string `json:"EndUser,omitempty" xml:"EndUser,omitempty"`
	Phone             *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListOfficeSiteUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListOfficeSiteUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponseBodyUsers) GetAssignedDesktopNumber() *int32 {
	return s.AssignedDesktopNumber
}

func (s *ListOfficeSiteUsersResponseBodyUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListOfficeSiteUsersResponseBodyUsers) GetDisplayNameNew() *string {
	return s.DisplayNameNew
}

func (s *ListOfficeSiteUsersResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *ListOfficeSiteUsersResponseBodyUsers) GetEndUser() *string {
	return s.EndUser
}

func (s *ListOfficeSiteUsersResponseBodyUsers) GetPhone() *string {
	return s.Phone
}

func (s *ListOfficeSiteUsersResponseBodyUsers) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetAssignedDesktopNumber(v int32) *ListOfficeSiteUsersResponseBodyUsers {
	s.AssignedDesktopNumber = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetDisplayName(v string) *ListOfficeSiteUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetDisplayNameNew(v string) *ListOfficeSiteUsersResponseBodyUsers {
	s.DisplayNameNew = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetEmail(v string) *ListOfficeSiteUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetEndUser(v string) *ListOfficeSiteUsersResponseBodyUsers {
	s.EndUser = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetPhone(v string) *ListOfficeSiteUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetUserPrincipalName(v string) *ListOfficeSiteUsersResponseBodyUsers {
	s.UserPrincipalName = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
