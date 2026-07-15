// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListJobsResponseBody
	GetCode() *int32
	SetData(v *ListJobsResponseBodyData) *ListJobsResponseBody
	GetData() *ListJobsResponseBodyData
	SetMessage(v string) *ListJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobsResponseBody
	GetSuccess() *bool
}

type ListJobsResponseBody struct {
	// The request status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The node list information.
	Data *ListJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is returned only if an error occurs.
	//
	// example:
	//
	// namespace can not find namespace: 1a72ecb1-b4cc-400a-a71b-20cdec9b****, namespaceSource:null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB58B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListJobsResponseBody) GetData() *ListJobsResponseBodyData {
	return s.Data
}

func (s *ListJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobsResponseBody) SetCode(v int32) *ListJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobsResponseBody) SetData(v *ListJobsResponseBodyData) *ListJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListJobsResponseBody) SetMessage(v string) *ListJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsResponseBodyData struct {
	// The node list and node details.
	Jobs []*ListJobsResponseBodyDataJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyData) GetJobs() []*ListJobsResponseBodyDataJobs {
	return s.Jobs
}

func (s *ListJobsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListJobsResponseBodyData) SetJobs(v []*ListJobsResponseBodyDataJobs) *ListJobsResponseBodyData {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBodyData) SetPageNumber(v int32) *ListJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBodyData) SetPageSize(v int32) *ListJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBodyData) SetTotal(v int32) *ListJobsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListJobsResponseBodyData) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobsResponseBodyDataJobs struct {
	// The retry interval on error, in seconds. Default value: 30.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// The full path of the node interface class. This field is returned only when the node is of the Java type.
	//
	// example:
	//
	// com.alibaba.schedulerx.test.helloworld
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The script code content for Python, Shell, or Go node types.
	//
	// example:
	//
	// echo \\"hello\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The node description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The node execution mode. Valid values:
	//
	// - **standalone**: standalone
	//
	// - **broadcast**: broadcast
	//
	// - **parallel**: parallel computing
	//
	// - **grid**: memory grid
	//
	// - **batch**: grid computing
	//
	// - **shard**: shard
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// The full path of the JAR package in OSS.
	//
	// example:
	//
	// https:doc***.oss-cn-hangzhou.aliyuncs.com/sc-****-D-0.0.2-SNAPSHOT.jar
	JarUrl *string `json:"JarUrl,omitempty" xml:"JarUrl,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 99341
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The node monitoring information.
	JobMonitorInfo *ListJobsResponseBodyDataJobsJobMonitorInfo `json:"JobMonitorInfo,omitempty" xml:"JobMonitorInfo,omitempty" type:"Struct"`
	// The node type.
	//
	// example:
	//
	// java
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The advanced configuration. This is used only for parallel computing, memory grid, and grid computing.
	MapTaskXAttrs *ListJobsResponseBodyDataJobsMapTaskXAttrs `json:"MapTaskXAttrs,omitempty" xml:"MapTaskXAttrs,omitempty" type:"Struct"`
	// The maximum number of retries on error. Set this based on business requirements. Default value: 0.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrently running instances. Default value: 1. This means that if the previous trigger has not finished running, the next trigger will not be initiated even if the scheduled time has arrived.
	//
	// example:
	//
	// 1
	MaxConcurrency *string `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The node name.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user-defined parameters that can be obtained at runtime.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The node status. Valid values:
	//
	// - **1**: Enabled. The node can be triggered normally.
	//
	// - **0**: Disabled. The node will not be triggered.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time configuration information.
	TimeConfig *ListJobsResponseBodyDataJobsTimeConfig `json:"TimeConfig,omitempty" xml:"TimeConfig,omitempty" type:"Struct"`
	// The node extension field.
	//
	// example:
	//
	// {"pageSize":5,"queueSize":10,"consumerSize":5,"dispatcherSize":5,"taskMaxAttempt":0,"taskAttemptInterval":0,"globalConsumerSize":1000,"taskDispatchMode":"push"}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
}

func (s ListJobsResponseBodyDataJobs) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyDataJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobs) GetAttemptInterval() *int32 {
	return s.AttemptInterval
}

func (s *ListJobsResponseBodyDataJobs) GetClassName() *string {
	return s.ClassName
}

func (s *ListJobsResponseBodyDataJobs) GetContent() *string {
	return s.Content
}

func (s *ListJobsResponseBodyDataJobs) GetDescription() *string {
	return s.Description
}

func (s *ListJobsResponseBodyDataJobs) GetExecuteMode() *string {
	return s.ExecuteMode
}

func (s *ListJobsResponseBodyDataJobs) GetJarUrl() *string {
	return s.JarUrl
}

func (s *ListJobsResponseBodyDataJobs) GetJobId() *int64 {
	return s.JobId
}

func (s *ListJobsResponseBodyDataJobs) GetJobMonitorInfo() *ListJobsResponseBodyDataJobsJobMonitorInfo {
	return s.JobMonitorInfo
}

func (s *ListJobsResponseBodyDataJobs) GetJobType() *string {
	return s.JobType
}

func (s *ListJobsResponseBodyDataJobs) GetMapTaskXAttrs() *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	return s.MapTaskXAttrs
}

func (s *ListJobsResponseBodyDataJobs) GetMaxAttempt() *int32 {
	return s.MaxAttempt
}

func (s *ListJobsResponseBodyDataJobs) GetMaxConcurrency() *string {
	return s.MaxConcurrency
}

func (s *ListJobsResponseBodyDataJobs) GetName() *string {
	return s.Name
}

func (s *ListJobsResponseBodyDataJobs) GetParameters() *string {
	return s.Parameters
}

func (s *ListJobsResponseBodyDataJobs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListJobsResponseBodyDataJobs) GetStatus() *int32 {
	return s.Status
}

func (s *ListJobsResponseBodyDataJobs) GetTimeConfig() *ListJobsResponseBodyDataJobsTimeConfig {
	return s.TimeConfig
}

func (s *ListJobsResponseBodyDataJobs) GetXAttrs() *string {
	return s.XAttrs
}

func (s *ListJobsResponseBodyDataJobs) SetAttemptInterval(v int32) *ListJobsResponseBodyDataJobs {
	s.AttemptInterval = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetClassName(v string) *ListJobsResponseBodyDataJobs {
	s.ClassName = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetContent(v string) *ListJobsResponseBodyDataJobs {
	s.Content = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetDescription(v string) *ListJobsResponseBodyDataJobs {
	s.Description = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetExecuteMode(v string) *ListJobsResponseBodyDataJobs {
	s.ExecuteMode = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetJarUrl(v string) *ListJobsResponseBodyDataJobs {
	s.JarUrl = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetJobId(v int64) *ListJobsResponseBodyDataJobs {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetJobMonitorInfo(v *ListJobsResponseBodyDataJobsJobMonitorInfo) *ListJobsResponseBodyDataJobs {
	s.JobMonitorInfo = v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetJobType(v string) *ListJobsResponseBodyDataJobs {
	s.JobType = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetMapTaskXAttrs(v *ListJobsResponseBodyDataJobsMapTaskXAttrs) *ListJobsResponseBodyDataJobs {
	s.MapTaskXAttrs = v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetMaxAttempt(v int32) *ListJobsResponseBodyDataJobs {
	s.MaxAttempt = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetMaxConcurrency(v string) *ListJobsResponseBodyDataJobs {
	s.MaxConcurrency = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetName(v string) *ListJobsResponseBodyDataJobs {
	s.Name = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetParameters(v string) *ListJobsResponseBodyDataJobs {
	s.Parameters = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetStartTime(v int64) *ListJobsResponseBodyDataJobs {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetStatus(v int32) *ListJobsResponseBodyDataJobs {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetTimeConfig(v *ListJobsResponseBodyDataJobsTimeConfig) *ListJobsResponseBodyDataJobs {
	s.TimeConfig = v
	return s
}

func (s *ListJobsResponseBodyDataJobs) SetXAttrs(v string) *ListJobsResponseBodyDataJobs {
	s.XAttrs = &v
	return s
}

func (s *ListJobsResponseBodyDataJobs) Validate() error {
	if s.JobMonitorInfo != nil {
		if err := s.JobMonitorInfo.Validate(); err != nil {
			return err
		}
	}
	if s.MapTaskXAttrs != nil {
		if err := s.MapTaskXAttrs.Validate(); err != nil {
			return err
		}
	}
	if s.TimeConfig != nil {
		if err := s.TimeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsResponseBodyDataJobsJobMonitorInfo struct {
	// The contact information.
	ContactInfo []*ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// The alert switch and threshold configuration.
	MonitorConfig *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty" type:"Struct"`
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfo) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfo) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfo) GetContactInfo() []*ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo {
	return s.ContactInfo
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfo) GetMonitorConfig() *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	return s.MonitorConfig
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfo) SetContactInfo(v []*ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) *ListJobsResponseBodyDataJobsJobMonitorInfo {
	s.ContactInfo = v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfo) SetMonitorConfig(v *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) *ListJobsResponseBodyDataJobsJobMonitorInfo {
	s.MonitorConfig = v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfo) Validate() error {
	if s.ContactInfo != nil {
		for _, item := range s.ContactInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MonitorConfig != nil {
		if err := s.MonitorConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo struct {
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=**********
	Ding *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// user@mail.com
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

func (s ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) GetDing() *string {
	return s.Ding
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) GetUserMail() *string {
	return s.UserMail
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) GetUserName() *string {
	return s.UserName
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) GetUserPhone() *string {
	return s.UserPhone
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) SetDing(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo {
	s.Ding = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) SetUserMail(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo {
	s.UserMail = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) SetUserName(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo {
	s.UserName = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) SetUserPhone(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo {
	s.UserPhone = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig struct {
	// Specifies whether to enable the failure alert switch. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// Specifies whether to enable the no-available-machine alert.
	//
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The alert notification method. Currently, only sms is supported.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// The timeout threshold, in seconds. Default value: 7200.
	//
	// example:
	//
	// 12300
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to enable the timeout alert switch. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Specifies whether to enable the timeout termination switch for the current trigger. This is disabled by default. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) GetFailEnable() *bool {
	return s.FailEnable
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) GetMissWorkerEnable() *bool {
	return s.MissWorkerEnable
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) GetSendChannel() *string {
	return s.SendChannel
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) GetTimeout() *int64 {
	return s.Timeout
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) GetTimeoutEnable() *bool {
	return s.TimeoutEnable
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) GetTimeoutKillEnable() *bool {
	return s.TimeoutKillEnable
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetFailEnable(v bool) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.FailEnable = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetMissWorkerEnable(v bool) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.MissWorkerEnable = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetSendChannel(v string) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.SendChannel = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetTimeout(v int64) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.Timeout = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetTimeoutEnable(v bool) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.TimeoutEnable = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) SetTimeoutKillEnable(v bool) *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig {
	s.TimeoutKillEnable = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsJobMonitorInfoMonitorConfig) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyDataJobsMapTaskXAttrs struct {
	// The number of threads for a single trigger on a single machine. Default value: 5.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The number of subtask dispatch threads. Default value: 5.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The number of subtasks pulled per batch for a parallel node. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The upper limit of the subtask queue cache. Default value: 10000.
	//
	// example:
	//
	// 10000
	QueueSize *int32 `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	// The retry interval for a subtask on failure.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The number of retries for a subtask on failure.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
}

func (s ListJobsResponseBodyDataJobsMapTaskXAttrs) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsMapTaskXAttrs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) GetConsumerSize() *int32 {
	return s.ConsumerSize
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) GetDispatcherSize() *int32 {
	return s.DispatcherSize
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) GetQueueSize() *int32 {
	return s.QueueSize
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) GetTaskAttemptInterval() *int32 {
	return s.TaskAttemptInterval
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) GetTaskMaxAttempt() *int32 {
	return s.TaskMaxAttempt
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetConsumerSize(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.ConsumerSize = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetDispatcherSize(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.DispatcherSize = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetPageSize(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetQueueSize(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.QueueSize = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetTaskAttemptInterval(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.TaskAttemptInterval = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) SetTaskMaxAttempt(v int32) *ListJobsResponseBodyDataJobsMapTaskXAttrs {
	s.TaskMaxAttempt = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsMapTaskXAttrs) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyDataJobsTimeConfig struct {
	// The custom calendar that can be specified for the cron type.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The time offset that can be specified for the cron type, in seconds.
	//
	// example:
	//
	// 0
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The time expression. Valid values:
	//
	// - **api**: No time expression.
	//
	// - **fix_rate**: A specific fixed frequency value. For example, 30 indicates that the node is triggered every 30 seconds.
	//
	// - **cron**: A standard cron expression.
	//
	// - **second_delay**: A fixed delay in seconds before each execution (1s to 60s).
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time configuration type. Valid values:
	//
	// - **1**: cron
	//
	// - **3**: fix_rate
	//
	// - **4**: second_delay
	//
	// - **100**: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
}

func (s ListJobsResponseBodyDataJobsTimeConfig) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyDataJobsTimeConfig) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) GetCalendar() *string {
	return s.Calendar
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) GetDataOffset() *int32 {
	return s.DataOffset
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) GetTimeType() *int32 {
	return s.TimeType
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) SetCalendar(v string) *ListJobsResponseBodyDataJobsTimeConfig {
	s.Calendar = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) SetDataOffset(v int32) *ListJobsResponseBodyDataJobsTimeConfig {
	s.DataOffset = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) SetTimeExpression(v string) *ListJobsResponseBodyDataJobsTimeConfig {
	s.TimeExpression = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) SetTimeType(v int32) *ListJobsResponseBodyDataJobsTimeConfig {
	s.TimeType = &v
	return s
}

func (s *ListJobsResponseBodyDataJobsTimeConfig) Validate() error {
	return dara.Validate(s)
}
