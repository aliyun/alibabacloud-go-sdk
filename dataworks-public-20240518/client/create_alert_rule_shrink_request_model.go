// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *CreateAlertRuleShrinkRequest
	GetEnabled() *bool
	SetName(v string) *CreateAlertRuleShrinkRequest
	GetName() *string
	SetNotificationShrink(v string) *CreateAlertRuleShrinkRequest
	GetNotificationShrink() *string
	SetOwner(v string) *CreateAlertRuleShrinkRequest
	GetOwner() *string
	SetTriggerConditionShrink(v string) *CreateAlertRuleShrinkRequest
	GetTriggerConditionShrink() *string
}

type CreateAlertRuleShrinkRequest struct {
	// Indicates whether the rule is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// xm_create_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration for the alert notification.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The ID of the Alibaba Cloud account used by the owner of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 279114181716147735
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The alert triggering condition.
	//
	// This parameter is required.
	TriggerConditionShrink *string `json:"TriggerCondition,omitempty" xml:"TriggerCondition,omitempty"`
}

func (s CreateAlertRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAlertRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlertRuleShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateAlertRuleShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateAlertRuleShrinkRequest) GetTriggerConditionShrink() *string {
	return s.TriggerConditionShrink
}

func (s *CreateAlertRuleShrinkRequest) SetEnabled(v bool) *CreateAlertRuleShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreateAlertRuleShrinkRequest) SetName(v string) *CreateAlertRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAlertRuleShrinkRequest) SetNotificationShrink(v string) *CreateAlertRuleShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateAlertRuleShrinkRequest) SetOwner(v string) *CreateAlertRuleShrinkRequest {
	s.Owner = &v
	return s
}

func (s *CreateAlertRuleShrinkRequest) SetTriggerConditionShrink(v string) *CreateAlertRuleShrinkRequest {
	s.TriggerConditionShrink = &v
	return s
}

func (s *CreateAlertRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
