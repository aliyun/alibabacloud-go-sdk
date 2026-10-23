// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTodoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCancelledCount(v int64) *CancelTodoResponseBody
	GetCancelledCount() *int64
	SetCode(v string) *CancelTodoResponseBody
	GetCode() *string
	SetMessage(v string) *CancelTodoResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelTodoResponseBody
	GetRequestId() *string
	SetStatus(v string) *CancelTodoResponseBody
	GetStatus() *string
	SetTodoId(v string) *CancelTodoResponseBody
	GetTodoId() *string
}

type CancelTodoResponseBody struct {
	// The number of canceled to-do items.
	//
	// example:
	//
	// 1
	CancelledCount *int64 `json:"cancelledCount,omitempty" xml:"cancelledCount,omitempty"`
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The status after cancellation.
	//
	// example:
	//
	// CANCELLED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the canceled to-do item. This parameter is empty if the to-do item is located by workItemId.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s CancelTodoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelTodoResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTodoResponseBody) GetCancelledCount() *int64 {
	return s.CancelledCount
}

func (s *CancelTodoResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelTodoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelTodoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelTodoResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CancelTodoResponseBody) GetTodoId() *string {
	return s.TodoId
}

func (s *CancelTodoResponseBody) SetCancelledCount(v int64) *CancelTodoResponseBody {
	s.CancelledCount = &v
	return s
}

func (s *CancelTodoResponseBody) SetCode(v string) *CancelTodoResponseBody {
	s.Code = &v
	return s
}

func (s *CancelTodoResponseBody) SetMessage(v string) *CancelTodoResponseBody {
	s.Message = &v
	return s
}

func (s *CancelTodoResponseBody) SetRequestId(v string) *CancelTodoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTodoResponseBody) SetStatus(v string) *CancelTodoResponseBody {
	s.Status = &v
	return s
}

func (s *CancelTodoResponseBody) SetTodoId(v string) *CancelTodoResponseBody {
	s.TodoId = &v
	return s
}

func (s *CancelTodoResponseBody) Validate() error {
	return dara.Validate(s)
}
