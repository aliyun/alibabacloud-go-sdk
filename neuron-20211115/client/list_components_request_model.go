// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListComponentsRequest
	GetCompanyId() *int64
	SetEnv(v string) *ListComponentsRequest
	GetEnv() *string
	SetName(v string) *ListComponentsRequest
	GetName() *string
	SetOrderBy(v string) *ListComponentsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListComponentsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListComponentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComponentsRequest
	GetPageSize() *int32
	SetProductId(v int64) *ListComponentsRequest
	GetProductId() *int64
	SetTemplateId(v int64) *ListComponentsRequest
	GetTemplateId() *int64
	SetType(v string) *ListComponentsRequest
	GetType() *string
}

type ListComponentsRequest struct {
	// This parameter is required.
	//
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
	// 6
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// example:
	//
	// Redis.Lock
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListComponentsRequest) GetEnv() *string {
	return s.Env
}

func (s *ListComponentsRequest) GetName() *string {
	return s.Name
}

func (s *ListComponentsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListComponentsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListComponentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComponentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComponentsRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListComponentsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListComponentsRequest) GetType() *string {
	return s.Type
}

func (s *ListComponentsRequest) SetCompanyId(v int64) *ListComponentsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListComponentsRequest) SetEnv(v string) *ListComponentsRequest {
	s.Env = &v
	return s
}

func (s *ListComponentsRequest) SetName(v string) *ListComponentsRequest {
	s.Name = &v
	return s
}

func (s *ListComponentsRequest) SetOrderBy(v string) *ListComponentsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListComponentsRequest) SetOrderDirection(v string) *ListComponentsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListComponentsRequest) SetPageNumber(v int32) *ListComponentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComponentsRequest) SetPageSize(v int32) *ListComponentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentsRequest) SetProductId(v int64) *ListComponentsRequest {
	s.ProductId = &v
	return s
}

func (s *ListComponentsRequest) SetTemplateId(v int64) *ListComponentsRequest {
	s.TemplateId = &v
	return s
}

func (s *ListComponentsRequest) SetType(v string) *ListComponentsRequest {
	s.Type = &v
	return s
}

func (s *ListComponentsRequest) Validate() error {
	return dara.Validate(s)
}
