// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupMonitoringAgentProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertConfig(v []*CreateGroupMonitoringAgentProcessRequestAlertConfig) *CreateGroupMonitoringAgentProcessRequest
	GetAlertConfig() []*CreateGroupMonitoringAgentProcessRequestAlertConfig
	SetGroupId(v string) *CreateGroupMonitoringAgentProcessRequest
	GetGroupId() *string
	SetMatchExpress(v []*CreateGroupMonitoringAgentProcessRequestMatchExpress) *CreateGroupMonitoringAgentProcessRequest
	GetMatchExpress() []*CreateGroupMonitoringAgentProcessRequestMatchExpress
	SetMatchExpressFilterRelation(v string) *CreateGroupMonitoringAgentProcessRequest
	GetMatchExpressFilterRelation() *string
	SetProcessName(v string) *CreateGroupMonitoringAgentProcessRequest
	GetProcessName() *string
	SetRegionId(v string) *CreateGroupMonitoringAgentProcessRequest
	GetRegionId() *string
}

type CreateGroupMonitoringAgentProcessRequest struct {
	// The alert rule configurations.
	//
	// Valid values of N: 1 to 3.
	//
	// This parameter is required.
	AlertConfig []*CreateGroupMonitoringAgentProcessRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	// The ID of the application group.
	//
	// For more information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The expressions used to match instances.
	//
	// Valid values of N: 1 to 3.
	MatchExpress []*CreateGroupMonitoringAgentProcessRequestMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
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
	// The process name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGroupMonitoringAgentProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessRequest) GetAlertConfig() []*CreateGroupMonitoringAgentProcessRequestAlertConfig {
	return s.AlertConfig
}

func (s *CreateGroupMonitoringAgentProcessRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupMonitoringAgentProcessRequest) GetMatchExpress() []*CreateGroupMonitoringAgentProcessRequestMatchExpress {
	return s.MatchExpress
}

func (s *CreateGroupMonitoringAgentProcessRequest) GetMatchExpressFilterRelation() *string {
	return s.MatchExpressFilterRelation
}

func (s *CreateGroupMonitoringAgentProcessRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *CreateGroupMonitoringAgentProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetAlertConfig(v []*CreateGroupMonitoringAgentProcessRequestAlertConfig) *CreateGroupMonitoringAgentProcessRequest {
	s.AlertConfig = v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetGroupId(v string) *CreateGroupMonitoringAgentProcessRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetMatchExpress(v []*CreateGroupMonitoringAgentProcessRequestMatchExpress) *CreateGroupMonitoringAgentProcessRequest {
	s.MatchExpress = v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetMatchExpressFilterRelation(v string) *CreateGroupMonitoringAgentProcessRequest {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetProcessName(v string) *CreateGroupMonitoringAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetRegionId(v string) *CreateGroupMonitoringAgentProcessRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMonitoringAgentProcessRequestAlertConfig struct {
	// The operator that is used to compare the metric value with the threshold. Valid values:
	//
	// 	- GreaterThanOrEqualToThreshold: greater than or equal to the threshold
	//
	// 	- GreaterThanThreshold: greater than the threshold
	//
	// 	- LessThanOrEqualToThreshold: less than or equal to the threshold
	//
	// 	- LessThanThreshold: less than the threshold
	//
	// 	- NotEqualToThreshold: not equal to the threshold
	//
	// 	- GreaterThanYesterday: greater than the metric value at the same time yesterday
	//
	// 	- LessThanYesterday: less than the metric value at the same time yesterday
	//
	// 	- GreaterThanLastWeek: greater than the metric value at the same time last week
	//
	// 	- LessThanLastWeek: less than the metric value at the same time last week
	//
	// 	- GreaterThanLastPeriod: greater than the metric value in the last monitoring cycle
	//
	// 	- LessThanLastPeriod: less than the metric value in the previous monitoring cycle
	//
	// Valid values of N: 1 to 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The period of time during which the alert rule is effective.
	//
	// Valid values of N: 1 to 3.
	//
	// example:
	//
	// 00:00-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The alert level. Valid values:
	//
	// 	- critical (default)
	//
	// 	- warn
	//
	// 	- info
	//
	// Valid values of N: 1 to 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// warn
	EscalationsLevel *string `json:"EscalationsLevel,omitempty" xml:"EscalationsLevel,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 00:00-23:59
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The mute period during which new alert notifications are not sent even if the trigger conditions are met. Unit: seconds. Minimum value: 3600, which is equivalent to one hour. Default value: 86400, which is equivalent to one day.
	//
	// Valid values of N: 1 to 3.
	//
	// >  Only one alert notification is sent during a mute period even if the metric value exceeds the alert threshold during consecutive checks.
	//
	// example:
	//
	// 86400
	SilenceTime *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The statistical aggregation method that is used to calculate the metric values.
	//
	// Valid values of N: 1 to 3.
	//
	// >  Set the value to Average.
	//
	// This parameter is required.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The alert triggers.
	TargetList []*CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The alert threshold.
	//
	// Valid values of N: 1 to 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of times for which the threshold can be consecutively exceeded. Default value: 3.
	//
	// Valid values of N: 1 to 3.
	//
	// >  A metric triggers an alert only after the metric value reaches the threshold consecutively for the specified times.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Times *string `json:"Times,omitempty" xml:"Times,omitempty"`
	// The callback URL.
	//
	// Valid values of N: 1 to 3.
	//
	// example:
	//
	// http://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s CreateGroupMonitoringAgentProcessRequestAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetEscalationsLevel() *string {
	return s.EscalationsLevel
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetSilenceTime() *string {
	return s.SilenceTime
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetTargetList() []*CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	return s.TargetList
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetTimes() *string {
	return s.Times
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) GetWebhook() *string {
	return s.Webhook
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetComparisonOperator(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetEffectiveInterval(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.EffectiveInterval = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetEscalationsLevel(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.EscalationsLevel = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetNoEffectiveInterval(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.NoEffectiveInterval = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetSilenceTime(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetStatistics(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.Statistics = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetTargetList(v []*CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.TargetList = v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetThreshold(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetTimes(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.Times = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetWebhook(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.Webhook = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	//
	// For more information about how to query the ARN of a resource, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
	//
	// Format: `acs:{Service name abbreviation}:{regionId}:{userId}:/{Resource type}/{Resource name}/message`. Example: `acs:mns:cn-hangzhou:120886317861****:/queues/test123/message`. Fields:
	//
	// 	- {Service name abbreviation}: the abbreviation of the service name. Set the value to Simple Message Queue (formerly MNS) (SMQ).
	//
	// 	- {userId}: the ID of the Alibaba Cloud account.
	//
	// 	- {regionId}: the region ID of the SMQ queue or topic.
	//
	// 	- {Resource type}: the type of the resource that triggers the alert. Valid values:
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
	// For more information about how to obtain the ID of a resource for which alerts are triggered, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
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

func (s CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) GetArn() *string {
	return s.Arn
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) GetId() *string {
	return s.Id
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) GetJsonParams() *string {
	return s.JsonParams
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) GetLevel() *string {
	return s.Level
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) SetArn(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	s.Arn = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) SetId(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	s.Id = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) SetJsonParams(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	s.JsonParams = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) SetLevel(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList {
	s.Level = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMonitoringAgentProcessRequestMatchExpress struct {
	// The matching condition. Valid values:
	//
	// 	- all (default value): matches all
	//
	// 	- startWith: starts with a prefix
	//
	// 	- endWith: ends with a suffix
	//
	// 	- contains: contains
	//
	// 	- notContains: does not contain
	//
	// 	- equals: equals
	//
	// Valid values of N: 1 to 3.
	//
	// example:
	//
	// startWith
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The criteria based on which the instances are matched.
	//
	// Valid values of N: 1 to 3.
	//
	// > Set the value to name. The value name indicates that the instances are matched based on the instance name.
	//
	// example:
	//
	// name1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The keyword used to match the instance name.
	//
	// Valid values of N: 1 to 3.
	//
	// example:
	//
	// portalHost
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateGroupMonitoringAgentProcessRequestMatchExpress) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessRequestMatchExpress) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) GetFunction() *string {
	return s.Function
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) GetName() *string {
	return s.Name
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) GetValue() *string {
	return s.Value
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) SetFunction(v string) *CreateGroupMonitoringAgentProcessRequestMatchExpress {
	s.Function = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) SetName(v string) *CreateGroupMonitoringAgentProcessRequestMatchExpress {
	s.Name = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) SetValue(v string) *CreateGroupMonitoringAgentProcessRequestMatchExpress {
	s.Value = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) Validate() error {
	return dara.Validate(s)
}
