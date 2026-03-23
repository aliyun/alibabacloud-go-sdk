// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachmentUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachName(v string) *GetAttachmentUrlRequest
	GetAttachName() *string
	SetNoteId(v string) *GetAttachmentUrlRequest
	GetNoteId() *string
	SetTicketId(v string) *GetAttachmentUrlRequest
	GetTicketId() *string
}

type GetAttachmentUrlRequest struct {
	// This parameter is required.
	AttachName *string `json:"AttachName,omitempty" xml:"AttachName,omitempty"`
	// This parameter is required.
	NoteId *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// This parameter is required.
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s GetAttachmentUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttachmentUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAttachmentUrlRequest) GetAttachName() *string {
	return s.AttachName
}

func (s *GetAttachmentUrlRequest) GetNoteId() *string {
	return s.NoteId
}

func (s *GetAttachmentUrlRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *GetAttachmentUrlRequest) SetAttachName(v string) *GetAttachmentUrlRequest {
	s.AttachName = &v
	return s
}

func (s *GetAttachmentUrlRequest) SetNoteId(v string) *GetAttachmentUrlRequest {
	s.NoteId = &v
	return s
}

func (s *GetAttachmentUrlRequest) SetTicketId(v string) *GetAttachmentUrlRequest {
	s.TicketId = &v
	return s
}

func (s *GetAttachmentUrlRequest) Validate() error {
	return dara.Validate(s)
}
