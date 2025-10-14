// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupMonitoringAgentProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeGroupMonitoringAgentProcessResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeGroupMonitoringAgentProcessResponseBody
	GetMessage() *string
	SetPageNumber(v string) *DescribeGroupMonitoringAgentProcessResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeGroupMonitoringAgentProcessResponseBody
	GetPageSize() *string
	SetProcesses(v *DescribeGroupMonitoringAgentProcessResponseBodyProcesses) *DescribeGroupMonitoringAgentProcessResponseBody
	GetProcesses() *DescribeGroupMonitoringAgentProcessResponseBodyProcesses
	SetRequestId(v string) *DescribeGroupMonitoringAgentProcessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeGroupMonitoringAgentProcessResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeGroupMonitoringAgentProcessResponseBody
	GetTotal() *string
}

type DescribeGroupMonitoringAgentProcessResponseBody struct {
	// The HTTP status codes.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The process monitoring tasks.
	Processes *DescribeGroupMonitoringAgentProcessResponseBodyProcesses `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7985D471-3FA8-4EE9-8F4B-45C19DF3D36F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 28
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) GetProcesses() *DescribeGroupMonitoringAgentProcessResponseBodyProcesses {
	return s.Processes
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetCode(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetMessage(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetPageNumber(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetPageSize(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetProcesses(v *DescribeGroupMonitoringAgentProcessResponseBodyProcesses) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Processes = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetRequestId(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetSuccess(v bool) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetTotal(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) Validate() error {
	if s.Processes != nil {
		if err := s.Processes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcesses struct {
	Process []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess `json:"Process,omitempty" xml:"Process,omitempty" type:"Repeated"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcesses) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcesses) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcesses) GetProcess() []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	return s.Process
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcesses) SetProcess(v []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) *DescribeGroupMonitoringAgentProcessResponseBodyProcesses {
	s.Process = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcesses) Validate() error {
	if s.Process != nil {
		for _, item := range s.Process {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess struct {
	// The alert rule configurations.
	AlertConfig *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	// The ID of the application group.
	//
	// example:
	//
	// 12345
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the process monitoring task.
	//
	// example:
	//
	// 3F6150F9-45C7-43F9-9578-A58B2E72****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The matching conditions.
	//
	// >  Only the instances that meet the conditional expressions are monitored by the process monitoring task.
	MatchExpress *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Struct"`
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
	// example:
	//
	// sshd
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) GetAlertConfig() *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig {
	return s.AlertConfig
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) GetId() *string {
	return s.Id
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) GetMatchExpress() *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress {
	return s.MatchExpress
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) GetMatchExpressFilterRelation() *string {
	return s.MatchExpressFilterRelation
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetAlertConfig(v *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.AlertConfig = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetGroupId(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetId(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.Id = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetMatchExpress(v *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.MatchExpress = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetMatchExpressFilterRelation(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetProcessName(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.ProcessName = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) Validate() error {
	if s.AlertConfig != nil {
		if err := s.AlertConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MatchExpress != nil {
		if err := s.MatchExpress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig struct {
	AlertConfig []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) GetAlertConfig() []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	return s.AlertConfig
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) SetAlertConfig(v []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) Validate() error {
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

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig struct {
	// The comparison operator that is used to compare the metric value with the threshold. Valid values:
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
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The time period during which the alert rule is effective.
	//
	// example:
	//
	// 00:00-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The level of the alert. Valid values:
	//
	// 	- critical
	//
	// 	- warn
	//
	// 	- Info
	//
	// example:
	//
	// warn
	EscalationsLevel *string `json:"EscalationsLevel,omitempty" xml:"EscalationsLevel,omitempty"`
	// The time period during which the alert rule is ineffective.
	//
	// example:
	//
	// 00:00-23:59
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The mute period during which new alert notifications are not sent even if the trigger conditions are met. Unit: seconds. Minimum value: 3600, which is equivalent to one hour. Default value: 86400, which is equivalent to one day.
	//
	// >  Only one alert notification is sent during each mute period even if the metric value exceeds the alert threshold several times.
	//
	// example:
	//
	// 86400
	SilenceTime *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The method used to calculate metric values that trigger alerts.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The resources for which alerts are triggered.
	TargetList *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Struct"`
	// The alert threshold.
	//
	// example:
	//
	// 5
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of times for which the threshold can be consecutively exceeded.
	//
	// >  A metric triggers an alert only after the metric value reaches the threshold consecutively for the specified times.
	//
	// example:
	//
	// 3
	Times *string `json:"Times,omitempty" xml:"Times,omitempty"`
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
	//
	// example:
	//
	// http://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetEscalationsLevel() *string {
	return s.EscalationsLevel
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetSilenceTime() *string {
	return s.SilenceTime
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetTargetList() *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList {
	return s.TargetList
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetTimes() *string {
	return s.Times
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GetWebhook() *string {
	return s.Webhook
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetComparisonOperator(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetEffectiveInterval(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetEscalationsLevel(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.EscalationsLevel = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetNoEffectiveInterval(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetSilenceTime(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetStatistics(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.Statistics = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetTargetList(v *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.TargetList = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetThreshold(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetTimes(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.Times = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetWebhook(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.Webhook = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) Validate() error {
	if s.TargetList != nil {
		if err := s.TargetList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList struct {
	Target []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList) GetTarget() []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget {
	return s.Target
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList) SetTarget(v []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList {
	s.Target = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetList) Validate() error {
	if s.Target != nil {
		for _, item := range s.Target {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource. Format: acs:{Service name abbreviation}:{regionId}:{userId}:/{Resource type}/{Resource name}/message. Example: acs:mns:cn-hangzhou:120886317861\\*\\*\\*\\*:/queues/test123/message. Fields:
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
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters of the alert callback. The parameters are in the JSON format.
	//
	// example:
	//
	// {"customField1":"value1","customField2":"$.name"}
	JsonParmas *string `json:"JsonParmas,omitempty" xml:"JsonParmas,omitempty"`
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
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) GetArn() *string {
	return s.Arn
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) GetId() *string {
	return s.Id
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) GetJsonParmas() *string {
	return s.JsonParmas
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) GetLevel() *string {
	return s.Level
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) SetArn(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget {
	s.Arn = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) SetId(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget {
	s.Id = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) SetJsonParmas(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget {
	s.JsonParmas = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) SetLevel(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget {
	s.Level = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfigTargetListTarget) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress struct {
	MatchExpress []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) GetMatchExpress() []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress {
	return s.MatchExpress
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) SetMatchExpress(v []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress {
	s.MatchExpress = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) Validate() error {
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

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress struct {
	// The matching condition. Valid values:
	//
	// 	- all (default): matches all
	//
	// 	- startWith: starts with a prefix
	//
	// 	- endWith: ends with a suffix
	//
	// 	- contains: contains
	//
	// 	- notContains: excludes
	//
	// 	- equals: equals
	//
	// >  The matched instances are monitored by the process monitoring task.
	//
	// example:
	//
	// all
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The criteria based on which the instances are matched.
	//
	// >  Set the value to `name`. The value name indicates that the instances are matched based on the instance name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The keyword used to match the instance name.
	//
	// example:
	//
	// portalHost
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) GetFunction() *string {
	return s.Function
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) GetName() *string {
	return s.Name
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) GetValue() *string {
	return s.Value
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) SetFunction(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress {
	s.Function = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) SetName(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress {
	s.Name = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) SetValue(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress {
	s.Value = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) Validate() error {
	return dara.Validate(s)
}
