// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateAlertRuleShrinkRequest
	GetEnabled() *bool
	SetId(v int64) *UpdateAlertRuleShrinkRequest
	GetId() *int64
	SetName(v string) *UpdateAlertRuleShrinkRequest
	GetName() *string
	SetNotificationShrink(v string) *UpdateAlertRuleShrinkRequest
	GetNotificationShrink() *string
	SetOwner(v string) *UpdateAlertRuleShrinkRequest
	GetOwner() *string
	SetTriggerConditionShrink(v string) *UpdateAlertRuleShrinkRequest
	GetTriggerConditionShrink() *string
}

type UpdateAlertRuleShrinkRequest struct {
	// Specifies whether to enable the rule.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 105412
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// collection_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration for the alert notification.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The ID of the Alibaba Cloud account used by the owner of the rule.
	//
	// example:
	//
	// 193379****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The alert triggering condition.
	TriggerConditionShrink *string `json:"TriggerCondition,omitempty" xml:"TriggerCondition,omitempty"`
}

func (s UpdateAlertRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAlertRuleShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateAlertRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAlertRuleShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *UpdateAlertRuleShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateAlertRuleShrinkRequest) GetTriggerConditionShrink() *string {
	return s.TriggerConditionShrink
}

func (s *UpdateAlertRuleShrinkRequest) SetEnabled(v bool) *UpdateAlertRuleShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateAlertRuleShrinkRequest) SetId(v int64) *UpdateAlertRuleShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateAlertRuleShrinkRequest) SetName(v string) *UpdateAlertRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlertRuleShrinkRequest) SetNotificationShrink(v string) *UpdateAlertRuleShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *UpdateAlertRuleShrinkRequest) SetOwner(v string) *UpdateAlertRuleShrinkRequest {
	s.Owner = &v
	return s
}

func (s *UpdateAlertRuleShrinkRequest) SetTriggerConditionShrink(v string) *UpdateAlertRuleShrinkRequest {
	s.TriggerConditionShrink = &v
	return s
}

func (s *UpdateAlertRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
