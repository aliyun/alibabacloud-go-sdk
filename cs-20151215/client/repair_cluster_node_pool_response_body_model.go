// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepairClusterNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RepairClusterNodePoolResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RepairClusterNodePoolResponseBody
	GetTaskId() *string
}

type RepairClusterNodePoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// db82195b-75a8-40e5-9be4-16f1829dc624
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// T-613b19bbd160ad4928000001
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s RepairClusterNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RepairClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *RepairClusterNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RepairClusterNodePoolResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RepairClusterNodePoolResponseBody) SetRequestId(v string) *RepairClusterNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *RepairClusterNodePoolResponseBody) SetTaskId(v string) *RepairClusterNodePoolResponseBody {
	s.TaskId = &v
	return s
}

func (s *RepairClusterNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
