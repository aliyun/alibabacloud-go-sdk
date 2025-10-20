// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListRolesResponseBody
	GetNextPageToken() *string
	SetRoles(v []*Role) *ListRolesResponseBody
	GetRoles() []*Role
}

type ListRolesResponseBody struct {
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Roles         []*Role `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListRolesResponseBody) GetRoles() []*Role {
	return s.Roles
}

func (s *ListRolesResponseBody) SetNextPageToken(v string) *ListRolesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRolesResponseBody) SetRoles(v []*Role) *ListRolesResponseBody {
	s.Roles = v
	return s
}

func (s *ListRolesResponseBody) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
