// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleSend interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v *AlertRuleAction) *AlertRuleSend
	GetAction() *AlertRuleAction
	SetNotification(v *AlertRuleNotification) *AlertRuleSend
	GetNotification() *AlertRuleNotification
	SetNotifyStrategies(v []*string) *AlertRuleSend
	GetNotifyStrategies() []*string
	SetRcaConfig(v *AlertRuleRcaConfig) *AlertRuleSend
	GetRcaConfig() *AlertRuleRcaConfig
	SetSendToArms(v bool) *AlertRuleSend
	GetSendToArms() *bool
}

type AlertRuleSend struct {
	// The alert action integration configuration.
	Action *AlertRuleAction `json:"action,omitempty" xml:"action,omitempty"`
	// The alert notification configuration.
	Notification *AlertRuleNotification `json:"notification,omitempty" xml:"notification,omitempty"`
	// The list of notification policies that define the notification methods or Policy Name values to use for different alert states, such as triggered and recover.
	NotifyStrategies []*string `json:"notifyStrategies,omitempty" xml:"notifyStrategies,omitempty" type:"Repeated"`
	// The root cause analysis (RCA) configuration. After this feature is enabled, root cause analysis is automatically performed when an alert is triggered.
	RcaConfig *AlertRuleRcaConfig `json:"rcaConfig,omitempty" xml:"rcaConfig,omitempty"`
	// Indicates whether event delivery of alert events to ARMS alert management is enabled.
	//
	// example:
	//
	// true
	SendToArms *bool `json:"sendToArms,omitempty" xml:"sendToArms,omitempty"`
}

func (s AlertRuleSend) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleSend) GoString() string {
	return s.String()
}

func (s *AlertRuleSend) GetAction() *AlertRuleAction {
	return s.Action
}

func (s *AlertRuleSend) GetNotification() *AlertRuleNotification {
	return s.Notification
}

func (s *AlertRuleSend) GetNotifyStrategies() []*string {
	return s.NotifyStrategies
}

func (s *AlertRuleSend) GetRcaConfig() *AlertRuleRcaConfig {
	return s.RcaConfig
}

func (s *AlertRuleSend) GetSendToArms() *bool {
	return s.SendToArms
}

func (s *AlertRuleSend) SetAction(v *AlertRuleAction) *AlertRuleSend {
	s.Action = v
	return s
}

func (s *AlertRuleSend) SetNotification(v *AlertRuleNotification) *AlertRuleSend {
	s.Notification = v
	return s
}

func (s *AlertRuleSend) SetNotifyStrategies(v []*string) *AlertRuleSend {
	s.NotifyStrategies = v
	return s
}

func (s *AlertRuleSend) SetRcaConfig(v *AlertRuleRcaConfig) *AlertRuleSend {
	s.RcaConfig = v
	return s
}

func (s *AlertRuleSend) SetSendToArms(v bool) *AlertRuleSend {
	s.SendToArms = &v
	return s
}

func (s *AlertRuleSend) Validate() error {
	if s.Action != nil {
		if err := s.Action.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.RcaConfig != nil {
		if err := s.RcaConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
