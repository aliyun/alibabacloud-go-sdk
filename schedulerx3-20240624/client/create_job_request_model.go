// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateJobRequest
	GetAppName() *string
	SetAttemptInterval(v int32) *CreateJobRequest
	GetAttemptInterval() *int32
	SetCalendar(v string) *CreateJobRequest
	GetCalendar() *string
	SetChildJobId(v string) *CreateJobRequest
	GetChildJobId() *string
	SetClusterId(v string) *CreateJobRequest
	GetClusterId() *string
	SetDescription(v string) *CreateJobRequest
	GetDescription() *string
	SetExecutorBlockStrategy(v int32) *CreateJobRequest
	GetExecutorBlockStrategy() *int32
	SetJobHandler(v string) *CreateJobRequest
	GetJobHandler() *string
	SetJobType(v string) *CreateJobRequest
	GetJobType() *string
	SetMaxAttempt(v int32) *CreateJobRequest
	GetMaxAttempt() *int32
	SetMaxConcurrency(v int32) *CreateJobRequest
	GetMaxConcurrency() *int32
	SetName(v string) *CreateJobRequest
	GetName() *string
	SetNoticeConfig(v *CreateJobRequestNoticeConfig) *CreateJobRequest
	GetNoticeConfig() *CreateJobRequestNoticeConfig
	SetNoticeContacts(v []*CreateJobRequestNoticeContacts) *CreateJobRequest
	GetNoticeContacts() []*CreateJobRequestNoticeContacts
	SetParameters(v string) *CreateJobRequest
	GetParameters() *string
	SetPriority(v int32) *CreateJobRequest
	GetPriority() *int32
	SetRouteStrategy(v int32) *CreateJobRequest
	GetRouteStrategy() *int32
	SetScript(v string) *CreateJobRequest
	GetScript() *string
	SetStartTime(v int64) *CreateJobRequest
	GetStartTime() *int64
	SetStatus(v int32) *CreateJobRequest
	GetStatus() *int32
	SetTimeExpression(v string) *CreateJobRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *CreateJobRequest
	GetTimeType() *int32
	SetTimezone(v string) *CreateJobRequest
	GetTimezone() *string
	SetWeight(v int32) *CreateJobRequest
	GetWeight() *int32
}

type CreateJobRequest struct {
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
	Name           *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	NoticeConfig   *CreateJobRequestNoticeConfig     `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty" type:"Struct"`
	NoticeContacts []*CreateJobRequestNoticeContacts `json:"NoticeContacts,omitempty" xml:"NoticeContacts,omitempty" type:"Repeated"`
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
	RouteStrategy *int32  `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	Script        *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// example:
	//
	// 1701310327000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateJobRequest) GetAttemptInterval() *int32 {
	return s.AttemptInterval
}

func (s *CreateJobRequest) GetCalendar() *string {
	return s.Calendar
}

func (s *CreateJobRequest) GetChildJobId() *string {
	return s.ChildJobId
}

func (s *CreateJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateJobRequest) GetExecutorBlockStrategy() *int32 {
	return s.ExecutorBlockStrategy
}

func (s *CreateJobRequest) GetJobHandler() *string {
	return s.JobHandler
}

func (s *CreateJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateJobRequest) GetMaxAttempt() *int32 {
	return s.MaxAttempt
}

func (s *CreateJobRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *CreateJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateJobRequest) GetNoticeConfig() *CreateJobRequestNoticeConfig {
	return s.NoticeConfig
}

func (s *CreateJobRequest) GetNoticeContacts() []*CreateJobRequestNoticeContacts {
	return s.NoticeContacts
}

func (s *CreateJobRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateJobRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateJobRequest) GetRouteStrategy() *int32 {
	return s.RouteStrategy
}

func (s *CreateJobRequest) GetScript() *string {
	return s.Script
}

func (s *CreateJobRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateJobRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateJobRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *CreateJobRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *CreateJobRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateJobRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateJobRequest) SetAppName(v string) *CreateJobRequest {
	s.AppName = &v
	return s
}

func (s *CreateJobRequest) SetAttemptInterval(v int32) *CreateJobRequest {
	s.AttemptInterval = &v
	return s
}

func (s *CreateJobRequest) SetCalendar(v string) *CreateJobRequest {
	s.Calendar = &v
	return s
}

func (s *CreateJobRequest) SetChildJobId(v string) *CreateJobRequest {
	s.ChildJobId = &v
	return s
}

func (s *CreateJobRequest) SetClusterId(v string) *CreateJobRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobRequest) SetDescription(v string) *CreateJobRequest {
	s.Description = &v
	return s
}

func (s *CreateJobRequest) SetExecutorBlockStrategy(v int32) *CreateJobRequest {
	s.ExecutorBlockStrategy = &v
	return s
}

func (s *CreateJobRequest) SetJobHandler(v string) *CreateJobRequest {
	s.JobHandler = &v
	return s
}

func (s *CreateJobRequest) SetJobType(v string) *CreateJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateJobRequest) SetMaxAttempt(v int32) *CreateJobRequest {
	s.MaxAttempt = &v
	return s
}

func (s *CreateJobRequest) SetMaxConcurrency(v int32) *CreateJobRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *CreateJobRequest) SetName(v string) *CreateJobRequest {
	s.Name = &v
	return s
}

func (s *CreateJobRequest) SetNoticeConfig(v *CreateJobRequestNoticeConfig) *CreateJobRequest {
	s.NoticeConfig = v
	return s
}

func (s *CreateJobRequest) SetNoticeContacts(v []*CreateJobRequestNoticeContacts) *CreateJobRequest {
	s.NoticeContacts = v
	return s
}

func (s *CreateJobRequest) SetParameters(v string) *CreateJobRequest {
	s.Parameters = &v
	return s
}

func (s *CreateJobRequest) SetPriority(v int32) *CreateJobRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobRequest) SetRouteStrategy(v int32) *CreateJobRequest {
	s.RouteStrategy = &v
	return s
}

func (s *CreateJobRequest) SetScript(v string) *CreateJobRequest {
	s.Script = &v
	return s
}

func (s *CreateJobRequest) SetStartTime(v int64) *CreateJobRequest {
	s.StartTime = &v
	return s
}

func (s *CreateJobRequest) SetStatus(v int32) *CreateJobRequest {
	s.Status = &v
	return s
}

func (s *CreateJobRequest) SetTimeExpression(v string) *CreateJobRequest {
	s.TimeExpression = &v
	return s
}

func (s *CreateJobRequest) SetTimeType(v int32) *CreateJobRequest {
	s.TimeType = &v
	return s
}

func (s *CreateJobRequest) SetTimezone(v string) *CreateJobRequest {
	s.Timezone = &v
	return s
}

func (s *CreateJobRequest) SetWeight(v int32) *CreateJobRequest {
	s.Weight = &v
	return s
}

func (s *CreateJobRequest) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestNoticeConfig struct {
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// example:
	//
	// 1
	FailLimitTimes *int32 `json:"FailLimitTimes,omitempty" xml:"FailLimitTimes,omitempty"`
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// example:
	//
	// mail
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// example:
	//
	// true
	SuccessNotice *bool `json:"SuccessNotice,omitempty" xml:"SuccessNotice,omitempty"`
	// example:
	//
	// 30
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

func (s CreateJobRequestNoticeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestNoticeConfig) GoString() string {
	return s.String()
}

func (s *CreateJobRequestNoticeConfig) GetFailEnable() *bool {
	return s.FailEnable
}

func (s *CreateJobRequestNoticeConfig) GetFailLimitTimes() *int32 {
	return s.FailLimitTimes
}

func (s *CreateJobRequestNoticeConfig) GetMissWorkerEnable() *bool {
	return s.MissWorkerEnable
}

func (s *CreateJobRequestNoticeConfig) GetSendChannel() *string {
	return s.SendChannel
}

func (s *CreateJobRequestNoticeConfig) GetSuccessNotice() *bool {
	return s.SuccessNotice
}

func (s *CreateJobRequestNoticeConfig) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateJobRequestNoticeConfig) GetTimeoutEnable() *bool {
	return s.TimeoutEnable
}

func (s *CreateJobRequestNoticeConfig) GetTimeoutKillEnable() *bool {
	return s.TimeoutKillEnable
}

func (s *CreateJobRequestNoticeConfig) SetFailEnable(v bool) *CreateJobRequestNoticeConfig {
	s.FailEnable = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetFailLimitTimes(v int32) *CreateJobRequestNoticeConfig {
	s.FailLimitTimes = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetMissWorkerEnable(v bool) *CreateJobRequestNoticeConfig {
	s.MissWorkerEnable = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetSendChannel(v string) *CreateJobRequestNoticeConfig {
	s.SendChannel = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetSuccessNotice(v bool) *CreateJobRequestNoticeConfig {
	s.SuccessNotice = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetTimeout(v int64) *CreateJobRequestNoticeConfig {
	s.Timeout = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetTimeoutEnable(v bool) *CreateJobRequestNoticeConfig {
	s.TimeoutEnable = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) SetTimeoutKillEnable(v bool) *CreateJobRequestNoticeConfig {
	s.TimeoutKillEnable = &v
	return s
}

func (s *CreateJobRequestNoticeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestNoticeContacts struct {
	// example:
	//
	// 1
	ContactType *int32 `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// example:
	//
	// xiaoming
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateJobRequestNoticeContacts) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestNoticeContacts) GoString() string {
	return s.String()
}

func (s *CreateJobRequestNoticeContacts) GetContactType() *int32 {
	return s.ContactType
}

func (s *CreateJobRequestNoticeContacts) GetName() *string {
	return s.Name
}

func (s *CreateJobRequestNoticeContacts) SetContactType(v int32) *CreateJobRequestNoticeContacts {
	s.ContactType = &v
	return s
}

func (s *CreateJobRequestNoticeContacts) SetName(v string) *CreateJobRequestNoticeContacts {
	s.Name = &v
	return s
}

func (s *CreateJobRequestNoticeContacts) Validate() error {
	return dara.Validate(s)
}
