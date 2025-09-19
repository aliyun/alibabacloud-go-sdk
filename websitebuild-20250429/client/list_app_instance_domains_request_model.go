// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstanceDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAppInstanceDomainsRequest
	GetBizId() *string
	SetMaxResults(v int32) *ListAppInstanceDomainsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppInstanceDomainsRequest
	GetNextToken() *string
	SetOrderColumn(v string) *ListAppInstanceDomainsRequest
	GetOrderColumn() *string
	SetOrderType(v string) *ListAppInstanceDomainsRequest
	GetOrderType() *string
	SetPageNum(v int32) *ListAppInstanceDomainsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAppInstanceDomainsRequest
	GetPageSize() *int32
}

type ListAppInstanceDomainsRequest struct {
	// example:
	//
	// WD20250718165839000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// CreationTime
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 0
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppInstanceDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAppInstanceDomainsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppInstanceDomainsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppInstanceDomainsRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListAppInstanceDomainsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListAppInstanceDomainsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAppInstanceDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstanceDomainsRequest) SetBizId(v string) *ListAppInstanceDomainsRequest {
	s.BizId = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) SetMaxResults(v int32) *ListAppInstanceDomainsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) SetNextToken(v string) *ListAppInstanceDomainsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) SetOrderColumn(v string) *ListAppInstanceDomainsRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) SetOrderType(v string) *ListAppInstanceDomainsRequest {
	s.OrderType = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) SetPageNum(v int32) *ListAppInstanceDomainsRequest {
	s.PageNum = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) SetPageSize(v int32) *ListAppInstanceDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) Validate() error {
	return dara.Validate(s)
}
