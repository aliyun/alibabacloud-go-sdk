// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeProjectsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeProjectsRequest
	GetNextToken() *string
	SetQuery(v string) *DescribeProjectsRequest
	GetQuery() *string
	SetSortBy(v string) *DescribeProjectsRequest
	GetSortBy() *string
	SetSortOrder(v string) *DescribeProjectsRequest
	GetSortOrder() *string
}

type DescribeProjectsRequest struct {
	// The number of entries per page. Default value: 10. Maximum value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Do not specify this parameter or set it to an empty string for the first page. For subsequent pages, pass the nextToken value from the previous response without any modification. If the nextToken value in the response is empty, the last page has been reached.
	//
	// example:
	//
	// eyJ0IjoiMjAyNi0wNy0xNlQwNzo1MzozOC4wMjFaIiwiaSI6MTAwMDQ0OH0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The keyword used for fuzzy match by project name or prompt.
	//
	// example:
	//
	// project
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The field by which to sort the results. Default value: last_scan_time.
	//
	// Valid values:
	//
	// 	- last_scan_time: the time when a task was last created.
	//
	// 	- created_at: the time when the project was created.
	//
	// 	- updated_at: the time when the project was last modified.
	//
	// example:
	//
	// last_scan_time
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The sort order. Default value: desc.
	//
	// Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
}

func (s DescribeProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProjectsRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeProjectsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeProjectsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *DescribeProjectsRequest) SetMaxResults(v int32) *DescribeProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProjectsRequest) SetNextToken(v string) *DescribeProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeProjectsRequest) SetQuery(v string) *DescribeProjectsRequest {
	s.Query = &v
	return s
}

func (s *DescribeProjectsRequest) SetSortBy(v string) *DescribeProjectsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeProjectsRequest) SetSortOrder(v string) *DescribeProjectsRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeProjectsRequest) Validate() error {
	return dara.Validate(s)
}
