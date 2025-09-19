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
}

type CreateSuspEventNoteRequest struct {
	// The ID of the alert event to which you want to add remarks. You can call the [DescribeSuspEvents](https://help.aliyun.com/document_detail/251497.html) operation to query the IDs of alert events.
	//
	// This parameter is required.
	//
	// example:
	//
	// 668931
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

func (s *CreateSuspEventNoteRequest) SetEventId(v int64) *CreateSuspEventNoteRequest {
	s.EventId = &v
	return s
}

func (s *CreateSuspEventNoteRequest) SetNote(v string) *CreateSuspEventNoteRequest {
	s.Note = &v
	return s
}

func (s *CreateSuspEventNoteRequest) Validate() error {
	return dara.Validate(s)
}
