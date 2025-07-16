// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNodesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodesShrinkRequest
	GetNextToken() *string
	SetParentNodeId(v string) *ListNodesShrinkRequest
	GetParentNodeId() *string
	SetTenantContextShrink(v string) *ListNodesShrinkRequest
	GetTenantContextShrink() *string
	SetWithPermissionRole(v bool) *ListNodesShrinkRequest
	GetWithPermissionRole() *bool
}

type ListNodesShrinkRequest struct {
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MNDoBb60VLBPraakI1Ywxyyn8lemrZQ3
	ParentNodeId        *string `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// false
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
}

func (s ListNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodesShrinkRequest) GetParentNodeId() *string {
	return s.ParentNodeId
}

func (s *ListNodesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListNodesShrinkRequest) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *ListNodesShrinkRequest) SetMaxResults(v int32) *ListNodesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodesShrinkRequest) SetNextToken(v string) *ListNodesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodesShrinkRequest) SetParentNodeId(v string) *ListNodesShrinkRequest {
	s.ParentNodeId = &v
	return s
}

func (s *ListNodesShrinkRequest) SetTenantContextShrink(v string) *ListNodesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetWithPermissionRole(v bool) *ListNodesShrinkRequest {
	s.WithPermissionRole = &v
	return s
}

func (s *ListNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
