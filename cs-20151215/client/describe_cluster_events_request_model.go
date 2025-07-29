// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeClusterEventsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeClusterEventsRequest
	GetPageSize() *int64
	SetTaskId(v string) *DescribeClusterEventsRequest
	GetTaskId() *string
}

type DescribeClusterEventsRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The ID of the query task.
	//
	// example:
	//
	// T-xascadasd*****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s DescribeClusterEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterEventsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeClusterEventsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeClusterEventsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeClusterEventsRequest) SetPageNumber(v int64) *DescribeClusterEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterEventsRequest) SetPageSize(v int64) *DescribeClusterEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterEventsRequest) SetTaskId(v string) *DescribeClusterEventsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterEventsRequest) Validate() error {
	return dara.Validate(s)
}
