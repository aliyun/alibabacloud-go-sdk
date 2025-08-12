// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupMonitoringAgentProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertConfig(v []*ModifyGroupMonitoringAgentProcessRequestAlertConfig) *ModifyGroupMonitoringAgentProcessRequest
	GetAlertConfig() []*ModifyGroupMonitoringAgentProcessRequestAlertConfig
	SetGroupId(v string) *ModifyGroupMonitoringAgentProcessRequest
	GetGroupId() *string
	SetId(v string) *ModifyGroupMonitoringAgentProcessRequest
	GetId() *string
	SetMatchExpressFilterRelation(v string) *ModifyGroupMonitoringAgentProcessRequest
	GetMatchExpressFilterRelation() *string
	SetRegionId(v string) *ModifyGroupMonitoringAgentProcessRequest
	GetRegionId() *string
}

type ModifyGroupMonitoringAgentProcessRequest struct {
	// The alert rule configurations.
	//
	// This parameter is required.
	AlertConfig []*ModifyGroupMonitoringAgentProcessRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6780****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the process monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92E3065F-0980-4E31-9AA0-BA6****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The logical operator used between conditional expressions that are used to match instances. Valid values:
	//
	// 	- all
	//
	// 	- and
	//
	// 	- or
	//
	// example:
	//
	// and
	MatchExpressFilterRelation *string `json:"MatchExpressFilterRelation,omitempty" xml:"MatchExpressFilterRelation,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyGroupMonitoringAgentProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *ModifyGroupMonitoringAgentProcessRequest) GetAlertConfig() []*ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	return s.AlertConfig
}

func (s *ModifyGroupMonitoringAgentProcessRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyGroupMonitoringAgentProcessRequest) GetId() *string {
	return s.Id
}

func (s *ModifyGroupMonitoringAgentProcessRequest) GetMatchExpressFilterRelation() *string {
	return s.MatchExpressFilterRelation
}

func (s *ModifyGroupMonitoringAgentProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGroupMonitoringAgentProcessRequest) SetAlertConfig(v []*ModifyGroupMonitoringAgentProcessRequestAlertConfig) *ModifyGroupMonitoringAgentProcessRequest {
	s.AlertConfig = v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequest) SetGroupId(v string) *ModifyGroupMonitoringAgentProcessRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequest) SetId(v string) *ModifyGroupMonitoringAgentProcessRequest {
	s.Id = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequest) SetMatchExpressFilterRelation(v string) *ModifyGroupMonitoringAgentProcessRequest {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequest) SetRegionId(v string) *ModifyGroupMonitoringAgentProcessRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyGroupMonitoringAgentProcessRequestAlertConfig struct {
	// The comparison operator that is used to compare the metric value with the threshold. Valid values of N: 1 to 200. Valid values:
	//
	// 	- GreaterThanOrEqualToThreshold: greater than or equal to the threshold
	//
	// 	- GreaterThanThreshold: greater than the threshold
	//
	// 	- LessThanOrEqualToThreshold: less than or equal to the threshold
	//
	// 	- LessThanThreshold: less than the threshold.
	//
	// 	- NotEqualToThreshold: not equal to the threshold
	//
	// 	- GreaterThanYesterday: greater than the metric value at the same time yesterday.
	//
	// 	- LessThanYesterday: less than the metric value at the same time yesterday
	//
	// 	- GreaterThanLastWeek: greater than the metric value at the same time last week
	//
	// 	- LessThanLastWeek: less than the metric value at the same time last week
	//
	// 	- GreaterThanLastPeriod: greater than the metric value in the last monitoring cycle
	//
	// 	- LessThanLastPeriod: less than the metric value in the last monitoring cycle
	//
	// This parameter is required.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The time period during which the alert rule is effective. Valid values of N: 1 to 200.
	//
	// example:
	//
	// 00:00-22:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The level of the alert. Valid values of N: 1 to 200. Valid values:
	//
	// 	- critical (default value): critical
	//
	// 	- warn: warning
	//
	// 	- info: information
	//
	// This parameter is required.
	//
	// example:
	//
	// warn
	EscalationsLevel *string `json:"EscalationsLevel,omitempty" xml:"EscalationsLevel,omitempty"`
	// The time period during which the alert rule is ineffective. Valid values of N: 1 to 200.
	//
	// example:
	//
	// 23:00-23:59
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The mute period during which new alerts are not sent even if the trigger conditions are met. Valid values of N: 1 to 200.
	//
	// Unit: seconds. Minimum value: 3600, which is equivalent to one hour. Default value: 86400, which is equivalent to one day.
	//
	// >  Only one alert notification is sent during a mute period even if the metric value exceeds the alert threshold during consecutive checks.
	//
	// example:
	//
	// 86400
	SilenceTime *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The statistical aggregation method that is used to calculate the metric values. Valid values of N: 1 to 200.
	//
	// >  Set the value to Average.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The alert trigger.
	TargetList []*ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The alert threshold. Valid values of N: 1 to 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of times for which the threshold can be consecutively exceeded. Valid values of N: 1 to 200. Default value: 3.
	//
	// >  A metric triggers an alert only after the metric value reaches the threshold consecutively for the specified times.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Times *string `json:"Times,omitempty" xml:"Times,omitempty"`
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule. Valid values of N: 1 to 200.
	//
	// example:
	//
	// http://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s ModifyGroupMonitoringAgentProcessRequestAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupMonitoringAgentProcessRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetEscalationsLevel() *string {
	return s.EscalationsLevel
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetSilenceTime() *string {
	return s.SilenceTime
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetStatistics() *string {
	return s.Statistics
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetTargetList() []*ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	return s.TargetList
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetThreshold() *string {
	return s.Threshold
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetTimes() *string {
	return s.Times
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) GetWebhook() *string {
	return s.Webhook
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetComparisonOperator(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetEffectiveInterval(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.EffectiveInterval = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetEscalationsLevel(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.EscalationsLevel = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetNoEffectiveInterval(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.NoEffectiveInterval = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetSilenceTime(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetStatistics(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.Statistics = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetTargetList(v []*ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.TargetList = v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetThreshold(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.Threshold = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetTimes(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.Times = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetWebhook(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.Webhook = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	//
	// For information about how to obtain the ARN of a resource, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
	//
	// Format: `acs:{Service name abbreviation}:{regionId}:{userId}:/{Resource type}/{Resource name}/message`. Example: `acs:mns:cn-hangzhou:120886317861****:/queues/test123/message`. Fields:
	//
	// - {Service name abbreviation}: the abbreviation of the service name. Valid value: mns.
	//
	// - {userId}: the ID of the Alibaba Cloud account.
	//
	// - {regionId}: the region ID of the message queue or topic.
	//
	// - {Resource type}: the type of the resource for which alerts are triggered. Valid values:
	//
	//     - **queues*	-
	//
	//     - **topics*	-
	//
	// - {Resourcename}: the name of the resource.
	//
	//   - If the resource type is set to **queues**, the resource name is the name of the message queue.
	//
	//   - If the resource type is set to **topics**, the resource name is the name of the topic.`
	//
	// example:
	//
	// acs:mns:cn-hangzhou:120886317861****:/queues/test/message
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the resource for which alerts are triggered.
	//
	// For information about how to obtain the ID of a resource for which alerts are triggered, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
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
	// The level of the alert. Valid values:
	//
	// 	- INFO: information
	//
	// 	- WARN: warning
	//
	// 	- CRITICAL: critical
	//
	// example:
	//
	// ["INFO", "WARN", "CRITICAL"]
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) GoString() string {
	return s.String()
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) GetArn() *string {
	return s.Arn
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) GetId() *string {
	return s.Id
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) GetJsonParams() *string {
	return s.JsonParams
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) GetLevel() *string {
	return s.Level
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) SetArn(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	s.Arn = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) SetId(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	s.Id = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) SetJsonParams(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	s.JsonParams = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) SetLevel(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	s.Level = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList) Validate() error {
	return dara.Validate(s)
}
