// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorContactGroup interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *MonitorContactGroup
	GetAccountId() *string
	SetContacts(v []*MonitorContact) *MonitorContactGroup
	GetContacts() []*MonitorContact
	SetEnterpriseId(v int64) *MonitorContactGroup
	GetEnterpriseId() *int64
	SetGmtCreate(v string) *MonitorContactGroup
	GetGmtCreate() *string
	SetGmtModified(v string) *MonitorContactGroup
	GetGmtModified() *string
	SetId(v int64) *MonitorContactGroup
	GetId() *int64
	SetName(v string) *MonitorContactGroup
	GetName() *string
}

type MonitorContactGroup struct {
	// example:
	//
	// 121321412341
	AccountId *string           `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Contacts  []*MonitorContact `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// 2022-05-01T00:00:00.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2022-05-01T00:00:00.000Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yani
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s MonitorContactGroup) String() string {
	return dara.Prettify(s)
}

func (s MonitorContactGroup) GoString() string {
	return s.String()
}

func (s *MonitorContactGroup) GetAccountId() *string {
	return s.AccountId
}

func (s *MonitorContactGroup) GetContacts() []*MonitorContact {
	return s.Contacts
}

func (s *MonitorContactGroup) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *MonitorContactGroup) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MonitorContactGroup) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MonitorContactGroup) GetId() *int64 {
	return s.Id
}

func (s *MonitorContactGroup) GetName() *string {
	return s.Name
}

func (s *MonitorContactGroup) SetAccountId(v string) *MonitorContactGroup {
	s.AccountId = &v
	return s
}

func (s *MonitorContactGroup) SetContacts(v []*MonitorContact) *MonitorContactGroup {
	s.Contacts = v
	return s
}

func (s *MonitorContactGroup) SetEnterpriseId(v int64) *MonitorContactGroup {
	s.EnterpriseId = &v
	return s
}

func (s *MonitorContactGroup) SetGmtCreate(v string) *MonitorContactGroup {
	s.GmtCreate = &v
	return s
}

func (s *MonitorContactGroup) SetGmtModified(v string) *MonitorContactGroup {
	s.GmtModified = &v
	return s
}

func (s *MonitorContactGroup) SetId(v int64) *MonitorContactGroup {
	s.Id = &v
	return s
}

func (s *MonitorContactGroup) SetName(v string) *MonitorContactGroup {
	s.Name = &v
	return s
}

func (s *MonitorContactGroup) Validate() error {
	if s.Contacts != nil {
		for _, item := range s.Contacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
