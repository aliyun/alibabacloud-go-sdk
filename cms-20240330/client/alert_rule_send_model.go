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
	SetSendToArms(v bool) *AlertRuleSend
	GetSendToArms() *bool
}

type AlertRuleSend struct {
	Action           *AlertRuleAction       `json:"action,omitempty" xml:"action,omitempty"`
	Notification     *AlertRuleNotification `json:"notification,omitempty" xml:"notification,omitempty"`
	NotifyStrategies []*string              `json:"notifyStrategies,omitempty" xml:"notifyStrategies,omitempty" type:"Repeated"`
	SendToArms       *bool                  `json:"sendToArms,omitempty" xml:"sendToArms,omitempty"`
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
	return nil
}
