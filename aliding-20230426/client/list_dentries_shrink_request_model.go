// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDentriesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDentriesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDentriesShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *ListDentriesShrinkRequest
	GetOrder() *string
	SetOrderBy(v string) *ListDentriesShrinkRequest
	GetOrderBy() *string
	SetParentId(v string) *ListDentriesShrinkRequest
	GetParentId() *string
	SetSpaceId(v string) *ListDentriesShrinkRequest
	GetSpaceId() *string
	SetTenantContextShrink(v string) *ListDentriesShrinkRequest
	GetTenantContextShrink() *string
	SetWithThumbnail(v bool) *ListDentriesShrinkRequest
	GetWithThumbnail() *bool
}

type ListDentriesShrinkRequest struct {
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// MODIFIED_TIME
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 854xxxxx
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// true
	WithThumbnail *bool `json:"WithThumbnail,omitempty" xml:"WithThumbnail,omitempty"`
}

func (s ListDentriesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDentriesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDentriesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDentriesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDentriesShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDentriesShrinkRequest) GetParentId() *string {
	return s.ParentId
}

func (s *ListDentriesShrinkRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *ListDentriesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListDentriesShrinkRequest) GetWithThumbnail() *bool {
	return s.WithThumbnail
}

func (s *ListDentriesShrinkRequest) SetMaxResults(v int32) *ListDentriesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDentriesShrinkRequest) SetNextToken(v string) *ListDentriesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDentriesShrinkRequest) SetOrder(v string) *ListDentriesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListDentriesShrinkRequest) SetOrderBy(v string) *ListDentriesShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDentriesShrinkRequest) SetParentId(v string) *ListDentriesShrinkRequest {
	s.ParentId = &v
	return s
}

func (s *ListDentriesShrinkRequest) SetSpaceId(v string) *ListDentriesShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *ListDentriesShrinkRequest) SetTenantContextShrink(v string) *ListDentriesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListDentriesShrinkRequest) SetWithThumbnail(v bool) *ListDentriesShrinkRequest {
	s.WithThumbnail = &v
	return s
}

func (s *ListDentriesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
