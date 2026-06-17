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
	// None.
	//
	// This parameter is required.
	AlertConfigEscalationList []*CreateHostAvailabilityRequestAlertConfigEscalationList `json:"AlertConfigEscalationList,omitempty" xml:"AlertConfigEscalationList,omitempty" type:"Repeated"`
	// The alert trigger targets.
	AlertConfigTargetList []*CreateHostAvailabilityRequestAlertConfigTargetList `json:"AlertConfigTargetList,omitempty" xml:"AlertConfigTargetList,omitempty" type:"Repeated"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The list of ECS instances that initiate detection. Valid values of N: 1 to 21.
	//
	// > Set this parameter when `TaskScope` is set to `GROUP_SPEC_INSTANCE`.
	//
	// example:
	//
	// i-absdfkwl321****
	InstanceList []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the availability monitoring task. The name must be 4 to 100 characters in length and can contain letters, digits, underscores (_), and Chinese characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
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
	// The detection type of the availability monitoring task. Valid values:
	//
	// - PING
	//
	// - TELNET
	//
	// - HTTP.
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

type CreateHostAvailabilityRequestAlertConfig struct {
	// 报警生效的结束时间。取值范围：0~23。
	//
	// 例如：`AlertConfig.StartTime`为0，`AlertConfig.EndTime`为22，表示报警生效时间为00:00:00至22:00:00。
	//
	// >如果报警不在生效时间内，则超过阈值也不会发送报警通知。
	//
	// example:
	//
	// 22
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 报警通知类型。取值：
	//
	// <props="china">- 2：电话+短信+邮件+钉钉机器人。
	//
	// <props="china">- 1：短信+邮件+钉钉机器人。
	//
	// <props="china">- 0：邮件+钉钉机器人。
	//
	//
	// <props="intl">0：邮件+钉钉机器人。
	//
	// <props="partner">0：邮件+钉钉机器人。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	NotifyType *int32 `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// 通道沉默时间。单位：秒，默认值：86400（1天）。
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// 报警生效的开始时间。取值范围：0~23。
	//
	// 例如：`AlertConfig.StartTime`为0，`AlertConfig.EndTime`为22，表示报警生效时间为00:00:00至22:00:00。
	//
	// >如果报警不在生效时间内，则超过阈值也不会发送报警通知。
	//
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// URL回调地址。
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
	// HTTP请求的Header。格式为`参数名:参数`，多个参数之间用回车符分隔，例如：
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
	// token:testTokenValue
	HttpHeader *string `json:"HttpHeader,omitempty" xml:"HttpHeader,omitempty"`
	// 探测类型的方法。取值：
	//
	// - GET
	//
	// - POST
	//
	// - HEAD
	//
	// >如果任务的探测类型为HTTP，则需要设置该参数。
	//
	// example:
	//
	// GET
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// 匹配HTTP响应内容的报警规则。取值：
	//
	// - true：如果HTTP响应内容包含设置的报警规则，则报警。
	//
	// - false：如果HTTP响应内容不包含设置的报警规则，则报警。
	//
	// >如果任务的探测类型为HTTP，则该参数生效。
	//
	// example:
	//
	// true
	HttpNegative *bool `json:"HttpNegative,omitempty" xml:"HttpNegative,omitempty"`
	// HTTP探测类型探测请求的Post内容。
	//
	// example:
	//
	// params1=paramsValue1
	HttpPostContent *string `json:"HttpPostContent,omitempty" xml:"HttpPostContent,omitempty"`
	// HTTP探测类型的响应字符集。
	//
	// > 仅支持UTF-8。
	//
	// example:
	//
	// UTF-8
	HttpResponseCharset *string `json:"HttpResponseCharset,omitempty" xml:"HttpResponseCharset,omitempty"`
	// 匹配响应的内容。
	//
	// example:
	//
	// ok
	HttpResponseMatchContent *string `json:"HttpResponseMatchContent,omitempty" xml:"HttpResponseMatchContent,omitempty"`
	// HTTP、Telnet探测类型的探测URI地址。
	//
	// example:
	//
	// https://www.aliyun.com
	//
	// telnet://127.0.0.1:80
	HttpURI *string `json:"HttpURI,omitempty" xml:"HttpURI,omitempty"`
	// 探测频率。单位：秒。取值：15、30、60、120、300、900、1800和3600。
	//
	// > 仅3.5.1及以上版本的云监控插件支持该参数。
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// 探测的域名或地址。
	//
	// >如果探测任务类型为PING，则需要设置该参数。
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
	// The statistical method for the alert. Valid values of N: 1 to 21. The valid values vary based on the metric:
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
	// > The statistical method for status code metrics is the raw value (Value). The statistical method for latency or packet loss rate metrics is the average value (Average).
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
	// - TelnetStatus: Telnet status code.
	//
	// - TelnetLatency: Telnet latency.
	//
	// - PingLostRate: Ping packet loss rate.
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
	// - `=`.
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
	// The Alibaba Cloud Resource Name (ARN) of the resource. Format: `acs:{AbbreviatedServiceName}:{regionId}:{userId}:/{ResourceType}/{ResourceName}/message`. Example: `acs:mns:ap-southeast-1:120886317861****:/queues/test123/message`. The following list describes the parameters:
	//
	// - {AbbreviatedServiceName}: Only Simple Message Queue (formerly MNS) is supported.
	//
	// - {userId}: The Alibaba Cloud account ID.
	//
	// - {regionId}: The region where the Simple Message Queue (formerly MNS) queue or topic resides.
	//
	// - {ResourceType}: The type of the resource that accepts alerts. Valid values:
	//
	//   - **queues**: queue.
	//
	//   - **topics**: topic.
	//
	// - {ResourceName}: The name of the resource.
	//
	//   - If the resource type is **queues**, the resource name is the queue name.
	//
	//   - If the resource type is **topics**, the resource name is the topic name.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:120886317861****:/queues/test/message
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
