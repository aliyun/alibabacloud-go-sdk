// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDriveSpacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDriveSpacesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDriveSpacesShrinkRequest
	GetNextToken() *string
	SetSpaceType(v string) *ListDriveSpacesShrinkRequest
	GetSpaceType() *string
	SetTenantContextShrink(v string) *ListDriveSpacesShrinkRequest
	GetTenantContextShrink() *string
}

type ListDriveSpacesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// fekaf
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// org
	SpaceType           *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s ListDriveSpacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDriveSpacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDriveSpacesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDriveSpacesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDriveSpacesShrinkRequest) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListDriveSpacesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListDriveSpacesShrinkRequest) SetMaxResults(v int32) *ListDriveSpacesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDriveSpacesShrinkRequest) SetNextToken(v string) *ListDriveSpacesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDriveSpacesShrinkRequest) SetSpaceType(v string) *ListDriveSpacesShrinkRequest {
	s.SpaceType = &v
	return s
}

func (s *ListDriveSpacesShrinkRequest) SetTenantContextShrink(v string) *ListDriveSpacesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListDriveSpacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
