// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuditNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNoteId(v string) *DeleteAuditNoteRequest
	GetNoteId() *string
	SetWorkspaceId(v string) *DeleteAuditNoteRequest
	GetWorkspaceId() *string
}

type DeleteAuditNoteRequest struct {
	NoteId *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteAuditNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuditNoteRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuditNoteRequest) GetNoteId() *string {
	return s.NoteId
}

func (s *DeleteAuditNoteRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteAuditNoteRequest) SetNoteId(v string) *DeleteAuditNoteRequest {
	s.NoteId = &v
	return s
}

func (s *DeleteAuditNoteRequest) SetWorkspaceId(v string) *DeleteAuditNoteRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteAuditNoteRequest) Validate() error {
	return dara.Validate(s)
}
