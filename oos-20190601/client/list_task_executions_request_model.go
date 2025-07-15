// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDateAfter(v string) *ListTaskExecutionsRequest
	GetEndDateAfter() *string
	SetEndDateBefore(v string) *ListTaskExecutionsRequest
	GetEndDateBefore() *string
	SetExecutionId(v string) *ListTaskExecutionsRequest
	GetExecutionId() *string
	SetIncludeChildTaskExecution(v bool) *ListTaskExecutionsRequest
	GetIncludeChildTaskExecution() *bool
	SetMaxResults(v int32) *ListTaskExecutionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTaskExecutionsRequest
	GetNextToken() *string
	SetParentTaskExecutionId(v string) *ListTaskExecutionsRequest
	GetParentTaskExecutionId() *string
	SetRegionId(v string) *ListTaskExecutionsRequest
	GetRegionId() *string
	SetSortField(v string) *ListTaskExecutionsRequest
	GetSortField() *string
	SetSortOrder(v string) *ListTaskExecutionsRequest
	GetSortOrder() *string
	SetStartDateAfter(v string) *ListTaskExecutionsRequest
	GetStartDateAfter() *string
	SetStartDateBefore(v string) *ListTaskExecutionsRequest
	GetStartDateBefore() *string
	SetStatus(v string) *ListTaskExecutionsRequest
	GetStatus() *string
	SetTaskAction(v string) *ListTaskExecutionsRequest
	GetTaskAction() *string
	SetTaskExecutionId(v string) *ListTaskExecutionsRequest
	GetTaskExecutionId() *string
	SetTaskName(v string) *ListTaskExecutionsRequest
	GetTaskName() *string
}

type ListTaskExecutionsRequest struct {
	// The execution ID of the task.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateAfter *string `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	// Specifies to query task executions that stop running at or later than the specified time.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateBefore *string `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	// The status of the execution. Valid values: Running, Started, Success, Failed, Waiting, Cancelled, Pending, and Skipped.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The number of entries to return on each page. Valid values: 20 to 100. Default value: 50.
	//
	// example:
	//
	// false
	IncludeChildTaskExecution *bool `json:"IncludeChildTaskExecution,omitempty" xml:"IncludeChildTaskExecution,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Sorts the task executions to query. Valid values:
	//
	// 	- **StartDate**: specifies that the task executions are sorted based on the time when they are created. This is the default value.
	//
	// 	- **EndDate**: specifies that the task executions are sorted based on the time when the time when they stop running.
	//
	// 	- **Status**: specifies that the task executions are sorted based on their statuses.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to show the child nodes in the loop task. Default value: False.
	//
	// example:
	//
	// task-exec-xxx
	ParentTaskExecutionId *string `json:"ParentTaskExecutionId,omitempty" xml:"ParentTaskExecutionId,omitempty"`
	// The ID of the execution.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The order in which you want to sort the task executions to query. Valid values:
	//
	// 	- **Ascending**: ascending order.
	//
	// 	- **Descending**: descending order. This is the default value.
	//
	// example:
	//
	// StartDate
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// Specifies to query task executions that stop running at or before the specified time.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateAfter *string `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	// Specifies to query task executions that start to run at or later than the specified time.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateBefore *string `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	// Specifies to query task executions that start to run at or before the specified time.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The execution ID of the parent node. In a loop task, set this parameter to the execution ID of the parent node.
	//
	// example:
	//
	// ACS::Sleep
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// task-exec-xxx
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The action of the task.
	//
	// example:
	//
	// describeInstance
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListTaskExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsRequest) GetEndDateAfter() *string {
	return s.EndDateAfter
}

func (s *ListTaskExecutionsRequest) GetEndDateBefore() *string {
	return s.EndDateBefore
}

func (s *ListTaskExecutionsRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *ListTaskExecutionsRequest) GetIncludeChildTaskExecution() *bool {
	return s.IncludeChildTaskExecution
}

func (s *ListTaskExecutionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTaskExecutionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTaskExecutionsRequest) GetParentTaskExecutionId() *string {
	return s.ParentTaskExecutionId
}

func (s *ListTaskExecutionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTaskExecutionsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListTaskExecutionsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTaskExecutionsRequest) GetStartDateAfter() *string {
	return s.StartDateAfter
}

func (s *ListTaskExecutionsRequest) GetStartDateBefore() *string {
	return s.StartDateBefore
}

func (s *ListTaskExecutionsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTaskExecutionsRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *ListTaskExecutionsRequest) GetTaskExecutionId() *string {
	return s.TaskExecutionId
}

func (s *ListTaskExecutionsRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTaskExecutionsRequest) SetEndDateAfter(v string) *ListTaskExecutionsRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetEndDateBefore(v string) *ListTaskExecutionsRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetExecutionId(v string) *ListTaskExecutionsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetIncludeChildTaskExecution(v bool) *ListTaskExecutionsRequest {
	s.IncludeChildTaskExecution = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetMaxResults(v int32) *ListTaskExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetNextToken(v string) *ListTaskExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetParentTaskExecutionId(v string) *ListTaskExecutionsRequest {
	s.ParentTaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetRegionId(v string) *ListTaskExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetSortField(v string) *ListTaskExecutionsRequest {
	s.SortField = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetSortOrder(v string) *ListTaskExecutionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetStartDateAfter(v string) *ListTaskExecutionsRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetStartDateBefore(v string) *ListTaskExecutionsRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetStatus(v string) *ListTaskExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetTaskAction(v string) *ListTaskExecutionsRequest {
	s.TaskAction = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetTaskExecutionId(v string) *ListTaskExecutionsRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetTaskName(v string) *ListTaskExecutionsRequest {
	s.TaskName = &v
	return s
}

func (s *ListTaskExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
