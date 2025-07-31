// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *DeleteNodeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *DeleteNodeResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *DeleteNodeResponseBody
	GetTaskId() *int32
}

type DeleteNodeResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 111111111111111
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9F9BDE64-BF30-41F3-BD29-C19CE4AB3404
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 111111111
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DeleteNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNodeResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteNodeResponseBody) SetOrderId(v string) *DeleteNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *DeleteNodeResponseBody) SetRequestId(v string) *DeleteNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeResponseBody) SetTaskId(v int32) *DeleteNodeResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
