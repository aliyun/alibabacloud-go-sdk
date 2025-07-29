// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveNodePoolNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveNodePoolNodesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RemoveNodePoolNodesResponseBody
	GetTaskId() *string
}

type RemoveNodePoolNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A9891419-D125-4D89-AFCA-68846675E2F7
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// T-62a944794ee141074400****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s RemoveNodePoolNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveNodePoolNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveNodePoolNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveNodePoolNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RemoveNodePoolNodesResponseBody) SetRequestId(v string) *RemoveNodePoolNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveNodePoolNodesResponseBody) SetTaskId(v string) *RemoveNodePoolNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *RemoveNodePoolNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
