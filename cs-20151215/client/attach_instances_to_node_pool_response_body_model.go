// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesToNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachInstancesToNodePoolResponseBody
	GetRequestId() *string
	SetTaskId(v string) *AttachInstancesToNodePoolResponseBody
	GetTaskId() *string
}

type AttachInstancesToNodePoolResponseBody struct {
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
	// T-5a54309c80282e39ea00002f
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s AttachInstancesToNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesToNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesToNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachInstancesToNodePoolResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *AttachInstancesToNodePoolResponseBody) SetRequestId(v string) *AttachInstancesToNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachInstancesToNodePoolResponseBody) SetTaskId(v string) *AttachInstancesToNodePoolResponseBody {
	s.TaskId = &v
	return s
}

func (s *AttachInstancesToNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
