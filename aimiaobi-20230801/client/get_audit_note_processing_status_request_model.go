// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditNoteProcessingStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetAuditNoteProcessingStatusRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetAuditNoteProcessingStatusRequest
	GetWorkspaceId() *string
}

type GetAuditNoteProcessingStatusRequest struct {
	// Task ID returned by the SubmitAuditNote operation. Save this ID securely. It uniquely identifies your custom rule library task.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx_Default_129873419274
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Unique identifier of your Model Studio workspace. To get this ID, see [Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAuditNoteProcessingStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuditNoteProcessingStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAuditNoteProcessingStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAuditNoteProcessingStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAuditNoteProcessingStatusRequest) SetTaskId(v string) *GetAuditNoteProcessingStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetAuditNoteProcessingStatusRequest) SetWorkspaceId(v string) *GetAuditNoteProcessingStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetAuditNoteProcessingStatusRequest) Validate() error {
	return dara.Validate(s)
}
