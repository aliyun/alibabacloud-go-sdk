// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroupId(v int64) *UpdateJobRequest
	GetAppGroupId() *int64
	SetAppName(v string) *UpdateJobRequest
	GetAppName() *string
	SetAttemptInterval(v int32) *UpdateJobRequest
	GetAttemptInterval() *int32
	SetCalendar(v string) *UpdateJobRequest
	GetCalendar() *string
	SetChildJobId(v string) *UpdateJobRequest
	GetChildJobId() *string
	SetClusterId(v string) *UpdateJobRequest
	GetClusterId() *string
	SetDependentStrategy(v int32) *UpdateJobRequest
	GetDependentStrategy() *int32
	SetDescription(v string) *UpdateJobRequest
	GetDescription() *string
	SetExecutorBlockStrategy(v int32) *UpdateJobRequest
	GetExecutorBlockStrategy() *int32
	SetJobHandler(v string) *UpdateJobRequest
	GetJobHandler() *string
	SetJobId(v int64) *UpdateJobRequest
	GetJobId() *int64
	SetLabel(v string) *UpdateJobRequest
	GetLabel() *string
	SetMaxAttempt(v int32) *UpdateJobRequest
	GetMaxAttempt() *int32
	SetMaxConcurrency(v int32) *UpdateJobRequest
	GetMaxConcurrency() *int32
	SetName(v string) *UpdateJobRequest
	GetName() *string
	SetNoticeConfig(v *UpdateJobRequestNoticeConfig) *UpdateJobRequest
	GetNoticeConfig() *UpdateJobRequestNoticeConfig
	SetNoticeContacts(v []*UpdateJobRequestNoticeContacts) *UpdateJobRequest
	GetNoticeContacts() []*UpdateJobRequestNoticeContacts
	SetParameters(v string) *UpdateJobRequest
	GetParameters() *string
	SetPriority(v int32) *UpdateJobRequest
	GetPriority() *int32
	SetRouteStrategy(v int32) *UpdateJobRequest
	GetRouteStrategy() *int32
	SetScript(v string) *UpdateJobRequest
	GetScript() *string
	SetStartTime(v int64) *UpdateJobRequest
	GetStartTime() *int64
	SetStartTimeType(v string) *UpdateJobRequest
	GetStartTimeType() *string
	SetTimeExpression(v string) *UpdateJobRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *UpdateJobRequest
	GetTimeType() *int32
	SetTimezone(v string) *UpdateJobRequest
	GetTimezone() *string
	SetWeight(v int32) *UpdateJobRequest
	GetWeight() *int32
	SetXAttrs(v string) *UpdateJobRequest
	GetXAttrs() *string
}

type UpdateJobRequest struct {
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
	NoticeConfig *UpdateJobRequestNoticeConfig `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty" type:"Struct"`
	// The notification contact configuration.
	NoticeContacts []*UpdateJobRequestNoticeContacts `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty" type:"Repeated"`
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

func (s UpdateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *UpdateJobRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateJobRequest) GetAttemptInterval() *int32 {
	return s.AttemptInterval
}

func (s *UpdateJobRequest) GetCalendar() *string {
	return s.Calendar
}

func (s *UpdateJobRequest) GetChildJobId() *string {
	return s.ChildJobId
}

func (s *UpdateJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateJobRequest) GetDependentStrategy() *int32 {
	return s.DependentStrategy
}

func (s *UpdateJobRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateJobRequest) GetExecutorBlockStrategy() *int32 {
	return s.ExecutorBlockStrategy
}

func (s *UpdateJobRequest) GetJobHandler() *string {
	return s.JobHandler
}

func (s *UpdateJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *UpdateJobRequest) GetLabel() *string {
	return s.Label
}

func (s *UpdateJobRequest) GetMaxAttempt() *int32 {
	return s.MaxAttempt
}

func (s *UpdateJobRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *UpdateJobRequest) GetName() *string {
	return s.Name
}

func (s *UpdateJobRequest) GetNoticeConfig() *UpdateJobRequestNoticeConfig {
	return s.NoticeConfig
}

func (s *UpdateJobRequest) GetNoticeContacts() []*UpdateJobRequestNoticeContacts {
	return s.NoticeContacts
}

func (s *UpdateJobRequest) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateJobRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateJobRequest) GetRouteStrategy() *int32 {
	return s.RouteStrategy
}

func (s *UpdateJobRequest) GetScript() *string {
	return s.Script
}

func (s *UpdateJobRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateJobRequest) GetStartTimeType() *string {
	return s.StartTimeType
}

func (s *UpdateJobRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *UpdateJobRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *UpdateJobRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateJobRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateJobRequest) GetXAttrs() *string {
	return s.XAttrs
}

func (s *UpdateJobRequest) SetAppGroupId(v int64) *UpdateJobRequest {
	s.AppGroupId = &v
	return s
}

func (s *UpdateJobRequest) SetAppName(v string) *UpdateJobRequest {
	s.AppName = &v
	return s
}

func (s *UpdateJobRequest) SetAttemptInterval(v int32) *UpdateJobRequest {
	s.AttemptInterval = &v
	return s
}

func (s *UpdateJobRequest) SetCalendar(v string) *UpdateJobRequest {
	s.Calendar = &v
	return s
}

func (s *UpdateJobRequest) SetChildJobId(v string) *UpdateJobRequest {
	s.ChildJobId = &v
	return s
}

func (s *UpdateJobRequest) SetClusterId(v string) *UpdateJobRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateJobRequest) SetDependentStrategy(v int32) *UpdateJobRequest {
	s.DependentStrategy = &v
	return s
}

func (s *UpdateJobRequest) SetDescription(v string) *UpdateJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateJobRequest) SetExecutorBlockStrategy(v int32) *UpdateJobRequest {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *UpdateJobRequest) SetJobHandler(v string) *UpdateJobRequest {
	s.JobHandler = &v
	return s
}

func (s *UpdateJobRequest) SetJobId(v int64) *UpdateJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateJobRequest) SetLabel(v string) *UpdateJobRequest {
	s.Label = &v
	return s
}

func (s *UpdateJobRequest) SetMaxAttempt(v int32) *UpdateJobRequest {
	s.MaxAttempt = &v
	return s
}

func (s *UpdateJobRequest) SetMaxConcurrency(v int32) *UpdateJobRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateJobRequest) SetName(v string) *UpdateJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateJobRequest) SetNoticeConfig(v *UpdateJobRequestNoticeConfig) *UpdateJobRequest {
	s.NoticeConfig = v
	return s
}

func (s *UpdateJobRequest) SetNoticeContacts(v []*UpdateJobRequestNoticeContacts) *UpdateJobRequest {
	s.NoticeContacts = v
	return s
}

func (s *UpdateJobRequest) SetParameters(v string) *UpdateJobRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateJobRequest) SetPriority(v int32) *UpdateJobRequest {
	s.Priority = &v
	return s
}

func (s *UpdateJobRequest) SetRouteStrategy(v int32) *UpdateJobRequest {
	s.RouteStrategy = &v
	return s
}

func (s *UpdateJobRequest) SetScript(v string) *UpdateJobRequest {
	s.Script = &v
	return s
}

func (s *UpdateJobRequest) SetStartTime(v int64) *UpdateJobRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateJobRequest) SetStartTimeType(v string) *UpdateJobRequest {
	s.StartTimeType = &v
	return s
}

func (s *UpdateJobRequest) SetTimeExpression(v string) *UpdateJobRequest {
	s.TimeExpression = &v
	return s
}

func (s *UpdateJobRequest) SetTimeType(v int32) *UpdateJobRequest {
	s.TimeType = &v
	return s
}

func (s *UpdateJobRequest) SetTimezone(v string) *UpdateJobRequest {
	s.Timezone = &v
	return s
}

func (s *UpdateJobRequest) SetWeight(v int32) *UpdateJobRequest {
	s.Weight = &v
	return s
}

func (s *UpdateJobRequest) SetXAttrs(v string) *UpdateJobRequest {
	s.XAttrs = &v
	return s
}

func (s *UpdateJobRequest) Validate() error {
	if s.NoticeConfig != nil {
		if err := s.NoticeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NoticeContacts != nil {
		for _, item := range s.NoticeContacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateJobRequestNoticeConfig struct {
	// The early completion threshold, in seconds.
	//
	// example:
	//
	// 30
	EndEarly *int32 `json:"EndEarly,omitempty" xml:"EndEarly,omitempty"`
	// Specifies whether to enable the early completion alert.
	EndEarlyEnable *bool `json:"EndEarlyEnable,omitempty" xml:"EndEarlyEnable,omitempty"`
	// Specifies whether to enable the failure alert. Valid values:
	//
	// - **true**: Enables the failure alert.
	//
	// - **false**: Disables the failure alert.
	//
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// The number of consecutive failures.
	//
	// > An alert is sent only when the number of consecutive failures exceeds the configured value.
	//
	// example:
	//
	// true
	FailLimitTimes *int32 `json:"FailLimitTimes,omitempty" xml:"FailLimitTimes,omitempty"`
	// Specifies whether to enable the no-available-machine alert. Valid values:
	//
	// - **true**: Enables the no-available-machine alert.
	//
	// - **false**: Disables the no-available-machine alert.
	//
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The notification channel. Valid values:
	//
	//  - sms: SMS
	//
	//  - phone: phone call
	//
	// - mail: email
	//
	// - webhook: webhook
	//
	// > Separate multiple notification channels with commas.
	//
	// example:
	//
	// webhook,sms,mail,phone
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// Specifies whether to enable the success notification. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	SuccessNotice *bool `json:"SuccessNotice,omitempty" xml:"SuccessNotice,omitempty"`
	// The node execution timeout period, in seconds.
	//
	// example:
	//
	// 90
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to enable the timeout alert. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Specifies whether to enable the timeout termination for the current trigger. Valid values:
	//
	// - **true**: Enables the timeout termination.
	//
	// - **false**: Disables the timeout termination.
	//
	// example:
	//
	// true
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
}

func (s UpdateJobRequestNoticeConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRequestNoticeConfig) GoString() string {
	return s.String()
}

func (s *UpdateJobRequestNoticeConfig) GetEndEarly() *int32 {
	return s.EndEarly
}

func (s *UpdateJobRequestNoticeConfig) GetEndEarlyEnable() *bool {
	return s.EndEarlyEnable
}

func (s *UpdateJobRequestNoticeConfig) GetFailEnable() *bool {
	return s.FailEnable
}

func (s *UpdateJobRequestNoticeConfig) GetFailLimitTimes() *int32 {
	return s.FailLimitTimes
}

func (s *UpdateJobRequestNoticeConfig) GetMissWorkerEnable() *bool {
	return s.MissWorkerEnable
}

func (s *UpdateJobRequestNoticeConfig) GetSendChannel() *string {
	return s.SendChannel
}

func (s *UpdateJobRequestNoticeConfig) GetSuccessNotice() *bool {
	return s.SuccessNotice
}

func (s *UpdateJobRequestNoticeConfig) GetTimeout() *int64 {
	return s.Timeout
}

func (s *UpdateJobRequestNoticeConfig) GetTimeoutEnable() *bool {
	return s.TimeoutEnable
}

func (s *UpdateJobRequestNoticeConfig) GetTimeoutKillEnable() *bool {
	return s.TimeoutKillEnable
}

func (s *UpdateJobRequestNoticeConfig) SetEndEarly(v int32) *UpdateJobRequestNoticeConfig {
	s.EndEarly = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetEndEarlyEnable(v bool) *UpdateJobRequestNoticeConfig {
	s.EndEarlyEnable = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetFailEnable(v bool) *UpdateJobRequestNoticeConfig {
	s.FailEnable = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetFailLimitTimes(v int32) *UpdateJobRequestNoticeConfig {
	s.FailLimitTimes = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetMissWorkerEnable(v bool) *UpdateJobRequestNoticeConfig {
	s.MissWorkerEnable = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetSendChannel(v string) *UpdateJobRequestNoticeConfig {
	s.SendChannel = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetSuccessNotice(v bool) *UpdateJobRequestNoticeConfig {
	s.SuccessNotice = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetTimeout(v int64) *UpdateJobRequestNoticeConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetTimeoutEnable(v bool) *UpdateJobRequestNoticeConfig {
	s.TimeoutEnable = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) SetTimeoutKillEnable(v bool) *UpdateJobRequestNoticeConfig {
	s.TimeoutKillEnable = &v
	return s
}

func (s *UpdateJobRequestNoticeConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateJobRequestNoticeContacts struct {
	// The contact type.
	//
	// >Default configurations: 1.
	//
	// example:
	//
	// 1
	ContactType *int32 `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// The contact name.
	//
	// example:
	//
	// xiaoming
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateJobRequestNoticeContacts) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRequestNoticeContacts) GoString() string {
	return s.String()
}

func (s *UpdateJobRequestNoticeContacts) GetContactType() *int32 {
	return s.ContactType
}

func (s *UpdateJobRequestNoticeContacts) GetName() *string {
	return s.Name
}

func (s *UpdateJobRequestNoticeContacts) SetContactType(v int32) *UpdateJobRequestNoticeContacts {
	s.ContactType = &v
	return s
}

func (s *UpdateJobRequestNoticeContacts) SetName(v string) *UpdateJobRequestNoticeContacts {
	s.Name = &v
	return s
}

func (s *UpdateJobRequestNoticeContacts) Validate() error {
	return dara.Validate(s)
}
