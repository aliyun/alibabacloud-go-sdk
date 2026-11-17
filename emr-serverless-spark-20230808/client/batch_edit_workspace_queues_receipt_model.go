// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEditWorkspaceQueuesReceipt interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptedQueues(v []*string) *BatchEditWorkspaceQueuesReceipt
	GetAcceptedQueues() []*string
	SetSkippedQueues(v []*string) *BatchEditWorkspaceQueuesReceipt
	GetSkippedQueues() []*string
	SetWorkflowInstanceId(v int64) *BatchEditWorkspaceQueuesReceipt
	GetWorkflowInstanceId() *int64
	SetWorkspaceId(v string) *BatchEditWorkspaceQueuesReceipt
	GetWorkspaceId() *string
}

type BatchEditWorkspaceQueuesReceipt struct {
	// The target states of these queues are progressively written during the workflow.
	AcceptedQueues []*string `json:"acceptedQueues,omitempty" xml:"acceptedQueues,omitempty" type:"Repeated"`
	// The current values are already equal to the target values, so no write operations are required.
	SkippedQueues []*string `json:"skippedQueues,omitempty" xml:"skippedQueues,omitempty" type:"Repeated"`
	// A null value indicates that all queues in this request are already in the target state and no workflow is started (idempotent no-op).
	//
	// example:
	//
	// 9527
	WorkflowInstanceId *int64 `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s BatchEditWorkspaceQueuesReceipt) String() string {
	return dara.Prettify(s)
}

func (s BatchEditWorkspaceQueuesReceipt) GoString() string {
	return s.String()
}

func (s *BatchEditWorkspaceQueuesReceipt) GetAcceptedQueues() []*string {
	return s.AcceptedQueues
}

func (s *BatchEditWorkspaceQueuesReceipt) GetSkippedQueues() []*string {
	return s.SkippedQueues
}

func (s *BatchEditWorkspaceQueuesReceipt) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *BatchEditWorkspaceQueuesReceipt) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchEditWorkspaceQueuesReceipt) SetAcceptedQueues(v []*string) *BatchEditWorkspaceQueuesReceipt {
	s.AcceptedQueues = v
	return s
}

func (s *BatchEditWorkspaceQueuesReceipt) SetSkippedQueues(v []*string) *BatchEditWorkspaceQueuesReceipt {
	s.SkippedQueues = v
	return s
}

func (s *BatchEditWorkspaceQueuesReceipt) SetWorkflowInstanceId(v int64) *BatchEditWorkspaceQueuesReceipt {
	s.WorkflowInstanceId = &v
	return s
}

func (s *BatchEditWorkspaceQueuesReceipt) SetWorkspaceId(v string) *BatchEditWorkspaceQueuesReceipt {
	s.WorkspaceId = &v
	return s
}

func (s *BatchEditWorkspaceQueuesReceipt) Validate() error {
	return dara.Validate(s)
}
