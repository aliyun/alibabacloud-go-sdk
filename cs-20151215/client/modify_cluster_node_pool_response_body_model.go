// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodepoolId(v string) *ModifyClusterNodePoolResponseBody
	GetNodepoolId() *string
	SetRequestId(v string) *ModifyClusterNodePoolResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyClusterNodePoolResponseBody
	GetTaskId() *string
}

type ModifyClusterNodePoolResponseBody struct {
	// The node pool ID.
	//
	// example:
	//
	// np737c3ac1ac684703b9e10673aa2c****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 687C5BAA-D103-4993-884B-C35E4314****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// T-5fd211e924e1d00787000293
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ModifyClusterNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolResponseBody) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *ModifyClusterNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterNodePoolResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyClusterNodePoolResponseBody) SetNodepoolId(v string) *ModifyClusterNodePoolResponseBody {
	s.NodepoolId = &v
	return s
}

func (s *ModifyClusterNodePoolResponseBody) SetRequestId(v string) *ModifyClusterNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterNodePoolResponseBody) SetTaskId(v string) *ModifyClusterNodePoolResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyClusterNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
