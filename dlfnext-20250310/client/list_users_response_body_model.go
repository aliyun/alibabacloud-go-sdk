// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListUsersResponseBody
	GetNextPageToken() *string
	SetUsers(v []*User) *ListUsersResponseBody
	GetUsers() []*User
}

type ListUsersResponseBody struct {
	// example:
	//
	// token!
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Users         []*User `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListUsersResponseBody) GetUsers() []*User {
	return s.Users
}

func (s *ListUsersResponseBody) SetNextPageToken(v string) *ListUsersResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*User) *ListUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
