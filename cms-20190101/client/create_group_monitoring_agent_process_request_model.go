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
	// The list of alert rule configurations.
	//
	// Valid values of N: 1 to 3.
	//
	// This parameter is required.
	AlertConfig []*CreateGroupMonitoringAgentProcessRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	// The application group ID.
	//
	// For information about how to obtain the application group ID, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3607****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The list of match conditional expressions.
	//
	// Valid values of N: 1 to 3.
	MatchExpress []*CreateGroupMonitoringAgentProcessRequestMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
	// The logical relation for matching instances. Valid values:
	//
	// - all: all.
	//
	// - and: AND.
	//
	// - or: OR.
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
	// Process_Java
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
	if s.AlertConfig != nil {
		for _, item := range s.AlertConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MatchExpress != nil {
		for _, item := range s.MatchExpress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateGroupMonitoringAgentProcessRequestAlertConfig struct {
	// The comparison operator for the threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to.
	//
	// - GreaterThanThreshold: greater than.
	//
	// - LessThanOrEqualToThreshold: less than or equal to.
	//
	// - LessThanThreshold: less than.
	//
	// - NotEqualToThreshold: not equal to.
	//
	// - GreaterThanYesterday: greater than the value at the same time yesterday.
	//
	// - LessThanYesterday: less than the value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the value at the same time last week.
	//
	// - LessThanLastWeek: less than the value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the value in the previous period.
	//
	// - LessThanLastPeriod: less than the value in the previous period.
	//
	// Valid values of N: 1 to 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The effective period of the alert rule.
	//
	// Valid values of N: 1 to 3.
	//
	// example:
	//
	// 07:00-22:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The alert level. Valid values:
	//
	// - critical: critical. This is the default value.
	//
	// - warn: warning.
	//
	// - info: informational.
	//
	// Valid values of N: 1 to 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// warn
	EscalationsLevel *string `json:"EscalationsLevel,omitempty" xml:"EscalationsLevel,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and no longer requires attention.
	//
	// example:
	//
	// 00:00-05:30
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The mute period. Unit: seconds. Minimum value: 3600, which indicates 1 hour. Default value: 86400, which indicates 1 day.
	//
	// Valid values of N: 1 to 3.
	//
	// > If the monitoring data continuously exceeds the alert rule threshold, the system sends only one alert notification within each mute period.
	//
	// example:
	//
	// 86400
	SilenceTime *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The statistical method for alerts.
	//
	// Valid values of N: 1 to 3.
	//
	// > Only Average is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The list of alert targets.
	TargetList []*CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The alert threshold.
	//
	// Valid values of N: 1 to 3.
	//
	// Unit: count.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of consecutive times the alert level is triggered. Default value: 3.
	//
	// Valid values of N: 1 to 3.
	//
	// > An alert is triggered only when the alert level is reached for the specified number of consecutive times and the alert threshold is met.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Times *string `json:"Times,omitempty" xml:"Times,omitempty"`
	// The webhook URL for alert callbacks.
	//
	// Valid values of N: 1 to 3.
	//
	// example:
	//
	// https://www.aliyun.com
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
	if s.TargetList != nil {
		for _, item := range s.TargetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateGroupMonitoringAgentProcessRequestAlertConfigTargetList struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	//
	// For information about how to obtain the ARN of a resource, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
	//
	// The ARN is in the following format: `acs:{service abbreviation}:{regionId}:{userId}:/{resource type}/{resource name}/message`. Example: `acs:mns:cn-hangzhou:120886317861****:/queues/test123/message`. The parameters are described as follows:
	//
	// - {service abbreviation}: the abbreviation of the cloud service. Only Simple Message Queue (formerly MNS) is supported.
	//
	// - {userId}: the Alibaba Cloud account ID.
	//
	// - {regionId}: the region where the Simple Message Queue (formerly MNS) queue or topic resides.
	//
	// - {resource type}: the type of the resource that receives alert notifications. Valid values:
	//
	//
	//
	//     - **queues**: queue.
	//
	//     - **topics**: topic.
	//
	// - {resource name}: the name of the resource.
	//
	//     - If the resource type is **queues**, the resource name is the queue name.
	//
	//     - If the resource type is **topics**, the resource name is the topic name.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:120886317861****:/queues/test/message
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the alert target.
	//
	// For information about how to obtain the ID of an alert target, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
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
	// - INFO: informational.
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
	// The match condition. Valid values:
	//
	// - all: matches all. This is the default value.
	//
	// - startWith: matches the prefix.
	//
	// - endWith: matches the suffix.
	//
	// - contains: contains the specified string.
	//
	// - notContains: does not contain the specified string.
	//
	// - equals: exactly matches.
	//
	// Valid values of N: 1 to 3.
	//
	// example:
	//
	// startWith
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The type of the match condition.
	//
	// Valid values of N: 1 to 3.
	//
	// > Only Name, which indicates the instance name, is supported.
	//
	// example:
	//
	// name1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the matched instance.
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
