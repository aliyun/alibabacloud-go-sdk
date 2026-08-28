// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRiskCheckResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRiskCheckResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRiskCheckResultsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListRiskCheckResultsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRiskCheckResultsRequest
	GetPageSize() *int32
}

type ListRiskCheckResultsRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// token-xxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListRiskCheckResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRiskCheckResultsRequest) GoString() string {
	return s.String()
}

func (s *ListRiskCheckResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRiskCheckResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRiskCheckResultsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRiskCheckResultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRiskCheckResultsRequest) SetMaxResults(v int32) *ListRiskCheckResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRiskCheckResultsRequest) SetNextToken(v string) *ListRiskCheckResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRiskCheckResultsRequest) SetPageNumber(v int32) *ListRiskCheckResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRiskCheckResultsRequest) SetPageSize(v int32) *ListRiskCheckResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRiskCheckResultsRequest) Validate() error {
	return dara.Validate(s)
}
