// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorNotifyStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetChannels(v []*string) *MonitorNotifyStrategy
	GetChannels() []*string
	SetContactGroupIds(v []*int64) *MonitorNotifyStrategy
	GetContactGroupIds() []*int64
	SetContactGroups(v []*MonitorContactGroup) *MonitorNotifyStrategy
	GetContactGroups() []*MonitorContactGroup
	SetContactIds(v []*int64) *MonitorNotifyStrategy
	GetContactIds() []*int64
	SetContacts(v []*MonitorContact) *MonitorNotifyStrategy
	GetContacts() []*MonitorContact
	SetEndTime(v string) *MonitorNotifyStrategy
	GetEndTime() *string
	SetId(v int64) *MonitorNotifyStrategy
	GetId() *int64
	SetName(v string) *MonitorNotifyStrategy
	GetName() *string
	SetRequestId(v string) *MonitorNotifyStrategy
	GetRequestId() *string
	SetStartTime(v string) *MonitorNotifyStrategy
	GetStartTime() *string
	SetWebhookIds(v []*int64) *MonitorNotifyStrategy
	GetWebhookIds() []*int64
	SetWebhooks(v []*MonitorWebhook) *MonitorNotifyStrategy
	GetWebhooks() []*MonitorWebhook
}

type MonitorNotifyStrategy struct {
	Channels        []*string              `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	ContactGroupIds []*int64               `json:"contactGroupIds,omitempty" xml:"contactGroupIds,omitempty" type:"Repeated"`
	ContactGroups   []*MonitorContactGroup `json:"contactGroups,omitempty" xml:"contactGroups,omitempty" type:"Repeated"`
	ContactIds      []*int64               `json:"contactIds,omitempty" xml:"contactIds,omitempty" type:"Repeated"`
	Contacts        []*MonitorContact      `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// example:
	//
	// 1439
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 告警策略1
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 00:00
	StartTime  *string           `json:"startTime,omitempty" xml:"startTime,omitempty"`
	WebhookIds []*int64          `json:"webhookIds,omitempty" xml:"webhookIds,omitempty" type:"Repeated"`
	Webhooks   []*MonitorWebhook `json:"webhooks,omitempty" xml:"webhooks,omitempty" type:"Repeated"`
}

func (s MonitorNotifyStrategy) String() string {
	return dara.Prettify(s)
}

func (s MonitorNotifyStrategy) GoString() string {
	return s.String()
}

func (s *MonitorNotifyStrategy) GetChannels() []*string {
	return s.Channels
}

func (s *MonitorNotifyStrategy) GetContactGroupIds() []*int64 {
	return s.ContactGroupIds
}

func (s *MonitorNotifyStrategy) GetContactGroups() []*MonitorContactGroup {
	return s.ContactGroups
}

func (s *MonitorNotifyStrategy) GetContactIds() []*int64 {
	return s.ContactIds
}

func (s *MonitorNotifyStrategy) GetContacts() []*MonitorContact {
	return s.Contacts
}

func (s *MonitorNotifyStrategy) GetEndTime() *string {
	return s.EndTime
}

func (s *MonitorNotifyStrategy) GetId() *int64 {
	return s.Id
}

func (s *MonitorNotifyStrategy) GetName() *string {
	return s.Name
}

func (s *MonitorNotifyStrategy) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorNotifyStrategy) GetStartTime() *string {
	return s.StartTime
}

func (s *MonitorNotifyStrategy) GetWebhookIds() []*int64 {
	return s.WebhookIds
}

func (s *MonitorNotifyStrategy) GetWebhooks() []*MonitorWebhook {
	return s.Webhooks
}

func (s *MonitorNotifyStrategy) SetChannels(v []*string) *MonitorNotifyStrategy {
	s.Channels = v
	return s
}

func (s *MonitorNotifyStrategy) SetContactGroupIds(v []*int64) *MonitorNotifyStrategy {
	s.ContactGroupIds = v
	return s
}

func (s *MonitorNotifyStrategy) SetContactGroups(v []*MonitorContactGroup) *MonitorNotifyStrategy {
	s.ContactGroups = v
	return s
}

func (s *MonitorNotifyStrategy) SetContactIds(v []*int64) *MonitorNotifyStrategy {
	s.ContactIds = v
	return s
}

func (s *MonitorNotifyStrategy) SetContacts(v []*MonitorContact) *MonitorNotifyStrategy {
	s.Contacts = v
	return s
}

func (s *MonitorNotifyStrategy) SetEndTime(v string) *MonitorNotifyStrategy {
	s.EndTime = &v
	return s
}

func (s *MonitorNotifyStrategy) SetId(v int64) *MonitorNotifyStrategy {
	s.Id = &v
	return s
}

func (s *MonitorNotifyStrategy) SetName(v string) *MonitorNotifyStrategy {
	s.Name = &v
	return s
}

func (s *MonitorNotifyStrategy) SetRequestId(v string) *MonitorNotifyStrategy {
	s.RequestId = &v
	return s
}

func (s *MonitorNotifyStrategy) SetStartTime(v string) *MonitorNotifyStrategy {
	s.StartTime = &v
	return s
}

func (s *MonitorNotifyStrategy) SetWebhookIds(v []*int64) *MonitorNotifyStrategy {
	s.WebhookIds = v
	return s
}

func (s *MonitorNotifyStrategy) SetWebhooks(v []*MonitorWebhook) *MonitorNotifyStrategy {
	s.Webhooks = v
	return s
}

func (s *MonitorNotifyStrategy) Validate() error {
	if s.ContactGroups != nil {
		for _, item := range s.ContactGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Contacts != nil {
		for _, item := range s.Contacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Webhooks != nil {
		for _, item := range s.Webhooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
