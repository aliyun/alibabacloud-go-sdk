// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteClusterNodesResponseBody
	GetClusterId() *string
	SetRequestId(v string) *DeleteClusterNodesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DeleteClusterNodesResponseBody
	GetTaskId() *string
}

type DeleteClusterNodesResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// c104d5d5f301c4e2a8ee578c37bc****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A9891419-D125-4D89-AFCA-68846675E2F7
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// T-60fea8ad2e277f0879000ae9
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s DeleteClusterNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteClusterNodesResponseBody) SetClusterId(v string) *DeleteClusterNodesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterNodesResponseBody) SetRequestId(v string) *DeleteClusterNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterNodesResponseBody) SetTaskId(v string) *DeleteClusterNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteClusterNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
