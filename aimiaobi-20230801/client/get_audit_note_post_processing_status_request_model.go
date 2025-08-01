// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditNotePostProcessingStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetAuditNotePostProcessingStatusRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetAuditNotePostProcessingStatusRequest
	GetWorkspaceId() *string
}

type GetAuditNotePostProcessingStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx_Default_12847192741412
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAuditNotePostProcessingStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuditNotePostProcessingStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAuditNotePostProcessingStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAuditNotePostProcessingStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAuditNotePostProcessingStatusRequest) SetTaskId(v string) *GetAuditNotePostProcessingStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusRequest) SetWorkspaceId(v string) *GetAuditNotePostProcessingStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusRequest) Validate() error {
	return dara.Validate(s)
}
