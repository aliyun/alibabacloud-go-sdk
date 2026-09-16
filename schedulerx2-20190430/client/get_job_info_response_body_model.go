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
	// The return code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the specified node.
	Data *GetJobInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is returned only when an error occurs.
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
	// Indicates whether the node details were retrieved. Valid values:
	//
	// - **true**: The node details were retrieved.
	//
	// - **false**: The node details failed to be retrieved.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobInfoResponseBodyData struct {
	// The node configuration information.
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
	if s.JobConfigInfo != nil {
		if err := s.JobConfigInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobInfoResponseBodyDataJobConfigInfo struct {
	// The retry interval on failure. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// The full path of the node interface class. This field is available only for Java-type nodes.
	//
	// example:
	//
	// com.alibaba.test.helloword
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The script content for script-type nodes.
	//
	// example:
	//
	// echo "clear" > /home/admin/edas-container/logs/catalina.out
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The node description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1789454134000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The node execution mode. Valid values:
	//
	// - **standalone**: standalone
	//
	// - **broadcatst**: broadcast
	//
	// - **parallel**: parallel computing
	//
	// - **grid**: in-memory grid
	//
	// - **batch**: grid computing
	//
	// - **shard**: shard
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// The full path of the file uploaded to Object Storage Service (OSS).
	//
	// If you select JAR package execution, you can upload the corresponding JAR package to this OSS path.
	//
	// example:
	//
	// https://test.oss-cn-hangzhou.aliyuncs.com/schedulerX/test.jar
	JarUrl *string `json:"JarUrl,omitempty" xml:"JarUrl,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 538039
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The node monitoring information.
	JobMonitorInfo *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo `json:"JobMonitorInfo,omitempty" xml:"JobMonitorInfo,omitempty" type:"Struct"`
	// The node type.
	//
	// example:
	//
	// java
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The advanced configuration. This configuration is available only for parallel computing, in-memory grid, and grid computing modes.
	MapTaskXAttrs *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs `json:"MapTaskXAttrs,omitempty" xml:"MapTaskXAttrs,omitempty" type:"Struct"`
	// The maximum number of retries on failure. Set this parameter based on your business requirements. Default value: 0.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrently running instances. Default value: 1. A value of 1 indicates that if the previous trigger has not finished running, the next trigger is skipped even if the scheduled time has arrived.
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
	// The node status. Valid values:
	//
	// - **1**: Enabled. The node can be triggered normally.
	//
	// - **0**: Disabled. The node is not triggered.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time configuration information.
	TimeConfig *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig `json:"TimeConfig,omitempty" xml:"TimeConfig,omitempty" type:"Struct"`
	// The extended fields of the node.
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

func (s *GetJobInfoResponseBodyDataJobConfigInfo) GetEndTime() *int64 {
	return s.EndTime
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

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetEndTime(v int64) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.EndTime = &v
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

type GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo struct {
	// The contact information.
	ContactInfo []*GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// The alert switch and threshold configuration.
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

type GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo struct {
	// The webhook URL of DingTalk.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=XXXXXX
	Ding *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// user@demo.com
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
	// Specifies whether to enable the alert for no available machines.
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
	// The timeout threshold. Unit: seconds. Default value: 7200.
	//
	// example:
	//
	// 12300
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to enable the timeout alert. Valid values:
	//
	// - **true**: Enables the timeout alert.
	//
	// - **false**: Disables the timeout alert.
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Specifies whether to terminate the current trigger upon timeout. This feature is disabled by default.
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
	// The number of threads for a single trigger on a single machine. Default value: 5.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The number of threads for subtask distribution. Default value: 5.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The number of subtasks pulled per request for parallel nodes. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The maximum number of subtasks that can be cached in the queue. Default value: 10000.
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
	// The maximum number of retries for a subtask on failure.
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
	// The custom calendar for the **cron*	- type. This parameter is optional.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The time offset for the **cron*	- type. Unit: seconds.
	//
	// example:
	//
	// 0
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The time expression. The following time expression types are supported:
	//
	// - **api**: No time expression is required.
	//
	// - **fix_rate**: A fixed frequency value. For example, 30 indicates that the node is triggered every 30 seconds.
	//
	// - **cron**: A standard cron expression.
	//
	// - **second_delay**: A fixed delay in seconds before each execution (valid range: 1s to 60s).
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
	// - **5**: one_time
	//
	// - **100**: api
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
