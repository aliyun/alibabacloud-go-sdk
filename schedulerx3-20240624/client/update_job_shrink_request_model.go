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
	SetDependentStrategy(v int32) *UpdateJobShrinkRequest
	GetDependentStrategy() *int32
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
	SetStartTimeType(v string) *UpdateJobShrinkRequest
	GetStartTimeType() *string
	SetTimeExpression(v string) *UpdateJobShrinkRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *UpdateJobShrinkRequest
	GetTimeType() *int32
	SetTimezone(v string) *UpdateJobShrinkRequest
	GetTimezone() *string
	SetWeight(v int32) *UpdateJobShrinkRequest
	GetWeight() *int32
	SetXAttrs(v string) *UpdateJobShrinkRequest
	GetXAttrs() *string
}

type UpdateJobShrinkRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The interval in seconds between retry attempts.
	//
	// example:
	//
	// 3
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// The custom calendar.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The client blocking strategy.
	//
	// - 1: Serial execution
	//
	// - 2: Ignore later schedules
	//
	// - 3: Overwrite earlier schedules
	//
	// example:
	//
	// 1,2
	ChildJobId *string `json:"ChildJobId,omitempty" xml:"ChildJobId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	DependentStrategy *int32 `json:"DependentStrategy,omitempty" xml:"DependentStrategy,omitempty"`
	// The job description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Notification contact configuration
	//
	// example:
	//
	// 1
	ExecutorBlockStrategy *int32 `json:"ExecutorBlockStrategy,omitempty" xml:"ExecutorBlockStrategy,omitempty"`
	// The job handler name.
	//
	// example:
	//
	// testJobVoidHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum number of retries for a failed job.
	//
	// example:
	//
	// 3
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrent job instances.
	//
	// > This parameter defines the maximum number of instances for a single job that can run concurrently. A value of `1` prevents duplicate execution. If this limit is exceeded, the scheduler skips the current job.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The job name.
	//
	// example:
	//
	// test-job
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Time zone
	//
	// > The default is the time zone of the SchedulerX server.
	NoticeConfigShrink *string `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty"`
	// Notification configuration
	NoticeContactsShrink *string `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty"`
	// The job parameters.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The job execution priority. Valid values:
	//
	// - `1`: Low
	//
	// - `5`: Medium
	//
	// - `10`: High
	//
	// - `15`: Very High
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The routing strategy. Valid values:
	//
	// - `1`: round-robin
	//
	// - `2`: random
	//
	// - `3`: first
	//
	// - `4`: last
	//
	// - `5`: least frequently used
	//
	// - `6`: least recently used
	//
	// - `7`: consistent hashing
	//
	// - `8`: sharded broadcast
	//
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// The script content for non-BEAN jobs.
	//
	// example:
	//
	// echo "hello world"
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The type of the start time.
	//
	// example:
	//
	// 1716902187
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task execution priority. The following values are supported:
	//
	// - 1: Low
	//
	// - 5: Medium
	//
	// - 10: High
	//
	// - 15: Very High
	//
	// example:
	//
	// 1
	StartTimeType *string `json:"StartTimeType,omitempty" xml:"StartTimeType,omitempty"`
	// The time expression. The expression format depends on the `TimeType`.
	//
	// - `none`: Leave this parameter empty.
	//
	// - `cron`: Specify a standard cron expression. Online validation is supported.
	//
	// - `api`: Leave this parameter empty.
	//
	// - `fixed_rate`: An integer that represents a fixed interval in seconds. For example, `30` triggers the job every 30 seconds.
	//
	// - `one_time`: A single execution time, specified in the `yyyy-MM-dd HH:mm:ss` format or as a timestamp in milliseconds. For example, "2022-10-10 10:10:00".
	//
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// - `-1`: none
	//
	// - `1`: cron
	//
	// - `3`: fixed_rate
	//
	// - `5`: one_time
	//
	// - `100`: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The start time of the schedule.
	//
	// example:
	//
	// Hongkong
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The ID of the child job. Separate multiple IDs with a comma.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// example:
	//
	// {"reponseMode":"streaming"}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
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

func (s *UpdateJobShrinkRequest) GetDependentStrategy() *int32 {
	return s.DependentStrategy
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

func (s *UpdateJobShrinkRequest) GetStartTimeType() *string {
	return s.StartTimeType
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

func (s *UpdateJobShrinkRequest) GetXAttrs() *string {
	return s.XAttrs
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

func (s *UpdateJobShrinkRequest) SetDependentStrategy(v int32) *UpdateJobShrinkRequest {
	s.DependentStrategy = &v
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

func (s *UpdateJobShrinkRequest) SetStartTimeType(v string) *UpdateJobShrinkRequest {
	s.StartTimeType = &v
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

func (s *UpdateJobShrinkRequest) SetXAttrs(v string) *UpdateJobShrinkRequest {
	s.XAttrs = &v
	return s
}

func (s *UpdateJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
