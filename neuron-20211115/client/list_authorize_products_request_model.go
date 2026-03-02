// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizeProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListAuthorizeProductsRequest
	GetCompanyId() *int64
	SetOrderBy(v string) *ListAuthorizeProductsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListAuthorizeProductsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListAuthorizeProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizeProductsRequest
	GetPageSize() *int32
}

type ListAuthorizeProductsRequest struct {
	// example:
	//
	// 41
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// gmtModified
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
}

func (s ListAuthorizeProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizeProductsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizeProductsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListAuthorizeProductsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListAuthorizeProductsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListAuthorizeProductsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizeProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizeProductsRequest) SetCompanyId(v int64) *ListAuthorizeProductsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListAuthorizeProductsRequest) SetOrderBy(v string) *ListAuthorizeProductsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAuthorizeProductsRequest) SetOrderDirection(v string) *ListAuthorizeProductsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListAuthorizeProductsRequest) SetPageNumber(v int32) *ListAuthorizeProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizeProductsRequest) SetPageSize(v int32) *ListAuthorizeProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizeProductsRequest) Validate() error {
	return dara.Validate(s)
}
