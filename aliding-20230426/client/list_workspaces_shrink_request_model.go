// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspacesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspacesShrinkRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListWorkspacesShrinkRequest
	GetOrderBy() *string
	SetTeamId(v string) *ListWorkspacesShrinkRequest
	GetTeamId() *string
	SetTenantContextShrink(v string) *ListWorkspacesShrinkRequest
	GetTenantContextShrink() *string
	SetWithPermissionRole(v bool) *ListWorkspacesShrinkRequest
	GetWithPermissionRole() *bool
}

type ListWorkspacesShrinkRequest struct {
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 123123
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// VIEW_TIME_DESC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// qweqwe
	TeamId              *string `json:"TeamId,omitempty" xml:"TeamId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// true
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
}

func (s ListWorkspacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListWorkspacesShrinkRequest) GetTeamId() *string {
	return s.TeamId
}

func (s *ListWorkspacesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListWorkspacesShrinkRequest) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *ListWorkspacesShrinkRequest) SetMaxResults(v int32) *ListWorkspacesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetNextToken(v string) *ListWorkspacesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetOrderBy(v string) *ListWorkspacesShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetTeamId(v string) *ListWorkspacesShrinkRequest {
	s.TeamId = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetTenantContextShrink(v string) *ListWorkspacesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetWithPermissionRole(v bool) *ListWorkspacesShrinkRequest {
	s.WithPermissionRole = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
