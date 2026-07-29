// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateServiceTaskResponseBody
	GetTaskId() *string
}

type CreateServiceTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the created task. Use this ID for subsequent query or deletion operations.
	//
	// example:
	//
	// a1b2c3d4-e5f6-7890-abcd-ef1234567890
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateServiceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateServiceTaskResponseBody) SetRequestId(v string) *CreateServiceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceTaskResponseBody) SetTaskId(v string) *CreateServiceTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateServiceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
