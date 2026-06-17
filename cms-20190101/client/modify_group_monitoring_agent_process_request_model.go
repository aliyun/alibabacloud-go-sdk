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
	// The configurations of the alert rule.
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
	// The ID of the process monitoring job for the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92E3065F-0980-4E31-9AA0-BA6****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is deprecated. You can ignore it.
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
	if s.AlertConfig != nil {
		for _, item := range s.AlertConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyGroupMonitoringAgentProcessRequestAlertConfig struct {
	// The comparison operator for the threshold of the Critical alert level. The value of N can be 1 to 200. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to
	//
	// - GreaterThanThreshold: greater than
	//
	// - LessThanOrEqualToThreshold: less than or equal to
	//
	// - LessThanThreshold: less than
	//
	// - NotEqualToThreshold: not equal to
	//
	// - GreaterThanYesterday: greater than the value at the same time yesterday
	//
	// - LessThanYesterday: less than the value at the same time yesterday
	//
	// - GreaterThanLastWeek: greater than the value at the same time last week
	//
	// - LessThanLastWeek: less than the value at the same time last week
	//
	// - GreaterThanLastPeriod: greater than the value in the last monitoring cycle
	//
	// - LessThanLastPeriod: less than the value in the last monitoring cycle
	//
	// This parameter is required.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The time period when the alert rule is effective. The value of N can be 1 to 200.
	//
	// example:
	//
	// 00:00-22:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The alert level. The value of N can be 1 to 200. Valid values:
	//
	// - critical (default): critical
	//
	// - warn: warning
	//
	// - info: information
	//
	// This parameter is required.
	//
	// example:
	//
	// warn
	EscalationsLevel *string `json:"EscalationsLevel,omitempty" xml:"EscalationsLevel,omitempty"`
	// This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// 00:00-05:30
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The mute period. The value of N can be 1 to 200.
	//
	// Unit: seconds. Minimum value: 3600. Default value: 86400.
	//
	// > If monitoring data continuously exceeds the alert threshold, an alert notification is sent only once during each mute period.
	//
	// example:
	//
	// 86400
	SilenceTime *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The statistical method for alerts. The value of N can be 1 to 200.
	//
	// > Only Average is supported.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// None.
	TargetList []*ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The alert threshold. The value of N can be 1 to 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of consecutive times that the alert level is reached. The value of N can be 1 to 200. Default value: 3.
	//
	// > An alert is triggered only when the alert level is reached the specified number of consecutive times and the threshold is met.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Times *string `json:"Times,omitempty" xml:"Times,omitempty"`
	// The callback URL. A POST request is sent to this URL when an alert is triggered. The value of N can be 1 to 200.
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

type ModifyGroupMonitoringAgentProcessRequestAlertConfigTargetList struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	//
	// For more information, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
	//
	// The ARN of a resource is in the following format: `acs:{product-abbreviation}:{regionId}:{userId}:/{resource-type}/{resource-name}/message`. For example: `acs:mns:cn-hangzhou:120886317861****:/queues/test123/message`. The parameters are described as follows:
	//
	// - {product-abbreviation}: Currently, only Simple Message Queue (formerly MNS) is supported.
	//
	// - {userId}: The ID of your Alibaba Cloud account.
	//
	// - {regionId}: The region where the Simple Message Queue (formerly MNS) queue or subject is located.
	//
	// - {resource-type}: The type of the resource that receives alerts. Valid values:
	//
	//   - **queues**: a queue.
	//
	//   - **topics**: a subject.
	//
	// - {resource-name}: The name of the resource.
	//
	//   - If the resource type is **queues**, the resource name is the queue name.
	//
	//   - If the resource type is **topics**, the resource name is the subject name.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:120886317861****:/queues/test/message
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the alert-triggered target.
	//
	// For more information, see [DescribeMetricRuleTargets](https://help.aliyun.com/document_detail/121592.html).
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
	// - INFO: information
	//
	// - WARN: warning
	//
	// - CRITICAL: critical
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
