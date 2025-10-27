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
	// The interval of retries after a job failure. Default value: 30. Unit: seconds.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// If you set TimeType to 1 (cron), you can specify calendar days.
	//
	// example:
	//
	// Business days
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The full path of the job interface class.
	//
	// This field is available only when you set the job type to java. In this case, you must enter a full path.
	//
	// example:
	//
	// com.alibaba.test.helloworld
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The number of threads that are triggered by a single worker at a time. Default value: 5. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The information about the alert contact.
	ContactInfo []*UpdateJobRequestContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// The script content. This parameter is required when you set the job type to python, shell, go, or k8s.
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
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of task distribution threads. Default value: 5. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The execution mode of the job. Valid values:
	//
	// 	- **Stand-alone operation**: standalone
	//
	// 	- **Broadcast run**: broadcatst
	//
	// 	- **Visual MapReduce**: parallel
	//
	// 	- **MapReduce**: batch
	//
	// 	- **Shard run**: shard
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// Specifies whether to turn on Failure alarm. If the switch is turned on, an alert will be generated upon a failure. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// The number of consecutive failures. An alert will be received if the number of consecutive failures reaches the value of this parameter.
	//
	// example:
	//
	// 1
	FailTimes *int32 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the job ID on the Task Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum number of retries after a job failure. This parameter is specified based on your business requirements.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrent instances. Default value: 1. The default value indicates that only one instance is allowed to run at a time. When an instance is running, another instance is not triggered even if the scheduled time for running the instance is reached.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// Specifies whether to turn on No machine alarm available. If the switch is turned on, an alert will be generated when no machine is available for running the job. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The job name.
	//
	// example:
	//
	// helloword
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The namespace source. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The number of tasks that can be pulled at a time. Default value: 100. This parameter is an advanced configuration item of the MapReduce job.
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
	// The maximum number of tasks that can be queued. Default value: 10000. This parameter is an advanced configuration item of the MapReduce job.
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
	// The method that is used to send alerts. Only Short Message Service (SMS) is supported.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// Specifies whether to turn on Successful notice. If the switch is turned on, a notice will be sent when a job succeeds.
	//
	// example:
	//
	// false
	SuccessNoticeEnable *bool `json:"SuccessNoticeEnable,omitempty" xml:"SuccessNoticeEnable,omitempty"`
	// The interval of retries after a task failure. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The job mode. Valid values: push and pull. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// push
	TaskDispatchMode *string `json:"TaskDispatchMode,omitempty" xml:"TaskDispatchMode,omitempty"`
	// The number of retries after a task failure. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
	// Custom task template for the k8s task type.
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
	// example:
	//
	// 30
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fix_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **100**: api
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
	// Specifies whether to turn on Timeout alarm. If the switch is turned on, an alert will be generated upon a timeout. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Specifies whether to turn on Timeout termination. If the switch is turned on, the job will be terminated upon a timeout. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
	// Time zone.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// If you set JobType to k8s, this parameter is required. xxljob task: {"resource":"job"} shell task: {"image":"busybox","resource":"shell"}
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
	// userA
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The mobile phone number of the alert contact.
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
