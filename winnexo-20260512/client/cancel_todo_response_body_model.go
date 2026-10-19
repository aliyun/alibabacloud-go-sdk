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
	// 作废条数
	//
	// example:
	//
	// 1
	CancelledCount *int64 `json:"cancelledCount,omitempty" xml:"cancelledCount,omitempty"`
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 作废后状态
	//
	// example:
	//
	// CANCELLED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 作废的待办ID（按 workItemId 定位时为空）
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
