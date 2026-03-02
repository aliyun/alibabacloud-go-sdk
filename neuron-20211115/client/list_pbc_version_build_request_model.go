// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcVersionBuildRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ListPbcVersionBuildRequest
	GetAccountId() *string
	SetCompanyId(v int64) *ListPbcVersionBuildRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *ListPbcVersionBuildRequest
	GetMarketId() *int64
	SetOrderBy(v string) *ListPbcVersionBuildRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPbcVersionBuildRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPbcVersionBuildRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPbcVersionBuildRequest
	GetPageSize() *int32
}

type ListPbcVersionBuildRequest struct {
	AccountId      *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CompanyId      *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPbcVersionBuildRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPbcVersionBuildRequest) GoString() string {
	return s.String()
}

func (s *ListPbcVersionBuildRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ListPbcVersionBuildRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListPbcVersionBuildRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListPbcVersionBuildRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPbcVersionBuildRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPbcVersionBuildRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPbcVersionBuildRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPbcVersionBuildRequest) SetAccountId(v string) *ListPbcVersionBuildRequest {
	s.AccountId = &v
	return s
}

func (s *ListPbcVersionBuildRequest) SetCompanyId(v int64) *ListPbcVersionBuildRequest {
	s.CompanyId = &v
	return s
}

func (s *ListPbcVersionBuildRequest) SetMarketId(v int64) *ListPbcVersionBuildRequest {
	s.MarketId = &v
	return s
}

func (s *ListPbcVersionBuildRequest) SetOrderBy(v string) *ListPbcVersionBuildRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPbcVersionBuildRequest) SetOrderDirection(v string) *ListPbcVersionBuildRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPbcVersionBuildRequest) SetPageNumber(v int32) *ListPbcVersionBuildRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPbcVersionBuildRequest) SetPageSize(v int32) *ListPbcVersionBuildRequest {
	s.PageSize = &v
	return s
}

func (s *ListPbcVersionBuildRequest) Validate() error {
	return dara.Validate(s)
}
