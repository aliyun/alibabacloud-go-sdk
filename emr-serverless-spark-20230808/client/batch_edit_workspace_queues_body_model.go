// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEditWorkspaceQueuesBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueues(v []*BatchEditQueueSpec) *BatchEditWorkspaceQueuesBody
	GetQueues() []*BatchEditQueueSpec
	SetWorkspaceId(v string) *BatchEditWorkspaceQueuesBody
	GetWorkspaceId() *string
}

type BatchEditWorkspaceQueuesBody struct {
	// The list of queues to update. Each entry independently specifies the absolute target state of a queue. Queue names must be unique within the same request.
	//
	// This parameter is required.
	Queues []*BatchEditQueueSpec `json:"queues,omitempty" xml:"queues,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s BatchEditWorkspaceQueuesBody) String() string {
	return dara.Prettify(s)
}

func (s BatchEditWorkspaceQueuesBody) GoString() string {
	return s.String()
}

func (s *BatchEditWorkspaceQueuesBody) GetQueues() []*BatchEditQueueSpec {
	return s.Queues
}

func (s *BatchEditWorkspaceQueuesBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchEditWorkspaceQueuesBody) SetQueues(v []*BatchEditQueueSpec) *BatchEditWorkspaceQueuesBody {
	s.Queues = v
	return s
}

func (s *BatchEditWorkspaceQueuesBody) SetWorkspaceId(v string) *BatchEditWorkspaceQueuesBody {
	s.WorkspaceId = &v
	return s
}

func (s *BatchEditWorkspaceQueuesBody) Validate() error {
	if s.Queues != nil {
		for _, item := range s.Queues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
