// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSuspEventNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v int64) *CreateSuspEventNoteRequest
	GetEventId() *int64
	SetNote(v string) *CreateSuspEventNoteRequest
	GetNote() *string
	SetResourceDirectoryAccountId(v int64) *CreateSuspEventNoteRequest
	GetResourceDirectoryAccountId() *int64
}

type CreateSuspEventNoteRequest struct {
	// The ID of the security alert event to which you want to add a note. Call [DescribeSuspEvents](https://help.aliyun.com/document_detail/251497.html) to obtain the ID of the alert event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 668931
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The note to add.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ignore
	Note                       *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ResourceDirectoryAccountId *int64  `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s CreateSuspEventNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSuspEventNoteRequest) GoString() string {
	return s.String()
}

func (s *CreateSuspEventNoteRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *CreateSuspEventNoteRequest) GetNote() *string {
	return s.Note
}

func (s *CreateSuspEventNoteRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *CreateSuspEventNoteRequest) SetEventId(v int64) *CreateSuspEventNoteRequest {
	s.EventId = &v
	return s
}

func (s *CreateSuspEventNoteRequest) SetNote(v string) *CreateSuspEventNoteRequest {
	s.Note = &v
	return s
}

func (s *CreateSuspEventNoteRequest) SetResourceDirectoryAccountId(v int64) *CreateSuspEventNoteRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *CreateSuspEventNoteRequest) Validate() error {
	return dara.Validate(s)
}
