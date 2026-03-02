// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListComponentTemplatesRequest
	GetCompanyId() *int64
	SetName(v string) *ListComponentTemplatesRequest
	GetName() *string
	SetOrderBy(v string) *ListComponentTemplatesRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListComponentTemplatesRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListComponentTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComponentTemplatesRequest
	GetPageSize() *int32
	SetProductId(v int64) *ListComponentTemplatesRequest
	GetProductId() *int64
	SetType(v string) *ListComponentTemplatesRequest
	GetType() *string
}

type ListComponentTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// order
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
	// 12
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// Redis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListComponentTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComponentTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListComponentTemplatesRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListComponentTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListComponentTemplatesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListComponentTemplatesRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListComponentTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComponentTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComponentTemplatesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListComponentTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListComponentTemplatesRequest) SetCompanyId(v int64) *ListComponentTemplatesRequest {
	s.CompanyId = &v
	return s
}

func (s *ListComponentTemplatesRequest) SetName(v string) *ListComponentTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListComponentTemplatesRequest) SetOrderBy(v string) *ListComponentTemplatesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListComponentTemplatesRequest) SetOrderDirection(v string) *ListComponentTemplatesRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListComponentTemplatesRequest) SetPageNumber(v int32) *ListComponentTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComponentTemplatesRequest) SetPageSize(v int32) *ListComponentTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentTemplatesRequest) SetProductId(v int64) *ListComponentTemplatesRequest {
	s.ProductId = &v
	return s
}

func (s *ListComponentTemplatesRequest) SetType(v string) *ListComponentTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListComponentTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
