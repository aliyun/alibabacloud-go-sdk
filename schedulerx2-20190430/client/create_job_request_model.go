// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttemptInterval(v int32) *CreateJobRequest
	GetAttemptInterval() *int32
	SetCalendar(v string) *CreateJobRequest
	GetCalendar() *string
	SetClassName(v string) *CreateJobRequest
	GetClassName() *string
	SetConsumerSize(v int32) *CreateJobRequest
	GetConsumerSize() *int32
	SetContactInfo(v []*CreateJobRequestContactInfo) *CreateJobRequest
	GetContactInfo() []*CreateJobRequestContactInfo
	SetContent(v string) *CreateJobRequest
	GetContent() *string
	SetDataOffset(v int32) *CreateJobRequest
	GetDataOffset() *int32
	SetDescription(v string) *CreateJobRequest
	GetDescription() *string
	SetDispatcherSize(v int32) *CreateJobRequest
	GetDispatcherSize() *int32
	SetExecuteMode(v string) *CreateJobRequest
	GetExecuteMode() *string
	SetFailEnable(v bool) *CreateJobRequest
	GetFailEnable() *bool
	SetFailTimes(v int32) *CreateJobRequest
	GetFailTimes() *int32
	SetGroupId(v string) *CreateJobRequest
	GetGroupId() *string
	SetJobType(v string) *CreateJobRequest
	GetJobType() *string
	SetMaxAttempt(v int32) *CreateJobRequest
	GetMaxAttempt() *int32
	SetMaxConcurrency(v int32) *CreateJobRequest
	GetMaxConcurrency() *int32
	SetMissWorkerEnable(v bool) *CreateJobRequest
	GetMissWorkerEnable() *bool
	SetName(v string) *CreateJobRequest
	GetName() *string
	SetNamespace(v string) *CreateJobRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *CreateJobRequest
	GetNamespaceSource() *string
	SetPageSize(v int32) *CreateJobRequest
	GetPageSize() *int32
	SetParameters(v string) *CreateJobRequest
	GetParameters() *string
	SetPriority(v int32) *CreateJobRequest
	GetPriority() *int32
	SetQueueSize(v int32) *CreateJobRequest
	GetQueueSize() *int32
	SetRegionId(v string) *CreateJobRequest
	GetRegionId() *string
	SetSendChannel(v string) *CreateJobRequest
	GetSendChannel() *string
	SetStatus(v int32) *CreateJobRequest
	GetStatus() *int32
	SetSuccessNoticeEnable(v bool) *CreateJobRequest
	GetSuccessNoticeEnable() *bool
	SetTaskAttemptInterval(v int32) *CreateJobRequest
	GetTaskAttemptInterval() *int32
	SetTaskMaxAttempt(v int32) *CreateJobRequest
	GetTaskMaxAttempt() *int32
	SetTimeExpression(v string) *CreateJobRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *CreateJobRequest
	GetTimeType() *int32
	SetTimeout(v int64) *CreateJobRequest
	GetTimeout() *int64
	SetTimeoutEnable(v bool) *CreateJobRequest
	GetTimeoutEnable() *bool
	SetTimeoutKillEnable(v bool) *CreateJobRequest
	GetTimeoutKillEnable() *bool
	SetTimezone(v string) *CreateJobRequest
	GetTimezone() *string
	SetXAttrs(v string) *CreateJobRequest
	GetXAttrs() *string
}

type CreateJobRequest struct {
	// The time interval between retry attempts in case of a job failure. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// If you set TimeType to 1 (cron), you can specify calendar days.
	//
	// example:
	//
	// This parameter is not supported. You do not need to specify this parameter.
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The full path of the job interface class.
	//
	// This parameter is available only when you set JobType to java. You must enter a full path.
	//
	// example:
	//
	// com.alibaba.schedulerx.test.helloworld
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The number of threads that a single worker triggers simultaneously. You can specify this parameter for MapReduce jobs. Default value: 5.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The information about the alert contact.
	ContactInfo []*CreateJobRequestContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// The script content. This parameter is required when you set JobType to python, shell, go, or k8s.
	//
	// example:
	//
	// echo \\"hello\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// If you set TimeType to 1 (cron), you can specify a time offset. Unit: seconds.
	//
	// example:
	//
	// 2400
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The job description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of task distribution threads. This parameter is an advanced configuration item of the MapReduce job. Default value: 5.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The execution mode of the job. Valid values:
	//
	// 	- **Stand-alone operation**
	//
	// 	- **Broadcast run**
	//
	// 	- **Visual MapReduce**
	//
	// 	- **MapReduce**
	//
	// 	- **Shard run**
	//
	// This parameter is required.
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// Specifies whether to trigger an alert when a job fails. Valid values:
	//
	// 	- **true**: triggers an alert when a job fails.
	//
	// 	- **false**: does not trigger an alert when a job fails.
	//
	// example:
	//
	// false
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// The maximum number of consecutive failures before an alert is triggered. An alert will be triggered if the number of consecutive failures reaches the value of this parameter.
	//
	// example:
	//
	// 2
	FailTimes *int32 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job type. Valid values:
	//
	// 	- java
	//
	// 	- python
	//
	// 	- shell
	//
	// 	- go
	//
	// 	- http
	//
	// 	- xxljob
	//
	// 	- dataworks
	//
	// 	- k8s
	//
	// 	- springschedule
	//
	// This parameter is required.
	//
	// example:
	//
	// java
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The maximum number of retry attempts in case of a job failure. Specify this parameter based on your business requirements. Default value: 0.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrent instances. By default, only one instance can run at a time. When an instance is running, the next instance is not triggered even if the scheduled start time arrives.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// Specifies whether to generate an alert if no machines are available to run the job. Valid values:
	//
	// 	- **true**: generates an alert if no machines are available to run the job.
	//
	// 	- **false**: does not generate an alert if no machines are available to run the job.
	//
	// example:
	//
	// false
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The job name.
	//
	// This parameter is required.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. You must specify this parameter only if the namespace is provided by a third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The number of entries per page. You can specify this parameter for MapReduce jobs. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user-defined parameters that you can obtain when the job is running.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Priority   *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The maximum capacity of the task queue. You can specify this parameter for MapReduce jobs. Default value: 10000.
	//
	// example:
	//
	// 10000
	QueueSize *int32 `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method that is used to send alerts. Set the value to sms. Default value: sms.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// Specifies whether to enable the job. If this parameter is set to 0, the job is disabled. If this parameter is set to 1, the job is enabled. Default value: 1.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether to send notifications for successfully running the job.
	//
	// example:
	//
	// false
	SuccessNoticeEnable *bool `json:"SuccessNoticeEnable,omitempty" xml:"SuccessNoticeEnable,omitempty"`
	// The time interval between retry attempts in case of a job failure. This parameter is an advanced configuration item of the MapReduce job. Default value: 0.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The maximum number of retry attempts in case of a job failure. This parameter is an advanced configuration item of the MapReduce job. Default value: 0.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
	// The time expression. Specify the time expression based on the value of TimeType:
	//
	// 	- If you set TimeType to **1*	- (cron), specify this parameter to a standard CRON expression.
	//
	// 	- If you set TimeType to **100*	- (api), no time expression is required.
	//
	// 	- If you set TimeType to **3*	- (fixed_rate), specify this parameter to a fixed frequency in seconds. For example, if you set this parameter to 30, the system triggers a job every 30 seconds.
	//
	// 	- If you set TimeType to **4*	- (second_delay), specify this parameter to a fixed delay after which the job is triggered. Valid values: 1 to 60. Unit: seconds.
	//
	// 	- If you set TimeType to **5*	- (one_time), specify this parameter to a specific time point at which the job is triggered. The time is in the format of yyyy-MM-dd HH:mm:ss, such as 2022-10-10 10:10:00, or a timestamp in milliseconds.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fixed_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **5**: one_time
	//
	// 	- **100**: api
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The timeout threshold. Unit: seconds. Default value: 7200.
	//
	// example:
	//
	// 7200
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to enable the timeout alert feature. If the feature is enabled, an alert will be triggered upon a timeout. Valid values:
	//
	// 	- **true**: enables the timeout alert feature.
	//
	// 	- **false**: disables the timeout alert feature.
	//
	// example:
	//
	// false
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Specifies whether to enable the timeout termination feature. If the feature is enabled, a job will automatically be terminated if it times out. Valid values:
	//
	// 	- **true**: enables the timeout termination feature.
	//
	// 	- **false**: disables the timeout termination feature.
	//
	// example:
	//
	// false
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
	// Time zone.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The extended attributes. If you set JobType to k8s, this parameter is required. For a job whose resource type is Job-YAML, set this parameter to {"resource":"job"}. For a job whose resource type is Shell-Script, set this parameter to {"image":"busybox","resource":"shell"}.
	//
	// example:
	//
	// {"resource":"job"}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
}

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetAttemptInterval() *int32 {
	return s.AttemptInterval
}

func (s *CreateJobRequest) GetCalendar() *string {
	return s.Calendar
}

func (s *CreateJobRequest) GetClassName() *string {
	return s.ClassName
}

func (s *CreateJobRequest) GetConsumerSize() *int32 {
	return s.ConsumerSize
}

func (s *CreateJobRequest) GetContactInfo() []*CreateJobRequestContactInfo {
	return s.ContactInfo
}

func (s *CreateJobRequest) GetContent() *string {
	return s.Content
}

func (s *CreateJobRequest) GetDataOffset() *int32 {
	return s.DataOffset
}

func (s *CreateJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateJobRequest) GetDispatcherSize() *int32 {
	return s.DispatcherSize
}

func (s *CreateJobRequest) GetExecuteMode() *string {
	return s.ExecuteMode
}

func (s *CreateJobRequest) GetFailEnable() *bool {
	return s.FailEnable
}

func (s *CreateJobRequest) GetFailTimes() *int32 {
	return s.FailTimes
}

func (s *CreateJobRequest) GetGroupId() *string {
	return s.GroupId
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

func (s *CreateJobRequest) GetMissWorkerEnable() *bool {
	return s.MissWorkerEnable
}

func (s *CreateJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateJobRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *CreateJobRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CreateJobRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateJobRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateJobRequest) GetQueueSize() *int32 {
	return s.QueueSize
}

func (s *CreateJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateJobRequest) GetSendChannel() *string {
	return s.SendChannel
}

func (s *CreateJobRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateJobRequest) GetSuccessNoticeEnable() *bool {
	return s.SuccessNoticeEnable
}

func (s *CreateJobRequest) GetTaskAttemptInterval() *int32 {
	return s.TaskAttemptInterval
}

func (s *CreateJobRequest) GetTaskMaxAttempt() *int32 {
	return s.TaskMaxAttempt
}

func (s *CreateJobRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *CreateJobRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *CreateJobRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateJobRequest) GetTimeoutEnable() *bool {
	return s.TimeoutEnable
}

func (s *CreateJobRequest) GetTimeoutKillEnable() *bool {
	return s.TimeoutKillEnable
}

func (s *CreateJobRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateJobRequest) GetXAttrs() *string {
	return s.XAttrs
}

func (s *CreateJobRequest) SetAttemptInterval(v int32) *CreateJobRequest {
	s.AttemptInterval = &v
	return s
}

func (s *CreateJobRequest) SetCalendar(v string) *CreateJobRequest {
	s.Calendar = &v
	return s
}

func (s *CreateJobRequest) SetClassName(v string) *CreateJobRequest {
	s.ClassName = &v
	return s
}

func (s *CreateJobRequest) SetConsumerSize(v int32) *CreateJobRequest {
	s.ConsumerSize = &v
	return s
}

func (s *CreateJobRequest) SetContactInfo(v []*CreateJobRequestContactInfo) *CreateJobRequest {
	s.ContactInfo = v
	return s
}

func (s *CreateJobRequest) SetContent(v string) *CreateJobRequest {
	s.Content = &v
	return s
}

func (s *CreateJobRequest) SetDataOffset(v int32) *CreateJobRequest {
	s.DataOffset = &v
	return s
}

func (s *CreateJobRequest) SetDescription(v string) *CreateJobRequest {
	s.Description = &v
	return s
}

func (s *CreateJobRequest) SetDispatcherSize(v int32) *CreateJobRequest {
	s.DispatcherSize = &v
	return s
}

func (s *CreateJobRequest) SetExecuteMode(v string) *CreateJobRequest {
	s.ExecuteMode = &v
	return s
}

func (s *CreateJobRequest) SetFailEnable(v bool) *CreateJobRequest {
	s.FailEnable = &v
	return s
}

func (s *CreateJobRequest) SetFailTimes(v int32) *CreateJobRequest {
	s.FailTimes = &v
	return s
}

func (s *CreateJobRequest) SetGroupId(v string) *CreateJobRequest {
	s.GroupId = &v
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

func (s *CreateJobRequest) SetMissWorkerEnable(v bool) *CreateJobRequest {
	s.MissWorkerEnable = &v
	return s
}

func (s *CreateJobRequest) SetName(v string) *CreateJobRequest {
	s.Name = &v
	return s
}

func (s *CreateJobRequest) SetNamespace(v string) *CreateJobRequest {
	s.Namespace = &v
	return s
}

func (s *CreateJobRequest) SetNamespaceSource(v string) *CreateJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *CreateJobRequest) SetPageSize(v int32) *CreateJobRequest {
	s.PageSize = &v
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

func (s *CreateJobRequest) SetQueueSize(v int32) *CreateJobRequest {
	s.QueueSize = &v
	return s
}

func (s *CreateJobRequest) SetRegionId(v string) *CreateJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateJobRequest) SetSendChannel(v string) *CreateJobRequest {
	s.SendChannel = &v
	return s
}

func (s *CreateJobRequest) SetStatus(v int32) *CreateJobRequest {
	s.Status = &v
	return s
}

func (s *CreateJobRequest) SetSuccessNoticeEnable(v bool) *CreateJobRequest {
	s.SuccessNoticeEnable = &v
	return s
}

func (s *CreateJobRequest) SetTaskAttemptInterval(v int32) *CreateJobRequest {
	s.TaskAttemptInterval = &v
	return s
}

func (s *CreateJobRequest) SetTaskMaxAttempt(v int32) *CreateJobRequest {
	s.TaskMaxAttempt = &v
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

func (s *CreateJobRequest) SetTimeout(v int64) *CreateJobRequest {
	s.Timeout = &v
	return s
}

func (s *CreateJobRequest) SetTimeoutEnable(v bool) *CreateJobRequest {
	s.TimeoutEnable = &v
	return s
}

func (s *CreateJobRequest) SetTimeoutKillEnable(v bool) *CreateJobRequest {
	s.TimeoutKillEnable = &v
	return s
}

func (s *CreateJobRequest) SetTimezone(v string) *CreateJobRequest {
	s.Timezone = &v
	return s
}

func (s *CreateJobRequest) SetXAttrs(v string) *CreateJobRequest {
	s.XAttrs = &v
	return s
}

func (s *CreateJobRequest) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestContactInfo struct {
	// The webhook URL of the DingTalk chatbot.[](https://open.dingtalk.com/document/org/application-types)
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=**********
	Ding *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// test***@***.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// Tom
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The mobile number of the alert contact.
	//
	// example:
	//
	// 1381111****
	UserPhone *string `json:"UserPhone,omitempty" xml:"UserPhone,omitempty"`
}

func (s CreateJobRequestContactInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestContactInfo) GoString() string {
	return s.String()
}

func (s *CreateJobRequestContactInfo) GetDing() *string {
	return s.Ding
}

func (s *CreateJobRequestContactInfo) GetUserMail() *string {
	return s.UserMail
}

func (s *CreateJobRequestContactInfo) GetUserName() *string {
	return s.UserName
}

func (s *CreateJobRequestContactInfo) GetUserPhone() *string {
	return s.UserPhone
}

func (s *CreateJobRequestContactInfo) SetDing(v string) *CreateJobRequestContactInfo {
	s.Ding = &v
	return s
}

func (s *CreateJobRequestContactInfo) SetUserMail(v string) *CreateJobRequestContactInfo {
	s.UserMail = &v
	return s
}

func (s *CreateJobRequestContactInfo) SetUserName(v string) *CreateJobRequestContactInfo {
	s.UserName = &v
	return s
}

func (s *CreateJobRequestContactInfo) SetUserPhone(v string) *CreateJobRequestContactInfo {
	s.UserPhone = &v
	return s
}

func (s *CreateJobRequestContactInfo) Validate() error {
	return dara.Validate(s)
}
