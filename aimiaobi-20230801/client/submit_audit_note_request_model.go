// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAuditNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *SubmitAuditNoteRequest
	GetFileKey() *string
	SetNoteId(v string) *SubmitAuditNoteRequest
	GetNoteId() *string
	SetWorkspaceId(v string) *SubmitAuditNoteRequest
	GetWorkspaceId() *string
}

type SubmitAuditNoteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// oss://default/your/file/key
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	NoteId  *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitAuditNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAuditNoteRequest) GoString() string {
	return s.String()
}

func (s *SubmitAuditNoteRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *SubmitAuditNoteRequest) GetNoteId() *string {
	return s.NoteId
}

func (s *SubmitAuditNoteRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitAuditNoteRequest) SetFileKey(v string) *SubmitAuditNoteRequest {
	s.FileKey = &v
	return s
}

func (s *SubmitAuditNoteRequest) SetNoteId(v string) *SubmitAuditNoteRequest {
	s.NoteId = &v
	return s
}

func (s *SubmitAuditNoteRequest) SetWorkspaceId(v string) *SubmitAuditNoteRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitAuditNoteRequest) Validate() error {
	return dara.Validate(s)
}
