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
}

type UpdateJobRequest struct {
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
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	Name           *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	NoticeConfig   *UpdateJobRequestNoticeConfig     `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty" type:"Struct"`
	NoticeContacts []*UpdateJobRequestNoticeContacts `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty" type:"Repeated"`
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
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// example:
	//
	// echo "hello world"
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// example:
	//
	// 1716902187
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	StartTimeType *string `json:"StartTimeType,omitempty" xml:"StartTimeType,omitempty"`
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
	// example:
	//
	// 30
	EndEarly       *int32 `json:"EndEarly,omitempty" xml:"EndEarly,omitempty"`
	EndEarlyEnable *bool  `json:"EndEarlyEnable,omitempty" xml:"EndEarlyEnable,omitempty"`
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// example:
	//
	// true
	FailLimitTimes *int32 `json:"FailLimitTimes,omitempty" xml:"FailLimitTimes,omitempty"`
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// example:
	//
	// webhook,sms,mail,phone
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// example:
	//
	// true
	SuccessNotice *bool `json:"SuccessNotice,omitempty" xml:"SuccessNotice,omitempty"`
	// example:
	//
	// 90
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
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
	// example:
	//
	// 1
	ContactType *int32 `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
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
