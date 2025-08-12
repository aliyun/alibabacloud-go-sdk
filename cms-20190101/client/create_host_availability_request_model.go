// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostAvailabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertConfig(v *CreateHostAvailabilityRequestAlertConfig) *CreateHostAvailabilityRequest
	GetAlertConfig() *CreateHostAvailabilityRequestAlertConfig
	SetTaskOption(v *CreateHostAvailabilityRequestTaskOption) *CreateHostAvailabilityRequest
	GetTaskOption() *CreateHostAvailabilityRequestTaskOption
	SetAlertConfigEscalationList(v []*CreateHostAvailabilityRequestAlertConfigEscalationList) *CreateHostAvailabilityRequest
	GetAlertConfigEscalationList() []*CreateHostAvailabilityRequestAlertConfigEscalationList
	SetAlertConfigTargetList(v []*CreateHostAvailabilityRequestAlertConfigTargetList) *CreateHostAvailabilityRequest
	GetAlertConfigTargetList() []*CreateHostAvailabilityRequestAlertConfigTargetList
	SetGroupId(v int64) *CreateHostAvailabilityRequest
	GetGroupId() *int64
	SetInstanceList(v []*string) *CreateHostAvailabilityRequest
	GetInstanceList() []*string
	SetRegionId(v string) *CreateHostAvailabilityRequest
	GetRegionId() *string
	SetTaskName(v string) *CreateHostAvailabilityRequest
	GetTaskName() *string
	SetTaskScope(v string) *CreateHostAvailabilityRequest
	GetTaskScope() *string
	SetTaskType(v string) *CreateHostAvailabilityRequest
	GetTaskType() *string
}

type CreateHostAvailabilityRequest struct {
	AlertConfig *CreateHostAvailabilityRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	TaskOption  *CreateHostAvailabilityRequestTaskOption  `json:"TaskOption,omitempty" xml:"TaskOption,omitempty" type:"Struct"`
	// None
	//
	// This parameter is required.
	AlertConfigEscalationList []*CreateHostAvailabilityRequestAlertConfigEscalationList `json:"AlertConfigEscalationList,omitempty" xml:"AlertConfigEscalationList,omitempty" type:"Repeated"`
	// The resources for which alerts are triggered.
	AlertConfigTargetList []*CreateHostAvailabilityRequestAlertConfigTargetList `json:"AlertConfigTargetList,omitempty" xml:"AlertConfigTargetList,omitempty" type:"Repeated"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ECS instances that are monitored. Valid values of N: 1 to 21.
	//
	// > This parameter must be specified when `TaskScope` is set to `GROUP_SPEC_INSTANCE`.
	//
	// example:
	//
	// i-absdfkwl321****
	InstanceList []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the availability monitoring task. The name must be 4 to 100 characters in length, and can contain letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The range of instances that are monitored by the availability monitoring task. Valid values:
	//
	// 	- GROUP: All ECS instances in the application group are monitored.
	//
	// 	- GROUP_SPEC_INSTANCE: Specified ECS instances in the application group are monitored. The TaskScope parameter must be used in combination with the InstanceList parameter. The InstanceList parameter specifies the ECS instances to be monitored.
	//
	// example:
	//
	// GROUP
	TaskScope *string `json:"TaskScope,omitempty" xml:"TaskScope,omitempty"`
	// The monitoring type of the availability monitoring task. Valid values:
	//
	// 	- PING
	//
	// 	- TELNET
	//
	// 	- HTTP
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateHostAvailabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityRequest) GetAlertConfig() *CreateHostAvailabilityRequestAlertConfig {
	return s.AlertConfig
}

func (s *CreateHostAvailabilityRequest) GetTaskOption() *CreateHostAvailabilityRequestTaskOption {
	return s.TaskOption
}

func (s *CreateHostAvailabilityRequest) GetAlertConfigEscalationList() []*CreateHostAvailabilityRequestAlertConfigEscalationList {
	return s.AlertConfigEscalationList
}

func (s *CreateHostAvailabilityRequest) GetAlertConfigTargetList() []*CreateHostAvailabilityRequestAlertConfigTargetList {
	return s.AlertConfigTargetList
}

func (s *CreateHostAvailabilityRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateHostAvailabilityRequest) GetInstanceList() []*string {
	return s.InstanceList
}

func (s *CreateHostAvailabilityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHostAvailabilityRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateHostAvailabilityRequest) GetTaskScope() *string {
	return s.TaskScope
}

func (s *CreateHostAvailabilityRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateHostAvailabilityRequest) SetAlertConfig(v *CreateHostAvailabilityRequestAlertConfig) *CreateHostAvailabilityRequest {
	s.AlertConfig = v
	return s
}

func (s *CreateHostAvailabilityRequest) SetTaskOption(v *CreateHostAvailabilityRequestTaskOption) *CreateHostAvailabilityRequest {
	s.TaskOption = v
	return s
}

func (s *CreateHostAvailabilityRequest) SetAlertConfigEscalationList(v []*CreateHostAvailabilityRequestAlertConfigEscalationList) *CreateHostAvailabilityRequest {
	s.AlertConfigEscalationList = v
	return s
}

func (s *CreateHostAvailabilityRequest) SetAlertConfigTargetList(v []*CreateHostAvailabilityRequestAlertConfigTargetList) *CreateHostAvailabilityRequest {
	s.AlertConfigTargetList = v
	return s
}

func (s *CreateHostAvailabilityRequest) SetGroupId(v int64) *CreateHostAvailabilityRequest {
	s.GroupId = &v
	return s
}

func (s *CreateHostAvailabilityRequest) SetInstanceList(v []*string) *CreateHostAvailabilityRequest {
	s.InstanceList = v
	return s
}

func (s *CreateHostAvailabilityRequest) SetRegionId(v string) *CreateHostAvailabilityRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHostAvailabilityRequest) SetTaskName(v string) *CreateHostAvailabilityRequest {
	s.TaskName = &v
	return s
}

func (s *CreateHostAvailabilityRequest) SetTaskScope(v string) *CreateHostAvailabilityRequest {
	s.TaskScope = &v
	return s
}

func (s *CreateHostAvailabilityRequest) SetTaskType(v string) *CreateHostAvailabilityRequest {
	s.TaskType = &v
	return s
}

func (s *CreateHostAvailabilityRequest) Validate() error {
	return dara.Validate(s)
}

type CreateHostAvailabilityRequestAlertConfig struct {
	// The end of the time range during which the alert rule is effective. Valid values: 0 to 23.
	//
	// For example, if the `AlertConfig.StartTime` parameter is set to 0 and the `AlertConfig.EndTime` parameter is set to 22, the alert rule is effective from 00:00:00 to 22:00:00.
	//
	// > Alert notifications are sent based on the specified threshold only if the alert rule is effective.
	//
	// example:
	//
	// 22
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The alert notification methods. Valid values:
	//
	// 0: Alert notifications are sent by using emails and DingTalk chatbots.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	NotifyType *int32 `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The mute period during which new alerts are not sent even if the trigger conditions are met. Unit: seconds. Default value: 86400. The default value indicates one day.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The beginning of the time range during which the alert rule is effective. Valid values: 0 to 23.
	//
	// For example, if the `AlertConfig.StartTime` parameter is set to 0 and the `AlertConfig.EndTime` parameter is set to 22, the alert rule is effective from 00:00:00 to 22:00:00.
	//
	// > Alert notifications are sent based on the specified threshold only if the alert rule is effective.
	//
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// https://www.aliyun.com/webhook.json
	WebHook *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s CreateHostAvailabilityRequestAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAvailabilityRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityRequestAlertConfig) GetEndTime() *int32 {
	return s.EndTime
}

func (s *CreateHostAvailabilityRequestAlertConfig) GetNotifyType() *int32 {
	return s.NotifyType
}

func (s *CreateHostAvailabilityRequestAlertConfig) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *CreateHostAvailabilityRequestAlertConfig) GetStartTime() *int32 {
	return s.StartTime
}

func (s *CreateHostAvailabilityRequestAlertConfig) GetWebHook() *string {
	return s.WebHook
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetEndTime(v int32) *CreateHostAvailabilityRequestAlertConfig {
	s.EndTime = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetNotifyType(v int32) *CreateHostAvailabilityRequestAlertConfig {
	s.NotifyType = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetSilenceTime(v int32) *CreateHostAvailabilityRequestAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetStartTime(v int32) *CreateHostAvailabilityRequestAlertConfig {
	s.StartTime = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetWebHook(v string) *CreateHostAvailabilityRequestAlertConfig {
	s.WebHook = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfig) Validate() error {
	return dara.Validate(s)
}

type CreateHostAvailabilityRequestTaskOption struct {
	// The header of the HTTP request. Format: `Parameter name:Parameter value`. Separate multiple parameters with carriage return characters. Example:
	//
	//     params1:value1
	//
	//     params2:value2
	//
	// example:
	//
	// token:testTokenValue
	HttpHeader *string `json:"HttpHeader,omitempty" xml:"HttpHeader,omitempty"`
	// The HTTP request method. Valid values:
	//
	// 	- GET
	//
	// 	- POST
	//
	// 	- HEAD
	//
	// > This parameter must be specified when TaskType is set to HTTP.
	//
	// example:
	//
	// GET
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The method to trigger an alert. The alert can be triggered based on whether the specified alert rule is included in the response body. Valid values:
	//
	// 	- true: If the HTTP response body includes the alert rule, an alert is triggered.
	//
	// 	- false: If the HTTP response does not include the alert rule, an alert is triggered.
	//
	// > This parameter must be specified when TaskType is set to HTTP.
	//
	// example:
	//
	// true
	HttpNegative *bool `json:"HttpNegative,omitempty" xml:"HttpNegative,omitempty"`
	// The content of the HTTP POST request.
	//
	// example:
	//
	// params1=paramsValue1
	HttpPostContent *string `json:"HttpPostContent,omitempty" xml:"HttpPostContent,omitempty"`
	// The character set that is used in the HTTP response.
	//
	// > Only UTF-8 is supported.
	//
	// example:
	//
	// UTF-8
	HttpResponseCharset *string `json:"HttpResponseCharset,omitempty" xml:"HttpResponseCharset,omitempty"`
	// The response to the HTTP request.
	//
	// example:
	//
	// ok
	HttpResponseMatchContent *string `json:"HttpResponseMatchContent,omitempty" xml:"HttpResponseMatchContent,omitempty"`
	// The URI that you want to monitor. This parameter is required if the TaskType parameter is set to HTTP or Telnet.
	//
	// example:
	//
	// https://www.aliyun.com
	HttpURI *string `json:"HttpURI,omitempty" xml:"HttpURI,omitempty"`
	// The interval at which detection requests are sent. Unit: seconds. Valid values: 15, 30, 60, 120, 300, 900, 1800, and 3600.
	//
	// > This parameter is available only for the CloudMonitor agent V3.5.1 or later.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The domain name or IP address that you want to monitor.
	//
	// >  This parameter is required if the TaskType parameter is set to PING.
	//
	// example:
	//
	// www.aliyun.com
	TelnetOrPingHost *string `json:"TelnetOrPingHost,omitempty" xml:"TelnetOrPingHost,omitempty"`
}

func (s CreateHostAvailabilityRequestTaskOption) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAvailabilityRequestTaskOption) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityRequestTaskOption) GetHttpHeader() *string {
	return s.HttpHeader
}

func (s *CreateHostAvailabilityRequestTaskOption) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *CreateHostAvailabilityRequestTaskOption) GetHttpNegative() *bool {
	return s.HttpNegative
}

func (s *CreateHostAvailabilityRequestTaskOption) GetHttpPostContent() *string {
	return s.HttpPostContent
}

func (s *CreateHostAvailabilityRequestTaskOption) GetHttpResponseCharset() *string {
	return s.HttpResponseCharset
}

func (s *CreateHostAvailabilityRequestTaskOption) GetHttpResponseMatchContent() *string {
	return s.HttpResponseMatchContent
}

func (s *CreateHostAvailabilityRequestTaskOption) GetHttpURI() *string {
	return s.HttpURI
}

func (s *CreateHostAvailabilityRequestTaskOption) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateHostAvailabilityRequestTaskOption) GetTelnetOrPingHost() *string {
	return s.TelnetOrPingHost
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpHeader(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpHeader = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpMethod(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpMethod = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpNegative(v bool) *CreateHostAvailabilityRequestTaskOption {
	s.HttpNegative = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpPostContent(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpPostContent = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpResponseCharset(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpResponseCharset = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpResponseMatchContent(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpResponseMatchContent = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpURI(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpURI = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetInterval(v int32) *CreateHostAvailabilityRequestTaskOption {
	s.Interval = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetTelnetOrPingHost(v string) *CreateHostAvailabilityRequestTaskOption {
	s.TelnetOrPingHost = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) Validate() error {
	return dara.Validate(s)
}

type CreateHostAvailabilityRequestAlertConfigEscalationList struct {
	// The method used to calculate the metric values that trigger alerts. Valid values of N: 1 to 21. Valid values:
	//
	// 	- HttpStatus: Value
	//
	// 	- HttpLatency: Average
	//
	// 	- TelnetStatus: Value
	//
	// 	- TelnetLatency: Average
	//
	// 	- PingLostRate: Average
	//
	// > The value Value indicates the original value and is used for metrics such as status codes. The value Average indicates the average value and is used for metrics such as the latency and packet loss rate.
	//
	// example:
	//
	// Value
	Aggregate *string `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
	// The metric for which the alert feature is enabled. Valid values of N: 1 to 21. Valid values:
	//
	// 	- HttpStatus: HTTP status code
	//
	// 	- HttpLatency: HTTP response time
	//
	// 	- TelnetStatus: Telnet status code
	//
	// 	- TelnetLatency: Telnet response time
	//
	// 	- PingLostRate: Ping packet loss rate
	//
	// This parameter is required.
	//
	// example:
	//
	// HttpStatus
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The comparison operator that is used in the alert rule. Valid values of N: 1 to 21. Valid values:
	//
	// 	- `>`
	//
	// 	- `>=`
	//
	// 	- `<`
	//
	// 	- `<=`
	//
	// 	- `=`
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before an alert is triggered. Valid values of N: 1 to 21.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
	// The alert threshold. Valid values of N: 1 to 21.
	//
	// example:
	//
	// 90
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHostAvailabilityRequestAlertConfigEscalationList) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAvailabilityRequestAlertConfigEscalationList) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) GetAggregate() *string {
	return s.Aggregate
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) GetOperator() *string {
	return s.Operator
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) GetTimes() *int32 {
	return s.Times
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) GetValue() *string {
	return s.Value
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetAggregate(v string) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.Aggregate = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetMetricName(v string) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.MetricName = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetOperator(v string) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.Operator = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetTimes(v int32) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.Times = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetValue(v string) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.Value = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) Validate() error {
	return dara.Validate(s)
}

type CreateHostAvailabilityRequestAlertConfigTargetList struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource. Format: `acs:{Service name abbreviation}:{regionId}:{userId}:/{Resource type}/{Resource name}/message`. Example: `acs:mns:cn-hangzhou:120886317861****:/queues/test123/message`. Fields:
	//
	// 	- {Service name abbreviation}: the abbreviation of the service name. Set the value to Simple Message Queue (formerly MNS) (SMQ).
	//
	// 	- {userId}: the ID of the Alibaba Cloud account.
	//
	// 	- {regionId}: the region ID of the SMQ queue or topic.
	//
	// 	- {Resource type}: the type of the resource for which alerts are triggered. Valid values:
	//
	//     	- **queues**
	//
	//     	- **topics**
	//
	// 	- {Resource name}: the resource name.
	//
	//     	- If the resource type is **queues**, the resource name is the queue name.
	//
	//     	- If the resource type is **topics**, the resource name is the topic name.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:120886317861****:/queues/test/message
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the resource for which alerts are triggered.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters of the alert callback. The parameters are in the JSON format.
	//
	// example:
	//
	// {"customField1":"value1","customField2":"$.name"}
	JsonParams *string `json:"JsonParams,omitempty" xml:"JsonParams,omitempty"`
	// The alert level. Valid values:
	//
	// 	- INFO
	//
	// 	- WARN
	//
	// 	- CRITICAL
	//
	// example:
	//
	// ["INFO", "WARN", "CRITICAL"]
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s CreateHostAvailabilityRequestAlertConfigTargetList) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAvailabilityRequestAlertConfigTargetList) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityRequestAlertConfigTargetList) GetArn() *string {
	return s.Arn
}

func (s *CreateHostAvailabilityRequestAlertConfigTargetList) GetId() *string {
	return s.Id
}

func (s *CreateHostAvailabilityRequestAlertConfigTargetList) GetJsonParams() *string {
	return s.JsonParams
}

func (s *CreateHostAvailabilityRequestAlertConfigTargetList) GetLevel() *string {
	return s.Level
}

func (s *CreateHostAvailabilityRequestAlertConfigTargetList) SetArn(v string) *CreateHostAvailabilityRequestAlertConfigTargetList {
	s.Arn = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigTargetList) SetId(v string) *CreateHostAvailabilityRequestAlertConfigTargetList {
	s.Id = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigTargetList) SetJsonParams(v string) *CreateHostAvailabilityRequestAlertConfigTargetList {
	s.JsonParams = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigTargetList) SetLevel(v string) *CreateHostAvailabilityRequestAlertConfigTargetList {
	s.Level = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigTargetList) Validate() error {
	return dara.Validate(s)
}
