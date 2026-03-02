// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpLanesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListPdpLanesRequest
	GetCompanyId() *int64
	SetEnv(v string) *ListPdpLanesRequest
	GetEnv() *string
	SetInletServiceId(v int64) *ListPdpLanesRequest
	GetInletServiceId() *int64
	SetKeyword(v string) *ListPdpLanesRequest
	GetKeyword() *string
	SetOrderBy(v string) *ListPdpLanesRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpLanesRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpLanesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpLanesRequest
	GetPageSize() *int32
	SetProductId(v int64) *ListPdpLanesRequest
	GetProductId() *int64
}

type ListPdpLanesRequest struct {
	// example:
	//
	// 4
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 1
	InletServiceId *int64 `json:"inletServiceId,omitempty" xml:"inletServiceId,omitempty"`
	// example:
	//
	// yunmall
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
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
	// example:
	//
	// 1193
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
}

func (s ListPdpLanesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpLanesRequest) GoString() string {
	return s.String()
}

func (s *ListPdpLanesRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListPdpLanesRequest) GetEnv() *string {
	return s.Env
}

func (s *ListPdpLanesRequest) GetInletServiceId() *int64 {
	return s.InletServiceId
}

func (s *ListPdpLanesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListPdpLanesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpLanesRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpLanesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpLanesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpLanesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListPdpLanesRequest) SetCompanyId(v int64) *ListPdpLanesRequest {
	s.CompanyId = &v
	return s
}

func (s *ListPdpLanesRequest) SetEnv(v string) *ListPdpLanesRequest {
	s.Env = &v
	return s
}

func (s *ListPdpLanesRequest) SetInletServiceId(v int64) *ListPdpLanesRequest {
	s.InletServiceId = &v
	return s
}

func (s *ListPdpLanesRequest) SetKeyword(v string) *ListPdpLanesRequest {
	s.Keyword = &v
	return s
}

func (s *ListPdpLanesRequest) SetOrderBy(v string) *ListPdpLanesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpLanesRequest) SetOrderDirection(v string) *ListPdpLanesRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpLanesRequest) SetPageNumber(v int32) *ListPdpLanesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpLanesRequest) SetPageSize(v int32) *ListPdpLanesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpLanesRequest) SetProductId(v int64) *ListPdpLanesRequest {
	s.ProductId = &v
	return s
}

func (s *ListPdpLanesRequest) Validate() error {
	return dara.Validate(s)
}
