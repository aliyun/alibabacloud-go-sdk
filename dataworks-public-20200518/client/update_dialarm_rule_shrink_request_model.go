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
	SetDescription(v string) *UpdateDIAlarmRuleShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdateDIAlarmRuleShrinkRequest
	GetEnabled() *bool
	SetMetricType(v string) *UpdateDIAlarmRuleShrinkRequest
	GetMetricType() *string
	SetNotificationSettingsShrink(v string) *UpdateDIAlarmRuleShrinkRequest
	GetNotificationSettingsShrink() *string
	SetTriggerConditionsShrink(v string) *UpdateDIAlarmRuleShrinkRequest
	GetTriggerConditionsShrink() *string
}

type UpdateDIAlarmRuleShrinkRequest struct {
	// The alert rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34982
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The description of the alert rule.
	//
	// example:
	//
	// mysql synchronizes to hologres heartbeat alert
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the alert rule. By default, the alert rule is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The metric type in the alert rule. Valid values:
	//
	// 	- Heartbeat
	//
	// 	- FailoverCount
	//
	// 	- Delay
	//
	// This parameter is required.
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The alert notification settings.
	//
	// This parameter is required.
	NotificationSettingsShrink *string `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty"`
	// The conditions that are used to trigger the alert rule.
	//
	// This parameter is required.
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

func (s *UpdateDIAlarmRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateDIAlarmRuleShrinkRequest) GetMetricType() *string {
	return s.MetricType
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

func (s *UpdateDIAlarmRuleShrinkRequest) SetDescription(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetEnabled(v bool) *UpdateDIAlarmRuleShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetMetricType(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.MetricType = &v
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
