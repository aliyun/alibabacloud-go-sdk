// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetLogEventRequest
	GetAppName() *string
	SetClusterId(v string) *GetLogEventRequest
	GetClusterId() *string
	SetEndTime(v int64) *GetLogEventRequest
	GetEndTime() *int64
	SetEvent(v string) *GetLogEventRequest
	GetEvent() *string
	SetEventType(v string) *GetLogEventRequest
	GetEventType() *string
	SetJobExecutionId(v int64) *GetLogEventRequest
	GetJobExecutionId() *int64
	SetJobName(v string) *GetLogEventRequest
	GetJobName() *string
	SetKeyword(v string) *GetLogEventRequest
	GetKeyword() *string
	SetPageNum(v int32) *GetLogEventRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetLogEventRequest
	GetPageSize() *int32
	SetReverse(v bool) *GetLogEventRequest
	GetReverse() *bool
	SetStartTime(v int64) *GetLogEventRequest
	GetStartTime() *int64
	SetWorkflowExecutionId(v int64) *GetLogEventRequest
	GetWorkflowExecutionId() *int64
	SetWorkflowName(v string) *GetLogEventRequest
	GetWorkflowName() *string
}

type GetLogEventRequest struct {
	// example:
	//
	// xxl-job-executor-perf-test-241
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-684d02ee5a6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1721636220
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// INFO
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// JOB
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 101
	JobExecutionId *int64 `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// test
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// test_partition_tbl
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
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
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1721268302000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1450568762586578000
	WorkflowExecutionId *int64 `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
	// example:
	//
	// 流程001
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s GetLogEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogEventRequest) GoString() string {
	return s.String()
}

func (s *GetLogEventRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetLogEventRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetLogEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetLogEventRequest) GetEvent() *string {
	return s.Event
}

func (s *GetLogEventRequest) GetEventType() *string {
	return s.EventType
}

func (s *GetLogEventRequest) GetJobExecutionId() *int64 {
	return s.JobExecutionId
}

func (s *GetLogEventRequest) GetJobName() *string {
	return s.JobName
}

func (s *GetLogEventRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetLogEventRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetLogEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetLogEventRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *GetLogEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetLogEventRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *GetLogEventRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *GetLogEventRequest) SetAppName(v string) *GetLogEventRequest {
	s.AppName = &v
	return s
}

func (s *GetLogEventRequest) SetClusterId(v string) *GetLogEventRequest {
	s.ClusterId = &v
	return s
}

func (s *GetLogEventRequest) SetEndTime(v int64) *GetLogEventRequest {
	s.EndTime = &v
	return s
}

func (s *GetLogEventRequest) SetEvent(v string) *GetLogEventRequest {
	s.Event = &v
	return s
}

func (s *GetLogEventRequest) SetEventType(v string) *GetLogEventRequest {
	s.EventType = &v
	return s
}

func (s *GetLogEventRequest) SetJobExecutionId(v int64) *GetLogEventRequest {
	s.JobExecutionId = &v
	return s
}

func (s *GetLogEventRequest) SetJobName(v string) *GetLogEventRequest {
	s.JobName = &v
	return s
}

func (s *GetLogEventRequest) SetKeyword(v string) *GetLogEventRequest {
	s.Keyword = &v
	return s
}

func (s *GetLogEventRequest) SetPageNum(v int32) *GetLogEventRequest {
	s.PageNum = &v
	return s
}

func (s *GetLogEventRequest) SetPageSize(v int32) *GetLogEventRequest {
	s.PageSize = &v
	return s
}

func (s *GetLogEventRequest) SetReverse(v bool) *GetLogEventRequest {
	s.Reverse = &v
	return s
}

func (s *GetLogEventRequest) SetStartTime(v int64) *GetLogEventRequest {
	s.StartTime = &v
	return s
}

func (s *GetLogEventRequest) SetWorkflowExecutionId(v int64) *GetLogEventRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *GetLogEventRequest) SetWorkflowName(v string) *GetLogEventRequest {
	s.WorkflowName = &v
	return s
}

func (s *GetLogEventRequest) Validate() error {
	return dara.Validate(s)
}
