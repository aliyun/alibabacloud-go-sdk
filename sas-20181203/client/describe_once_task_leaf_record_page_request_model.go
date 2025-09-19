// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnceTaskLeafRecordPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeOnceTaskLeafRecordPageRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *DescribeOnceTaskLeafRecordPageRequest
	GetEndTime() *int64
	SetPageSize(v int32) *DescribeOnceTaskLeafRecordPageRequest
	GetPageSize() *int32
	SetRelateInfo(v bool) *DescribeOnceTaskLeafRecordPageRequest
	GetRelateInfo() *bool
	SetSource(v string) *DescribeOnceTaskLeafRecordPageRequest
	GetSource() *string
	SetStartTime(v int64) *DescribeOnceTaskLeafRecordPageRequest
	GetStartTime() *int64
	SetStatusList(v []*string) *DescribeOnceTaskLeafRecordPageRequest
	GetStatusList() []*string
	SetTaskId(v string) *DescribeOnceTaskLeafRecordPageRequest
	GetTaskId() *string
	SetTaskType(v string) *DescribeOnceTaskLeafRecordPageRequest
	GetTaskType() *string
}

type DescribeOnceTaskLeafRecordPageRequest struct {
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end timestamp of the sub-task.
	//
	// example:
	//
	// 1668064495000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return on each page. Default value: 20
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether extension information is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	RelateInfo *bool `json:"RelateInfo,omitempty" xml:"RelateInfo,omitempty"`
	// The source of the request.
	//
	// example:
	//
	// console_batch
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start timestamp of the sub-task.
	//
	// example:
	//
	// 1648438617000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status information.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The ID of the sub-task.
	//
	// example:
	//
	// 1471d8ebb96795b41ede090b9758****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the sub-task. Valid values:
	//
	// 	- **IMAGE_SCAN**: image scan task
	//
	// 	- **IMAGE_REGISTRY_PULL**: image asset synchronization task
	//
	// This parameter is required.
	//
	// example:
	//
	// CLIENT_PROBLEM_CHECK
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeOnceTaskLeafRecordPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskLeafRecordPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskLeafRecordPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOnceTaskLeafRecordPageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeOnceTaskLeafRecordPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOnceTaskLeafRecordPageRequest) GetRelateInfo() *bool {
	return s.RelateInfo
}

func (s *DescribeOnceTaskLeafRecordPageRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeOnceTaskLeafRecordPageRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeOnceTaskLeafRecordPageRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *DescribeOnceTaskLeafRecordPageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeOnceTaskLeafRecordPageRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeOnceTaskLeafRecordPageRequest) SetCurrentPage(v int32) *DescribeOnceTaskLeafRecordPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageRequest) SetEndTime(v int64) *DescribeOnceTaskLeafRecordPageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageRequest) SetPageSize(v int32) *DescribeOnceTaskLeafRecordPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageRequest) SetRelateInfo(v bool) *DescribeOnceTaskLeafRecordPageRequest {
	s.RelateInfo = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageRequest) SetSource(v string) *DescribeOnceTaskLeafRecordPageRequest {
	s.Source = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageRequest) SetStartTime(v int64) *DescribeOnceTaskLeafRecordPageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageRequest) SetStatusList(v []*string) *DescribeOnceTaskLeafRecordPageRequest {
	s.StatusList = v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageRequest) SetTaskId(v string) *DescribeOnceTaskLeafRecordPageRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageRequest) SetTaskType(v string) *DescribeOnceTaskLeafRecordPageRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageRequest) Validate() error {
	return dara.Validate(s)
}
