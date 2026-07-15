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
	SetStartTime(v int64) *CreateJobRequest
	GetStartTime() *int64
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
	// The retry interval on failure. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// The custom calendar. This parameter is available for the cron time type.
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The full path of the node interface class.
	//
	// This field is required only when you select the Java node type. Specify the full path.
	//
	// example:
	//
	// com.alibaba.schedulerx.test.helloworld
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The advanced configuration for parallel grid nodes. The number of threads triggered for a single execution on a single machine. Default value: 5.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The node contact information.
	//
	// 	Notice: This field is deprecated.</notice>
	ContactInfo []*CreateJobRequestContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// - If the node type is python, shell, or k8s, specify the corresponding script content.
	//
	// - If the node type is golang, the content format example is {"jobName":"HelloWorld"}.
	//
	// example:
	//
	// echo \\"hello\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time offset. Unit: seconds. This parameter is available for the cron time type.
	//
	// example:
	//
	// 2400
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The node description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The advanced configuration for parallel grid nodes. The number of subtask dispatch threads. Default value: 5.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The node execution mode. The following execution modes are supported:
	//
	// - **Standalone**: standalone
	//
	// - **Broadcast**: broadcast
	//
	// - **Visual MapReduce**: parallel
	//
	// - **MapReduce**: batch
	//
	// - **Sharding**: sharding
	//
	// This parameter is required.
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// Specifies whether to enable the failure alert. Valid values:
	//
	// - **true**: Enables the failure alert.
	//
	// - **false**: Disables the failure alert.
	//
	// example:
	//
	// false
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// The number of consecutive failures before an alert is triggered.
	//
	// example:
	//
	// 2
	FailTimes *int32 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	// The application ID. You can obtain the application ID on the Application Management page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The node type. The following node types are supported:
	//
	// - java
	//
	// - python
	//
	// - shell
	//
	// - go
	//
	// - http
	//
	// - xxljob
	//
	// - dataworks
	//
	// - k8s
	//
	// - springschedule
	//
	// This parameter is required.
	//
	// example:
	//
	// java
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The maximum number of retries on failure. Set this parameter based on your business requirements. Default value: 0.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrently running instances. Default value: 1. This means that if the previous trigger has not finished running, the next trigger is not performed even if the scheduled time arrives.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// Specifies whether to enable the no-available-machine alert. Valid values:
	//
	// - **true**: Enables the no-available-machine alert.
	//
	// - **false**: Disables the no-available-machine alert.
	//
	// example:
	//
	// false
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The node name.
	//
	// This parameter is required.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required only for special third-party users.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The advanced configuration for parallel grid nodes. The number of subtasks pulled in a single request. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user-defined parameters that can be obtained at runtime.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The node priority. Valid values:
	//
	// - **1**: low
	//
	// - **5**: medium
	//
	// - **10**: high
	//
	// - **15**: very high
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The advanced configuration for parallel grid nodes. The maximum cache size of the subtask queue. Default value: 10000.
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
	// The alert notification channel.
	//
	// - Use the default channel of the application group: default.
	//
	// - Specify a notification channel for the node: sms, mail, phone, or webhook.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The node status. Valid values: 0: disabled. 1: enabled. Default value: 1 (enabled).
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether to enable the success notification.
	//
	// example:
	//
	// false
	SuccessNoticeEnable *bool `json:"SuccessNoticeEnable,omitempty" xml:"SuccessNoticeEnable,omitempty"`
	// The advanced configuration for parallel grid nodes. The retry interval for a failed subtask. Default value: 0.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The advanced configuration for parallel grid nodes. The number of retries for a failed subtask. Default value: 0.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
	// The time expression. Set the time expression based on the selected time type.
	//
	// - **cron**: Specify a standard cron expression. Online verification is supported.
	//
	// - **api**: No time expression is required.
	//
	// - **fixed_rate**: Specify a fixed frequency value in seconds. For example, 30 indicates that the node is triggered every 30 seconds.
	//
	// - **second_delay**: Specify a fixed delay in seconds before each execution (1s to 60s).
	//
	// - **one_time**: Specify a time in the format of yyyy-MM-dd HH:mm:ss or a timestamp in milliseconds. For example, "2022-10-10 10:10:00".
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. The following time types are supported:
	//
	// - **cron**: 1
	//
	// - **fixed_rate**: 3
	//
	// - **second_delay**: 4
	//
	// - **one_time**: 5
	//
	// - **api**: 100
	//
	// - **none**: -1
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
	// Specifies whether to enable the timeout alert. Valid values:
	//
	// - **true**: Enables the timeout alert.
	//
	// - **false**: Disables the timeout alert.
	//
	// example:
	//
	// false
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Specifies whether to enable timeout termination. Valid values:
	//
	// - **true**: Enables timeout termination.
	//
	// - **false**: Disables timeout termination.
	//
	// example:
	//
	// false
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
	// The time zone.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// If the node type is k8s, configure this parameter.
	//
	// Job task: {"resource":"job"}
	//
	// Shell task: {"image":"busybox","resource":"shell"}
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

func (s *CreateJobRequest) GetStartTime() *int64 {
	return s.StartTime
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

func (s *CreateJobRequest) SetStartTime(v int64) *CreateJobRequest {
	s.StartTime = &v
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
	if s.ContactInfo != nil {
		for _, item := range s.ContactInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateJobRequestContactInfo struct {
	// The webhook URL of the DingTalk chatbot for the alert contact\\"s DingTalk group. References: [DingTalk development documentation](https://open.dingtalk.com/document/org/application-types).
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
	// John Smith
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The mobile phone number of the alert recipient.
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
