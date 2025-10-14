// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHostAvailabilityListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHostAvailabilityListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeHostAvailabilityListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHostAvailabilityListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHostAvailabilityListResponseBody
	GetSuccess() *bool
	SetTaskList(v *DescribeHostAvailabilityListResponseBodyTaskList) *DescribeHostAvailabilityListResponseBody
	GetTaskList() *DescribeHostAvailabilityListResponseBodyTaskList
	SetTotal(v int32) *DescribeHostAvailabilityListResponseBody
	GetTotal() *int32
}

type DescribeHostAvailabilityListResponseBody struct {
	// The HTTP status code.
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
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A288E86-45C3-4858-9DB0-6D85B10BD92A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the availability monitoring tasks.
	TaskList *DescribeHostAvailabilityListResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Struct"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHostAvailabilityListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHostAvailabilityListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHostAvailabilityListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHostAvailabilityListResponseBody) GetTaskList() *DescribeHostAvailabilityListResponseBodyTaskList {
	return s.TaskList
}

func (s *DescribeHostAvailabilityListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeHostAvailabilityListResponseBody) SetCode(v string) *DescribeHostAvailabilityListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetMessage(v string) *DescribeHostAvailabilityListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetRequestId(v string) *DescribeHostAvailabilityListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetSuccess(v bool) *DescribeHostAvailabilityListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetTaskList(v *DescribeHostAvailabilityListResponseBodyTaskList) *DescribeHostAvailabilityListResponseBody {
	s.TaskList = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetTotal(v int32) *DescribeHostAvailabilityListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) Validate() error {
	if s.TaskList != nil {
		if err := s.TaskList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHostAvailabilityListResponseBodyTaskList struct {
	NodeTaskConfig []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig `json:"NodeTaskConfig,omitempty" xml:"NodeTaskConfig,omitempty" type:"Repeated"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskList) GetNodeTaskConfig() []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	return s.NodeTaskConfig
}

func (s *DescribeHostAvailabilityListResponseBodyTaskList) SetNodeTaskConfig(v []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) *DescribeHostAvailabilityListResponseBodyTaskList {
	s.NodeTaskConfig = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskList) Validate() error {
	if s.NodeTaskConfig != nil {
		for _, item := range s.NodeTaskConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig struct {
	// The configurations of the alert rule.
	AlertConfig *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	// Indicates whether the availability monitoring task is disabled. Valid values:
	//
	// 	- true: The availability monitoring task is disabled.
	//
	// 	- false: The availability monitoring task is enabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 12345
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// Group_ECS
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the availability monitoring task.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ECS instances that are monitored.
	Instances *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The name of the availability monitoring task.
	//
	// example:
	//
	// ecs_instance
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The optional parameters of the availability monitoring task.
	TaskOption *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption `json:"TaskOption,omitempty" xml:"TaskOption,omitempty" type:"Struct"`
	// The range of instances that are monitored by the availability monitoring task. Valid values:
	//
	// 	- GROUP: All ECS instances in the application group are monitored.
	//
	// 	- GROUP_SPEC_INSTANCE: Specified ECS instances in the application group are monitored.
	//
	// example:
	//
	// GROUP
	TaskScope *string `json:"TaskScope,omitempty" xml:"TaskScope,omitempty"`
	// The task type. Valid values:
	//
	// 	- PING
	//
	// 	- TELNET
	//
	// 	- HTTP
	//
	// example:
	//
	// HTTP
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetAlertConfig() *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	return s.AlertConfig
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetId() *int64 {
	return s.Id
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetInstances() *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances {
	return s.Instances
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetTaskOption() *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	return s.TaskOption
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetTaskScope() *string {
	return s.TaskScope
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetAlertConfig(v *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetDisabled(v bool) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.Disabled = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetGroupId(v int64) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.GroupId = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetGroupName(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.GroupName = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetId(v int64) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.Id = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetInstances(v *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.Instances = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetTaskName(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.TaskName = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetTaskOption(v *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.TaskOption = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetTaskScope(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.TaskScope = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetTaskType(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.TaskType = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) Validate() error {
	if s.AlertConfig != nil {
		if err := s.AlertConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	if s.TaskOption != nil {
		if err := s.TaskOption.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig struct {
	// The end of the time period during which the alert rule is effective. Valid values: 0 to 23.
	//
	// For example, if the `AlertConfig.StartTime` parameter is set to 0 and the `AlertConfig.EndTime` parameter is set to 22, the alert rule is effective from 00:00:00 to 22:00:00.
	//
	// >  Alert notifications are sent based on the specified threshold only if the alert rule is effective.
	//
	// example:
	//
	// 22
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The trigger conditions of the alert rule.
	EscalationList *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList `json:"EscalationList,omitempty" xml:"EscalationList,omitempty" type:"Struct"`
	// The alert notification methods. Valid values:
	//
	// 	- 2: Alert notifications are sent by using emails and DingTalk chatbots.
	//
	// 	- 1: Alert notifications are sent by using emails and DingTalk chatbots.
	//
	// 	- 0: Alert notifications are sent by using emails and DingTalk chatbots.
	//
	// example:
	//
	// 0
	NotifyType *int32 `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The mute period during which new alerts are not sent even if the trigger conditions are met. Unit: seconds. Default value: 86400.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The beginning of the time period during which the alert rule is effective. Valid values: 0 to 23.
	//
	// For example, if the `AlertConfig.StartTime` parameter is set to 0 and the `AlertConfig.EndTime` parameter is set to 22, the alert rule is effective from 00:00:00 to 22:00:00.
	//
	// >  Alert notifications are sent based on the specified threshold only if the alert rule is effective.
	//
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The monitored resources.
	TargetList *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Struct"`
	// The callback URL.
	//
	// CloudMonitor pushes an alert notification to the specified callback URL by sending an HTTP POST request. Only the HTTP protocol is supported.
	//
	// example:
	//
	// https://www.aliyun.com
	WebHook *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) GetEndTime() *int32 {
	return s.EndTime
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) GetEscalationList() *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList {
	return s.EscalationList
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) GetNotifyType() *int32 {
	return s.NotifyType
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) GetStartTime() *int32 {
	return s.StartTime
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) GetTargetList() *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList {
	return s.TargetList
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) GetWebHook() *string {
	return s.WebHook
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetEndTime(v int32) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetEscalationList(v *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.EscalationList = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetNotifyType(v int32) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.NotifyType = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetSilenceTime(v int32) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetStartTime(v int32) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.StartTime = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetTargetList(v *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.TargetList = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetWebHook(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.WebHook = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) Validate() error {
	if s.EscalationList != nil {
		if err := s.EscalationList.Validate(); err != nil {
			return err
		}
	}
	if s.TargetList != nil {
		if err := s.TargetList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList struct {
	EscalationList []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList `json:"escalationList,omitempty" xml:"escalationList,omitempty" type:"Repeated"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) GetEscalationList() []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	return s.EscalationList
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) SetEscalationList(v []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList {
	s.EscalationList = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) Validate() error {
	if s.EscalationList != nil {
		for _, item := range s.EscalationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList struct {
	// The method used to calculate metric values that trigger alerts. Valid values:
	//
	// 	- Value: the value of the HTTP status code
	//
	// 	- Average: the average HTTP response time
	//
	// 	- Value: the value of the Telnet status code
	//
	// 	- TelnetLatency: the average Telnet response time
	//
	// 	- Average: the average Ping packet loss rate
	//
	// example:
	//
	// Value
	Aggregate *string `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
	// The name of the metric. Valid values:
	//
	// 	- HttpStatus
	//
	// 	- HttpLatency
	//
	// 	- TelnetStatus
	//
	// 	- TelnetLatency
	//
	// 	- PingLostRate
	//
	// example:
	//
	// HttpStatus
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The comparison operator that is used in the alert rule. Valid values:
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
	// =
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The consecutive number of times for which the metric value is measured before an alert is triggered.
	//
	// example:
	//
	// 3
	Times *string `json:"Times,omitempty" xml:"Times,omitempty"`
	// The alert threshold.
	//
	// example:
	//
	// 400
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) GetAggregate() *string {
	return s.Aggregate
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) GetOperator() *string {
	return s.Operator
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) GetTimes() *string {
	return s.Times
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) GetValue() *string {
	return s.Value
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetAggregate(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.Aggregate = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetMetricName(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.MetricName = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetOperator(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.Operator = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetTimes(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.Times = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetValue(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.Value = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) Validate() error {
	return dara.Validate(s)
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList struct {
	Target []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList) GetTarget() []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget {
	return s.Target
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList) SetTarget(v []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList {
	s.Target = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetList) Validate() error {
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

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget struct {
	// The Alibaba Cloud Resource Name (ARN) of the function.
	//
	// Format: `arn:acs:${Service}:${Region}:${Account}:${ResourceType}/${ResourceId}`. Fields:
	//
	// 	- Service: the service code
	//
	// 	- Region: the region ID
	//
	// 	- Account: the ID of the Alibaba Cloud account
	//
	// 	- ResourceType: the resource type
	//
	// 	- ResourceId: the resource ID.
	//
	// example:
	//
	// acs:mns:cn-hangzhou:17754132319*****:/queues/test/messages
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the resource that triggers the alert.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-formatted parameters of the alert callback.
	//
	// example:
	//
	// {"key1":"value1"}
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
	// INFO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) GetArn() *string {
	return s.Arn
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) GetId() *string {
	return s.Id
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) GetJsonParams() *string {
	return s.JsonParams
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) GetLevel() *string {
	return s.Level
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) SetArn(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget {
	s.Arn = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) SetId(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget {
	s.Id = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) SetJsonParams(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget {
	s.JsonParams = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) SetLevel(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget {
	s.Level = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigTargetListTarget) Validate() error {
	return dara.Validate(s)
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances struct {
	Instance []*string `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) GetInstance() []*string {
	return s.Instance
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) SetInstance(v []*string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances {
	s.Instance = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption struct {
	// The response to the HTTP request.
	//
	// example:
	//
	// OK
	HttpKeyword *string `json:"HttpKeyword,omitempty" xml:"HttpKeyword,omitempty"`
	// The HTTP request method. Valid values:
	//
	// 	- GET
	//
	// 	- POST
	//
	// 	- HEAD
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
	// example:
	//
	// UTF-8
	HttpResponseCharset *string `json:"HttpResponseCharset,omitempty" xml:"HttpResponseCharset,omitempty"`
	// The URI that you want to monitor. If the TaskType parameter is set to HTTP, this parameter is required.
	//
	// example:
	//
	// https://www.aliyun.com
	HttpURI *string `json:"HttpURI,omitempty" xml:"HttpURI,omitempty"`
	// The interval at which detection requests are sent. Unit: seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The domain name or IP address that you want to monitor.
	//
	// example:
	//
	// ssh.aliyun.com
	TelnetOrPingHost *string `json:"TelnetOrPingHost,omitempty" xml:"TelnetOrPingHost,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GetHttpKeyword() *string {
	return s.HttpKeyword
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GetHttpNegative() *bool {
	return s.HttpNegative
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GetHttpPostContent() *string {
	return s.HttpPostContent
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GetHttpResponseCharset() *string {
	return s.HttpResponseCharset
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GetHttpURI() *string {
	return s.HttpURI
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GetTelnetOrPingHost() *string {
	return s.TelnetOrPingHost
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpKeyword(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpKeyword = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpMethod(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpMethod = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpNegative(v bool) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpNegative = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpPostContent(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpPostContent = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpResponseCharset(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpResponseCharset = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpURI(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpURI = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetInterval(v int32) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.Interval = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetTelnetOrPingHost(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.TelnetOrPingHost = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) Validate() error {
	return dara.Validate(s)
}
