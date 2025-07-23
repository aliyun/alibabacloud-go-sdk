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
	Content    *string                     `json:"content,omitempty" xml:"content,omitempty"`
	Format     *string                     `json:"format,omitempty" xml:"format,omitempty"`
	IncidentId *string                     `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	NoteId     *string                     `json:"noteId,omitempty" xml:"noteId,omitempty"`
	Operator   *IncidentNoteStructOperator `json:"operator,omitempty" xml:"operator,omitempty" type:"Struct"`
	Time       *int64                      `json:"time,omitempty" xml:"time,omitempty"`
	Type       *string                     `json:"type,omitempty" xml:"type,omitempty"`
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
	return dara.Validate(s)
}

type IncidentNoteStructOperator struct {
	Contact   *string `json:"contact,omitempty" xml:"contact,omitempty"`
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
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
