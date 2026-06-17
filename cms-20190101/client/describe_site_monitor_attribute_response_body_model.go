// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSiteMonitorAttributeResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeSiteMonitorAttributeResponseBody
	GetMessage() *string
	SetMetricRules(v *DescribeSiteMonitorAttributeResponseBodyMetricRules) *DescribeSiteMonitorAttributeResponseBody
	GetMetricRules() *DescribeSiteMonitorAttributeResponseBodyMetricRules
	SetRequestId(v string) *DescribeSiteMonitorAttributeResponseBody
	GetRequestId() *string
	SetSiteMonitors(v *DescribeSiteMonitorAttributeResponseBodySiteMonitors) *DescribeSiteMonitorAttributeResponseBody
	GetSiteMonitors() *DescribeSiteMonitorAttributeResponseBodySiteMonitors
	SetSuccess(v bool) *DescribeSiteMonitorAttributeResponseBody
	GetSuccess() *bool
}

type DescribeSiteMonitorAttributeResponseBody struct {
	// The status code.
	//
	// >The value 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message     *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	MetricRules *DescribeSiteMonitorAttributeResponseBodyMetricRules `json:"MetricRules,omitempty" xml:"MetricRules,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D17DF650-7EBD-54D0-903A-1D4E624D7402
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the monitoring task.
	SiteMonitors *DescribeSiteMonitorAttributeResponseBodySiteMonitors `json:"SiteMonitors,omitempty" xml:"SiteMonitors,omitempty" type:"Struct"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSiteMonitorAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSiteMonitorAttributeResponseBody) GetMetricRules() *DescribeSiteMonitorAttributeResponseBodyMetricRules {
	return s.MetricRules
}

func (s *DescribeSiteMonitorAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteMonitorAttributeResponseBody) GetSiteMonitors() *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	return s.SiteMonitors
}

func (s *DescribeSiteMonitorAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetCode(v string) *DescribeSiteMonitorAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetMessage(v string) *DescribeSiteMonitorAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetMetricRules(v *DescribeSiteMonitorAttributeResponseBodyMetricRules) *DescribeSiteMonitorAttributeResponseBody {
	s.MetricRules = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetRequestId(v string) *DescribeSiteMonitorAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetSiteMonitors(v *DescribeSiteMonitorAttributeResponseBodySiteMonitors) *DescribeSiteMonitorAttributeResponseBody {
	s.SiteMonitors = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetSuccess(v bool) *DescribeSiteMonitorAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) Validate() error {
	if s.MetricRules != nil {
		if err := s.MetricRules.Validate(); err != nil {
			return err
		}
	}
	if s.SiteMonitors != nil {
		if err := s.SiteMonitors.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodyMetricRules struct {
	MetricRule []*DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule `json:"MetricRule,omitempty" xml:"MetricRule,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodyMetricRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodyMetricRules) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRules) GetMetricRule() []*DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	return s.MetricRule
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRules) SetMetricRule(v []*DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) *DescribeSiteMonitorAttributeResponseBodyMetricRules {
	s.MetricRule = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRules) Validate() error {
	if s.MetricRule != nil {
		for _, item := range s.MetricRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule struct {
	ActionEnable       *string `json:"ActionEnable,omitempty" xml:"ActionEnable,omitempty"`
	AlarmActions       *string `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Dimensions         *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EvaluationCount    *string `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Expression         *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Level              *string `json:"Level,omitempty" xml:"Level,omitempty"`
	MetricName         *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OkActions          *string `json:"OkActions,omitempty" xml:"OkActions,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RuleId             *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	StateValue         *string `json:"StateValue,omitempty" xml:"StateValue,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetActionEnable() *string {
	return s.ActionEnable
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetAlarmActions() *string {
	return s.AlarmActions
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetEvaluationCount() *string {
	return s.EvaluationCount
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetExpression() *string {
	return s.Expression
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetLevel() *string {
	return s.Level
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetOkActions() *string {
	return s.OkActions
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetPeriod() *string {
	return s.Period
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetStateValue() *string {
	return s.StateValue
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetActionEnable(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.ActionEnable = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetAlarmActions(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.AlarmActions = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetComparisonOperator(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetDimensions(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Dimensions = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetEvaluationCount(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetExpression(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Expression = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetLevel(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Level = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetMetricName(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.MetricName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetNamespace(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Namespace = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetOkActions(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.OkActions = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetPeriod(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Period = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetRuleId(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.RuleId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetRuleName(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.RuleName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetStateValue(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.StateValue = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetStatistics(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Statistics = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetThreshold(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Threshold = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitors struct {
	// The monitored address of the monitoring task.
	//
	// example:
	//
	// https://aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The type of detection point. Default value: PC.
	//
	// Valid values:
	//
	// - PC: wired network.
	//
	// - MOBILE: mobile network.
	//
	// example:
	//
	// PC
	AgentGroup *string `json:"AgentGroup,omitempty" xml:"AgentGroup,omitempty"`
	// The custom monitoring schedule. You can select a time range from Monday to Sunday for monitoring.
	CustomSchedule *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule `json:"CustomSchedule,omitempty" xml:"CustomSchedule,omitempty" type:"Struct"`
	// The monitoring interval. Unit: minutes. Valid values: 1, 5, 15, 30, and 60.
	//
	// example:
	//
	// 1
	Interval  *string                                                        `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspCities *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities `json:"IspCities,omitempty" xml:"IspCities,omitempty" type:"Struct"`
	// The extended options. Each monitoring type has different extended options. For more information, see [CreateSiteMonitor](https://help.aliyun.com/document_detail/115048.html).
	OptionJson *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson `json:"OptionJson,omitempty" xml:"OptionJson,omitempty" type:"Struct"`
	// The ID of the monitoring task.
	//
	// example:
	//
	// cc641dff-c19d-45f3-ad0a-818a0c4f****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the monitoring task.
	//
	// example:
	//
	// test123
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The status of the monitoring task. Valid values:
	//
	// - 1: Enabled.
	//
	// - 2: Disabled.
	//
	// example:
	//
	// 1
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The type of the monitoring task. Site monitoring task types include HTTP(S), PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	//
	// example:
	//
	// HTTP
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The VPC configuration for the internal network monitoring task.
	VpcConfig *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty" type:"Struct"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitors) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitors) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetAddress() *string {
	return s.Address
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetAgentGroup() *string {
	return s.AgentGroup
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetCustomSchedule() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule {
	return s.CustomSchedule
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetIspCities() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities {
	return s.IspCities
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetOptionJson() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	return s.OptionJson
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) GetVpcConfig() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig {
	return s.VpcConfig
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetAddress(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.Address = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetAgentGroup(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.AgentGroup = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetCustomSchedule(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.CustomSchedule = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetInterval(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.Interval = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetIspCities(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.IspCities = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetOptionJson(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.OptionJson = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetTaskId(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetTaskName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.TaskName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetTaskState(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.TaskState = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetTaskType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.TaskType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetVpcConfig(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.VpcConfig = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) Validate() error {
	if s.CustomSchedule != nil {
		if err := s.CustomSchedule.Validate(); err != nil {
			return err
		}
	}
	if s.IspCities != nil {
		if err := s.IspCities.Validate(); err != nil {
			return err
		}
	}
	if s.OptionJson != nil {
		if err := s.OptionJson.Validate(); err != nil {
			return err
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule struct {
	Days *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays `json:"days,omitempty" xml:"days,omitempty" type:"Struct"`
	// The custom monitoring end time.
	//
	// Unit: hours.
	//
	// example:
	//
	// 18
	EndHour *int32 `json:"end_hour,omitempty" xml:"end_hour,omitempty"`
	// The custom monitoring start time.
	//
	// Unit: hours.
	//
	// example:
	//
	// 8
	StartHour *int32 `json:"start_hour,omitempty" xml:"start_hour,omitempty"`
	// The time zone for custom monitoring.
	//
	// example:
	//
	// local
	TimeZone *string `json:"time_zone,omitempty" xml:"time_zone,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) GetDays() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays {
	return s.Days
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) GetEndHour() *int32 {
	return s.EndHour
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) GetStartHour() *int32 {
	return s.StartHour
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) SetDays(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule {
	s.Days = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) SetEndHour(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule {
	s.EndHour = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) SetStartHour(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule {
	s.StartHour = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) SetTimeZone(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule {
	s.TimeZone = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomSchedule) Validate() error {
	if s.Days != nil {
		if err := s.Days.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays struct {
	Days []*int32 `json:"days,omitempty" xml:"days,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays) GetDays() []*int32 {
	return s.Days
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays) SetDays(v []*int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays {
	s.Days = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsCustomScheduleDays) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities struct {
	IspCity []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity `json:"IspCity,omitempty" xml:"IspCity,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) GetIspCity() []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	return s.IspCity
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) SetIspCity(v []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities {
	s.IspCity = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) Validate() error {
	if s.IspCity != nil {
		for _, item := range s.IspCity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity struct {
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	Isp      *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	IspName  *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) GetCity() *string {
	return s.City
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) GetCityName() *string {
	return s.CityName
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) GetIsp() *string {
	return s.Isp
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) GetIspName() *string {
	return s.IspName
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) GetType() *string {
	return s.Type
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) SetCity(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	s.City = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) SetCityName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	s.CityName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) SetIsp(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	s.Isp = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) SetIspName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	s.IspName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) SetType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	s.Type = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson struct {
	Assertions *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions `json:"assertions,omitempty" xml:"assertions,omitempty" type:"Struct"`
	// The number of retries after a DNS failure.
	//
	// example:
	//
	// 3
	Attempts *int64 `json:"attempts,omitempty" xml:"attempts,omitempty"`
	// The authentication information.
	AuthInfo       *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo       `json:"auth_info,omitempty" xml:"auth_info,omitempty" type:"Struct"`
	BlockedUrlList *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList `json:"blocked_url_list,omitempty" xml:"blocked_url_list,omitempty" type:"Struct"`
	BrowserHeaders *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders `json:"browser_headers,omitempty" xml:"browser_headers,omitempty" type:"Struct"`
	BrowserHosts   *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts   `json:"browser_hosts,omitempty" xml:"browser_hosts,omitempty" type:"Struct"`
	BrowserInfo    *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo    `json:"browser_info,omitempty" xml:"browser_info,omitempty" type:"Struct"`
	// Specifies whether to ignore certificate errors. Valid values:
	//
	// - false: Does not ignore certificate errors.
	//
	// - true: Ignores certificate errors.
	//
	// example:
	//
	// false
	BrowserInsecure *bool `json:"browser_insecure,omitempty" xml:"browser_insecure,omitempty"`
	// The browser monitoring version. Valid values:
	//
	// - 1: Single-page monitoring.
	//
	// - 2: Multi-page monitoring.
	//
	// example:
	//
	// 1
	BrowserTaskVersion *string                                                                        `json:"browser_task_version,omitempty" xml:"browser_task_version,omitempty"`
	ConfigVariables    *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables `json:"config_variables,omitempty" xml:"config_variables,omitempty" type:"Struct"`
	// The cookie for the HTTP request.
	//
	// example:
	//
	// lang=en
	Cookie *string `json:"cookie,omitempty" xml:"cookie,omitempty"`
	// Specifies whether to enable automatic MTR network diagnostics after a task failure. Valid values:
	//
	// - false: Disabled.
	//
	// - true: Enabled.
	//
	// example:
	//
	// false
	DiagnosisMtr *bool `json:"diagnosis_mtr,omitempty" xml:"diagnosis_mtr,omitempty"`
	// Specifies whether to enable automatic PING network latency detection after a task failure. Valid values:
	//
	// - false: Disabled.
	//
	// - true: Enabled.
	//
	// example:
	//
	// false
	DiagnosisPing *bool `json:"diagnosis_ping,omitempty" xml:"diagnosis_ping,omitempty"`
	// The DNS hijacking configuration list.
	//
	// example:
	//
	// www.taobao.com:www.taobao.com.danuoyi.tbcache.com
	DnsHijackWhitelist *string `json:"dns_hijack_whitelist,omitempty" xml:"dns_hijack_whitelist,omitempty"`
	// The DNS matching rule. Valid values:
	//
	// - IN_DNS: The expected aliases or IP addresses are all included in the DNS response.
	//
	// - DNS_IN: All DNS responses are included in the expected aliases or IP addresses.
	//
	// - EQUAL: The DNS response exactly matches the expected aliases or IP addresses.
	//
	// - ANY: The DNS response and the expected aliases or IP addresses have an intersection.
	//
	// example:
	//
	// IN_DNS
	DnsMatchRule *string `json:"dns_match_rule,omitempty" xml:"dns_match_rule,omitempty"`
	// The IP address of the DNS server.
	//
	// > This parameter applies only to the DNS monitoring type.
	//
	// example:
	//
	// 192.168.XX.XX
	DnsServer *string `json:"dns_server,omitempty" xml:"dns_server,omitempty"`
	// The DNS resolution type. This parameter applies only to the DNS monitoring type. Valid values:
	//
	// - A: Specifies the IP address corresponding to a hostname or domain name.
	//
	// - CNAME: Maps multiple domain names to another domain name.
	//
	// - NS: Specifies the DNS server that resolves a domain name.
	//
	// - MX: Points a domain name to a mail server address.
	//
	// - TXT: A description of the hostname or domain name. The text length is limited to 512 bytes and is typically used for SPF (Sender Policy Framework) records for anti-spam purposes.
	//
	// example:
	//
	// A
	DnsType *string `json:"dns_type,omitempty" xml:"dns_type,omitempty"`
	// Specifies whether the WebSocket task is allowed to return no message or an empty message. Valid values:
	//
	// - false (default): Not allowed.
	//
	// - true: Allowed.
	//
	// example:
	//
	// false
	EmptyMessage *bool `json:"empty_message,omitempty" xml:"empty_message,omitempty"`
	// Specifies whether to enable packet capture for this task.
	//
	// example:
	//
	// true
	EnablePacketCapture  *bool                                                                               `json:"enable_packet_capture,omitempty" xml:"enable_packet_capture,omitempty"`
	ExpectExistString    *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString    `json:"expect_exist_string,omitempty" xml:"expect_exist_string,omitempty" type:"Struct"`
	ExpectNonExistString *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString `json:"expect_non_exist_string,omitempty" xml:"expect_non_exist_string,omitempty" type:"Struct"`
	// The alias or address to be resolved.
	//
	// > This parameter applies only to the DNS monitoring type.
	//
	// example:
	//
	// dns_server
	ExpectValue *string `json:"expect_value,omitempty" xml:"expect_value,omitempty"`
	// The packet loss rate.
	//
	// > This parameter applies only to the PING monitoring type.
	//
	// example:
	//
	// 0.4
	FailureRate *float32 `json:"failure_rate,omitempty" xml:"failure_rate,omitempty"`
	// The HTTP request header.
	//
	// example:
	//
	// testKey:testValue
	Header *string `json:"header,omitempty" xml:"header,omitempty"`
	// The number of hops for traceroute diagnostics when a PING task fails.
	//
	// example:
	//
	// 20
	Hops *int32 `json:"hops,omitempty" xml:"hops,omitempty"`
	// The custom host for HTTP tasks. The format is ip1,ip2:address. Multiple mappings can be configured. The left side of the colon contains A records or CNAMEs that the domain name can be resolved to, separated by commas. The right side of the colon is the domain name.
	//
	// example:
	//
	// 127.0.0.1:www.aliyun.com
	HostBinding *string `json:"host_binding,omitempty" xml:"host_binding,omitempty"`
	// Specifies how the custom host takes effect. Valid values: 0 (random) and 1 (round-robin).
	//
	// example:
	//
	// 0
	HostBindingType *int32 `json:"host_binding_type,omitempty" xml:"host_binding_type,omitempty"`
	// The HTTP request method. Valid values:
	//
	// - get
	//
	// - post
	//
	// - head.
	//
	// example:
	//
	// get
	HttpMethod *string `json:"http_method,omitempty" xml:"http_method,omitempty"`
	// The timeout period for a single PING request using the ICMP protocol. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	IcmpTimeoutMillis *int32 `json:"icmp_timeout_millis,omitempty" xml:"icmp_timeout_millis,omitempty"`
	// The network type of the task. Valid values: v4, v6, and auto. Default value: v4.
	//
	// example:
	//
	// v4
	IpNetwork *string `json:"ip_network,omitempty" xml:"ip_network,omitempty"`
	// Specifies whether to decode and store the password using Base64. Valid values:
	//
	// - true: The password is decoded and stored using Base64.
	//
	// - false: The password is not decoded and stored using Base64.
	//
	// example:
	//
	// false
	IsBase64Encode *string `json:"isBase64Encode,omitempty" xml:"isBase64Encode,omitempty"`
	// Specifies whether alert rules are included. Valid values:
	//
	// - 0: Yes.
	//
	// - 1: No.
	//
	// example:
	//
	// 1
	MatchRule *int32 `json:"match_rule,omitempty" xml:"match_rule,omitempty"`
	// The maximum TLS version.
	//
	// example:
	//
	// tlsv1.3
	MaxTlsVersion *string `json:"max_tls_version,omitempty" xml:"max_tls_version,omitempty"`
	// The minimum TLS version. TLS 1.2 and later are supported by default. TLS 1.0 and 1.1 are disabled. To support these versions, modify the configuration.
	//
	// example:
	//
	// tlsv1.2
	MinTlsVersion *string `json:"min_tls_version,omitempty" xml:"min_tls_version,omitempty"`
	// The password for SMTP, POP3, or FTP monitoring types.
	//
	// example:
	//
	// 123****
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The number of PING packets for the PING monitoring type.
	//
	// example:
	//
	// 29
	PingNum *int32 `json:"ping_num,omitempty" xml:"ping_num,omitempty"`
	// The PING port. This parameter applies to TCP PING.
	//
	// example:
	//
	// 80
	PingPort *int32 `json:"ping_port,omitempty" xml:"ping_port,omitempty"`
	// The PING protocol type. Valid values:
	//
	// - icmp
	//
	// - tcp
	//
	// - udp.
	//
	// example:
	//
	// icmp
	PingType *string `json:"ping_type,omitempty" xml:"ping_type,omitempty"`
	// The port for TCP, UDP, SMTP, or POP3 monitoring types.
	//
	// example:
	//
	// 110
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The certificate file name of the private certificate.
	//
	// example:
	//
	// cert.pem
	PrivateCrtFileName *string `json:"private_crt_file_name,omitempty" xml:"private_crt_file_name,omitempty"`
	// The monitoring protocol.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Specifies whether the browser monitoring task uses the QUIC protocol. Valid values:
	//
	// - true: Uses the QUIC protocol.
	//
	// - false: Does not use the QUIC protocol.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	QuicEnabled *bool                                                                     `json:"quic_enabled,omitempty" xml:"quic_enabled,omitempty"`
	QuicTarget  *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget `json:"quic_target,omitempty" xml:"quic_target,omitempty" type:"Struct"`
	// The request content for the HTTP monitoring type.
	//
	// example:
	//
	// aa=bb
	RequestContent *string `json:"request_content,omitempty" xml:"request_content,omitempty"`
	// The format of the HTTP request content. Valid values:
	//
	// - hex: hexadecimal.
	//
	// - txt: text.
	//
	// example:
	//
	// txt
	RequestFormat *string `json:"request_format,omitempty" xml:"request_format,omitempty"`
	// The expected response content to match.
	//
	// example:
	//
	// txt
	ResponseContent *string `json:"response_content,omitempty" xml:"response_content,omitempty"`
	// The format of the HTTP response content. Valid values:
	//
	// - hex: hexadecimal.
	//
	// - txt: text.
	//
	// example:
	//
	// txt
	ResponseFormat *string `json:"response_format,omitempty" xml:"response_format,omitempty"`
	// The number of retries after a monitoring failure.
	//
	// example:
	//
	// 0
	RetryDelay *int32 `json:"retry_delay,omitempty" xml:"retry_delay,omitempty"`
	// This parameter takes effect for SMTP monitoring tasks. Set this parameter to 1 to use a secure connection. Default value: 0.
	//
	// example:
	//
	// 0
	SafeLink *int32 `json:"safe_link,omitempty" xml:"safe_link,omitempty"`
	// Specifies whether to enable page screenshots.
	//
	// example:
	//
	// false
	ScreenShot *bool `json:"screen_shot,omitempty" xml:"screen_shot,omitempty"`
	// For browser monitoring tasks, specifies whether to scroll to the bottom of the page after it is opened.
	//
	// example:
	//
	// false
	ScrollEnd *bool `json:"scroll_end,omitempty" xml:"scroll_end,omitempty"`
	// The Server Name Indication (SNI).
	//
	// example:
	//
	// www.aliyun.com
	ServerName *string                                                              `json:"server_name,omitempty" xml:"server_name,omitempty"`
	Steps      *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Struct"`
	StrictMode *bool                                                                `json:"strict_mode,omitempty" xml:"strict_mode,omitempty"`
	// The supported cipher suites.
	//
	// example:
	//
	// TLS_AES_256_GCM_SHA384
	SupportedCipherSuits *string `json:"supported_cipher_suits,omitempty" xml:"supported_cipher_suits,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// example:
	//
	// 3
	TimeOut *int64 `json:"time_out,omitempty" xml:"time_out,omitempty"`
	// The deployment region of the target application when integrating with Managed Service for OpenTelemetry.
	//
	// example:
	//
	// cn-hangzhou
	TraceRegion *string `json:"trace_region,omitempty" xml:"trace_region,omitempty"`
	// Settings for the Tracing Analysis protocol used when integrating with Managed Service for OpenTelemetry.
	//
	// Valid values:
	//
	// - OpenTelemetry
	//
	// - Zipkin
	//
	// - Jaeger.
	//
	// example:
	//
	// OpenTelemetry
	TraceType                     *string                                                                                      `json:"trace_type,omitempty" xml:"trace_type,omitempty"`
	TrafficHijackElementBlacklist *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist `json:"traffic_hijack_element_blacklist,omitempty" xml:"traffic_hijack_element_blacklist,omitempty" type:"Struct"`
	// When a redirect occurs, if the number of resources loaded by the browser exceeds this value, traffic hijacking is considered to have occurred. When this value is 0, no verification is performed. Default value: 0.
	//
	// example:
	//
	// 0
	TrafficHijackElementCount     *int32                                                                                       `json:"traffic_hijack_element_count,omitempty" xml:"traffic_hijack_element_count,omitempty"`
	TrafficHijackElementWhitelist *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist `json:"traffic_hijack_element_whitelist,omitempty" xml:"traffic_hijack_element_whitelist,omitempty" type:"Struct"`
	// Specifies whether to use a private certificate.
	UsePrivateCrt *bool `json:"use_private_crt,omitempty" xml:"use_private_crt,omitempty"`
	// Specifies whether to use an SSL connection when performing a TCP task.
	//
	// example:
	//
	// tlsv1.0
	UseSsl *bool `json:"use_ssl,omitempty" xml:"use_ssl,omitempty"`
	// The username for FTP, SMTP, or POP3.
	//
	// example:
	//
	// testUser
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// The additional wait time after the page is opened in a browser monitoring task.
	//
	// example:
	//
	// 3
	WaitTimeAfterCompletion *int32 `json:"waitTime_after_completion,omitempty" xml:"waitTime_after_completion,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetAssertions() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions {
	return s.Assertions
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetAttempts() *int64 {
	return s.Attempts
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetAuthInfo() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	return s.AuthInfo
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetBlockedUrlList() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList {
	return s.BlockedUrlList
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetBrowserHeaders() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders {
	return s.BrowserHeaders
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetBrowserHosts() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts {
	return s.BrowserHosts
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetBrowserInfo() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo {
	return s.BrowserInfo
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetBrowserInsecure() *bool {
	return s.BrowserInsecure
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetBrowserTaskVersion() *string {
	return s.BrowserTaskVersion
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetConfigVariables() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables {
	return s.ConfigVariables
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetDiagnosisMtr() *bool {
	return s.DiagnosisMtr
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetDiagnosisPing() *bool {
	return s.DiagnosisPing
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetDnsHijackWhitelist() *string {
	return s.DnsHijackWhitelist
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetDnsMatchRule() *string {
	return s.DnsMatchRule
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetDnsServer() *string {
	return s.DnsServer
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetDnsType() *string {
	return s.DnsType
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetEmptyMessage() *bool {
	return s.EmptyMessage
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetEnablePacketCapture() *bool {
	return s.EnablePacketCapture
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetExpectExistString() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString {
	return s.ExpectExistString
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetExpectNonExistString() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString {
	return s.ExpectNonExistString
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetExpectValue() *string {
	return s.ExpectValue
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetFailureRate() *float32 {
	return s.FailureRate
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetHeader() *string {
	return s.Header
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetHops() *int32 {
	return s.Hops
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetHostBinding() *string {
	return s.HostBinding
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetHostBindingType() *int32 {
	return s.HostBindingType
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetIcmpTimeoutMillis() *int32 {
	return s.IcmpTimeoutMillis
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetIpNetwork() *string {
	return s.IpNetwork
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetIsBase64Encode() *string {
	return s.IsBase64Encode
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetMatchRule() *int32 {
	return s.MatchRule
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetMaxTlsVersion() *string {
	return s.MaxTlsVersion
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetMinTlsVersion() *string {
	return s.MinTlsVersion
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetPassword() *string {
	return s.Password
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetPingNum() *int32 {
	return s.PingNum
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetPingPort() *int32 {
	return s.PingPort
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetPingType() *string {
	return s.PingType
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetPort() *int32 {
	return s.Port
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetPrivateCrtFileName() *string {
	return s.PrivateCrtFileName
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetQuicEnabled() *bool {
	return s.QuicEnabled
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetQuicTarget() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget {
	return s.QuicTarget
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetRequestContent() *string {
	return s.RequestContent
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetRequestFormat() *string {
	return s.RequestFormat
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetResponseContent() *string {
	return s.ResponseContent
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetResponseFormat() *string {
	return s.ResponseFormat
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetRetryDelay() *int32 {
	return s.RetryDelay
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetSafeLink() *int32 {
	return s.SafeLink
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetScreenShot() *bool {
	return s.ScreenShot
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetScrollEnd() *bool {
	return s.ScrollEnd
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetServerName() *string {
	return s.ServerName
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetSteps() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps {
	return s.Steps
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetStrictMode() *bool {
	return s.StrictMode
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetSupportedCipherSuits() *string {
	return s.SupportedCipherSuits
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetTimeOut() *int64 {
	return s.TimeOut
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetTraceRegion() *string {
	return s.TraceRegion
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetTraceType() *string {
	return s.TraceType
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetTrafficHijackElementBlacklist() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist {
	return s.TrafficHijackElementBlacklist
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetTrafficHijackElementCount() *int32 {
	return s.TrafficHijackElementCount
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetTrafficHijackElementWhitelist() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist {
	return s.TrafficHijackElementWhitelist
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetUsePrivateCrt() *bool {
	return s.UsePrivateCrt
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetUseSsl() *bool {
	return s.UseSsl
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetUsername() *string {
	return s.Username
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GetWaitTimeAfterCompletion() *int32 {
	return s.WaitTimeAfterCompletion
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetAssertions(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Assertions = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetAttempts(v int64) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Attempts = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetAuthInfo(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.AuthInfo = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetBlockedUrlList(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.BlockedUrlList = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetBrowserHeaders(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.BrowserHeaders = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetBrowserHosts(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.BrowserHosts = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetBrowserInfo(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.BrowserInfo = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetBrowserInsecure(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.BrowserInsecure = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetBrowserTaskVersion(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.BrowserTaskVersion = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetConfigVariables(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ConfigVariables = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetCookie(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Cookie = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetDiagnosisMtr(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.DiagnosisMtr = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetDiagnosisPing(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.DiagnosisPing = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetDnsHijackWhitelist(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.DnsHijackWhitelist = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetDnsMatchRule(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.DnsMatchRule = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetDnsServer(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.DnsServer = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetDnsType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.DnsType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetEmptyMessage(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.EmptyMessage = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetEnablePacketCapture(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.EnablePacketCapture = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetExpectExistString(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ExpectExistString = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetExpectNonExistString(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ExpectNonExistString = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetExpectValue(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ExpectValue = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetFailureRate(v float32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.FailureRate = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetHeader(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Header = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetHops(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Hops = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetHostBinding(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.HostBinding = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetHostBindingType(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.HostBindingType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetHttpMethod(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.HttpMethod = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetIcmpTimeoutMillis(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.IcmpTimeoutMillis = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetIpNetwork(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.IpNetwork = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetIsBase64Encode(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.IsBase64Encode = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetMatchRule(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.MatchRule = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetMaxTlsVersion(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.MaxTlsVersion = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetMinTlsVersion(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.MinTlsVersion = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetPassword(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Password = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetPingNum(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.PingNum = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetPingPort(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.PingPort = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetPingType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.PingType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetPort(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Port = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetPrivateCrtFileName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.PrivateCrtFileName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetProtocol(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Protocol = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetQuicEnabled(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.QuicEnabled = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetQuicTarget(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.QuicTarget = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetRequestContent(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.RequestContent = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetRequestFormat(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.RequestFormat = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetResponseContent(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ResponseContent = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetResponseFormat(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ResponseFormat = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetRetryDelay(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.RetryDelay = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetSafeLink(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.SafeLink = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetScreenShot(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ScreenShot = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetScrollEnd(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ScrollEnd = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetServerName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ServerName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetSteps(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Steps = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetStrictMode(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.StrictMode = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetSupportedCipherSuits(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.SupportedCipherSuits = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetTimeOut(v int64) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.TimeOut = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetTraceRegion(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.TraceRegion = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetTraceType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.TraceType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetTrafficHijackElementBlacklist(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.TrafficHijackElementBlacklist = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetTrafficHijackElementCount(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.TrafficHijackElementCount = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetTrafficHijackElementWhitelist(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.TrafficHijackElementWhitelist = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetUsePrivateCrt(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.UsePrivateCrt = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetUseSsl(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.UseSsl = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetUsername(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Username = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetWaitTimeAfterCompletion(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.WaitTimeAfterCompletion = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) Validate() error {
	if s.Assertions != nil {
		if err := s.Assertions.Validate(); err != nil {
			return err
		}
	}
	if s.AuthInfo != nil {
		if err := s.AuthInfo.Validate(); err != nil {
			return err
		}
	}
	if s.BlockedUrlList != nil {
		if err := s.BlockedUrlList.Validate(); err != nil {
			return err
		}
	}
	if s.BrowserHeaders != nil {
		if err := s.BrowserHeaders.Validate(); err != nil {
			return err
		}
	}
	if s.BrowserHosts != nil {
		if err := s.BrowserHosts.Validate(); err != nil {
			return err
		}
	}
	if s.BrowserInfo != nil {
		if err := s.BrowserInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ConfigVariables != nil {
		if err := s.ConfigVariables.Validate(); err != nil {
			return err
		}
	}
	if s.ExpectExistString != nil {
		if err := s.ExpectExistString.Validate(); err != nil {
			return err
		}
	}
	if s.ExpectNonExistString != nil {
		if err := s.ExpectNonExistString.Validate(); err != nil {
			return err
		}
	}
	if s.QuicTarget != nil {
		if err := s.QuicTarget.Validate(); err != nil {
			return err
		}
	}
	if s.Steps != nil {
		if err := s.Steps.Validate(); err != nil {
			return err
		}
	}
	if s.TrafficHijackElementBlacklist != nil {
		if err := s.TrafficHijackElementBlacklist.Validate(); err != nil {
			return err
		}
	}
	if s.TrafficHijackElementWhitelist != nil {
		if err := s.TrafficHijackElementWhitelist.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions struct {
	Assertions []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions `json:"assertions,omitempty" xml:"assertions,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions) GetAssertions() []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions {
	return s.Assertions
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions) SetAssertions(v []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions {
	s.Assertions = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertions) Validate() error {
	if s.Assertions != nil {
		for _, item := range s.Assertions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions struct {
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Property *string `json:"property,omitempty" xml:"property,omitempty"`
	Target   *string `json:"target,omitempty" xml:"target,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) GetOperator() *string {
	return s.Operator
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) GetProperty() *string {
	return s.Property
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) GetTarget() *string {
	return s.Target
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) GetType() *string {
	return s.Type
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) SetOperator(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions {
	s.Operator = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) SetProperty(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions {
	s.Property = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) SetTarget(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions {
	s.Target = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) SetType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions {
	s.Type = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAssertionsAssertions) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo struct {
	// Supported only in multi-step monitoring. The AccessKey ID used for Alibaba Cloud authentication. We recommend that you use encrypted storage.
	//
	// example:
	//
	// testAk
	AccessKeyId *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
	// Supported only in multi-step monitoring. The AccessKey secret used for Alibaba Cloud authentication. We recommend that you use encrypted storage.
	//
	// example:
	//
	// testSk
	AccessKeySecret *string `json:"access_key_secret,omitempty" xml:"access_key_secret,omitempty"`
	// Supported only in multi-step monitoring. The API action of the request when using Alibaba Cloud operations.
	//
	// example:
	//
	// example_api
	ApiAction *string `json:"api_action,omitempty" xml:"api_action,omitempty"`
	// Supported only in multi-step monitoring. The API version of the request when using Alibaba Cloud operations.
	//
	// example:
	//
	// 2019-01-01
	ApiVersion *string `json:"api_version,omitempty" xml:"api_version,omitempty"`
	// The OAuth 2.0 authentication style. Valid values: ROA and RPC.
	//
	// example:
	//
	// ROA
	AuthStyle *string `json:"auth_style,omitempty" xml:"auth_style,omitempty"`
	// The client ID used for client authentication in OAuth 2.0.
	//
	// example:
	//
	// client_id
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// The client secret used for client authentication in OAuth 2.0.
	//
	// example:
	//
	// client_secret
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	// The grant type used in OAuth 2.0 authentication. Valid values: client_credentials and password.
	//
	// example:
	//
	// password
	GrantType *string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	// The password used for HTTP Basic Authentication.
	//
	// example:
	//
	// test_password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// Supported only in multi-step monitoring. The region ID of the request when using Alibaba Cloud authentication.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string                                                                       `json:"region_id,omitempty" xml:"region_id,omitempty"`
	Scopes   *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Struct"`
	// The service name of the request when using AWS authentication.
	//
	// example:
	//
	// example_service_name
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// The session token used for AWS authentication.
	//
	// example:
	//
	// example_token
	SessionToken *string `json:"session_token,omitempty" xml:"session_token,omitempty"`
	// The authorization server URL in OAuth 2.0.
	//
	// example:
	//
	// https://www.example.com
	TokenUrl *string `json:"token_url,omitempty" xml:"token_url,omitempty"`
	// The authentication type. HTTP Basic Authentication is supported. Valid values: basic.
	//
	// example:
	//
	// basic
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether the key is stored in the client cookie for digest authentication.
	//
	// example:
	//
	// false
	UseCookieSessionKey *bool `json:"use_cookie_session_key,omitempty" xml:"use_cookie_session_key,omitempty"`
	// The username used for HTTP Basic Authentication.
	//
	// example:
	//
	// test_username
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// Supported only in multi-step monitoring. Specifies whether additional resources exist when using Alibaba Cloud authentication for this step.
	//
	// example:
	//
	// false
	WithAddonResources *bool `json:"with_addon_resources,omitempty" xml:"with_addon_resources,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetApiAction() *string {
	return s.ApiAction
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetAuthStyle() *string {
	return s.AuthStyle
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetGrantType() *string {
	return s.GrantType
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetPassword() *string {
	return s.Password
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetScopes() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes {
	return s.Scopes
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetSessionToken() *string {
	return s.SessionToken
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetTokenUrl() *string {
	return s.TokenUrl
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetType() *string {
	return s.Type
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetUseCookieSessionKey() *bool {
	return s.UseCookieSessionKey
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetUsername() *string {
	return s.Username
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) GetWithAddonResources() *bool {
	return s.WithAddonResources
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetAccessKeyId(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetAccessKeySecret(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.AccessKeySecret = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetApiAction(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.ApiAction = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetApiVersion(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.ApiVersion = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetAuthStyle(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.AuthStyle = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetClientId(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.ClientId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetClientSecret(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.ClientSecret = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetGrantType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.GrantType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetPassword(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.Password = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetRegionId(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetScopes(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.Scopes = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetServiceName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.ServiceName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetSessionToken(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.SessionToken = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetTokenUrl(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.TokenUrl = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.Type = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetUseCookieSessionKey(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.UseCookieSessionKey = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetUsername(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.Username = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) SetWithAddonResources(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo {
	s.WithAddonResources = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfo) Validate() error {
	if s.Scopes != nil {
		if err := s.Scopes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes struct {
	Scopes []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes) GetScopes() []*string {
	return s.Scopes
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes) SetScopes(v []*string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes {
	s.Scopes = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonAuthInfoScopes) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList struct {
	BlockedUrlList []*string `json:"blocked_url_list,omitempty" xml:"blocked_url_list,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList) GetBlockedUrlList() []*string {
	return s.BlockedUrlList
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList) SetBlockedUrlList(v []*string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList {
	s.BlockedUrlList = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBlockedUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders struct {
	BrowserHeaders []map[string]interface{} `json:"browser_headers,omitempty" xml:"browser_headers,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders) GetBrowserHeaders() []map[string]interface{} {
	return s.BrowserHeaders
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders) SetBrowserHeaders(v []map[string]interface{}) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders {
	s.BrowserHeaders = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts struct {
	BrowserHosts []*string `json:"browser_hosts,omitempty" xml:"browser_hosts,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts) GetBrowserHosts() []*string {
	return s.BrowserHosts
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts) SetBrowserHosts(v []*string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts {
	s.BrowserHosts = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserHosts) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo struct {
	BrowserInfo []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo `json:"browser_info,omitempty" xml:"browser_info,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo) GetBrowserInfo() []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo {
	return s.BrowserInfo
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo) SetBrowserInfo(v []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo {
	s.BrowserInfo = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfo) Validate() error {
	if s.BrowserInfo != nil {
		for _, item := range s.BrowserInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo struct {
	Browser *string `json:"browser,omitempty" xml:"browser,omitempty"`
	Device  *string `json:"device,omitempty" xml:"device,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo) GetBrowser() *string {
	return s.Browser
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo) GetDevice() *string {
	return s.Device
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo) SetBrowser(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo {
	s.Browser = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo) SetDevice(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo {
	s.Device = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonBrowserInfoBrowserInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables struct {
	ConfigVariables []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables `json:"config_variables,omitempty" xml:"config_variables,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables) GetConfigVariables() []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables {
	return s.ConfigVariables
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables) SetConfigVariables(v []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables {
	s.ConfigVariables = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariables) Validate() error {
	if s.ConfigVariables != nil {
		for _, item := range s.ConfigVariables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Secure *bool   `json:"secure,omitempty" xml:"secure,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) GetName() *string {
	return s.Name
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) GetSecure() *bool {
	return s.Secure
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) GetValue() *string {
	return s.Value
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) SetName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables {
	s.Name = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) SetSecure(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables {
	s.Secure = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) SetValue(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables {
	s.Value = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonConfigVariablesConfigVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString struct {
	ExpectExistString []*string `json:"expect_exist_string,omitempty" xml:"expect_exist_string,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString) GetExpectExistString() []*string {
	return s.ExpectExistString
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString) SetExpectExistString(v []*string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString {
	s.ExpectExistString = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectExistString) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString struct {
	ExpectNonExistString []*string `json:"expect_non_exist_string,omitempty" xml:"expect_non_exist_string,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString) GetExpectNonExistString() []*string {
	return s.ExpectNonExistString
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString) SetExpectNonExistString(v []*string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString {
	s.ExpectNonExistString = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonExpectNonExistString) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget struct {
	QuicTarget []*string `json:"quic_target,omitempty" xml:"quic_target,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget) GetQuicTarget() []*string {
	return s.QuicTarget
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget) SetQuicTarget(v []*string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget {
	s.QuicTarget = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonQuicTarget) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps struct {
	Steps []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps) GetSteps() []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	return s.Steps
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps) SetSteps(v []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps {
	s.Steps = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonSteps) Validate() error {
	if s.Steps != nil {
		for _, item := range s.Steps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps struct {
	AllowFailure       *bool                                                                                       `json:"allow_failure,omitempty" xml:"allow_failure,omitempty"`
	AutoExtractCookie  *bool                                                                                       `json:"auto_extract_cookie,omitempty" xml:"auto_extract_cookie,omitempty"`
	ExtractedVariables *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables `json:"extracted_variables,omitempty" xml:"extracted_variables,omitempty" type:"Struct"`
	IsCritical         *bool                                                                                       `json:"is_critical,omitempty" xml:"is_critical,omitempty"`
	// Deprecated
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	Option             *string `json:"option,omitempty" xml:"option,omitempty"`
	StepName           *string `json:"step_name,omitempty" xml:"step_name,omitempty"`
	StepType           *string `json:"step_type,omitempty" xml:"step_type,omitempty"`
	Url                *string `json:"url,omitempty" xml:"url,omitempty"`
	UseGeneratedCookie *bool   `json:"use_generated_cookie,omitempty" xml:"use_generated_cookie,omitempty"`
	WaitTimeInSecs     *int32  `json:"wait_time_in_secs,omitempty" xml:"wait_time_in_secs,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetAllowFailure() *bool {
	return s.AllowFailure
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetAutoExtractCookie() *bool {
	return s.AutoExtractCookie
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetExtractedVariables() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables {
	return s.ExtractedVariables
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetIsCritical() *bool {
	return s.IsCritical
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetName() *string {
	return s.Name
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetOption() *string {
	return s.Option
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetStepName() *string {
	return s.StepName
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetStepType() *string {
	return s.StepType
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetUrl() *string {
	return s.Url
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetUseGeneratedCookie() *bool {
	return s.UseGeneratedCookie
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) GetWaitTimeInSecs() *int32 {
	return s.WaitTimeInSecs
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetAllowFailure(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.AllowFailure = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetAutoExtractCookie(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.AutoExtractCookie = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetExtractedVariables(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.ExtractedVariables = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetIsCritical(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.IsCritical = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.Name = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetOption(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.Option = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetStepName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.StepName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetStepType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.StepType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetUrl(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.Url = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetUseGeneratedCookie(v bool) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.UseGeneratedCookie = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) SetWaitTimeInSecs(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps {
	s.WaitTimeInSecs = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsSteps) Validate() error {
	if s.ExtractedVariables != nil {
		if err := s.ExtractedVariables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables struct {
	ExtractedVariables []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables `json:"extracted_variables,omitempty" xml:"extracted_variables,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables) GetExtractedVariables() []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables {
	return s.ExtractedVariables
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables) SetExtractedVariables(v []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables {
	s.ExtractedVariables = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariables) Validate() error {
	if s.ExtractedVariables != nil {
		for _, item := range s.ExtractedVariables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables struct {
	ExtractedType *string                                                                                                             `json:"extracted_type,omitempty" xml:"extracted_type,omitempty"`
	Field         *string                                                                                                             `json:"field,omitempty" xml:"field,omitempty"`
	Name          *string                                                                                                             `json:"name,omitempty" xml:"name,omitempty"`
	Parser        *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser `json:"parser,omitempty" xml:"parser,omitempty" type:"Struct"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) GetExtractedType() *string {
	return s.ExtractedType
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) GetField() *string {
	return s.Field
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) GetName() *string {
	return s.Name
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) GetParser() *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser {
	return s.Parser
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) SetExtractedType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables {
	s.ExtractedType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) SetField(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables {
	s.Field = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) SetName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables {
	s.Name = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) SetParser(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables {
	s.Parser = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariables) Validate() error {
	if s.Parser != nil {
		if err := s.Parser.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser struct {
	ParserType *string `json:"parser_type,omitempty" xml:"parser_type,omitempty"`
	Value      *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser) GetParserType() *string {
	return s.ParserType
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser) GetValue() *string {
	return s.Value
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser) SetParserType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser {
	s.ParserType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser) SetValue(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser {
	s.Value = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonStepsStepsExtractedVariablesExtractedVariablesParser) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist struct {
	TrafficHijackElementBlacklist []*string `json:"traffic_hijack_element_blacklist,omitempty" xml:"traffic_hijack_element_blacklist,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist) GetTrafficHijackElementBlacklist() []*string {
	return s.TrafficHijackElementBlacklist
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist) SetTrafficHijackElementBlacklist(v []*string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist {
	s.TrafficHijackElementBlacklist = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementBlacklist) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist struct {
	TrafficHijackElementWhitelist []*string `json:"traffic_hijack_element_whitelist,omitempty" xml:"traffic_hijack_element_whitelist,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist) GetTrafficHijackElementWhitelist() []*string {
	return s.TrafficHijackElementWhitelist
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist) SetTrafficHijackElementWhitelist(v []*string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist {
	s.TrafficHijackElementWhitelist = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJsonTrafficHijackElementWhitelist) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig struct {
	// The region where the target site of the internal network monitoring task is located.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the security group associated with the internal network monitoring task.
	//
	// example:
	//
	// sg-xxxxxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the VPC associated with the internal network monitoring task.
	//
	// example:
	//
	// vpc-xxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch associated with the internal network monitoring task.
	//
	// example:
	//
	// vsw-xxxxxx
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) GetRegion() *string {
	return s.Region
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) SetRegion(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig {
	s.Region = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) SetSecurityGroupId(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) SetVpcId(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) SetVswitchId(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig {
	s.VswitchId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsVpcConfig) Validate() error {
	return dara.Validate(s)
}
