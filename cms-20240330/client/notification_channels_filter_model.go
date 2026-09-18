// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotificationChannelsFilter interface {
	dara.Model
	String() string
	GoString() string
	SetContainsContacts(v []*string) *NotificationChannelsFilter
	GetContainsContacts() []*string
	SetContainsCustomWebhooks(v []*string) *NotificationChannelsFilter
	GetContainsCustomWebhooks() []*string
	SetContainsDingWebhooks(v []*string) *NotificationChannelsFilter
	GetContainsDingWebhooks() []*string
	SetContainsFsWebhooks(v []*string) *NotificationChannelsFilter
	GetContainsFsWebhooks() []*string
	SetContainsGroups(v []*string) *NotificationChannelsFilter
	GetContainsGroups() []*string
	SetContainsSlackWebhooks(v []*string) *NotificationChannelsFilter
	GetContainsSlackWebhooks() []*string
	SetContainsWxWebhooks(v []*string) *NotificationChannelsFilter
	GetContainsWxWebhooks() []*string
}

type NotificationChannelsFilter struct {
	// The alert contact list of the rule contains any value in the array (OR semantics), corresponding to V1 notification.contacts.
	//
	// example:
	//
	// ["John","Jane"]
	ContainsContacts []*string `json:"containsContacts,omitempty" xml:"containsContacts,omitempty" type:"Repeated"`
	// The custom webhook list of the rule contains any value in the array (OR semantics), corresponding to V1 notification.customWebhooks.
	//
	// example:
	//
	// ["https://my-service.example.com/webhook/alert"]
	ContainsCustomWebhooks []*string `json:"containsCustomWebhooks,omitempty" xml:"containsCustomWebhooks,omitempty" type:"Repeated"`
	// The DingTalk webhook list of the rule contains any value in the array (OR semantics), corresponding to V1 notification.dingWebhooks.
	//
	// example:
	//
	// ["https://oapi.dingtalk.com/robot/send?access_token=abc123"]
	ContainsDingWebhooks []*string `json:"containsDingWebhooks,omitempty" xml:"containsDingWebhooks,omitempty" type:"Repeated"`
	// The Lark webhook list of the rule contains any value in the array (OR semantics), corresponding to V1 notification.fsWebhooks.
	//
	// example:
	//
	// ["https://open.feishu.cn/open-apis/bot/v2/hook/abc123"]
	ContainsFsWebhooks []*string `json:"containsFsWebhooks,omitempty" xml:"containsFsWebhooks,omitempty" type:"Repeated"`
	// The alert contact group list of the rule contains any value in the array (OR semantics), corresponding to V1 notification.groups.
	//
	// example:
	//
	// ["OpsTeam","SRETeam"]
	ContainsGroups []*string `json:"containsGroups,omitempty" xml:"containsGroups,omitempty" type:"Repeated"`
	// The Slack webhook list of the rule contains any value in the array (OR semantics), corresponding to V1 notification.slackWebhooks.
	//
	// example:
	//
	// ["https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXX"]
	ContainsSlackWebhooks []*string `json:"containsSlackWebhooks,omitempty" xml:"containsSlackWebhooks,omitempty" type:"Repeated"`
	// The WeCom webhook list of the rule contains any value in the array (OR semantics), corresponding to V1 notification.wxWebhooks.
	//
	// example:
	//
	// ["https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=abc123"]
	ContainsWxWebhooks []*string `json:"containsWxWebhooks,omitempty" xml:"containsWxWebhooks,omitempty" type:"Repeated"`
}

func (s NotificationChannelsFilter) String() string {
	return dara.Prettify(s)
}

func (s NotificationChannelsFilter) GoString() string {
	return s.String()
}

func (s *NotificationChannelsFilter) GetContainsContacts() []*string {
	return s.ContainsContacts
}

func (s *NotificationChannelsFilter) GetContainsCustomWebhooks() []*string {
	return s.ContainsCustomWebhooks
}

func (s *NotificationChannelsFilter) GetContainsDingWebhooks() []*string {
	return s.ContainsDingWebhooks
}

func (s *NotificationChannelsFilter) GetContainsFsWebhooks() []*string {
	return s.ContainsFsWebhooks
}

func (s *NotificationChannelsFilter) GetContainsGroups() []*string {
	return s.ContainsGroups
}

func (s *NotificationChannelsFilter) GetContainsSlackWebhooks() []*string {
	return s.ContainsSlackWebhooks
}

func (s *NotificationChannelsFilter) GetContainsWxWebhooks() []*string {
	return s.ContainsWxWebhooks
}

func (s *NotificationChannelsFilter) SetContainsContacts(v []*string) *NotificationChannelsFilter {
	s.ContainsContacts = v
	return s
}

func (s *NotificationChannelsFilter) SetContainsCustomWebhooks(v []*string) *NotificationChannelsFilter {
	s.ContainsCustomWebhooks = v
	return s
}

func (s *NotificationChannelsFilter) SetContainsDingWebhooks(v []*string) *NotificationChannelsFilter {
	s.ContainsDingWebhooks = v
	return s
}

func (s *NotificationChannelsFilter) SetContainsFsWebhooks(v []*string) *NotificationChannelsFilter {
	s.ContainsFsWebhooks = v
	return s
}

func (s *NotificationChannelsFilter) SetContainsGroups(v []*string) *NotificationChannelsFilter {
	s.ContainsGroups = v
	return s
}

func (s *NotificationChannelsFilter) SetContainsSlackWebhooks(v []*string) *NotificationChannelsFilter {
	s.ContainsSlackWebhooks = v
	return s
}

func (s *NotificationChannelsFilter) SetContainsWxWebhooks(v []*string) *NotificationChannelsFilter {
	s.ContainsWxWebhooks = v
	return s
}

func (s *NotificationChannelsFilter) Validate() error {
	return dara.Validate(s)
}
