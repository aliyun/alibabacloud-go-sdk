// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLibrarysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListLibrarysRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *ListLibrarysRequest
	GetMarketId() *int64
	SetName(v string) *ListLibrarysRequest
	GetName() *string
	SetOrderBy(v string) *ListLibrarysRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListLibrarysRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListLibrarysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLibrarysRequest
	GetPageSize() *int32
	SetProvider(v string) *ListLibrarysRequest
	GetProvider() *string
}

type ListLibrarysRequest struct {
	// This parameter is required.
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// This parameter is required.
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Provider       *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ListLibrarysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLibrarysRequest) GoString() string {
	return s.String()
}

func (s *ListLibrarysRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListLibrarysRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListLibrarysRequest) GetName() *string {
	return s.Name
}

func (s *ListLibrarysRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListLibrarysRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListLibrarysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLibrarysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLibrarysRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListLibrarysRequest) SetCompanyId(v int64) *ListLibrarysRequest {
	s.CompanyId = &v
	return s
}

func (s *ListLibrarysRequest) SetMarketId(v int64) *ListLibrarysRequest {
	s.MarketId = &v
	return s
}

func (s *ListLibrarysRequest) SetName(v string) *ListLibrarysRequest {
	s.Name = &v
	return s
}

func (s *ListLibrarysRequest) SetOrderBy(v string) *ListLibrarysRequest {
	s.OrderBy = &v
	return s
}

func (s *ListLibrarysRequest) SetOrderDirection(v string) *ListLibrarysRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListLibrarysRequest) SetPageNumber(v int32) *ListLibrarysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLibrarysRequest) SetPageSize(v int32) *ListLibrarysRequest {
	s.PageSize = &v
	return s
}

func (s *ListLibrarysRequest) SetProvider(v string) *ListLibrarysRequest {
	s.Provider = &v
	return s
}

func (s *ListLibrarysRequest) Validate() error {
	return dara.Validate(s)
}
