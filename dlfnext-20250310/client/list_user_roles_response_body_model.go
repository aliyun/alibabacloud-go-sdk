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
	// The pagination token used to retrieve the next page of data. If null is returned, the current page is the last page.
	//
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	// The roles.
	Roles []*Role `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
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
