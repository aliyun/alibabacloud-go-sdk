// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUserRolesRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListUserRolesRequest
	GetPageToken() *string
	SetUserPrincipal(v string) *ListUserRolesRequest
	GetUserPrincipal() *string
}

type ListUserRolesRequest struct {
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
	// acs:ram::[accountId]:user/user_name
	UserPrincipal *string `json:"userPrincipal,omitempty" xml:"userPrincipal,omitempty"`
}

func (s ListUserRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserRolesRequest) GoString() string {
	return s.String()
}

func (s *ListUserRolesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserRolesRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListUserRolesRequest) GetUserPrincipal() *string {
	return s.UserPrincipal
}

func (s *ListUserRolesRequest) SetMaxResults(v int32) *ListUserRolesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserRolesRequest) SetPageToken(v string) *ListUserRolesRequest {
	s.PageToken = &v
	return s
}

func (s *ListUserRolesRequest) SetUserPrincipal(v string) *ListUserRolesRequest {
	s.UserPrincipal = &v
	return s
}

func (s *ListUserRolesRequest) Validate() error {
	return dara.Validate(s)
}
