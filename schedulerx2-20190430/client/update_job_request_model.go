// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttemptInterval(v int32) *UpdateJobRequest
	GetAttemptInterval() *int32
	SetCalendar(v string) *UpdateJobRequest
	GetCalendar() *string
	SetClassName(v string) *UpdateJobRequest
	GetClassName() *string
	SetConsumerSize(v int32) *UpdateJobRequest
	GetConsumerSize() *int32
	SetContactInfo(v []*UpdateJobRequestContactInfo) *UpdateJobRequest
	GetContactInfo() []*UpdateJobRequestContactInfo
	SetContent(v string) *UpdateJobRequest
	GetContent() *string
	SetDataOffset(v int32) *UpdateJobRequest
	GetDataOffset() *int32
	SetDescription(v string) *UpdateJobRequest
	GetDescription() *string
	SetDispatcherSize(v int32) *UpdateJobRequest
	GetDispatcherSize() *int32
	SetExecuteMode(v string) *UpdateJobRequest
	GetExecuteMode() *string
	SetFailEnable(v bool) *UpdateJobRequest
	GetFailEnable() *bool
	SetFailTimes(v int32) *UpdateJobRequest
	GetFailTimes() *int32
	SetGroupId(v string) *UpdateJobRequest
	GetGroupId() *string
	SetJobId(v int64) *UpdateJobRequest
	GetJobId() *int64
	SetMaxAttempt(v int32) *UpdateJobRequest
	GetMaxAttempt() *int32
	SetMaxConcurrency(v int32) *UpdateJobRequest
	GetMaxConcurrency() *int32
	SetMissWorkerEnable(v bool) *UpdateJobRequest
	GetMissWorkerEnable() *bool
	SetName(v string) *UpdateJobRequest
	GetName() *string
	SetNamespace(v string) *UpdateJobRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *UpdateJobRequest
	GetNamespaceSource() *string
	SetPageSize(v int32) *UpdateJobRequest
	GetPageSize() *int32
	SetParameters(v string) *UpdateJobRequest
	GetParameters() *string
	SetPriority(v int32) *UpdateJobRequest
	GetPriority() *int32
	SetQueueSize(v int32) *UpdateJobRequest
	GetQueueSize() *int32
	SetRegionId(v string) *UpdateJobRequest
	GetRegionId() *string
	SetSendChannel(v string) *UpdateJobRequest
	GetSendChannel() *string
	SetStartTime(v int64) *UpdateJobRequest
	GetStartTime() *int64
	SetSuccessNoticeEnable(v bool) *UpdateJobRequest
	GetSuccessNoticeEnable() *bool
	SetTaskAttemptInterval(v int32) *UpdateJobRequest
	GetTaskAttemptInterval() *int32
	SetTaskDispatchMode(v string) *UpdateJobRequest
	GetTaskDispatchMode() *string
	SetTaskMaxAttempt(v int32) *UpdateJobRequest
	GetTaskMaxAttempt() *int32
	SetTemplate(v string) *UpdateJobRequest
	GetTemplate() *string
	SetTimeExpression(v string) *UpdateJobRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *UpdateJobRequest
	GetTimeType() *int32
	SetTimeout(v int64) *UpdateJobRequest
	GetTimeout() *int64
	SetTimeoutEnable(v bool) *UpdateJobRequest
	GetTimeoutEnable() *bool
	SetTimeoutKillEnable(v bool) *UpdateJobRequest
	GetTimeoutKillEnable() *bool
	SetTimezone(v string) *UpdateJobRequest
	GetTimezone() *string
	SetXAttrs(v string) *UpdateJobRequest
	GetXAttrs() *string
}

type UpdateJobRequest struct {
	// The retry interval on errors. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// The custom calendar that can be optionally specified for the cron type.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The full path of the node interface class.
	//
	// This field is required only for Java node types, and the full path must be specified.
	//
	// example:
	//
	// com.alibaba.test.helloworld
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The advanced configuration for parallel grid tasks. The number of threads for a single trigger on a single machine. Default value: 5.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The contact information for the node.
	//
	// 	Notice: This field is deprecated.</notice>
	ContactInfo []*UpdateJobRequestContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// - If the node type is python, shell, or k8s, specify the corresponding script content.
	//
	// - If the node type is golang, the content format example is {"jobName":"HelloWorld"}.
	//
	// example:
	//
	// echo \\"hello\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time offset that can be optionally specified for the cron type. Unit: seconds.
	//
	// example:
	//
	// 2400
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The node description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The advanced configuration for parallel grid tasks. The number of subtask dispatch threads. Default value: 5.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The node execution mode. Valid values:
	//
	// - **standalone**: standalone
	//
	// - **broadcatst**: broadcast
	//
	// - **parallel**: visual MapReduce
	//
	// - **batch**: MapReduce
	//
	// - **shard**: shard
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// Specifies whether to enable the failure alert. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// The number of consecutive failures before an alert is triggered.
	//
	// example:
	//
	// 1
	FailTimes *int32 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	// The application ID. You can obtain the application ID on the Application Management page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The node ID. You can obtain the node ID on the Task Management page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum number of retries on errors. Set this parameter based on your business requirements.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrently running instances. Default value: 1. This means that if the previous trigger has not finished running, the next trigger is not performed even if the scheduled time has arrived.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// Specifies whether to enable the no-available-machine alert. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The node name.
	//
	// example:
	//
	// helloword
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
	// The advanced configuration for parallel grid tasks. The number of subtasks pulled per request. Default value: 100.
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
	// The advanced configuration for parallel grid tasks. The maximum cache size of the subtask queue. Default value: 10000.
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
	// The alert notification method. Currently, only sms is supported.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether to enable the success notification.
	//
	// example:
	//
	// false
	SuccessNoticeEnable *bool `json:"SuccessNoticeEnable,omitempty" xml:"SuccessNoticeEnable,omitempty"`
	// The advanced configuration for parallel grid tasks. The retry interval for failed subtasks.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The advanced configuration for parallel grid tasks. Specifies the push model or pull model.
	//
	// example:
	//
	// push
	TaskDispatchMode *string `json:"TaskDispatchMode,omitempty" xml:"TaskDispatchMode,omitempty"`
	// The advanced configuration for parallel grid tasks. The number of retries for failed subtasks.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
	// The custom task template for k8s node types.
	//
	// example:
	//
	// apiVersion: v1
	//
	// kind: Pod
	//
	// metadata:
	//
	//   name: schedulerx-node-{JOB_ID}
	//
	//   namespace: {NAMESPACE}
	//
	// spec:
	//
	//   containers:
	//
	//   - name: node-job
	//
	//     image: node:16
	//
	//     imagePullPolicy: IfNotPresent
	//
	//     volumeMounts:
	//
	//     - name: script-node
	//
	//       mountPath: script/node
	//
	//     command: ["node", "script/node/node-{JOB_ID}.js"]
	//
	//   volumes:
	//
	//   - name: script-node
	//
	//     configMap:
	//
	//       name: schedulerx-configmap
	//
	//       items:
	//
	//       - key: schedulerx-node-{JOB_ID}
	//
	//         path: node-{JOB_ID}.js
	//
	//   restartPolicy: Never
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
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
	// example:
	//
	// 30
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time configuration type. Valid values:
	//
	// - **1**: cron
	//
	// - **3**: fix_rate
	//
	// - **4**: second_delay
	//
	// - **5**: one_time
	//
	// - **100**: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The timeout threshold. Unit: seconds.
	//
	// example:
	//
	// 7200
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to enable the timeout alert. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Specifies whether to enable the timeout termination for the current trigger. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
	// The time zone.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The parameter that must be configured for k8s node types.
	//
	// Job task: {"resource":"job"}
	//
	// Shell task: {"image":"busybox","resource":"shell"}
	//
	// example:
	//
	// {"resource":"shell","fileFormat":"unix","templateType":"customTemplate"}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
}

func (s UpdateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) GetAttemptInterval() *int32 {
	return s.AttemptInterval
}

func (s *UpdateJobRequest) GetCalendar() *string {
	return s.Calendar
}

func (s *UpdateJobRequest) GetClassName() *string {
	return s.ClassName
}

func (s *UpdateJobRequest) GetConsumerSize() *int32 {
	return s.ConsumerSize
}

func (s *UpdateJobRequest) GetContactInfo() []*UpdateJobRequestContactInfo {
	return s.ContactInfo
}

func (s *UpdateJobRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateJobRequest) GetDataOffset() *int32 {
	return s.DataOffset
}

func (s *UpdateJobRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateJobRequest) GetDispatcherSize() *int32 {
	return s.DispatcherSize
}

func (s *UpdateJobRequest) GetExecuteMode() *string {
	return s.ExecuteMode
}

func (s *UpdateJobRequest) GetFailEnable() *bool {
	return s.FailEnable
}

func (s *UpdateJobRequest) GetFailTimes() *int32 {
	return s.FailTimes
}

func (s *UpdateJobRequest) GetGroupId() *string {
	return s.GroupId
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

func (s *UpdateJobRequest) GetMissWorkerEnable() *bool {
	return s.MissWorkerEnable
}

func (s *UpdateJobRequest) GetName() *string {
	return s.Name
}

func (s *UpdateJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateJobRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *UpdateJobRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *UpdateJobRequest) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateJobRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateJobRequest) GetQueueSize() *int32 {
	return s.QueueSize
}

func (s *UpdateJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateJobRequest) GetSendChannel() *string {
	return s.SendChannel
}

func (s *UpdateJobRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateJobRequest) GetSuccessNoticeEnable() *bool {
	return s.SuccessNoticeEnable
}

func (s *UpdateJobRequest) GetTaskAttemptInterval() *int32 {
	return s.TaskAttemptInterval
}

func (s *UpdateJobRequest) GetTaskDispatchMode() *string {
	return s.TaskDispatchMode
}

func (s *UpdateJobRequest) GetTaskMaxAttempt() *int32 {
	return s.TaskMaxAttempt
}

func (s *UpdateJobRequest) GetTemplate() *string {
	return s.Template
}

func (s *UpdateJobRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *UpdateJobRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *UpdateJobRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *UpdateJobRequest) GetTimeoutEnable() *bool {
	return s.TimeoutEnable
}

func (s *UpdateJobRequest) GetTimeoutKillEnable() *bool {
	return s.TimeoutKillEnable
}

func (s *UpdateJobRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateJobRequest) GetXAttrs() *string {
	return s.XAttrs
}

func (s *UpdateJobRequest) SetAttemptInterval(v int32) *UpdateJobRequest {
	s.AttemptInterval = &v
	return s
}

func (s *UpdateJobRequest) SetCalendar(v string) *UpdateJobRequest {
	s.Calendar = &v
	return s
}

func (s *UpdateJobRequest) SetClassName(v string) *UpdateJobRequest {
	s.ClassName = &v
	return s
}

func (s *UpdateJobRequest) SetConsumerSize(v int32) *UpdateJobRequest {
	s.ConsumerSize = &v
	return s
}

func (s *UpdateJobRequest) SetContactInfo(v []*UpdateJobRequestContactInfo) *UpdateJobRequest {
	s.ContactInfo = v
	return s
}

func (s *UpdateJobRequest) SetContent(v string) *UpdateJobRequest {
	s.Content = &v
	return s
}

func (s *UpdateJobRequest) SetDataOffset(v int32) *UpdateJobRequest {
	s.DataOffset = &v
	return s
}

func (s *UpdateJobRequest) SetDescription(v string) *UpdateJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateJobRequest) SetDispatcherSize(v int32) *UpdateJobRequest {
	s.DispatcherSize = &v
	return s
}

func (s *UpdateJobRequest) SetExecuteMode(v string) *UpdateJobRequest {
	s.ExecuteMode = &v
	return s
}

func (s *UpdateJobRequest) SetFailEnable(v bool) *UpdateJobRequest {
	s.FailEnable = &v
	return s
}

func (s *UpdateJobRequest) SetFailTimes(v int32) *UpdateJobRequest {
	s.FailTimes = &v
	return s
}

func (s *UpdateJobRequest) SetGroupId(v string) *UpdateJobRequest {
	s.GroupId = &v
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

func (s *UpdateJobRequest) SetMissWorkerEnable(v bool) *UpdateJobRequest {
	s.MissWorkerEnable = &v
	return s
}

func (s *UpdateJobRequest) SetName(v string) *UpdateJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateJobRequest) SetNamespace(v string) *UpdateJobRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateJobRequest) SetNamespaceSource(v string) *UpdateJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *UpdateJobRequest) SetPageSize(v int32) *UpdateJobRequest {
	s.PageSize = &v
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

func (s *UpdateJobRequest) SetQueueSize(v int32) *UpdateJobRequest {
	s.QueueSize = &v
	return s
}

func (s *UpdateJobRequest) SetRegionId(v string) *UpdateJobRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateJobRequest) SetSendChannel(v string) *UpdateJobRequest {
	s.SendChannel = &v
	return s
}

func (s *UpdateJobRequest) SetStartTime(v int64) *UpdateJobRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateJobRequest) SetSuccessNoticeEnable(v bool) *UpdateJobRequest {
	s.SuccessNoticeEnable = &v
	return s
}

func (s *UpdateJobRequest) SetTaskAttemptInterval(v int32) *UpdateJobRequest {
	s.TaskAttemptInterval = &v
	return s
}

func (s *UpdateJobRequest) SetTaskDispatchMode(v string) *UpdateJobRequest {
	s.TaskDispatchMode = &v
	return s
}

func (s *UpdateJobRequest) SetTaskMaxAttempt(v int32) *UpdateJobRequest {
	s.TaskMaxAttempt = &v
	return s
}

func (s *UpdateJobRequest) SetTemplate(v string) *UpdateJobRequest {
	s.Template = &v
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

func (s *UpdateJobRequest) SetTimeout(v int64) *UpdateJobRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateJobRequest) SetTimeoutEnable(v bool) *UpdateJobRequest {
	s.TimeoutEnable = &v
	return s
}

func (s *UpdateJobRequest) SetTimeoutKillEnable(v bool) *UpdateJobRequest {
	s.TimeoutKillEnable = &v
	return s
}

func (s *UpdateJobRequest) SetTimezone(v string) *UpdateJobRequest {
	s.Timezone = &v
	return s
}

func (s *UpdateJobRequest) SetXAttrs(v string) *UpdateJobRequest {
	s.XAttrs = &v
	return s
}

func (s *UpdateJobRequest) Validate() error {
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

type UpdateJobRequestContactInfo struct {
	// The webhook URL of the DingTalk chatbot in the DingTalk group for alert contacts. References: [DingTalk development documentation](https://open.dingtalk.com/document/org/application-types).
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=**********
	Ding *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// test***@***.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The username.
	//
	// example:
	//
	// userA
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The mobile phone number of the user.
	//
	// example:
	//
	// 1381111****
	UserPhone *string `json:"UserPhone,omitempty" xml:"UserPhone,omitempty"`
}

func (s UpdateJobRequestContactInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRequestContactInfo) GoString() string {
	return s.String()
}

func (s *UpdateJobRequestContactInfo) GetDing() *string {
	return s.Ding
}

func (s *UpdateJobRequestContactInfo) GetUserMail() *string {
	return s.UserMail
}

func (s *UpdateJobRequestContactInfo) GetUserName() *string {
	return s.UserName
}

func (s *UpdateJobRequestContactInfo) GetUserPhone() *string {
	return s.UserPhone
}

func (s *UpdateJobRequestContactInfo) SetDing(v string) *UpdateJobRequestContactInfo {
	s.Ding = &v
	return s
}

func (s *UpdateJobRequestContactInfo) SetUserMail(v string) *UpdateJobRequestContactInfo {
	s.UserMail = &v
	return s
}

func (s *UpdateJobRequestContactInfo) SetUserName(v string) *UpdateJobRequestContactInfo {
	s.UserName = &v
	return s
}

func (s *UpdateJobRequestContactInfo) SetUserPhone(v string) *UpdateJobRequestContactInfo {
	s.UserPhone = &v
	return s
}

func (s *UpdateJobRequestContactInfo) Validate() error {
	return dara.Validate(s)
}
