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
	// The alert configuration list.
	//
	// This parameter is required.
	AlertConfigEscalationList []*ModifyHostAvailabilityRequestAlertConfigEscalationList `json:"AlertConfigEscalationList,omitempty" xml:"AlertConfigEscalationList,omitempty" type:"Repeated"`
	// The alert trigger targets.
	AlertConfigTargetList []*ModifyHostAvailabilityRequestAlertConfigTargetList `json:"AlertConfigTargetList,omitempty" xml:"AlertConfigTargetList,omitempty" type:"Repeated"`
	// The application group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The availability monitoring task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of ECS instances that initiate detection. Valid values of N: 1 to 21.
	//
	// > If `TaskScope` is set to `GROUP_SPEC_INSTANCE`, you must set this parameter.
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
	// The detection scope of the availability monitoring task. Valid values:
	//
	// - GROUP: uses all ECS instances in the current application group as detection probes.
	//
	// - GROUP_SPEC_INSTANCE: uses specified ECS instances in the current application group as detection probes. If you set this parameter to GROUP_SPEC_INSTANCE, you must also set InstanceList to specify the ECS instances that initiate detection.
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
	if s.AlertConfig != nil {
		if err := s.AlertConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TaskOption != nil {
		if err := s.TaskOption.Validate(); err != nil {
			return err
		}
	}
	if s.AlertConfigEscalationList != nil {
		for _, item := range s.AlertConfigEscalationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AlertConfigTargetList != nil {
		for _, item := range s.AlertConfigTargetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyHostAvailabilityRequestAlertConfig struct {
	// The end time when the alert takes effect.
	//
	// Unit: hours (h). Valid values: 0 to 23. For example, a value of 22 indicates 22:00:00.
	//
	// > If the alert is not within the effective period, alert notifications are not sent even if the threshold is exceeded.
	//
	// example:
	//
	// 22
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The alert notification type. Valid values:
	//
	// <props="china">- 2: phone call + text message + email + DingTalk chatbot.
	//
	// <props="china">- 1: text message + email + DingTalk chatbot.
	//
	// <props="china">- 0: email + DingTalk chatbot.
	//
	//
	// <props="intl">0: email + DingTalk chatbot.
	//
	// <props="partner">0: email + DingTalk chatbot.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	NotifyType *int32 `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The mute period. Unit: seconds. Default value: 86400 (1 day).
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The start time when the alert takes effect.
	//
	// Unit: hours (h). Valid values: 0 to 23. For example, a value of 0 indicates 00:00:00.
	//
	// > If the alert is not within the effective period, alert notifications are not sent even if the threshold is exceeded.
	//
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The webhook address.
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
	// The header of the HTTP request. Format: `Parameter:Value`. Separate multiple parameters with carriage returns. Example:
	//
	// ```
	//
	// params1:value1
	//
	// params2:value2
	//
	// ```
	//
	// example:
	//
	// params1:value1
	HttpHeader *string `json:"HttpHeader,omitempty" xml:"HttpHeader,omitempty"`
	// The method of the detection type. Valid values:
	//
	// - GET
	//
	// - POST
	//
	// - HEAD
	//
	// > If the detection task type is HTTP, you must set this parameter. For information about how to set the detection task type, see [CreateHostAvailability](https://help.aliyun.com/document_detail/115317.html).
	//
	// example:
	//
	// GET
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The alert rule for matching HTTP response content. Valid values:
	//
	// - true: An alert is triggered if the HTTP response content contains the specified alert rule.
	//
	// - false: An alert is triggered if the HTTP response content does not contain the specified alert rule.
	//
	// > If the detection task type is HTTP, you must set this parameter. For information about how to set the detection task type, see [CreateHostAvailability](https://help.aliyun.com/document_detail/115317.html).
	//
	// example:
	//
	// true
	HttpNegative *bool `json:"HttpNegative,omitempty" xml:"HttpNegative,omitempty"`
	// The POST content of the HTTP detection request.
	//
	// example:
	//
	// params1=value1
	HttpPostContent *string `json:"HttpPostContent,omitempty" xml:"HttpPostContent,omitempty"`
	// The response character set for the HTTP detection type.
	//
	// > Only UTF-8 is supported.
	//
	// example:
	//
	// UTF-8
	HttpResponseCharset *string `json:"HttpResponseCharset,omitempty" xml:"HttpResponseCharset,omitempty"`
	// The POST content of the HTTP detection response.
	//
	// example:
	//
	// ok
	HttpResponseMatchContent *string `json:"HttpResponseMatchContent,omitempty" xml:"HttpResponseMatchContent,omitempty"`
	// The URI for HTTP or Telnet detection types.
	//
	// example:
	//
	// https://www.aliyun.com
	//
	// telnet://127.0.0.1:80
	HttpURI *string `json:"HttpURI,omitempty" xml:"HttpURI,omitempty"`
	// The detection frequency. Unit: seconds. Valid values: 15, 30, 60, 120, 300, 900, 1800, and 3600.
	//
	// > Only CloudMonitor agent 3.5.1 or later supports this parameter.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The domain name or address to detect.
	//
	// > If the detection task type is PING, you must set this parameter. For information about how to set the detection task type, see [CreateHostAvailability](https://help.aliyun.com/document_detail/115317.html).
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
	// The statistical method for the alert. Valid values of N: 1 to 21. The statistical method varies by metric:
	//
	// - HttpStatus: Value.
	//
	// - HttpLatency: Average.
	//
	// - TelnetStatus: Value.
	//
	// - TelnetLatency: Average.
	//
	// - PingLostRate: Average.
	//
	// > The statistical method for status codes is the original value (Value). The statistical method for latency or packet loss rate is the average value (Average).
	//
	// example:
	//
	// Value
	Aggregate *string `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
	// The metric for the alert. Valid values of N: 1 to 21. Valid values:
	//
	// - HttpStatus: HTTP status code.
	//
	// - HttpLatency: HTTP latency.
	//
	// - TelnetStatus: TELNET status code.
	//
	// - TelnetLatency: TELNET latency.
	//
	// - PingLostRate: PING packet loss rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// HttpStatus
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The comparison operator for the alert rule. Valid values of N: 1 to 21. Valid values:
	//
	// - `>`
	//
	// - `>=`
	//
	// - `<`
	//
	// - `<=`
	//
	// - `=`
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The number of alert retries. Valid values of N: 1 to 21.
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
	// The resource ARN.
	//
	// Format: `acs:{AbbreviatedCloudServiceName}:{regionId}:{userId}:/{ResourceType}/{ResourceName}/message`. Example: `acs:mns:cn-hangzhou:120886317861****:/queues/test123/message`. The parameters are described as follows:
	//
	// - {AbbreviatedCloudServiceName}: Only Simple Message Queue (formerly MNS) is supported.
	//
	// - {userId}: The Alibaba Cloud account ID.
	//
	// - {regionId}: The region where Simple Message Queue (formerly MNS) or the topic resides.
	//
	// - {ResourceType}: The type of the resource that accepts alerts. Valid values:
	//
	//   - **queues**: queue.
	//
	//   - **topics**: topic.
	//
	// - {ResourceName}: The resource name.
	//
	//   - If the resource type is **queues**, the resource name is the queue name.
	//
	//   - If the resource type is **topics**, the resource name is the topic name.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:111:/queues/test/message
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the alert trigger target.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-formatted parameters for the alert callback.
	//
	// example:
	//
	// {"customField1":"value1","customField2":"$.name"}
	JsonParams *string `json:"JsonParams,omitempty" xml:"JsonParams,omitempty"`
	// The alert level. Valid values:
	//
	// - INFO: information.
	//
	// - WARN: warning.
	//
	// - CRITICAL: critical.
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
