// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListProductsRequest
	GetCompanyId() *int64
	SetKeyword(v string) *ListProductsRequest
	GetKeyword() *string
	SetOrderBy(v string) *ListProductsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListProductsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProductsRequest
	GetPageSize() *int32
}

type ListProductsRequest struct {
	// example:
	//
	// 41
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// A short description of struct
	//
	// example:
	//
	// yunmall
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// gmtCreate
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

func (s ListProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListProductsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProductsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListProductsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListProductsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductsRequest) SetCompanyId(v int64) *ListProductsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListProductsRequest) SetKeyword(v string) *ListProductsRequest {
	s.Keyword = &v
	return s
}

func (s *ListProductsRequest) SetOrderBy(v string) *ListProductsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListProductsRequest) SetOrderDirection(v string) *ListProductsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListProductsRequest) SetPageNumber(v int32) *ListProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProductsRequest) SetPageSize(v int32) *ListProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
