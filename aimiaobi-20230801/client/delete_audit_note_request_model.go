// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuditNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DeleteAuditNoteRequest
	GetWorkspaceId() *string
}

type DeleteAuditNoteRequest struct {
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

func (s *DeleteAuditNoteRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteAuditNoteRequest) SetWorkspaceId(v string) *DeleteAuditNoteRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteAuditNoteRequest) Validate() error {
	return dara.Validate(s)
}
