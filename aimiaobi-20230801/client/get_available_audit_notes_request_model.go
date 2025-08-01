// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailableAuditNotesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetAvailableAuditNotesRequest
	GetWorkspaceId() *string
}

type GetAvailableAuditNotesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAvailableAuditNotesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableAuditNotesRequest) GoString() string {
	return s.String()
}

func (s *GetAvailableAuditNotesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAvailableAuditNotesRequest) SetWorkspaceId(v string) *GetAvailableAuditNotesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetAvailableAuditNotesRequest) Validate() error {
	return dara.Validate(s)
}
