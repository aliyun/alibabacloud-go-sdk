// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodePoolComponentInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteNodePoolComponentInstanceResponseBody
	GetClusterId() *string
	SetRequestId(v string) *DeleteNodePoolComponentInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DeleteNodePoolComponentInstanceResponseBody
	GetTaskId() *string
}

type DeleteNodePoolComponentInstanceResponseBody struct {
	// example:
	//
	// c846d7d529e34413c9ab1****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// EB022AB1-4CF7-5BB6-B44A-38****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// T-696de321273bb00****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s DeleteNodePoolComponentInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodePoolComponentInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodePoolComponentInstanceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteNodePoolComponentInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNodePoolComponentInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteNodePoolComponentInstanceResponseBody) SetClusterId(v string) *DeleteNodePoolComponentInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceResponseBody) SetRequestId(v string) *DeleteNodePoolComponentInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceResponseBody) SetTaskId(v string) *DeleteNodePoolComponentInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
