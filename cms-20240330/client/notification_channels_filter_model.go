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
	ContainsContacts       []*string `json:"containsContacts,omitempty" xml:"containsContacts,omitempty" type:"Repeated"`
	ContainsCustomWebhooks []*string `json:"containsCustomWebhooks,omitempty" xml:"containsCustomWebhooks,omitempty" type:"Repeated"`
	ContainsDingWebhooks   []*string `json:"containsDingWebhooks,omitempty" xml:"containsDingWebhooks,omitempty" type:"Repeated"`
	ContainsFsWebhooks     []*string `json:"containsFsWebhooks,omitempty" xml:"containsFsWebhooks,omitempty" type:"Repeated"`
	ContainsGroups         []*string `json:"containsGroups,omitempty" xml:"containsGroups,omitempty" type:"Repeated"`
	ContainsSlackWebhooks  []*string `json:"containsSlackWebhooks,omitempty" xml:"containsSlackWebhooks,omitempty" type:"Repeated"`
	ContainsWxWebhooks     []*string `json:"containsWxWebhooks,omitempty" xml:"containsWxWebhooks,omitempty" type:"Repeated"`
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
