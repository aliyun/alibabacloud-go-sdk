// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspacesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspacesRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListWorkspacesRequest
	GetOrderBy() *string
	SetTeamId(v string) *ListWorkspacesRequest
	GetTeamId() *string
	SetTenantContext(v *ListWorkspacesRequestTenantContext) *ListWorkspacesRequest
	GetTenantContext() *ListWorkspacesRequestTenantContext
	SetWithPermissionRole(v bool) *ListWorkspacesRequest
	GetWithPermissionRole() *bool
}

type ListWorkspacesRequest struct {
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
	TeamId        *string                             `json:"TeamId,omitempty" xml:"TeamId,omitempty"`
	TenantContext *ListWorkspacesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// true
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListWorkspacesRequest) GetTeamId() *string {
	return s.TeamId
}

func (s *ListWorkspacesRequest) GetTenantContext() *ListWorkspacesRequestTenantContext {
	return s.TenantContext
}

func (s *ListWorkspacesRequest) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetOrderBy(v string) *ListWorkspacesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListWorkspacesRequest) SetTeamId(v string) *ListWorkspacesRequest {
	s.TeamId = &v
	return s
}

func (s *ListWorkspacesRequest) SetTenantContext(v *ListWorkspacesRequestTenantContext) *ListWorkspacesRequest {
	s.TenantContext = v
	return s
}

func (s *ListWorkspacesRequest) SetWithPermissionRole(v bool) *ListWorkspacesRequest {
	s.WithPermissionRole = &v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkspacesRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListWorkspacesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkspacesRequestTenantContext) SetTenantId(v string) *ListWorkspacesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListWorkspacesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
