// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *DeleteClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DeleteClusterResponseBody
	GetTaskId() *string
}

type DeleteClusterResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// cb95aa626a47740afbf6aa099b650****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 687C5BAA-D103-4993-884B-C35E4314****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// T-5a54309c80282e39ea****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteClusterResponseBody) SetClusterId(v string) *DeleteClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetTaskId(v string) *DeleteClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
