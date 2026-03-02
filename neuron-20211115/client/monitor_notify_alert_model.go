// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorNotifyAlert interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroups(v []*MonitorContactGroup) *MonitorNotifyAlert
	GetContactGroups() []*MonitorContactGroup
	SetContacts(v []*MonitorContact) *MonitorNotifyAlert
	GetContacts() []*MonitorContact
	SetEndTime(v string) *MonitorNotifyAlert
	GetEndTime() *string
	SetId(v int64) *MonitorNotifyAlert
	GetId() *int64
	SetName(v string) *MonitorNotifyAlert
	GetName() *string
	SetNotifyChannels(v []*string) *MonitorNotifyAlert
	GetNotifyChannels() []*string
	SetRuleNames(v []*string) *MonitorNotifyAlert
	GetRuleNames() []*string
	SetStartTime(v string) *MonitorNotifyAlert
	GetStartTime() *string
	SetType(v string) *MonitorNotifyAlert
	GetType() *string
	SetWebhooks(v []*MonitorWebhook) *MonitorNotifyAlert
	GetWebhooks() []*MonitorWebhook
}

type MonitorNotifyAlert struct {
	ContactGroups []*MonitorContactGroup `json:"contactGroups,omitempty" xml:"contactGroups,omitempty" type:"Repeated"`
	Contacts      []*MonitorContact      `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// example:
	//
	// 1439
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	Id             *int64    `json:"id,omitempty" xml:"id,omitempty"`
	Name           *string   `json:"name,omitempty" xml:"name,omitempty"`
	NotifyChannels []*string `json:"notifyChannels,omitempty" xml:"notifyChannels,omitempty" type:"Repeated"`
	RuleNames      []*string `json:"ruleNames,omitempty" xml:"ruleNames,omitempty" type:"Repeated"`
	// example:
	//
	// 00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 日志
	Type     *string           `json:"type,omitempty" xml:"type,omitempty"`
	Webhooks []*MonitorWebhook `json:"webhooks,omitempty" xml:"webhooks,omitempty" type:"Repeated"`
}

func (s MonitorNotifyAlert) String() string {
	return dara.Prettify(s)
}

func (s MonitorNotifyAlert) GoString() string {
	return s.String()
}

func (s *MonitorNotifyAlert) GetContactGroups() []*MonitorContactGroup {
	return s.ContactGroups
}

func (s *MonitorNotifyAlert) GetContacts() []*MonitorContact {
	return s.Contacts
}

func (s *MonitorNotifyAlert) GetEndTime() *string {
	return s.EndTime
}

func (s *MonitorNotifyAlert) GetId() *int64 {
	return s.Id
}

func (s *MonitorNotifyAlert) GetName() *string {
	return s.Name
}

func (s *MonitorNotifyAlert) GetNotifyChannels() []*string {
	return s.NotifyChannels
}

func (s *MonitorNotifyAlert) GetRuleNames() []*string {
	return s.RuleNames
}

func (s *MonitorNotifyAlert) GetStartTime() *string {
	return s.StartTime
}

func (s *MonitorNotifyAlert) GetType() *string {
	return s.Type
}

func (s *MonitorNotifyAlert) GetWebhooks() []*MonitorWebhook {
	return s.Webhooks
}

func (s *MonitorNotifyAlert) SetContactGroups(v []*MonitorContactGroup) *MonitorNotifyAlert {
	s.ContactGroups = v
	return s
}

func (s *MonitorNotifyAlert) SetContacts(v []*MonitorContact) *MonitorNotifyAlert {
	s.Contacts = v
	return s
}

func (s *MonitorNotifyAlert) SetEndTime(v string) *MonitorNotifyAlert {
	s.EndTime = &v
	return s
}

func (s *MonitorNotifyAlert) SetId(v int64) *MonitorNotifyAlert {
	s.Id = &v
	return s
}

func (s *MonitorNotifyAlert) SetName(v string) *MonitorNotifyAlert {
	s.Name = &v
	return s
}

func (s *MonitorNotifyAlert) SetNotifyChannels(v []*string) *MonitorNotifyAlert {
	s.NotifyChannels = v
	return s
}

func (s *MonitorNotifyAlert) SetRuleNames(v []*string) *MonitorNotifyAlert {
	s.RuleNames = v
	return s
}

func (s *MonitorNotifyAlert) SetStartTime(v string) *MonitorNotifyAlert {
	s.StartTime = &v
	return s
}

func (s *MonitorNotifyAlert) SetType(v string) *MonitorNotifyAlert {
	s.Type = &v
	return s
}

func (s *MonitorNotifyAlert) SetWebhooks(v []*MonitorWebhook) *MonitorNotifyAlert {
	s.Webhooks = v
	return s
}

func (s *MonitorNotifyAlert) Validate() error {
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
