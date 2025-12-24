// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListScheduleEventRequest
	GetAppName() *string
	SetClusterId(v string) *ListScheduleEventRequest
	GetClusterId() *string
	SetEndTime(v int64) *ListScheduleEventRequest
	GetEndTime() *int64
	SetEvent(v string) *ListScheduleEventRequest
	GetEvent() *string
	SetEventType(v string) *ListScheduleEventRequest
	GetEventType() *string
	SetJobExecutionId(v string) *ListScheduleEventRequest
	GetJobExecutionId() *string
	SetJobName(v string) *ListScheduleEventRequest
	GetJobName() *string
	SetKeyword(v string) *ListScheduleEventRequest
	GetKeyword() *string
	SetPageNum(v int32) *ListScheduleEventRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListScheduleEventRequest
	GetPageSize() *int32
	SetReverse(v bool) *ListScheduleEventRequest
	GetReverse() *bool
	SetStartTime(v int64) *ListScheduleEventRequest
	GetStartTime() *int64
	SetWorkflowExecutionId(v int64) *ListScheduleEventRequest
	GetWorkflowExecutionId() *int64
	SetWorkflowName(v string) *ListScheduleEventRequest
	GetWorkflowName() *string
}

type ListScheduleEventRequest struct {
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
	// 1728872796295
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// INFO
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// JOB | WORKFLOW
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// hello word
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
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1581317873000
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

func (s ListScheduleEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleEventRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleEventRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListScheduleEventRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListScheduleEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListScheduleEventRequest) GetEvent() *string {
	return s.Event
}

func (s *ListScheduleEventRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListScheduleEventRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *ListScheduleEventRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListScheduleEventRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListScheduleEventRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListScheduleEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScheduleEventRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListScheduleEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListScheduleEventRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *ListScheduleEventRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListScheduleEventRequest) SetAppName(v string) *ListScheduleEventRequest {
	s.AppName = &v
	return s
}

func (s *ListScheduleEventRequest) SetClusterId(v string) *ListScheduleEventRequest {
	s.ClusterId = &v
	return s
}

func (s *ListScheduleEventRequest) SetEndTime(v int64) *ListScheduleEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListScheduleEventRequest) SetEvent(v string) *ListScheduleEventRequest {
	s.Event = &v
	return s
}

func (s *ListScheduleEventRequest) SetEventType(v string) *ListScheduleEventRequest {
	s.EventType = &v
	return s
}

func (s *ListScheduleEventRequest) SetJobExecutionId(v string) *ListScheduleEventRequest {
	s.JobExecutionId = &v
	return s
}

func (s *ListScheduleEventRequest) SetJobName(v string) *ListScheduleEventRequest {
	s.JobName = &v
	return s
}

func (s *ListScheduleEventRequest) SetKeyword(v string) *ListScheduleEventRequest {
	s.Keyword = &v
	return s
}

func (s *ListScheduleEventRequest) SetPageNum(v int32) *ListScheduleEventRequest {
	s.PageNum = &v
	return s
}

func (s *ListScheduleEventRequest) SetPageSize(v int32) *ListScheduleEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduleEventRequest) SetReverse(v bool) *ListScheduleEventRequest {
	s.Reverse = &v
	return s
}

func (s *ListScheduleEventRequest) SetStartTime(v int64) *ListScheduleEventRequest {
	s.StartTime = &v
	return s
}

func (s *ListScheduleEventRequest) SetWorkflowExecutionId(v int64) *ListScheduleEventRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *ListScheduleEventRequest) SetWorkflowName(v string) *ListScheduleEventRequest {
	s.WorkflowName = &v
	return s
}

func (s *ListScheduleEventRequest) Validate() error {
	return dara.Validate(s)
}
