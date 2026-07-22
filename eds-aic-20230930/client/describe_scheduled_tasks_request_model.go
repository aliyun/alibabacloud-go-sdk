// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeScheduledTasksRequest
	GetInstanceIds() []*string
	SetMaxResults(v int32) *DescribeScheduledTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeScheduledTasksRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeScheduledTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScheduledTasksRequest
	GetPageSize() *int32
	SetScheduledIds(v []*string) *DescribeScheduledTasksRequest
	GetScheduledIds() []*string
	SetStatus(v string) *DescribeScheduledTasksRequest
	GetStatus() *string
	SetTaskName(v string) *DescribeScheduledTasksRequest
	GetTaskName() *string
}

type DescribeScheduledTasksRequest struct {
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that indicates the position from which to start reading. Leave this parameter empty to read from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scheduled task IDs used to filter results.
	//
	// example:
	//
	// ["scheduled-abcd1234"]
	ScheduledIds []*string `json:"ScheduledIds,omitempty" xml:"ScheduledIds,omitempty" type:"Repeated"`
	// The status used to filter results. Valid values: ACTIVE and DISABLED.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task name.
	//
	// example:
	//
	// Data synchronization.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeScheduledTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeScheduledTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeScheduledTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeScheduledTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScheduledTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScheduledTasksRequest) GetScheduledIds() []*string {
	return s.ScheduledIds
}

func (s *DescribeScheduledTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeScheduledTasksRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeScheduledTasksRequest) SetInstanceIds(v []*string) *DescribeScheduledTasksRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetMaxResults(v int32) *DescribeScheduledTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetNextToken(v string) *DescribeScheduledTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetPageNumber(v int32) *DescribeScheduledTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetPageSize(v int32) *DescribeScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledIds(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledIds = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetStatus(v string) *DescribeScheduledTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetTaskName(v string) *DescribeScheduledTasksRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeScheduledTasksRequest) Validate() error {
	return dara.Validate(s)
}
