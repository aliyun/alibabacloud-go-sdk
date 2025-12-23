// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateWorkflowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecType(v string) *OperateWorkflowInstanceRequest
	GetExecType() *string
	SetWorkflowInstanceId(v string) *OperateWorkflowInstanceRequest
	GetWorkflowInstanceId() *string
	SetWorkspaceId(v string) *OperateWorkflowInstanceRequest
	GetWorkspaceId() *string
}

type OperateWorkflowInstanceRequest struct {
	// example:
	//
	// PAUSE
	ExecType *string `json:"execType,omitempty" xml:"execType,omitempty"`
	// example:
	//
	// wi-l9o479p8rrx****
	WorkflowInstanceId *string `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	// example:
	//
	// w-lxyy60mpgpg****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s OperateWorkflowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *OperateWorkflowInstanceRequest) GetExecType() *string {
	return s.ExecType
}

func (s *OperateWorkflowInstanceRequest) GetWorkflowInstanceId() *string {
	return s.WorkflowInstanceId
}

func (s *OperateWorkflowInstanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *OperateWorkflowInstanceRequest) SetExecType(v string) *OperateWorkflowInstanceRequest {
	s.ExecType = &v
	return s
}

func (s *OperateWorkflowInstanceRequest) SetWorkflowInstanceId(v string) *OperateWorkflowInstanceRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *OperateWorkflowInstanceRequest) SetWorkspaceId(v string) *OperateWorkflowInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *OperateWorkflowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
