// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeClusterTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeClusterTasksRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeClusterTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterTasksRequest
	GetPageSize() *int32
}

type DescribeClusterTasksRequest struct {
	// The maximum number of results to return per request.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"max_results,omitempty" xml:"max_results,omitempty"`
	// The pagination token. Set the value to the next_token parameter value returned from the previous API call.
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of items per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s DescribeClusterTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeClusterTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeClusterTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterTasksRequest) SetMaxResults(v int32) *DescribeClusterTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeClusterTasksRequest) SetNextToken(v string) *DescribeClusterTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeClusterTasksRequest) SetPageNumber(v int32) *DescribeClusterTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterTasksRequest) SetPageSize(v int32) *DescribeClusterTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterTasksRequest) Validate() error {
	return dara.Validate(s)
}
