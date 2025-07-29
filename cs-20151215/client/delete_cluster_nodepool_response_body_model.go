// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodepoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClusterNodepoolResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DeleteClusterNodepoolResponseBody
	GetTaskId() *string
}

type DeleteClusterNodepoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7263C978-3DBD-4E06-B319-793B38A2F388
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// task IDs
	//
	// example:
	//
	// T-655ace947e0e6603af000004
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s DeleteClusterNodepoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodepoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodepoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterNodepoolResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteClusterNodepoolResponseBody) SetRequestId(v string) *DeleteClusterNodepoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterNodepoolResponseBody) SetTaskId(v string) *DeleteClusterNodepoolResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteClusterNodepoolResponseBody) Validate() error {
	return dara.Validate(s)
}
