// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentContactStruct interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v []*string) *IncidentContactStruct
	GetChannel() []*string
	SetContactId(v string) *IncidentContactStruct
	GetContactId() *string
	SetContactType(v string) *IncidentContactStruct
	GetContactType() *string
}

type IncidentContactStruct struct {
	Channel     []*string `json:"channel,omitempty" xml:"channel,omitempty" type:"Repeated"`
	ContactId   *string   `json:"contactId,omitempty" xml:"contactId,omitempty"`
	ContactType *string   `json:"contactType,omitempty" xml:"contactType,omitempty"`
}

func (s IncidentContactStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentContactStruct) GoString() string {
	return s.String()
}

func (s *IncidentContactStruct) GetChannel() []*string {
	return s.Channel
}

func (s *IncidentContactStruct) GetContactId() *string {
	return s.ContactId
}

func (s *IncidentContactStruct) GetContactType() *string {
	return s.ContactType
}

func (s *IncidentContactStruct) SetChannel(v []*string) *IncidentContactStruct {
	s.Channel = v
	return s
}

func (s *IncidentContactStruct) SetContactId(v string) *IncidentContactStruct {
	s.ContactId = &v
	return s
}

func (s *IncidentContactStruct) SetContactType(v string) *IncidentContactStruct {
	s.ContactType = &v
	return s
}

func (s *IncidentContactStruct) Validate() error {
	return dara.Validate(s)
}
