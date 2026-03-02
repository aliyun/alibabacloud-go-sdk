// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListPbcsRequest
	GetCompanyId() *int64
	SetDeveloperId(v string) *ListPbcsRequest
	GetDeveloperId() *string
	SetMarketId(v int64) *ListPbcsRequest
	GetMarketId() *int64
	SetName(v string) *ListPbcsRequest
	GetName() *string
	SetOrderBy(v string) *ListPbcsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPbcsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPbcsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPbcsRequest
	GetPageSize() *int32
}

type ListPbcsRequest struct {
	CompanyId   *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	DeveloperId *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
	// example:
	//
	// 1
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPbcsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPbcsRequest) GoString() string {
	return s.String()
}

func (s *ListPbcsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListPbcsRequest) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *ListPbcsRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListPbcsRequest) GetName() *string {
	return s.Name
}

func (s *ListPbcsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPbcsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPbcsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPbcsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPbcsRequest) SetCompanyId(v int64) *ListPbcsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListPbcsRequest) SetDeveloperId(v string) *ListPbcsRequest {
	s.DeveloperId = &v
	return s
}

func (s *ListPbcsRequest) SetMarketId(v int64) *ListPbcsRequest {
	s.MarketId = &v
	return s
}

func (s *ListPbcsRequest) SetName(v string) *ListPbcsRequest {
	s.Name = &v
	return s
}

func (s *ListPbcsRequest) SetOrderBy(v string) *ListPbcsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPbcsRequest) SetOrderDirection(v string) *ListPbcsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPbcsRequest) SetPageNumber(v int32) *ListPbcsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPbcsRequest) SetPageSize(v int32) *ListPbcsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPbcsRequest) Validate() error {
	return dara.Validate(s)
}
