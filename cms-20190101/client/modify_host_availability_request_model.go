// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostAvailabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertConfig(v *ModifyHostAvailabilityRequestAlertConfig) *ModifyHostAvailabilityRequest
	GetAlertConfig() *ModifyHostAvailabilityRequestAlertConfig
	SetTaskOption(v *ModifyHostAvailabilityRequestTaskOption) *ModifyHostAvailabilityRequest
	GetTaskOption() *ModifyHostAvailabilityRequestTaskOption
	SetAlertConfigEscalationList(v []*ModifyHostAvailabilityRequestAlertConfigEscalationList) *ModifyHostAvailabilityRequest
	GetAlertConfigEscalationList() []*ModifyHostAvailabilityRequestAlertConfigEscalationList
	SetAlertConfigTargetList(v []*ModifyHostAvailabilityRequestAlertConfigTargetList) *ModifyHostAvailabilityRequest
	GetAlertConfigTargetList() []*ModifyHostAvailabilityRequestAlertConfigTargetList
	SetGroupId(v int64) *ModifyHostAvailabilityRequest
	GetGroupId() *int64
	SetId(v int64) *ModifyHostAvailabilityRequest
	GetId() *int64
	SetInstanceList(v []*string) *ModifyHostAvailabilityRequest
	GetInstanceList() []*string
	SetRegionId(v string) *ModifyHostAvailabilityRequest
	GetRegionId() *string
	SetTaskName(v string) *ModifyHostAvailabilityRequest
	GetTaskName() *string
	SetTaskScope(v string) *ModifyHostAvailabilityRequest
	GetTaskScope() *string
}

type ModifyHostAvailabilityRequest struct {
	AlertConfig *ModifyHostAvailabilityRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	TaskOption  *ModifyHostAvailabilityRequestTaskOption  `json:"TaskOption,omitempty" xml:"TaskOption,omitempty" type:"Struct"`
	// The alert configurations.
	//
	// This parameter is required.
	AlertConfigEscalationList []*ModifyHostAvailabilityRequestAlertConfigEscalationList `json:"AlertConfigEscalationList,omitempty" xml:"AlertConfigEscalationList,omitempty" type:"Repeated"`
	// The information about the resources for which alerts are triggered.
	AlertConfigTargetList []*ModifyHostAvailabilityRequestAlertConfigTargetList `json:"AlertConfigTargetList,omitempty" xml:"AlertConfigTargetList,omitempty" type:"Repeated"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the availability monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ECS instances that are monitored. Valid values of N: 1 to 21.
	//
	// > This parameter must be specified when `TaskScope` is set to `GROUP_SPEC_INSTANCE`.
	//
	// example:
	//
	// i-absdfkwl321****
	InstanceList []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the availability monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// task2
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
}

func (s ModifyHostAvailabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityRequest) GetAlertConfig() *ModifyHostAvailabilityRequestAlertConfig {
	return s.AlertConfig
}

func (s *ModifyHostAvailabilityRequest) GetTaskOption() *ModifyHostAvailabilityRequestTaskOption {
	return s.TaskOption
}

func (s *ModifyHostAvailabilityRequest) GetAlertConfigEscalationList() []*ModifyHostAvailabilityRequestAlertConfigEscalationList {
	return s.AlertConfigEscalationList
}

func (s *ModifyHostAvailabilityRequest) GetAlertConfigTargetList() []*ModifyHostAvailabilityRequestAlertConfigTargetList {
	return s.AlertConfigTargetList
}

func (s *ModifyHostAvailabilityRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ModifyHostAvailabilityRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyHostAvailabilityRequest) GetInstanceList() []*string {
	return s.InstanceList
}

func (s *ModifyHostAvailabilityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHostAvailabilityRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifyHostAvailabilityRequest) GetTaskScope() *string {
	return s.TaskScope
}

func (s *ModifyHostAvailabilityRequest) SetAlertConfig(v *ModifyHostAvailabilityRequestAlertConfig) *ModifyHostAvailabilityRequest {
	s.AlertConfig = v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetTaskOption(v *ModifyHostAvailabilityRequestTaskOption) *ModifyHostAvailabilityRequest {
	s.TaskOption = v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetAlertConfigEscalationList(v []*ModifyHostAvailabilityRequestAlertConfigEscalationList) *ModifyHostAvailabilityRequest {
	s.AlertConfigEscalationList = v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetAlertConfigTargetList(v []*ModifyHostAvailabilityRequestAlertConfigTargetList) *ModifyHostAvailabilityRequest {
	s.AlertConfigTargetList = v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetGroupId(v int64) *ModifyHostAvailabilityRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetId(v int64) *ModifyHostAvailabilityRequest {
	s.Id = &v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetInstanceList(v []*string) *ModifyHostAvailabilityRequest {
	s.InstanceList = v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetRegionId(v string) *ModifyHostAvailabilityRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetTaskName(v string) *ModifyHostAvailabilityRequest {
	s.TaskName = &v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetTaskScope(v string) *ModifyHostAvailabilityRequest {
	s.TaskScope = &v
	return s
}

func (s *ModifyHostAvailabilityRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyHostAvailabilityRequestAlertConfig struct {
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

func (s ModifyHostAvailabilityRequestAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAvailabilityRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityRequestAlertConfig) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ModifyHostAvailabilityRequestAlertConfig) GetNotifyType() *int32 {
	return s.NotifyType
}

func (s *ModifyHostAvailabilityRequestAlertConfig) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *ModifyHostAvailabilityRequestAlertConfig) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ModifyHostAvailabilityRequestAlertConfig) GetWebHook() *string {
	return s.WebHook
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetEndTime(v int32) *ModifyHostAvailabilityRequestAlertConfig {
	s.EndTime = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetNotifyType(v int32) *ModifyHostAvailabilityRequestAlertConfig {
	s.NotifyType = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetSilenceTime(v int32) *ModifyHostAvailabilityRequestAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetStartTime(v int32) *ModifyHostAvailabilityRequestAlertConfig {
	s.StartTime = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetWebHook(v string) *ModifyHostAvailabilityRequestAlertConfig {
	s.WebHook = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyHostAvailabilityRequestTaskOption struct {
	// The header of the HTTP request. Format: `Parameter name:Parameter value`. Separate multiple parameters with carriage return characters. Example:
	//
	//     params1:value1
	//
	//     params2:value2
	//
	// example:
	//
	// params1:value1
	HttpHeader *string `json:"HttpHeader,omitempty" xml:"HttpHeader,omitempty"`
	// The HTTP request method. Valid values:
	//
	// 	- GET
	//
	// 	- POST
	//
	// 	- HEAD
	//
	// > This parameter must be specified when TaskType is set to HTTP. For more information about how to configure the TaskType parameter, see [CreateHostAvailability](https://help.aliyun.com/document_detail/115317.html).
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
	// > This parameter must be specified when TaskType is set to HTTP. For more information about how to configure the TaskType parameter, see [CreateHostAvailability](https://help.aliyun.com/document_detail/115317.html).
	//
	// example:
	//
	// true
	HttpNegative *bool `json:"HttpNegative,omitempty" xml:"HttpNegative,omitempty"`
	// The content of the HTTP POST request.
	//
	// example:
	//
	// params1=value1
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
	// >  This parameter is required if the TaskType parameter is set to PING. For more information about how to set the TaskType parameter, see [CreateHostAvailability](https://help.aliyun.com/document_detail/115317.html).
	//
	// example:
	//
	// www.aliyun.com
	TelnetOrPingHost *string `json:"TelnetOrPingHost,omitempty" xml:"TelnetOrPingHost,omitempty"`
}

func (s ModifyHostAvailabilityRequestTaskOption) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAvailabilityRequestTaskOption) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityRequestTaskOption) GetHttpHeader() *string {
	return s.HttpHeader
}

func (s *ModifyHostAvailabilityRequestTaskOption) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *ModifyHostAvailabilityRequestTaskOption) GetHttpNegative() *bool {
	return s.HttpNegative
}

func (s *ModifyHostAvailabilityRequestTaskOption) GetHttpPostContent() *string {
	return s.HttpPostContent
}

func (s *ModifyHostAvailabilityRequestTaskOption) GetHttpResponseCharset() *string {
	return s.HttpResponseCharset
}

func (s *ModifyHostAvailabilityRequestTaskOption) GetHttpResponseMatchContent() *string {
	return s.HttpResponseMatchContent
}

func (s *ModifyHostAvailabilityRequestTaskOption) GetHttpURI() *string {
	return s.HttpURI
}

func (s *ModifyHostAvailabilityRequestTaskOption) GetInterval() *int32 {
	return s.Interval
}

func (s *ModifyHostAvailabilityRequestTaskOption) GetTelnetOrPingHost() *string {
	return s.TelnetOrPingHost
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpHeader(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpHeader = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpMethod(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpMethod = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpNegative(v bool) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpNegative = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpPostContent(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpPostContent = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpResponseCharset(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpResponseCharset = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpResponseMatchContent(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpResponseMatchContent = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpURI(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpURI = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetInterval(v int32) *ModifyHostAvailabilityRequestTaskOption {
	s.Interval = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetTelnetOrPingHost(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.TelnetOrPingHost = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) Validate() error {
	return dara.Validate(s)
}

type ModifyHostAvailabilityRequestAlertConfigEscalationList struct {
	// The method used to calculate the metric values that trigger alerts. Valid values of N: 1 to 21. The value of this parameter varies based on the metric. The following items show the correspondence between metrics and calculation methods:
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
	// 3
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyHostAvailabilityRequestAlertConfigEscalationList) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAvailabilityRequestAlertConfigEscalationList) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) GetAggregate() *string {
	return s.Aggregate
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) GetMetricName() *string {
	return s.MetricName
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) GetOperator() *string {
	return s.Operator
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) GetTimes() *int32 {
	return s.Times
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) GetValue() *string {
	return s.Value
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetAggregate(v string) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.Aggregate = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetMetricName(v string) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.MetricName = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetOperator(v string) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.Operator = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetTimes(v int32) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.Times = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetValue(v string) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.Value = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) Validate() error {
	return dara.Validate(s)
}

type ModifyHostAvailabilityRequestAlertConfigTargetList struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	//
	// Format: `acs:{Service name abbreviation}:{regionId}:{userId}:/{Resource type}/{Resource name}/message`. Example: `acs:mns:cn-hangzhou:120886317861****:/queues/test123/message`. Fields:
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
	// acs:mns:cn-hangzhou:111:/queues/test/message
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

func (s ModifyHostAvailabilityRequestAlertConfigTargetList) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAvailabilityRequestAlertConfigTargetList) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityRequestAlertConfigTargetList) GetArn() *string {
	return s.Arn
}

func (s *ModifyHostAvailabilityRequestAlertConfigTargetList) GetId() *string {
	return s.Id
}

func (s *ModifyHostAvailabilityRequestAlertConfigTargetList) GetJsonParams() *string {
	return s.JsonParams
}

func (s *ModifyHostAvailabilityRequestAlertConfigTargetList) GetLevel() *string {
	return s.Level
}

func (s *ModifyHostAvailabilityRequestAlertConfigTargetList) SetArn(v string) *ModifyHostAvailabilityRequestAlertConfigTargetList {
	s.Arn = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigTargetList) SetId(v string) *ModifyHostAvailabilityRequestAlertConfigTargetList {
	s.Id = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigTargetList) SetJsonParams(v string) *ModifyHostAvailabilityRequestAlertConfigTargetList {
	s.JsonParams = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigTargetList) SetLevel(v string) *ModifyHostAvailabilityRequestAlertConfigTargetList {
	s.Level = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigTargetList) Validate() error {
	return dara.Validate(s)
}
