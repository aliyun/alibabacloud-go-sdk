// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIAlarmRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDIAlarmRuleShrinkRequest
	GetClientToken() *string
	SetDIJobId(v int64) *CreateDIAlarmRuleShrinkRequest
	GetDIJobId() *int64
	SetDescription(v string) *CreateDIAlarmRuleShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateDIAlarmRuleShrinkRequest
	GetEnabled() *bool
	SetMetricType(v string) *CreateDIAlarmRuleShrinkRequest
	GetMetricType() *string
	SetName(v string) *CreateDIAlarmRuleShrinkRequest
	GetName() *string
	SetNotificationSettingsShrink(v string) *CreateDIAlarmRuleShrinkRequest
	GetNotificationSettingsShrink() *string
	SetTriggerConditionsShrink(v string) *CreateDIAlarmRuleShrinkRequest
	GetTriggerConditionsShrink() *string
}

type CreateDIAlarmRuleShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the synchronization task with which the alert rule is associated.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The name of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// alartRule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert notification settings.
	//
	// This parameter is required.
	NotificationSettingsShrink *string `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty"`
	// The conditions that can trigger the alert rule.
	//
	// This parameter is required.
	TriggerConditionsShrink *string `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty"`
}

func (s CreateDIAlarmRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDIAlarmRuleShrinkRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *CreateDIAlarmRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDIAlarmRuleShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateDIAlarmRuleShrinkRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *CreateDIAlarmRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateDIAlarmRuleShrinkRequest) GetNotificationSettingsShrink() *string {
	return s.NotificationSettingsShrink
}

func (s *CreateDIAlarmRuleShrinkRequest) GetTriggerConditionsShrink() *string {
	return s.TriggerConditionsShrink
}

func (s *CreateDIAlarmRuleShrinkRequest) SetClientToken(v string) *CreateDIAlarmRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetDIJobId(v int64) *CreateDIAlarmRuleShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetDescription(v string) *CreateDIAlarmRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetEnabled(v bool) *CreateDIAlarmRuleShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetMetricType(v string) *CreateDIAlarmRuleShrinkRequest {
	s.MetricType = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetName(v string) *CreateDIAlarmRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetNotificationSettingsShrink(v string) *CreateDIAlarmRuleShrinkRequest {
	s.NotificationSettingsShrink = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetTriggerConditionsShrink(v string) *CreateDIAlarmRuleShrinkRequest {
	s.TriggerConditionsShrink = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
