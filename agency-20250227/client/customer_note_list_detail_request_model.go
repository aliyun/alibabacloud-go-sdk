// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteListDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNoteId(v int64) *CustomerNoteListDetailRequest
	GetNoteId() *int64
}

type CustomerNoteListDetailRequest struct {
	// The note ID.
	//
	// example:
	//
	// 1620737
	NoteId *int64 `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
}

func (s CustomerNoteListDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListDetailRequest) GoString() string {
	return s.String()
}

func (s *CustomerNoteListDetailRequest) GetNoteId() *int64 {
	return s.NoteId
}

func (s *CustomerNoteListDetailRequest) SetNoteId(v int64) *CustomerNoteListDetailRequest {
	s.NoteId = &v
	return s
}

func (s *CustomerNoteListDetailRequest) Validate() error {
	return dara.Validate(s)
}
