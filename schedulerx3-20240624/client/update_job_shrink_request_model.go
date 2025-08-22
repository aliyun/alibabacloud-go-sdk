// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateJobShrinkRequest
	GetAppName() *string
	SetAttemptInterval(v int32) *UpdateJobShrinkRequest
	GetAttemptInterval() *int32
	SetCalendar(v string) *UpdateJobShrinkRequest
	GetCalendar() *string
	SetChildJobId(v string) *UpdateJobShrinkRequest
	GetChildJobId() *string
	SetClusterId(v string) *UpdateJobShrinkRequest
	GetClusterId() *string
	SetDescription(v string) *UpdateJobShrinkRequest
	GetDescription() *string
	SetExecutorBlockStrategy(v int32) *UpdateJobShrinkRequest
	GetExecutorBlockStrategy() *int32
	SetJobHandler(v string) *UpdateJobShrinkRequest
	GetJobHandler() *string
	SetJobId(v int64) *UpdateJobShrinkRequest
	GetJobId() *int64
	SetMaxAttempt(v int32) *UpdateJobShrinkRequest
	GetMaxAttempt() *int32
	SetMaxConcurrency(v int32) *UpdateJobShrinkRequest
	GetMaxConcurrency() *int32
	SetName(v string) *UpdateJobShrinkRequest
	GetName() *string
	SetNoticeConfigShrink(v string) *UpdateJobShrinkRequest
	GetNoticeConfigShrink() *string
	SetNoticeContactsShrink(v string) *UpdateJobShrinkRequest
	GetNoticeContactsShrink() *string
	SetParameters(v string) *UpdateJobShrinkRequest
	GetParameters() *string
	SetPriority(v int32) *UpdateJobShrinkRequest
	GetPriority() *int32
	SetRouteStrategy(v int32) *UpdateJobShrinkRequest
	GetRouteStrategy() *int32
	SetScript(v string) *UpdateJobShrinkRequest
	GetScript() *string
	SetStartTime(v int64) *UpdateJobShrinkRequest
	GetStartTime() *int64
	SetTimeExpression(v string) *UpdateJobShrinkRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *UpdateJobShrinkRequest
	GetTimeType() *int32
	SetTimezone(v string) *UpdateJobShrinkRequest
	GetTimezone() *string
	SetWeight(v int32) *UpdateJobShrinkRequest
	GetWeight() *int32
}

type UpdateJobShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 3
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// example:
	//
	// workday
	Calendar   *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	ChildJobId *string `json:"ChildJobId,omitempty" xml:"ChildJobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	ExecutorBlockStrategy *int32 `json:"ExecutorBlockStrategy,omitempty" xml:"ExecutorBlockStrategy,omitempty"`
	// example:
	//
	// testJobVoidHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 3
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// test-job
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NoticeConfigShrink   *string `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty"`
	NoticeContactsShrink *string `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty"`
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32  `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	Script        *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// example:
	//
	// 1716902187
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// Asia/Beijing
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateJobShrinkRequest) GetAttemptInterval() *int32 {
	return s.AttemptInterval
}

func (s *UpdateJobShrinkRequest) GetCalendar() *string {
	return s.Calendar
}

func (s *UpdateJobShrinkRequest) GetChildJobId() *string {
	return s.ChildJobId
}

func (s *UpdateJobShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateJobShrinkRequest) GetExecutorBlockStrategy() *int32 {
	return s.ExecutorBlockStrategy
}

func (s *UpdateJobShrinkRequest) GetJobHandler() *string {
	return s.JobHandler
}

func (s *UpdateJobShrinkRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *UpdateJobShrinkRequest) GetMaxAttempt() *int32 {
	return s.MaxAttempt
}

func (s *UpdateJobShrinkRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *UpdateJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateJobShrinkRequest) GetNoticeConfigShrink() *string {
	return s.NoticeConfigShrink
}

func (s *UpdateJobShrinkRequest) GetNoticeContactsShrink() *string {
	return s.NoticeContactsShrink
}

func (s *UpdateJobShrinkRequest) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateJobShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateJobShrinkRequest) GetRouteStrategy() *int32 {
	return s.RouteStrategy
}

func (s *UpdateJobShrinkRequest) GetScript() *string {
	return s.Script
}

func (s *UpdateJobShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateJobShrinkRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *UpdateJobShrinkRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *UpdateJobShrinkRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateJobShrinkRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateJobShrinkRequest) SetAppName(v string) *UpdateJobShrinkRequest {
	s.AppName = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetAttemptInterval(v int32) *UpdateJobShrinkRequest {
	s.AttemptInterval = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetCalendar(v string) *UpdateJobShrinkRequest {
	s.Calendar = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetChildJobId(v string) *UpdateJobShrinkRequest {
	s.ChildJobId = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetClusterId(v string) *UpdateJobShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetDescription(v string) *UpdateJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetExecutorBlockStrategy(v int32) *UpdateJobShrinkRequest {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetJobHandler(v string) *UpdateJobShrinkRequest {
	s.JobHandler = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetJobId(v int64) *UpdateJobShrinkRequest {
	s.JobId = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetMaxAttempt(v int32) *UpdateJobShrinkRequest {
	s.MaxAttempt = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetMaxConcurrency(v int32) *UpdateJobShrinkRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetName(v string) *UpdateJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetNoticeConfigShrink(v string) *UpdateJobShrinkRequest {
	s.NoticeConfigShrink = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetNoticeContactsShrink(v string) *UpdateJobShrinkRequest {
	s.NoticeContactsShrink = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetParameters(v string) *UpdateJobShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetPriority(v int32) *UpdateJobShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetRouteStrategy(v int32) *UpdateJobShrinkRequest {
	s.RouteStrategy = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetScript(v string) *UpdateJobShrinkRequest {
	s.Script = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetStartTime(v int64) *UpdateJobShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetTimeExpression(v string) *UpdateJobShrinkRequest {
	s.TimeExpression = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetTimeType(v int32) *UpdateJobShrinkRequest {
	s.TimeType = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetTimezone(v string) *UpdateJobShrinkRequest {
	s.Timezone = &v
	return s
}

func (s *UpdateJobShrinkRequest) SetWeight(v int32) *UpdateJobShrinkRequest {
	s.Weight = &v
	return s
}

func (s *UpdateJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
