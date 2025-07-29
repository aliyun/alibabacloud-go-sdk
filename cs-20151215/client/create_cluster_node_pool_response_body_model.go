// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodepoolId(v string) *CreateClusterNodePoolResponseBody
	GetNodepoolId() *string
	SetRequestId(v string) *CreateClusterNodePoolResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateClusterNodePoolResponseBody
	GetTaskId() *string
}

type CreateClusterNodePoolResponseBody struct {
	// The node pool ID.
	//
	// example:
	//
	// np31da1b38983f4511b490fc62108a****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0527ac9a-c899-4341-a21a-****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// T-613b19bbd160ad492800****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateClusterNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponseBody) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *CreateClusterNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterNodePoolResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateClusterNodePoolResponseBody) SetNodepoolId(v string) *CreateClusterNodePoolResponseBody {
	s.NodepoolId = &v
	return s
}

func (s *CreateClusterNodePoolResponseBody) SetRequestId(v string) *CreateClusterNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterNodePoolResponseBody) SetTaskId(v string) *CreateClusterNodePoolResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateClusterNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
