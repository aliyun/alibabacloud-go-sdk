// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayQuotaRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListGatewayQuotaRulesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListGatewayQuotaRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListGatewayQuotaRulesRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListGatewayQuotaRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayQuotaRulesRequest
	GetPageSize() *int32
}

type ListGatewayQuotaRulesRequest struct {
	// The rule name keyword, used for fuzzy match.
	//
	// example:
	//
	// daily
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of records to retrieve in a single request. This parameter is not supported.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. This parameter is not supported.
	//
	// example:
	//
	// C4tM8BlBJwHSNyjWpGaci4/7dKNGp1JMgsKtvCagmtY=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListGatewayQuotaRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayQuotaRulesRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayQuotaRulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGatewayQuotaRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGatewayQuotaRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGatewayQuotaRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayQuotaRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayQuotaRulesRequest) SetKeyword(v string) *ListGatewayQuotaRulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListGatewayQuotaRulesRequest) SetMaxResults(v int32) *ListGatewayQuotaRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGatewayQuotaRulesRequest) SetNextToken(v string) *ListGatewayQuotaRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListGatewayQuotaRulesRequest) SetPageNumber(v int32) *ListGatewayQuotaRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayQuotaRulesRequest) SetPageSize(v int32) *ListGatewayQuotaRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayQuotaRulesRequest) Validate() error {
	return dara.Validate(s)
}
