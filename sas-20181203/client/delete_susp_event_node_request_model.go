// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSuspEventNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNoteId(v int64) *DeleteSuspEventNodeRequest
	GetNoteId() *int64
	SetResourceDirectoryAccountId(v int64) *DeleteSuspEventNodeRequest
	GetResourceDirectoryAccountId() *int64
}

type DeleteSuspEventNodeRequest struct {
	// The ID of the description.
	//
	// > You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to obtain the ID of the description by using the EventNotes field.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	NoteId                     *int64 `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s DeleteSuspEventNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSuspEventNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteSuspEventNodeRequest) GetNoteId() *int64 {
	return s.NoteId
}

func (s *DeleteSuspEventNodeRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DeleteSuspEventNodeRequest) SetNoteId(v int64) *DeleteSuspEventNodeRequest {
	s.NoteId = &v
	return s
}

func (s *DeleteSuspEventNodeRequest) SetResourceDirectoryAccountId(v int64) *DeleteSuspEventNodeRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DeleteSuspEventNodeRequest) Validate() error {
	return dara.Validate(s)
}
