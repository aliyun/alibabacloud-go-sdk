// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmAndPostProcessAuditNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *ConfirmAndPostProcessAuditNoteRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *ConfirmAndPostProcessAuditNoteRequest
	GetWorkspaceId() *string
}

type ConfirmAndPostProcessAuditNoteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx_Default_1241541251241
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ConfirmAndPostProcessAuditNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAndPostProcessAuditNoteRequest) GoString() string {
	return s.String()
}

func (s *ConfirmAndPostProcessAuditNoteRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ConfirmAndPostProcessAuditNoteRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ConfirmAndPostProcessAuditNoteRequest) SetTaskId(v string) *ConfirmAndPostProcessAuditNoteRequest {
	s.TaskId = &v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteRequest) SetWorkspaceId(v string) *ConfirmAndPostProcessAuditNoteRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteRequest) Validate() error {
	return dara.Validate(s)
}
