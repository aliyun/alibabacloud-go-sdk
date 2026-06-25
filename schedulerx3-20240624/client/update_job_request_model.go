// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobRequest interface {
	dara.Model
	String() string
	GoString() string
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
	NoticeConfig *UpdateJobRequestNoticeConfig `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty" type:"Struct"`
	// Notification configuration
	NoticeContacts []*UpdateJobRequestNoticeContacts `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty" type:"Repeated"`
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

func (s UpdateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
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
	// The threshold, in seconds, for a job to be considered as finishing early.
	//
	// example:
	//
	// 30
	EndEarly *int32 `json:"EndEarly,omitempty" xml:"EndEarly,omitempty"`
	// Indicates whether to enable an alarm when a job finishes earlier than expected. Set to `true` to enable the alarm, or `false` to disable it.
	EndEarlyEnable *bool `json:"EndEarlyEnable,omitempty" xml:"EndEarlyEnable,omitempty"`
	// Indicates whether to enable the failure alarm. Set to `true` to enable the alarm, or `false` to disable it.
	//
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// The number of consecutive failures required to trigger a failure alarm.
	//
	// example:
	//
	// true
	FailLimitTimes *int32 `json:"FailLimitTimes,omitempty" xml:"FailLimitTimes,omitempty"`
	// Indicates whether to enable an alarm if no workers are available. Set to `true` to enable the alarm, or `false` to disable it.
	//
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The notification channel. Valid values:
	//
	// \\- `sms`: sms
	//
	// \\- `phone`: voice call
	//
	// \\- `mail`: email
	//
	// \\- `webhook`: webhook
	//
	// \\> You can specify multiple channels, separated by commas.
	//
	// example:
	//
	// webhook,sms,mail,phone
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// Indicates whether to enable success notifications. Set to `true` to enable notifications, or `false` to disable them.
	//
	// example:
	//
	// true
	SuccessNotice *bool `json:"SuccessNotice,omitempty" xml:"SuccessNotice,omitempty"`
	// The job execution timeout in seconds.
	//
	// example:
	//
	// 90
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Indicates whether to enable the timeout alarm. Set to `true` to enable the alarm, or `false` to disable it.
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Indicates whether to terminate a timed-out job. Set to `true` to terminate the job, or `false` to let it continue.
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
	// \\> Default value: 1.
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
