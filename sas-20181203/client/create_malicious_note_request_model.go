// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaliciousNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v int64) *CreateMaliciousNoteRequest
	GetEventId() *int64
	SetNote(v string) *CreateMaliciousNoteRequest
	GetNote() *string
}

type CreateMaliciousNoteRequest struct {
	// The ID of the alert event to which you want to add remarks.
	//
	// >  You can call the [ListAgentlessMaliciousFiles](~~ListAgentlessMaliciousFiles~~) operation to obtain the ID of the alert event from the NoteId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80****
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The remarks that you want to add.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ignore
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s CreateMaliciousNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMaliciousNoteRequest) GoString() string {
	return s.String()
}

func (s *CreateMaliciousNoteRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *CreateMaliciousNoteRequest) GetNote() *string {
	return s.Note
}

func (s *CreateMaliciousNoteRequest) SetEventId(v int64) *CreateMaliciousNoteRequest {
	s.EventId = &v
	return s
}

func (s *CreateMaliciousNoteRequest) SetNote(v string) *CreateMaliciousNoteRequest {
	s.Note = &v
	return s
}

func (s *CreateMaliciousNoteRequest) Validate() error {
	return dara.Validate(s)
}
