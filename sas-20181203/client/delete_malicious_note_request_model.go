// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaliciousNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNoteId(v int64) *DeleteMaliciousNoteRequest
	GetNoteId() *int64
}

type DeleteMaliciousNoteRequest struct {
	// The ID of the remarks.
	//
	// >  You can call the [ListAgentlessMaliciousFiles](~~ListAgentlessMaliciousFiles~~) operation to obtain the ID from the NoteId parameter.
	//
	// example:
	//
	// 1
	NoteId *int64 `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
}

func (s DeleteMaliciousNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaliciousNoteRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaliciousNoteRequest) GetNoteId() *int64 {
	return s.NoteId
}

func (s *DeleteMaliciousNoteRequest) SetNoteId(v int64) *DeleteMaliciousNoteRequest {
	s.NoteId = &v
	return s
}

func (s *DeleteMaliciousNoteRequest) Validate() error {
	return dara.Validate(s)
}
