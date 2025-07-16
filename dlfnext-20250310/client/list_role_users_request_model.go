// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoleUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRoleUsersRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListRoleUsersRequest
	GetPageToken() *string
	SetRolePrincipal(v string) *ListRoleUsersRequest
	GetRolePrincipal() *string
}

type ListRoleUsersRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
}

func (s ListRoleUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRoleUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRoleUsersRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListRoleUsersRequest) GetRolePrincipal() *string {
	return s.RolePrincipal
}

func (s *ListRoleUsersRequest) SetMaxResults(v int32) *ListRoleUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRoleUsersRequest) SetPageToken(v string) *ListRoleUsersRequest {
	s.PageToken = &v
	return s
}

func (s *ListRoleUsersRequest) SetRolePrincipal(v string) *ListRoleUsersRequest {
	s.RolePrincipal = &v
	return s
}

func (s *ListRoleUsersRequest) Validate() error {
	return dara.Validate(s)
}
