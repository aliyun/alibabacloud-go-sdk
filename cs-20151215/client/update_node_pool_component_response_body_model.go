// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodePoolComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateNodePoolComponentResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UpdateNodePoolComponentResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateNodePoolComponentResponseBody
	GetTaskId() *string
}

type UpdateNodePoolComponentResponseBody struct {
	// example:
	//
	// c82e6987e2961451182edacd74faf****
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 49511F2D-D56A-5C24-B9AE-C8491E09B***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// T-67d7ec016ce37c0106000***
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateNodePoolComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateNodePoolComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodePoolComponentResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateNodePoolComponentResponseBody) SetClusterId(v string) *UpdateNodePoolComponentResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpdateNodePoolComponentResponseBody) SetRequestId(v string) *UpdateNodePoolComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodePoolComponentResponseBody) SetTaskId(v string) *UpdateNodePoolComponentResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateNodePoolComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
