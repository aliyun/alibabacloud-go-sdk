// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorNotifyObjectResult interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroups(v []*MonitorContactGroup) *MonitorNotifyObjectResult
	GetContactGroups() []*MonitorContactGroup
	SetContacts(v []*MonitorContact) *MonitorNotifyObjectResult
	GetContacts() []*MonitorContact
	SetRequestId(v string) *MonitorNotifyObjectResult
	GetRequestId() *string
	SetWebhooks(v []*MonitorWebhook) *MonitorNotifyObjectResult
	GetWebhooks() []*MonitorWebhook
}

type MonitorNotifyObjectResult struct {
	ContactGroups []*MonitorContactGroup `json:"contactGroups,omitempty" xml:"contactGroups,omitempty" type:"Repeated"`
	Contacts      []*MonitorContact      `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	RequestId     *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Webhooks      []*MonitorWebhook      `json:"webhooks,omitempty" xml:"webhooks,omitempty" type:"Repeated"`
}

func (s MonitorNotifyObjectResult) String() string {
	return dara.Prettify(s)
}

func (s MonitorNotifyObjectResult) GoString() string {
	return s.String()
}

func (s *MonitorNotifyObjectResult) GetContactGroups() []*MonitorContactGroup {
	return s.ContactGroups
}

func (s *MonitorNotifyObjectResult) GetContacts() []*MonitorContact {
	return s.Contacts
}

func (s *MonitorNotifyObjectResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorNotifyObjectResult) GetWebhooks() []*MonitorWebhook {
	return s.Webhooks
}

func (s *MonitorNotifyObjectResult) SetContactGroups(v []*MonitorContactGroup) *MonitorNotifyObjectResult {
	s.ContactGroups = v
	return s
}

func (s *MonitorNotifyObjectResult) SetContacts(v []*MonitorContact) *MonitorNotifyObjectResult {
	s.Contacts = v
	return s
}

func (s *MonitorNotifyObjectResult) SetRequestId(v string) *MonitorNotifyObjectResult {
	s.RequestId = &v
	return s
}

func (s *MonitorNotifyObjectResult) SetWebhooks(v []*MonitorWebhook) *MonitorNotifyObjectResult {
	s.Webhooks = v
	return s
}

func (s *MonitorNotifyObjectResult) Validate() error {
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
