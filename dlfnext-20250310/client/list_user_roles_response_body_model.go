// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListUserRolesResponseBody
	GetNextPageToken() *string
	SetRoles(v []*Role) *ListUserRolesResponseBody
	GetRoles() []*Role
}

type ListUserRolesResponseBody struct {
	// example:
	//
	// token!
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Roles         []*Role `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s ListUserRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserRolesResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListUserRolesResponseBody) GetRoles() []*Role {
	return s.Roles
}

func (s *ListUserRolesResponseBody) SetNextPageToken(v string) *ListUserRolesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListUserRolesResponseBody) SetRoles(v []*Role) *ListUserRolesResponseBody {
	s.Roles = v
	return s
}

func (s *ListUserRolesResponseBody) Validate() error {
	return dara.Validate(s)
}
