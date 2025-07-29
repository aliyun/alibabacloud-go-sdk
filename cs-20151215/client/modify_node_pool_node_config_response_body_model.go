// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolNodeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodepoolId(v string) *ModifyNodePoolNodeConfigResponseBody
	GetNodepoolId() *string
	SetRequestId(v string) *ModifyNodePoolNodeConfigResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyNodePoolNodeConfigResponseBody
	GetTaskId() *string
}

type ModifyNodePoolNodeConfigResponseBody struct {
	// The node pool ID.
	//
	// example:
	//
	// np737c3ac1ac684703b9e10673aa2c****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D7631D83-6E98-1949-B665-766A62xxxxxx
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// T-5fd211e924e1d00787xxxxxx
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ModifyNodePoolNodeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolNodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolNodeConfigResponseBody) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *ModifyNodePoolNodeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNodePoolNodeConfigResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyNodePoolNodeConfigResponseBody) SetNodepoolId(v string) *ModifyNodePoolNodeConfigResponseBody {
	s.NodepoolId = &v
	return s
}

func (s *ModifyNodePoolNodeConfigResponseBody) SetRequestId(v string) *ModifyNodePoolNodeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodePoolNodeConfigResponseBody) SetTaskId(v string) *ModifyNodePoolNodeConfigResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyNodePoolNodeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
