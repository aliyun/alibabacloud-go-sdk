// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateJobShrinkRequest
	GetAppName() *string
	SetAttemptInterval(v int32) *CreateJobShrinkRequest
	GetAttemptInterval() *int32
	SetCalendar(v string) *CreateJobShrinkRequest
	GetCalendar() *string
	SetChildJobId(v string) *CreateJobShrinkRequest
	GetChildJobId() *string
	SetClusterId(v string) *CreateJobShrinkRequest
	GetClusterId() *string
	SetCoordinateShrink(v string) *CreateJobShrinkRequest
	GetCoordinateShrink() *string
	SetDependentStrategy(v int32) *CreateJobShrinkRequest
	GetDependentStrategy() *int32
	SetDescription(v string) *CreateJobShrinkRequest
	GetDescription() *string
	SetExecutorBlockStrategy(v int32) *CreateJobShrinkRequest
	GetExecutorBlockStrategy() *int32
	SetJobHandler(v string) *CreateJobShrinkRequest
	GetJobHandler() *string
	SetJobType(v string) *CreateJobShrinkRequest
	GetJobType() *string
	SetMaxAttempt(v int32) *CreateJobShrinkRequest
	GetMaxAttempt() *int32
	SetMaxConcurrency(v int32) *CreateJobShrinkRequest
	GetMaxConcurrency() *int32
	SetName(v string) *CreateJobShrinkRequest
	GetName() *string
	SetNoticeConfigShrink(v string) *CreateJobShrinkRequest
	GetNoticeConfigShrink() *string
	SetNoticeContactsShrink(v string) *CreateJobShrinkRequest
	GetNoticeContactsShrink() *string
	SetParameters(v string) *CreateJobShrinkRequest
	GetParameters() *string
	SetPriority(v int32) *CreateJobShrinkRequest
	GetPriority() *int32
	SetRouteStrategy(v int32) *CreateJobShrinkRequest
	GetRouteStrategy() *int32
	SetScript(v string) *CreateJobShrinkRequest
	GetScript() *string
	SetStartTime(v int64) *CreateJobShrinkRequest
	GetStartTime() *int64
	SetStartTimeType(v int32) *CreateJobShrinkRequest
	GetStartTimeType() *int32
	SetStatus(v int32) *CreateJobShrinkRequest
	GetStatus() *int32
	SetTimeExpression(v string) *CreateJobShrinkRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *CreateJobShrinkRequest
	GetTimeType() *int32
	SetTimezone(v string) *CreateJobShrinkRequest
	GetTimezone() *string
	SetWeight(v int32) *CreateJobShrinkRequest
	GetWeight() *int32
}

type CreateJobShrinkRequest struct {
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
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// example:
	//
	// 1,2
	ChildJobId *string `json:"ChildJobId,omitempty" xml:"ChildJobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CoordinateShrink *string `json:"Coordinate,omitempty" xml:"Coordinate,omitempty"`
	// example:
	//
	// 1
	DependentStrategy *int32 `json:"DependentStrategy,omitempty" xml:"DependentStrategy,omitempty"`
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
	// xxljob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// 3
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// This parameter is required.
	//
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
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// example:
	//
	// echo "hello world"
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// example:
	//
	// 1701310327000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	StartTimeType *int32 `json:"StartTimeType,omitempty" xml:"StartTimeType,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateJobShrinkRequest) GetAttemptInterval() *int32 {
	return s.AttemptInterval
}

func (s *CreateJobShrinkRequest) GetCalendar() *string {
	return s.Calendar
}

func (s *CreateJobShrinkRequest) GetChildJobId() *string {
	return s.ChildJobId
}

func (s *CreateJobShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateJobShrinkRequest) GetCoordinateShrink() *string {
	return s.CoordinateShrink
}

func (s *CreateJobShrinkRequest) GetDependentStrategy() *int32 {
	return s.DependentStrategy
}

func (s *CreateJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateJobShrinkRequest) GetExecutorBlockStrategy() *int32 {
	return s.ExecutorBlockStrategy
}

func (s *CreateJobShrinkRequest) GetJobHandler() *string {
	return s.JobHandler
}

func (s *CreateJobShrinkRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateJobShrinkRequest) GetMaxAttempt() *int32 {
	return s.MaxAttempt
}

func (s *CreateJobShrinkRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *CreateJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateJobShrinkRequest) GetNoticeConfigShrink() *string {
	return s.NoticeConfigShrink
}

func (s *CreateJobShrinkRequest) GetNoticeContactsShrink() *string {
	return s.NoticeContactsShrink
}

func (s *CreateJobShrinkRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateJobShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateJobShrinkRequest) GetRouteStrategy() *int32 {
	return s.RouteStrategy
}

func (s *CreateJobShrinkRequest) GetScript() *string {
	return s.Script
}

func (s *CreateJobShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateJobShrinkRequest) GetStartTimeType() *int32 {
	return s.StartTimeType
}

func (s *CreateJobShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateJobShrinkRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *CreateJobShrinkRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *CreateJobShrinkRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateJobShrinkRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateJobShrinkRequest) SetAppName(v string) *CreateJobShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateJobShrinkRequest) SetAttemptInterval(v int32) *CreateJobShrinkRequest {
	s.AttemptInterval = &v
	return s
}

func (s *CreateJobShrinkRequest) SetCalendar(v string) *CreateJobShrinkRequest {
	s.Calendar = &v
	return s
}

func (s *CreateJobShrinkRequest) SetChildJobId(v string) *CreateJobShrinkRequest {
	s.ChildJobId = &v
	return s
}

func (s *CreateJobShrinkRequest) SetClusterId(v string) *CreateJobShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobShrinkRequest) SetCoordinateShrink(v string) *CreateJobShrinkRequest {
	s.CoordinateShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetDependentStrategy(v int32) *CreateJobShrinkRequest {
	s.DependentStrategy = &v
	return s
}

func (s *CreateJobShrinkRequest) SetDescription(v string) *CreateJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateJobShrinkRequest) SetExecutorBlockStrategy(v int32) *CreateJobShrinkRequest {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobHandler(v string) *CreateJobShrinkRequest {
	s.JobHandler = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobType(v string) *CreateJobShrinkRequest {
	s.JobType = &v
	return s
}

func (s *CreateJobShrinkRequest) SetMaxAttempt(v int32) *CreateJobShrinkRequest {
	s.MaxAttempt = &v
	return s
}

func (s *CreateJobShrinkRequest) SetMaxConcurrency(v int32) *CreateJobShrinkRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *CreateJobShrinkRequest) SetName(v string) *CreateJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateJobShrinkRequest) SetNoticeConfigShrink(v string) *CreateJobShrinkRequest {
	s.NoticeConfigShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetNoticeContactsShrink(v string) *CreateJobShrinkRequest {
	s.NoticeContactsShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetParameters(v string) *CreateJobShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *CreateJobShrinkRequest) SetPriority(v int32) *CreateJobShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobShrinkRequest) SetRouteStrategy(v int32) *CreateJobShrinkRequest {
	s.RouteStrategy = &v
	return s
}

func (s *CreateJobShrinkRequest) SetScript(v string) *CreateJobShrinkRequest {
	s.Script = &v
	return s
}

func (s *CreateJobShrinkRequest) SetStartTime(v int64) *CreateJobShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateJobShrinkRequest) SetStartTimeType(v int32) *CreateJobShrinkRequest {
	s.StartTimeType = &v
	return s
}

func (s *CreateJobShrinkRequest) SetStatus(v int32) *CreateJobShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateJobShrinkRequest) SetTimeExpression(v string) *CreateJobShrinkRequest {
	s.TimeExpression = &v
	return s
}

func (s *CreateJobShrinkRequest) SetTimeType(v int32) *CreateJobShrinkRequest {
	s.TimeType = &v
	return s
}

func (s *CreateJobShrinkRequest) SetTimezone(v string) *CreateJobShrinkRequest {
	s.Timezone = &v
	return s
}

func (s *CreateJobShrinkRequest) SetWeight(v int32) *CreateJobShrinkRequest {
	s.Weight = &v
	return s
}

func (s *CreateJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
