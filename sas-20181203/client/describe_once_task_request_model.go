// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeOnceTaskRequest
	GetCurrentPage() *int32
	SetEndTimeQuery(v int64) *DescribeOnceTaskRequest
	GetEndTimeQuery() *int64
	SetPageSize(v int32) *DescribeOnceTaskRequest
	GetPageSize() *int32
	SetRootTaskId(v string) *DescribeOnceTaskRequest
	GetRootTaskId() *string
	SetSource(v string) *DescribeOnceTaskRequest
	GetSource() *string
	SetStartTimeQuery(v int64) *DescribeOnceTaskRequest
	GetStartTimeQuery() *int64
	SetTaskId(v string) *DescribeOnceTaskRequest
	GetTaskId() *string
	SetTaskType(v string) *DescribeOnceTaskRequest
	GetTaskType() *string
}

type DescribeOnceTaskRequest struct {
	// The page number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The timestamp when the root task ends. Unit: milliseconds.
	//
	// example:
	//
	// 1651766520000
	EndTimeQuery *int64 `json:"EndTimeQuery,omitempty" xml:"EndTimeQuery,omitempty"`
	// The number of client tasks per page in a paged query. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The root task ID.
	//
	// > **TaskType*	- and **RootTaskId*	- cannot both be empty.
	//
	// example:
	//
	// bb5d657479bba5e1d308b6c9e85c9174
	RootTaskId *string `json:"RootTaskId,omitempty" xml:"RootTaskId,omitempty"`
	// The node source. Valid values include but are not limited to:
	//
	// - **schedule**: automatic scheduling of vulnerability scanning
	//
	// - **console**: one-click detection from the vulnerability scanning console
	//
	// example:
	//
	// console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The timestamp when the root task starts. Unit: milliseconds.
	//
	// example:
	//
	// 1651737301000
	StartTimeQuery *int64 `json:"StartTimeQuery,omitempty" xml:"StartTimeQuery,omitempty"`
	// The task ID.
	//
	// example:
	//
	// d7b2acf8d362742123e4a84e1bf8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The node type. Valid values:
	//
	// - **VUL_CHECK_TASK**: vulnerability scanning node
	//
	// - **CLIENT_PROBLEM_CHECK**: client node
	//
	// - **CLIENT_DEV_OPS**: cloud O&M node
	//
	// - **ASSET_SECURITY_CHECK**: asset information collection node
	//
	// > **TaskType*	- and **RootTaskId*	- cannot both be empty.
	//
	// example:
	//
	// CLIENT_PROBLEM_CHECK
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeOnceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOnceTaskRequest) GetEndTimeQuery() *int64 {
	return s.EndTimeQuery
}

func (s *DescribeOnceTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOnceTaskRequest) GetRootTaskId() *string {
	return s.RootTaskId
}

func (s *DescribeOnceTaskRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeOnceTaskRequest) GetStartTimeQuery() *int64 {
	return s.StartTimeQuery
}

func (s *DescribeOnceTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeOnceTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeOnceTaskRequest) SetCurrentPage(v int32) *DescribeOnceTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOnceTaskRequest) SetEndTimeQuery(v int64) *DescribeOnceTaskRequest {
	s.EndTimeQuery = &v
	return s
}

func (s *DescribeOnceTaskRequest) SetPageSize(v int32) *DescribeOnceTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOnceTaskRequest) SetRootTaskId(v string) *DescribeOnceTaskRequest {
	s.RootTaskId = &v
	return s
}

func (s *DescribeOnceTaskRequest) SetSource(v string) *DescribeOnceTaskRequest {
	s.Source = &v
	return s
}

func (s *DescribeOnceTaskRequest) SetStartTimeQuery(v int64) *DescribeOnceTaskRequest {
	s.StartTimeQuery = &v
	return s
}

func (s *DescribeOnceTaskRequest) SetTaskId(v string) *DescribeOnceTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeOnceTaskRequest) SetTaskType(v string) *DescribeOnceTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeOnceTaskRequest) Validate() error {
	return dara.Validate(s)
}
