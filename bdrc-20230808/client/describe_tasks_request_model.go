// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTasksRequest
	GetNextToken() *string
	SetTaskStatus(v string) *DescribeTasksRequest
	GetTaskStatus() *string
	SetTaskType(v string) *DescribeTasksRequest
	GetTaskType() *string
}

type DescribeTasksRequest struct {
	// The maximum number of entries to return per page. The valid range is 10 to 500. If this parameter is omitted or its value is less than 10, a default value of 10 is used. Values greater than 500 are treated as 500.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. To retrieve the next page of results, set this parameter to the `NextToken` value from the response of the previous API call. For more information, see the API description above.
	//
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies the status of tasks to query. If this parameter is omitted, the API returns tasks in all states. Valid values: `CREATED`, `RUNNING`, `COMPLETE`, `FAILED`, `EXPIRED`, and `CANCELED`.
	//
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// UPDATE_RESOURCES
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTasksRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTasksRequest) SetMaxResults(v int32) *DescribeTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTasksRequest) SetNextToken(v string) *DescribeTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskStatus(v string) *DescribeTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskType(v string) *DescribeTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeTasksRequest) Validate() error {
	return dara.Validate(s)
}
