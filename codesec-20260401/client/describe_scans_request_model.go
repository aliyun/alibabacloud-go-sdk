// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *DescribeScansRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeScansRequest
	GetNextToken() *string
	SetStatus(v string) *DescribeScansRequest
	GetStatus() *string
	SetTaskName(v string) *DescribeScansRequest
	GetTaskName() *string
}

type DescribeScansRequest struct {
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Do not specify this parameter or set it to an empty string for the first page. For subsequent pages, pass the nextToken value from the previous response without any modification. If the nextToken value in the response is empty, the last page has been reached.
	//
	// example:
	//
	// eyJ0IjoiMjAyNi0wNy0xNlQwNzo1MzozOC4wMjFaIiwiaSI6MTAwMDQ0OH0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The task status. Valid values:
	//
	// 	- running: Running.
	//
	// 	- completed: Completed.
	//
	// 	- failed: Failed.
	//
	// example:
	//
	// completed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task name.
	//
	// example:
	//
	// name
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s DescribeScansRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScansRequest) GoString() string {
	return s.String()
}

func (s *DescribeScansRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeScansRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeScansRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeScansRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeScansRequest) SetMaxResults(v int64) *DescribeScansRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeScansRequest) SetNextToken(v string) *DescribeScansRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeScansRequest) SetStatus(v string) *DescribeScansRequest {
	s.Status = &v
	return s
}

func (s *DescribeScansRequest) SetTaskName(v string) *DescribeScansRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeScansRequest) Validate() error {
	return dara.Validate(s)
}
