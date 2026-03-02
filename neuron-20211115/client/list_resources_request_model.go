// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListResourcesRequest
	GetCompanyId() *int64
	SetEnv(v string) *ListResourcesRequest
	GetEnv() *string
	SetName(v string) *ListResourcesRequest
	GetName() *string
	SetOrderBy(v string) *ListResourcesRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListResourcesRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourcesRequest
	GetPageSize() *int32
	SetProductId(v int64) *ListResourcesRequest
	GetProductId() *int64
	SetType(v string) *ListResourcesRequest
	GetType() *string
}

type ListResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 62
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// snowberg-staging
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// gmt_create
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// Redis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListResourcesRequest) GetEnv() *string {
	return s.Env
}

func (s *ListResourcesRequest) GetName() *string {
	return s.Name
}

func (s *ListResourcesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListResourcesRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListResourcesRequest) GetType() *string {
	return s.Type
}

func (s *ListResourcesRequest) SetCompanyId(v int64) *ListResourcesRequest {
	s.CompanyId = &v
	return s
}

func (s *ListResourcesRequest) SetEnv(v string) *ListResourcesRequest {
	s.Env = &v
	return s
}

func (s *ListResourcesRequest) SetName(v string) *ListResourcesRequest {
	s.Name = &v
	return s
}

func (s *ListResourcesRequest) SetOrderBy(v string) *ListResourcesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListResourcesRequest) SetOrderDirection(v string) *ListResourcesRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetProductId(v int64) *ListResourcesRequest {
	s.ProductId = &v
	return s
}

func (s *ListResourcesRequest) SetType(v string) *ListResourcesRequest {
	s.Type = &v
	return s
}

func (s *ListResourcesRequest) Validate() error {
	return dara.Validate(s)
}
