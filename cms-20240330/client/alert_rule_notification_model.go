// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleNotification interface {
	dara.Model
	String() string
	GoString() string
	SetContacts(v []*string) *AlertRuleNotification
	GetContacts() []*string
	SetCustomWebhooks(v []*string) *AlertRuleNotification
	GetCustomWebhooks() []*string
	SetDingCoolAppWebhooks(v []*string) *AlertRuleNotification
	GetDingCoolAppWebhooks() []*string
	SetDingWebhooks(v []*string) *AlertRuleNotification
	GetDingWebhooks() []*string
	SetFsWebhooks(v []*string) *AlertRuleNotification
	GetFsWebhooks() []*string
	SetGroups(v []*string) *AlertRuleNotification
	GetGroups() []*string
	SetNotifyTime(v *AlertRuleTimeSpan) *AlertRuleNotification
	GetNotifyTime() *AlertRuleTimeSpan
	SetQwencloudContacts(v map[string]map[string]interface{}) *AlertRuleNotification
	GetQwencloudContacts() map[string]map[string]interface{}
	SetSendOk(v bool) *AlertRuleNotification
	GetSendOk() *bool
	SetSeverityNotifications(v map[string]*SeverityNotifyConfig) *AlertRuleNotification
	GetSeverityNotifications() map[string]*SeverityNotifyConfig
	SetSilenceTime(v int64) *AlertRuleNotification
	GetSilenceTime() *int64
	SetSlackWebhooks(v []*string) *AlertRuleNotification
	GetSlackWebhooks() []*string
	SetWxWebhooks(v []*string) *AlertRuleNotification
	GetWxWebhooks() []*string
}

type AlertRuleNotification struct {
	// The list of contact IDs.
	Contacts []*string `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// The list of custom webhook Notification Recipient IDs.
	CustomWebhooks []*string `json:"customWebhooks,omitempty" xml:"customWebhooks,omitempty" type:"Repeated"`
	// The list of DingTalk Cool App webhook Notification Recipient IDs.
	DingCoolAppWebhooks []*string `json:"dingCoolAppWebhooks,omitempty" xml:"dingCoolAppWebhooks,omitempty" type:"Repeated"`
	// The list of DingTalk webhook Notification Recipient IDs.
	DingWebhooks []*string `json:"dingWebhooks,omitempty" xml:"dingWebhooks,omitempty" type:"Repeated"`
	// The list of Lark webhook Notification Recipient IDs.
	FsWebhooks []*string `json:"fsWebhooks,omitempty" xml:"fsWebhooks,omitempty" type:"Repeated"`
	// The list of contact group IDs.
	Groups []*string `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// The notification time period. Notifications are sent only during this time period.
	NotifyTime            *AlertRuleTimeSpan                `json:"notifyTime,omitempty" xml:"notifyTime,omitempty"`
	QwencloudContacts     map[string]map[string]interface{} `json:"qwencloudContacts,omitempty" xml:"qwencloudContacts,omitempty"`
	SendOk                *bool                             `json:"sendOk,omitempty" xml:"sendOk,omitempty"`
	SeverityNotifications map[string]*SeverityNotifyConfig  `json:"severityNotifications,omitempty" xml:"severityNotifications,omitempty"`
	// The notification mute duration, in seconds.
	//
	// example:
	//
	// 86400
	SilenceTime *int64 `json:"silenceTime,omitempty" xml:"silenceTime,omitempty"`
	// The list of Slack webhook Notification Recipient IDs.
	SlackWebhooks []*string `json:"slackWebhooks,omitempty" xml:"slackWebhooks,omitempty" type:"Repeated"`
	// The list of WeChat webhook Notification Recipient IDs.
	WxWebhooks []*string `json:"wxWebhooks,omitempty" xml:"wxWebhooks,omitempty" type:"Repeated"`
}

func (s AlertRuleNotification) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleNotification) GoString() string {
	return s.String()
}

func (s *AlertRuleNotification) GetContacts() []*string {
	return s.Contacts
}

func (s *AlertRuleNotification) GetCustomWebhooks() []*string {
	return s.CustomWebhooks
}

func (s *AlertRuleNotification) GetDingCoolAppWebhooks() []*string {
	return s.DingCoolAppWebhooks
}

func (s *AlertRuleNotification) GetDingWebhooks() []*string {
	return s.DingWebhooks
}

func (s *AlertRuleNotification) GetFsWebhooks() []*string {
	return s.FsWebhooks
}

func (s *AlertRuleNotification) GetGroups() []*string {
	return s.Groups
}

func (s *AlertRuleNotification) GetNotifyTime() *AlertRuleTimeSpan {
	return s.NotifyTime
}

func (s *AlertRuleNotification) GetQwencloudContacts() map[string]map[string]interface{} {
	return s.QwencloudContacts
}

func (s *AlertRuleNotification) GetSendOk() *bool {
	return s.SendOk
}

func (s *AlertRuleNotification) GetSeverityNotifications() map[string]*SeverityNotifyConfig {
	return s.SeverityNotifications
}

func (s *AlertRuleNotification) GetSilenceTime() *int64 {
	return s.SilenceTime
}

func (s *AlertRuleNotification) GetSlackWebhooks() []*string {
	return s.SlackWebhooks
}

func (s *AlertRuleNotification) GetWxWebhooks() []*string {
	return s.WxWebhooks
}

func (s *AlertRuleNotification) SetContacts(v []*string) *AlertRuleNotification {
	s.Contacts = v
	return s
}

func (s *AlertRuleNotification) SetCustomWebhooks(v []*string) *AlertRuleNotification {
	s.CustomWebhooks = v
	return s
}

func (s *AlertRuleNotification) SetDingCoolAppWebhooks(v []*string) *AlertRuleNotification {
	s.DingCoolAppWebhooks = v
	return s
}

func (s *AlertRuleNotification) SetDingWebhooks(v []*string) *AlertRuleNotification {
	s.DingWebhooks = v
	return s
}

func (s *AlertRuleNotification) SetFsWebhooks(v []*string) *AlertRuleNotification {
	s.FsWebhooks = v
	return s
}

func (s *AlertRuleNotification) SetGroups(v []*string) *AlertRuleNotification {
	s.Groups = v
	return s
}

func (s *AlertRuleNotification) SetNotifyTime(v *AlertRuleTimeSpan) *AlertRuleNotification {
	s.NotifyTime = v
	return s
}

func (s *AlertRuleNotification) SetQwencloudContacts(v map[string]map[string]interface{}) *AlertRuleNotification {
	s.QwencloudContacts = v
	return s
}

func (s *AlertRuleNotification) SetSendOk(v bool) *AlertRuleNotification {
	s.SendOk = &v
	return s
}

func (s *AlertRuleNotification) SetSeverityNotifications(v map[string]*SeverityNotifyConfig) *AlertRuleNotification {
	s.SeverityNotifications = v
	return s
}

func (s *AlertRuleNotification) SetSilenceTime(v int64) *AlertRuleNotification {
	s.SilenceTime = &v
	return s
}

func (s *AlertRuleNotification) SetSlackWebhooks(v []*string) *AlertRuleNotification {
	s.SlackWebhooks = v
	return s
}

func (s *AlertRuleNotification) SetWxWebhooks(v []*string) *AlertRuleNotification {
	s.WxWebhooks = v
	return s
}

func (s *AlertRuleNotification) Validate() error {
	if s.NotifyTime != nil {
		if err := s.NotifyTime.Validate(); err != nil {
			return err
		}
	}
	return nil
}
