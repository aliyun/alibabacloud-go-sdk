// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTestTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceTestTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateServiceTestTaskResponseBody
	GetTaskId() *string
}

type CreateServiceTestTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DF0F666F-FBBC-55C3-A368-C955DE7B4839
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// stt-568c2c5a687a409b977e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateServiceTestTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceTestTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceTestTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateServiceTestTaskResponseBody) SetRequestId(v string) *CreateServiceTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceTestTaskResponseBody) SetTaskId(v string) *CreateServiceTestTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateServiceTestTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
