// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRolesRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListRolesRequest
	GetPageToken() *string
	SetRoleName(v string) *ListRolesRequest
	GetRoleName() *string
}

type ListRolesRequest struct {
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
	// role_name
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s ListRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRolesRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListRolesRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *ListRolesRequest) SetMaxResults(v int32) *ListRolesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRolesRequest) SetPageToken(v string) *ListRolesRequest {
	s.PageToken = &v
	return s
}

func (s *ListRolesRequest) SetRoleName(v string) *ListRolesRequest {
	s.RoleName = &v
	return s
}

func (s *ListRolesRequest) Validate() error {
	return dara.Validate(s)
}
