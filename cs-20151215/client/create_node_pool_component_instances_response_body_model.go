// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodePoolComponentInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNodePoolComponentInstancesResponseBody
	GetClusterId() *string
	SetRequestId(v string) *CreateNodePoolComponentInstancesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateNodePoolComponentInstancesResponseBody
	GetTaskId() *string
}

type CreateNodePoolComponentInstancesResponseBody struct {
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

func (s CreateNodePoolComponentInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodePoolComponentInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodePoolComponentInstancesResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNodePoolComponentInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodePoolComponentInstancesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateNodePoolComponentInstancesResponseBody) SetClusterId(v string) *CreateNodePoolComponentInstancesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateNodePoolComponentInstancesResponseBody) SetRequestId(v string) *CreateNodePoolComponentInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodePoolComponentInstancesResponseBody) SetTaskId(v string) *CreateNodePoolComponentInstancesResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateNodePoolComponentInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
