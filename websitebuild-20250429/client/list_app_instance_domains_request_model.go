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
	SetDomainKeyword(v string) *ListAppInstanceDomainsRequest
	GetDomainKeyword() *string
	SetDomainType(v string) *ListAppInstanceDomainsRequest
	GetDomainType() *string
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
	SetStatusQueryMode(v string) *ListAppInstanceDomainsRequest
	GetStatusQueryMode() *string
}

type ListAppInstanceDomainsRequest struct {
	// The business ID.
	//
	// example:
	//
	// WD20250718165839000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The domain name keyword for fuzzy match. This parameter is optional.
	//
	// example:
	//
	// example.com
	DomainKeyword *string `json:"DomainKeyword,omitempty" xml:"DomainKeyword,omitempty"`
	// The domain name management type. Valid values: CUSTOM and PLATFORM_PREFIX. If this parameter is not specified, the default value is CUSTOM.
	//
	// example:
	//
	// CUSTOM
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The number of entries to return for a single query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query. If no next query is available, this value is empty.
	//
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field used for sorting.
	//
	// example:
	//
	// CreationTime
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The sort order. Valid values: ASC and DESC.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 0
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query mode for the domain name status. Valid values: BASIC and REALTIME. BASIC returns only the persisted status. REALTIME refreshes the status from external services. If this parameter is not specified, the default value is REALTIME.
	//
	// example:
	//
	// REALTIME
	StatusQueryMode *string `json:"StatusQueryMode,omitempty" xml:"StatusQueryMode,omitempty"`
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

func (s *ListAppInstanceDomainsRequest) GetDomainKeyword() *string {
	return s.DomainKeyword
}

func (s *ListAppInstanceDomainsRequest) GetDomainType() *string {
	return s.DomainType
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

func (s *ListAppInstanceDomainsRequest) GetStatusQueryMode() *string {
	return s.StatusQueryMode
}

func (s *ListAppInstanceDomainsRequest) SetBizId(v string) *ListAppInstanceDomainsRequest {
	s.BizId = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) SetDomainKeyword(v string) *ListAppInstanceDomainsRequest {
	s.DomainKeyword = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) SetDomainType(v string) *ListAppInstanceDomainsRequest {
	s.DomainType = &v
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

func (s *ListAppInstanceDomainsRequest) SetStatusQueryMode(v string) *ListAppInstanceDomainsRequest {
	s.StatusQueryMode = &v
	return s
}

func (s *ListAppInstanceDomainsRequest) Validate() error {
	return dara.Validate(s)
}
