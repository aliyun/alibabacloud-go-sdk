// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentNoteStruct interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *IncidentNoteStruct
	GetContent() *string
	SetFormat(v string) *IncidentNoteStruct
	GetFormat() *string
	SetIncidentId(v string) *IncidentNoteStruct
	GetIncidentId() *string
	SetNoteId(v string) *IncidentNoteStruct
	GetNoteId() *string
	SetOperator(v *IncidentNoteStructOperator) *IncidentNoteStruct
	GetOperator() *IncidentNoteStructOperator
	SetTime(v int64) *IncidentNoteStruct
	GetTime() *int64
	SetType(v string) *IncidentNoteStruct
	GetType() *string
}

type IncidentNoteStruct struct {
	// Content.
	//
	// example:
	//
	// The operations team has been contacted and is currently investigating the issue.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Content format.
	//
	// example:
	//
	// markdown
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// Incident ID.
	//
	// example:
	//
	// incident-001
	IncidentId *string `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	// Note ID.
	//
	// example:
	//
	// note-001
	NoteId *string `json:"noteId,omitempty" xml:"noteId,omitempty"`
	// Operator.
	Operator *IncidentNoteStructOperator `json:"operator,omitempty" xml:"operator,omitempty" type:"Struct"`
	// Time.
	//
	// example:
	//
	// 1741234567890
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// Type.
	//
	// example:
	//
	// apm
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IncidentNoteStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentNoteStruct) GoString() string {
	return s.String()
}

func (s *IncidentNoteStruct) GetContent() *string {
	return s.Content
}

func (s *IncidentNoteStruct) GetFormat() *string {
	return s.Format
}

func (s *IncidentNoteStruct) GetIncidentId() *string {
	return s.IncidentId
}

func (s *IncidentNoteStruct) GetNoteId() *string {
	return s.NoteId
}

func (s *IncidentNoteStruct) GetOperator() *IncidentNoteStructOperator {
	return s.Operator
}

func (s *IncidentNoteStruct) GetTime() *int64 {
	return s.Time
}

func (s *IncidentNoteStruct) GetType() *string {
	return s.Type
}

func (s *IncidentNoteStruct) SetContent(v string) *IncidentNoteStruct {
	s.Content = &v
	return s
}

func (s *IncidentNoteStruct) SetFormat(v string) *IncidentNoteStruct {
	s.Format = &v
	return s
}

func (s *IncidentNoteStruct) SetIncidentId(v string) *IncidentNoteStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentNoteStruct) SetNoteId(v string) *IncidentNoteStruct {
	s.NoteId = &v
	return s
}

func (s *IncidentNoteStruct) SetOperator(v *IncidentNoteStructOperator) *IncidentNoteStruct {
	s.Operator = v
	return s
}

func (s *IncidentNoteStruct) SetTime(v int64) *IncidentNoteStruct {
	s.Time = &v
	return s
}

func (s *IncidentNoteStruct) SetType(v string) *IncidentNoteStruct {
	s.Type = &v
	return s
}

func (s *IncidentNoteStruct) Validate() error {
	if s.Operator != nil {
		if err := s.Operator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IncidentNoteStructOperator struct {
	// Contact.
	//
	// example:
	//
	// {\\"mobile_phone_num\\":\\"153xxxx8040\\",\\"mobile_country_code\\":\\"86\\",\\"email\\":\\"flightxxg@dida.com\\"}
	Contact *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// Contact ID.
	//
	// example:
	//
	// user-12345
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	// Name.
	//
	// example:
	//
	// Zhang San.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// User ID.
	//
	// example:
	//
	// 4123456
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentNoteStructOperator) String() string {
	return dara.Prettify(s)
}

func (s IncidentNoteStructOperator) GoString() string {
	return s.String()
}

func (s *IncidentNoteStructOperator) GetContact() *string {
	return s.Contact
}

func (s *IncidentNoteStructOperator) GetContactId() *string {
	return s.ContactId
}

func (s *IncidentNoteStructOperator) GetName() *string {
	return s.Name
}

func (s *IncidentNoteStructOperator) GetUserId() *int64 {
	return s.UserId
}

func (s *IncidentNoteStructOperator) SetContact(v string) *IncidentNoteStructOperator {
	s.Contact = &v
	return s
}

func (s *IncidentNoteStructOperator) SetContactId(v string) *IncidentNoteStructOperator {
	s.ContactId = &v
	return s
}

func (s *IncidentNoteStructOperator) SetName(v string) *IncidentNoteStructOperator {
	s.Name = &v
	return s
}

func (s *IncidentNoteStructOperator) SetUserId(v int64) *IncidentNoteStructOperator {
	s.UserId = &v
	return s
}

func (s *IncidentNoteStructOperator) Validate() error {
	return dara.Validate(s)
}
