// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppRiskEventByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAiAppRiskEventByPageRequest
	GetCurrentPage() *int32
	SetMaxResults(v int32) *ListAiAppRiskEventByPageRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAiAppRiskEventByPageRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListAiAppRiskEventByPageRequest
	GetPageSize() *int32
	SetQuery(v string) *ListAiAppRiskEventByPageRequest
	GetQuery() *string
	SetRegionId(v string) *ListAiAppRiskEventByPageRequest
	GetRegionId() *string
}

type ListAiAppRiskEventByPageRequest struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of results to return per request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// example:
	//
	// 1a320d468c75e987765861ec6d10f8cd3aea63fac9610c5c
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query parameter.
	//
	// example:
	//
	// {\\"Lang\\":\\"zh\\"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAiAppRiskEventByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppRiskEventByPageRequest) GoString() string {
	return s.String()
}

func (s *ListAiAppRiskEventByPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAiAppRiskEventByPageRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAiAppRiskEventByPageRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAiAppRiskEventByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiAppRiskEventByPageRequest) GetQuery() *string {
	return s.Query
}

func (s *ListAiAppRiskEventByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAiAppRiskEventByPageRequest) SetCurrentPage(v int32) *ListAiAppRiskEventByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAiAppRiskEventByPageRequest) SetMaxResults(v int32) *ListAiAppRiskEventByPageRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAiAppRiskEventByPageRequest) SetNextToken(v string) *ListAiAppRiskEventByPageRequest {
	s.NextToken = &v
	return s
}

func (s *ListAiAppRiskEventByPageRequest) SetPageSize(v int32) *ListAiAppRiskEventByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAiAppRiskEventByPageRequest) SetQuery(v string) *ListAiAppRiskEventByPageRequest {
	s.Query = &v
	return s
}

func (s *ListAiAppRiskEventByPageRequest) SetRegionId(v string) *ListAiAppRiskEventByPageRequest {
	s.RegionId = &v
	return s
}

func (s *ListAiAppRiskEventByPageRequest) Validate() error {
	return dara.Validate(s)
}
