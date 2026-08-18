// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodePoolComponentInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateNodePoolComponentInstanceResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UpdateNodePoolComponentInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateNodePoolComponentInstanceResponseBody
	GetTaskId() *string
}

type UpdateNodePoolComponentInstanceResponseBody struct {
	// example:
	//
	// c2230fxxxxx
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// xxxx
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// T-xxxx
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpdateNodePoolComponentInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentInstanceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateNodePoolComponentInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodePoolComponentInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateNodePoolComponentInstanceResponseBody) SetClusterId(v string) *UpdateNodePoolComponentInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceResponseBody) SetRequestId(v string) *UpdateNodePoolComponentInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceResponseBody) SetTaskId(v string) *UpdateNodePoolComponentInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
