// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleNotificationFilter interface {
	dara.Model
	String() string
	GoString() string
	SetContacts(v []*string) *AlertRuleNotificationFilter
	GetContacts() []*string
	SetCustomWebhooks(v []*string) *AlertRuleNotificationFilter
	GetCustomWebhooks() []*string
	SetDingWebhooks(v []*string) *AlertRuleNotificationFilter
	GetDingWebhooks() []*string
	SetFsWebhooks(v []*string) *AlertRuleNotificationFilter
	GetFsWebhooks() []*string
	SetGroups(v []*string) *AlertRuleNotificationFilter
	GetGroups() []*string
	SetSlackWebhooks(v []*string) *AlertRuleNotificationFilter
	GetSlackWebhooks() []*string
	SetWxWebhooks(v []*string) *AlertRuleNotificationFilter
	GetWxWebhooks() []*string
}

type AlertRuleNotificationFilter struct {
	Contacts       []*string `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	CustomWebhooks []*string `json:"customWebhooks,omitempty" xml:"customWebhooks,omitempty" type:"Repeated"`
	DingWebhooks   []*string `json:"dingWebhooks,omitempty" xml:"dingWebhooks,omitempty" type:"Repeated"`
	FsWebhooks     []*string `json:"fsWebhooks,omitempty" xml:"fsWebhooks,omitempty" type:"Repeated"`
	Groups         []*string `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	SlackWebhooks  []*string `json:"slackWebhooks,omitempty" xml:"slackWebhooks,omitempty" type:"Repeated"`
	WxWebhooks     []*string `json:"wxWebhooks,omitempty" xml:"wxWebhooks,omitempty" type:"Repeated"`
}

func (s AlertRuleNotificationFilter) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleNotificationFilter) GoString() string {
	return s.String()
}

func (s *AlertRuleNotificationFilter) GetContacts() []*string {
	return s.Contacts
}

func (s *AlertRuleNotificationFilter) GetCustomWebhooks() []*string {
	return s.CustomWebhooks
}

func (s *AlertRuleNotificationFilter) GetDingWebhooks() []*string {
	return s.DingWebhooks
}

func (s *AlertRuleNotificationFilter) GetFsWebhooks() []*string {
	return s.FsWebhooks
}

func (s *AlertRuleNotificationFilter) GetGroups() []*string {
	return s.Groups
}

func (s *AlertRuleNotificationFilter) GetSlackWebhooks() []*string {
	return s.SlackWebhooks
}

func (s *AlertRuleNotificationFilter) GetWxWebhooks() []*string {
	return s.WxWebhooks
}

func (s *AlertRuleNotificationFilter) SetContacts(v []*string) *AlertRuleNotificationFilter {
	s.Contacts = v
	return s
}

func (s *AlertRuleNotificationFilter) SetCustomWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.CustomWebhooks = v
	return s
}

func (s *AlertRuleNotificationFilter) SetDingWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.DingWebhooks = v
	return s
}

func (s *AlertRuleNotificationFilter) SetFsWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.FsWebhooks = v
	return s
}

func (s *AlertRuleNotificationFilter) SetGroups(v []*string) *AlertRuleNotificationFilter {
	s.Groups = v
	return s
}

func (s *AlertRuleNotificationFilter) SetSlackWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.SlackWebhooks = v
	return s
}

func (s *AlertRuleNotificationFilter) SetWxWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.WxWebhooks = v
	return s
}

func (s *AlertRuleNotificationFilter) Validate() error {
	return dara.Validate(s)
}
