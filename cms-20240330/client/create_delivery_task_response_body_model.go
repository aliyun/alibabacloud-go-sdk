// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDeliveryTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateDeliveryTaskResponseBody
	GetTaskId() *string
}

type CreateDeliveryTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 4CB23A2E-B426-5D4B-9AA2-6C7A508D954B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 98e367d8fb8cc83b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeliveryTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDeliveryTaskResponseBody) SetRequestId(v string) *CreateDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeliveryTaskResponseBody) SetTaskId(v string) *CreateDeliveryTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
