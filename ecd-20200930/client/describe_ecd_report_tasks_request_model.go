// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEcdReportTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *DescribeEcdReportTasksRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeEcdReportTasksRequest
	GetPageSize() *int32
	SetStatus(v []*string) *DescribeEcdReportTasksRequest
	GetStatus() []*string
	SetSubType(v string) *DescribeEcdReportTasksRequest
	GetSubType() *string
	SetTaskId(v string) *DescribeEcdReportTasksRequest
	GetTaskId() *string
	SetTaskType(v string) *DescribeEcdReportTasksRequest
	GetTaskType() *string
}

type DescribeEcdReportTasksRequest struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20。
	PageSize *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status   []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// example:
	//
	// DESKTOP
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// ret-sfkdsjfi*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// RESOURCE_REPORT
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeEcdReportTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEcdReportTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeEcdReportTasksRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeEcdReportTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEcdReportTasksRequest) GetStatus() []*string {
	return s.Status
}

func (s *DescribeEcdReportTasksRequest) GetSubType() *string {
	return s.SubType
}

func (s *DescribeEcdReportTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeEcdReportTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeEcdReportTasksRequest) SetPageNum(v int32) *DescribeEcdReportTasksRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeEcdReportTasksRequest) SetPageSize(v int32) *DescribeEcdReportTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEcdReportTasksRequest) SetStatus(v []*string) *DescribeEcdReportTasksRequest {
	s.Status = v
	return s
}

func (s *DescribeEcdReportTasksRequest) SetSubType(v string) *DescribeEcdReportTasksRequest {
	s.SubType = &v
	return s
}

func (s *DescribeEcdReportTasksRequest) SetTaskId(v string) *DescribeEcdReportTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeEcdReportTasksRequest) SetTaskType(v string) *DescribeEcdReportTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeEcdReportTasksRequest) Validate() error {
	return dara.Validate(s)
}
