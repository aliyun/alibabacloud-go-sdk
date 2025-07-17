// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIAlarmRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIAlarmRuleId(v int64) *UpdateDIAlarmRuleShrinkRequest
	GetDIAlarmRuleId() *int64
	SetDIJobId(v int64) *UpdateDIAlarmRuleShrinkRequest
	GetDIJobId() *int64
	SetDescription(v string) *UpdateDIAlarmRuleShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdateDIAlarmRuleShrinkRequest
	GetEnabled() *bool
	SetId(v int64) *UpdateDIAlarmRuleShrinkRequest
	GetId() *int64
	SetMetricType(v string) *UpdateDIAlarmRuleShrinkRequest
	GetMetricType() *string
	SetName(v string) *UpdateDIAlarmRuleShrinkRequest
	GetName() *string
	SetNotificationSettingsShrink(v string) *UpdateDIAlarmRuleShrinkRequest
	GetNotificationSettingsShrink() *string
	SetTriggerConditionsShrink(v string) *UpdateDIAlarmRuleShrinkRequest
	GetTriggerConditionsShrink() *string
}

type UpdateDIAlarmRuleShrinkRequest struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 34982
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 1
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the alert rule.
	//
	// example:
	//
	// The description of the alert rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the alert rule. By default, the alert rule is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The alert rule Id
	//
	// example:
	//
	// 34982
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metric type in the alert rule. Valid values:
	//
	// 	- Heartbeat
	//
	// 	- FailoverCount
	//
	// 	- Delay
	//
	// 	- DdlReport
	//
	// 	- ResourceUtilization
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// alarm_rule_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert notification settings.
	NotificationSettingsShrink *string `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty"`
	// The conditions that can trigger the alert rule.
	TriggerConditionsShrink *string `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty"`
}

func (s UpdateDIAlarmRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIAlarmRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetDIAlarmRuleId() *int64 {
	return s.DIAlarmRuleId
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetNotificationSettingsShrink() *string {
	return s.NotificationSettingsShrink
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetTriggerConditionsShrink() *string {
	return s.TriggerConditionsShrink
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetDIAlarmRuleId(v int64) *UpdateDIAlarmRuleShrinkRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetDIJobId(v int64) *UpdateDIAlarmRuleShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetDescription(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetEnabled(v bool) *UpdateDIAlarmRuleShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetId(v int64) *UpdateDIAlarmRuleShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetMetricType(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.MetricType = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetName(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetNotificationSettingsShrink(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.NotificationSettingsShrink = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetTriggerConditionsShrink(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.TriggerConditionsShrink = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
