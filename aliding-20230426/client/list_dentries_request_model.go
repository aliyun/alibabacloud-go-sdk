// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDentriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDentriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDentriesRequest
	GetNextToken() *string
	SetOrder(v string) *ListDentriesRequest
	GetOrder() *string
	SetOrderBy(v string) *ListDentriesRequest
	GetOrderBy() *string
	SetParentId(v string) *ListDentriesRequest
	GetParentId() *string
	SetSpaceId(v string) *ListDentriesRequest
	GetSpaceId() *string
	SetTenantContext(v *ListDentriesRequestTenantContext) *ListDentriesRequest
	GetTenantContext() *ListDentriesRequestTenantContext
	SetWithThumbnail(v bool) *ListDentriesRequest
	GetWithThumbnail() *bool
}

type ListDentriesRequest struct {
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
	SpaceId       *string                           `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContext *ListDentriesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// true
	WithThumbnail *bool `json:"WithThumbnail,omitempty" xml:"WithThumbnail,omitempty"`
}

func (s ListDentriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesRequest) GoString() string {
	return s.String()
}

func (s *ListDentriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDentriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDentriesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDentriesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDentriesRequest) GetParentId() *string {
	return s.ParentId
}

func (s *ListDentriesRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *ListDentriesRequest) GetTenantContext() *ListDentriesRequestTenantContext {
	return s.TenantContext
}

func (s *ListDentriesRequest) GetWithThumbnail() *bool {
	return s.WithThumbnail
}

func (s *ListDentriesRequest) SetMaxResults(v int32) *ListDentriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDentriesRequest) SetNextToken(v string) *ListDentriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDentriesRequest) SetOrder(v string) *ListDentriesRequest {
	s.Order = &v
	return s
}

func (s *ListDentriesRequest) SetOrderBy(v string) *ListDentriesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDentriesRequest) SetParentId(v string) *ListDentriesRequest {
	s.ParentId = &v
	return s
}

func (s *ListDentriesRequest) SetSpaceId(v string) *ListDentriesRequest {
	s.SpaceId = &v
	return s
}

func (s *ListDentriesRequest) SetTenantContext(v *ListDentriesRequestTenantContext) *ListDentriesRequest {
	s.TenantContext = v
	return s
}

func (s *ListDentriesRequest) SetWithThumbnail(v bool) *ListDentriesRequest {
	s.WithThumbnail = &v
	return s
}

func (s *ListDentriesRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDentriesRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListDentriesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListDentriesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListDentriesRequestTenantContext) SetTenantId(v string) *ListDentriesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListDentriesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
