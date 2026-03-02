// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDependLibrarysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *ListDependLibrarysRequest
	GetApplicant() *string
	SetMarketId(v int64) *ListDependLibrarysRequest
	GetMarketId() *int64
	SetOrderBy(v string) *ListDependLibrarysRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListDependLibrarysRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListDependLibrarysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDependLibrarysRequest
	GetPageSize() *int32
}

type ListDependLibrarysRequest struct {
	Applicant      *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDependLibrarysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDependLibrarysRequest) GoString() string {
	return s.String()
}

func (s *ListDependLibrarysRequest) GetApplicant() *string {
	return s.Applicant
}

func (s *ListDependLibrarysRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListDependLibrarysRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDependLibrarysRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDependLibrarysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDependLibrarysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDependLibrarysRequest) SetApplicant(v string) *ListDependLibrarysRequest {
	s.Applicant = &v
	return s
}

func (s *ListDependLibrarysRequest) SetMarketId(v int64) *ListDependLibrarysRequest {
	s.MarketId = &v
	return s
}

func (s *ListDependLibrarysRequest) SetOrderBy(v string) *ListDependLibrarysRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDependLibrarysRequest) SetOrderDirection(v string) *ListDependLibrarysRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDependLibrarysRequest) SetPageNumber(v int32) *ListDependLibrarysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDependLibrarysRequest) SetPageSize(v int32) *ListDependLibrarysRequest {
	s.PageSize = &v
	return s
}

func (s *ListDependLibrarysRequest) Validate() error {
	return dara.Validate(s)
}
