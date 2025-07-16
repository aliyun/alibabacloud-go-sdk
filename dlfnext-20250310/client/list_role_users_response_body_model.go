// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoleUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListRoleUsersResponseBody
	GetNextPageToken() *string
	SetUsers(v []*User) *ListRoleUsersResponseBody
	GetUsers() []*User
}

type ListRoleUsersResponseBody struct {
	// example:
	//
	// token!
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Users         []*User `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListRoleUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoleUsersResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListRoleUsersResponseBody) GetUsers() []*User {
	return s.Users
}

func (s *ListRoleUsersResponseBody) SetNextPageToken(v string) *ListRoleUsersResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRoleUsersResponseBody) SetUsers(v []*User) *ListRoleUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListRoleUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
