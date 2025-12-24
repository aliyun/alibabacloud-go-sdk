// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListJobExecutionsRequest
	GetAppName() *string
	SetClusterId(v string) *ListJobExecutionsRequest
	GetClusterId() *string
	SetEndTime(v string) *ListJobExecutionsRequest
	GetEndTime() *string
	SetJobExecutionId(v string) *ListJobExecutionsRequest
	GetJobExecutionId() *string
	SetJobId(v int64) *ListJobExecutionsRequest
	GetJobId() *int64
	SetJobName(v string) *ListJobExecutionsRequest
	GetJobName() *string
	SetPageNum(v int32) *ListJobExecutionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListJobExecutionsRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListJobExecutionsRequest
	GetStartTime() *string
	SetStatus(v int32) *ListJobExecutionsRequest
	GetStatus() *int32
	SetWorkflowExecutionId(v int64) *ListJobExecutionsRequest
	GetWorkflowExecutionId() *int64
}

type ListJobExecutionsRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 2024-11-12 20:50:56
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2024-11-12 20:50:55
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 100
	WorkflowExecutionId *int64 `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
}

func (s ListJobExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListJobExecutionsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListJobExecutionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListJobExecutionsRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *ListJobExecutionsRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListJobExecutionsRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListJobExecutionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListJobExecutionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobExecutionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobExecutionsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListJobExecutionsRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *ListJobExecutionsRequest) SetAppName(v string) *ListJobExecutionsRequest {
	s.AppName = &v
	return s
}

func (s *ListJobExecutionsRequest) SetClusterId(v string) *ListJobExecutionsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobExecutionsRequest) SetEndTime(v string) *ListJobExecutionsRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobExecutionsRequest) SetJobExecutionId(v string) *ListJobExecutionsRequest {
	s.JobExecutionId = &v
	return s
}

func (s *ListJobExecutionsRequest) SetJobId(v int64) *ListJobExecutionsRequest {
	s.JobId = &v
	return s
}

func (s *ListJobExecutionsRequest) SetJobName(v string) *ListJobExecutionsRequest {
	s.JobName = &v
	return s
}

func (s *ListJobExecutionsRequest) SetPageNum(v int32) *ListJobExecutionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListJobExecutionsRequest) SetPageSize(v int32) *ListJobExecutionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobExecutionsRequest) SetStartTime(v string) *ListJobExecutionsRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobExecutionsRequest) SetStatus(v int32) *ListJobExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListJobExecutionsRequest) SetWorkflowExecutionId(v int64) *ListJobExecutionsRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *ListJobExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
