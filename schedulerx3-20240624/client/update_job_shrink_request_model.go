// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroupId(v int64) *UpdateJobShrinkRequest
	GetAppGroupId() *int64
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
	SetLabel(v string) *UpdateJobShrinkRequest
	GetLabel() *string
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
	// The application ID.
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The retry interval upon node failure.
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
	// The child node IDs, separated by commas.
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
	// The dependency strategy.
	//
	// example:
	//
	// 1
	DependentStrategy *int32 `json:"DependentStrategy,omitempty" xml:"DependentStrategy,omitempty"`
	// The node description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The client blocking strategy. Valid values:
	//
	// - 1: serial execution on a single machine
	//
	// - 2: ignore subsequent scheduling
	//
	// - 3: override previous scheduling
	//
	// example:
	//
	// 1
	ExecutorBlockStrategy *int32 `json:"ExecutorBlockStrategy,omitempty" xml:"ExecutorBlockStrategy,omitempty"`
	// The jobhandler name.
	//
	// example:
	//
	// testJobVoidHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The node label information.
	//
	// example:
	//
	// {key:value}
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum number of retry attempts upon node failure.
	//
	// example:
	//
	// 3
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum concurrency of the node.
	//
	// >The maximum number of instances that can run simultaneously for the same node. A value of 1 indicates that repeated execution is not allowed. If the concurrency limit is exceeded, the current scheduling is skipped.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The node name.
	//
	// example:
	//
	// test-job
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The notification configuration.
	NoticeConfigShrink *string `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty"`
	// The notification contact configuration.
	NoticeContactsShrink *string `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty"`
	// The node parameters.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The execution priority of the node. Valid values:
	//
	// - 1: low
	//
	// - 5: medium
	//
	// - 10: high
	//
	// - 15: very high
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The routing strategy. Valid values:
	//
	// - 1: round robin
	//
	// - 2: random
	//
	// - 3: first
	//
	// - 4: last
	//
	// - 5: least frequently used
	//
	// - 6: least recently used
	//
	// - 7: consistent hashing
	//
	// - 8: shard broadcast
	//
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// The script configured for non-BEAN nodes.
	//
	// example:
	//
	// echo "hello world"
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The scheduling start time.
	//
	// example:
	//
	// 1716902187
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The start time type.
	//
	// example:
	//
	// 1
	StartTimeType *string `json:"StartTimeType,omitempty" xml:"StartTimeType,omitempty"`
	// The time expression. Set the time expression based on the selected time type.
	//
	// - none: No value is required.
	//
	// - cron: Enter a standard cron expression. Online verification is supported.
	//
	// - api: No value is required.
	//
	// - fixed_rate: Enter a fixed frequency value in seconds. For example, 30 indicates that the node is triggered every 30 seconds.
	//
	// - one_time: Enter a scheduling time in the format of yyyy-MM-dd HH:mm:ss or a timestamp in milliseconds. For example, "2022-10-10 10:10:00".
	//
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// - -1: none
	//
	// - 1: cron
	//
	// - 3: fix_rate
	//
	// - 5: one_time
	//
	// - 100: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The time zone.
	//
	// > Default value: the time zone of the SchedulerX server.
	//
	// example:
	//
	// Hongkong
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The node weight.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The extended properties of the node.
	//
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

func (s *UpdateJobShrinkRequest) GetAppGroupId() *int64 {
	return s.AppGroupId
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

func (s *UpdateJobShrinkRequest) GetLabel() *string {
	return s.Label
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

func (s *UpdateJobShrinkRequest) SetAppGroupId(v int64) *UpdateJobShrinkRequest {
	s.AppGroupId = &v
	return s
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

func (s *UpdateJobShrinkRequest) SetLabel(v string) *UpdateJobShrinkRequest {
	s.Label = &v
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
