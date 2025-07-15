// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAuditResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetSmartAuditResultRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetSmartAuditResultRequest
	GetWorkspaceId() *string
}

type GetSmartAuditResultRequest struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetSmartAuditResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAuditResultRequest) GoString() string {
	return s.String()
}

func (s *GetSmartAuditResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSmartAuditResultRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSmartAuditResultRequest) SetTaskId(v string) *GetSmartAuditResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetSmartAuditResultRequest) SetWorkspaceId(v string) *GetSmartAuditResultRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetSmartAuditResultRequest) Validate() error {
	return dara.Validate(s)
}
