// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJobInfoResponseBody
	GetCode() *int32
	SetData(v *GetJobInfoResponseBodyData) *GetJobInfoResponseBody
	GetData() *GetJobInfoResponseBodyData
	SetMessage(v string) *GetJobInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJobInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobInfoResponseBody
	GetSuccess() *bool
}

type GetJobInfoResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the job.
	Data *GetJobInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned only if an error occurs.
	//
	// example:
	//
	// jobid: 92583 not match groupId: testSchedulerx.defaultGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the job details were obtained. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJobInfoResponseBody) GetData() *GetJobInfoResponseBodyData {
	return s.Data
}

func (s *GetJobInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJobInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobInfoResponseBody) SetCode(v int32) *GetJobInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobInfoResponseBody) SetData(v *GetJobInfoResponseBodyData) *GetJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetJobInfoResponseBody) SetMessage(v string) *GetJobInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobInfoResponseBody) SetRequestId(v string) *GetJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobInfoResponseBody) SetSuccess(v bool) *GetJobInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyData struct {
	// The configurations of the job.
	JobConfigInfo *GetJobInfoResponseBodyDataJobConfigInfo `json:"JobConfigInfo,omitempty" xml:"JobConfigInfo,omitempty" type:"Struct"`
}

func (s GetJobInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyData) GetJobConfigInfo() *GetJobInfoResponseBodyDataJobConfigInfo {
	return s.JobConfigInfo
}

func (s *GetJobInfoResponseBodyData) SetJobConfigInfo(v *GetJobInfoResponseBodyDataJobConfigInfo) *GetJobInfoResponseBodyData {
	s.JobConfigInfo = v
	return s
}

func (s *GetJobInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyDataJobConfigInfo struct {
	// The interval at which the system retried to run the job after a job failure. Default value: 30. Unit: seconds.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// The full path of the job interface class. This parameter is returned only for jobs whose job type is Java.
	//
	// example:
	//
	// com.alibaba.test.helloword
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The script of a script job.
	//
	// example:
	//
	// echo "clear" > /home/admin/edas-container/logs/catalina.out
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the job.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode of the job. Valid values:
	//
	// 	- **Stand-alone operation**: standalone
	//
	// 	- **Broadcast run**: broadcast
	//
	// 	- **Visual MapReduce**: parallel
	//
	// 	- **MapReduce**: batch
	//
	// 	- **Shard run**: sharding
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// The full path used to upload files to Object Storage Service (OSS).
	//
	// If you use a JAR package, you can upload the JAR package to this OSS path.
	//
	// example:
	//
	// https://test.oss-cn-hangzhou.aliyuncs.com/schedulerX/test.jar
	JarUrl *string `json:"JarUrl,omitempty" xml:"JarUrl,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 538039
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The monitoring information of the job.
	JobMonitorInfo *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo `json:"JobMonitorInfo,omitempty" xml:"JobMonitorInfo,omitempty" type:"Struct"`
	// The job type.
	//
	// example:
	//
	// java
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The advanced configurations of the job.
	MapTaskXAttrs *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs `json:"MapTaskXAttrs,omitempty" xml:"MapTaskXAttrs,omitempty" type:"Struct"`
	// The maximum number of retries after a job failure. This parameter was specified based on your business requirements. Default value: 0.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrent instances. Default value: 1. The default value indicates that if the last triggered instance is running, the next instance is not triggered even if the scheduled point in time for running the next instance is reached.
	//
	// example:
	//
	// 1
	MaxConcurrency *string `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The job name.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user-defined parameters that you can obtain when the job is running.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Indicates whether the job was enabled. Valid values:
	//
	// 	- **1**: The job was enabled and could be triggered.
	//
	// 	- **0**: The job was disabled and could not be triggered.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time configurations.
	TimeConfig *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig `json:"TimeConfig,omitempty" xml:"TimeConfig,omitempty" type:"Struct"`
	// The extended fields.
	//
	// example:
	//
	// {"pageSize":5,"queueSize":10,"consumerSize":5,"dispatcherSize":5,"taskMaxAttempt":0,"taskAttemptInterval":0,"globalConsumerSize":1000,"taskDispatchMode":"push"}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfo) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfo) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetAttemptInterval() *int32 {
	return s.AttemptInterval
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetClassName() *string {
	return s.ClassName
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetContent() *string {
	return s.Content
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetDescription() *string {
	return s.Description
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetExecuteMode() *string {
	return s.ExecuteMode
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetJarUrl() *string {
	return s.JarUrl
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetJobId() *int64 {
	return s.JobId
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetJobMonitorInfo() *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo {
	return s.JobMonitorInfo
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetJobType() *string {
	return s.JobType
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetMapTaskXAttrs() *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	return s.MapTaskXAttrs
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetMaxAttempt() *int32 {
	return s.MaxAttempt
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetMaxConcurrency() *string {
	return s.MaxConcurrency
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetName() *string {
	return s.Name
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetParameters() *string {
	return s.Parameters
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetStatus() *int32 {
	return s.Status
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetTimeConfig() *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig {
	return s.TimeConfig
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetXAttrs() *string {
	return s.XAttrs
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetAttemptInterval(v int32) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.AttemptInterval = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetClassName(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.ClassName = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetContent(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Content = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetDescription(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Description = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetExecuteMode(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.ExecuteMode = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetJarUrl(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.JarUrl = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetJobId(v int64) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.JobId = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetJobMonitorInfo(v *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.JobMonitorInfo = v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetJobType(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.JobType = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetMapTaskXAttrs(v *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.MapTaskXAttrs = v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetMaxAttempt(v int32) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.MaxAttempt = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetMaxConcurrency(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.MaxConcurrency = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetName(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Name = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetParameters(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Parameters = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetStatus(v int32) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.Status = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetTimeConfig(v *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.TimeConfig = v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetXAttrs(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.XAttrs = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo struct {
	// The alert contact Information.
	ContactInfo []*GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// The configurations of the alerting features and the alert thresholds.
	MonitorConfig *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty" type:"Struct"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) GetContactInfo() []*GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo {
	return s.ContactInfo
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) GetMonitorConfig() *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	return s.MonitorConfig
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) SetContactInfo(v []*GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo {
	s.ContactInfo = v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) SetMonitorConfig(v *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo {
	s.MonitorConfig = v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo struct {
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=XXXXXX
	Ding *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// user@demo.com
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

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) GetDing() *string {
	return s.Ding
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) GetUserMail() *string {
	return s.UserMail
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) GetUserName() *string {
	return s.UserName
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) GetUserPhone() *string {
	return s.UserPhone
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) SetDing(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo {
	s.Ding = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) SetUserMail(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo {
	s.UserMail = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) SetUserName(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo {
	s.UserName = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) SetUserPhone(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo {
	s.UserPhone = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig struct {
	// Indicates whether the Failure alarm switch was turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// Indicates whether the No machine alarm available switch was turned on.
	//
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The method used to send alerts. Only Short Message Service (SMS) is supported.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// The timeout threshold. Default value: 7200. Unit: seconds.
	//
	// example:
	//
	// 12300
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Indicates whether the Timeout alarm switch was turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Indicates whether the Timeout termination switch was turned on. The switch is turned off by default.
	//
	// example:
	//
	// true
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) GetFailEnable() *bool {
	return s.FailEnable
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) GetMissWorkerEnable() *bool {
	return s.MissWorkerEnable
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) GetSendChannel() *string {
	return s.SendChannel
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) GetTimeoutEnable() *bool {
	return s.TimeoutEnable
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) GetTimeoutKillEnable() *bool {
	return s.TimeoutKillEnable
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetFailEnable(v bool) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.FailEnable = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetMissWorkerEnable(v bool) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.MissWorkerEnable = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetSendChannel(v string) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.SendChannel = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetTimeout(v int64) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.Timeout = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetTimeoutEnable(v bool) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.TimeoutEnable = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) SetTimeoutKillEnable(v bool) *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig {
	s.TimeoutKillEnable = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoMonitorConfig) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs struct {
	// The number of threads that were triggered by a single worker at a time. Default value: 5.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The number of task distribution threads. Default value: 5.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The number of tasks that were pulled by a parallel job at a time. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The maximum number of tasks that can be queued. Default value: 10000.
	//
	// example:
	//
	// 10000
	QueueSize *int32 `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	// The interval at which the system retried to run the task after a task failure.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The number of retries after a task failure.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) GetConsumerSize() *int32 {
	return s.ConsumerSize
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) GetDispatcherSize() *int32 {
	return s.DispatcherSize
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) GetQueueSize() *int32 {
	return s.QueueSize
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) GetTaskAttemptInterval() *int32 {
	return s.TaskAttemptInterval
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) GetTaskMaxAttempt() *int32 {
	return s.TaskMaxAttempt
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetConsumerSize(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.ConsumerSize = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetDispatcherSize(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.DispatcherSize = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetPageSize(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.PageSize = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetQueueSize(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.QueueSize = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetTaskAttemptInterval(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.TaskAttemptInterval = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) SetTaskMaxAttempt(v int32) *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs {
	s.TaskMaxAttempt = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyDataJobConfigInfoTimeConfig struct {
	// Custom calendar days specified if TimeType is set to **1*	- (cron).
	//
	// example:
	//
	// Business days
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The time offset specified if TimeType is set to **1*	- (cron). Unit: seconds.
	//
	// example:
	//
	// 0
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The time expression specified based on the value of TimeType:
	//
	// 	- If TimeType is set to **100*	- (api), no time expression is required.
	//
	// 	- If TimeType is set to **3*	- (fix_rate), this parameter value indicates the specific and fixed frequency. For example, if the value is 30, the system triggers a job every 30 seconds.
	//
	// 	- If TimeType is set to **1*	- (cron), this parameter value indicates the standard CRON expression used to specify the time when to schedule the job.
	//
	// 	- If TimeType is set to **4*	- (second_delay), this parameter value indicates the fixed delay after which the job is triggered. Valid values: 1 to 60. Unit: seconds.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fix_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **5**: one_time
	//
	// 	- **100**: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) GetCalendar() *string {
	return s.Calendar
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) GetDataOffset() *int32 {
	return s.DataOffset
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) GetTimeType() *int32 {
	return s.TimeType
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) SetCalendar(v string) *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig {
	s.Calendar = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) SetDataOffset(v int32) *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig {
	s.DataOffset = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) SetTimeExpression(v string) *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig {
	s.TimeExpression = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) SetTimeType(v int32) *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig {
	s.TimeType = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig) Validate() error {
	return dara.Validate(s)
}
