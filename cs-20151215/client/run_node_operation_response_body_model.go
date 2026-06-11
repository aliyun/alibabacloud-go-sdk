// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNodeOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RunNodeOperationResponseBody
	GetClusterId() *string
	SetRequestId(v string) *RunNodeOperationResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RunNodeOperationResponseBody
	GetTaskId() *string
}

type RunNodeOperationResponseBody struct {
	// example:
	//
	// c2230fxxxxx
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// T-xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunNodeOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunNodeOperationResponseBody) GoString() string {
	return s.String()
}

func (s *RunNodeOperationResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *RunNodeOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunNodeOperationResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RunNodeOperationResponseBody) SetClusterId(v string) *RunNodeOperationResponseBody {
	s.ClusterId = &v
	return s
}

func (s *RunNodeOperationResponseBody) SetRequestId(v string) *RunNodeOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunNodeOperationResponseBody) SetTaskId(v string) *RunNodeOperationResponseBody {
	s.TaskId = &v
	return s
}

func (s *RunNodeOperationResponseBody) Validate() error {
	return dara.Validate(s)
}
