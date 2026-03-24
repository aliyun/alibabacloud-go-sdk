// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContactForIncidentView interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *ContactForIncidentView
	GetContactId() *string
	SetName(v string) *ContactForIncidentView
	GetName() *string
}

type ContactForIncidentView struct {
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ContactForIncidentView) String() string {
	return dara.Prettify(s)
}

func (s ContactForIncidentView) GoString() string {
	return s.String()
}

func (s *ContactForIncidentView) GetContactId() *string {
	return s.ContactId
}

func (s *ContactForIncidentView) GetName() *string {
	return s.Name
}

func (s *ContactForIncidentView) SetContactId(v string) *ContactForIncidentView {
	s.ContactId = &v
	return s
}

func (s *ContactForIncidentView) SetName(v string) *ContactForIncidentView {
	s.Name = &v
	return s
}

func (s *ContactForIncidentView) Validate() error {
	return dara.Validate(s)
}
