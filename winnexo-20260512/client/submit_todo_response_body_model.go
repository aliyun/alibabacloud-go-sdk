// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTodoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitTodoResponseBody
	GetCode() *string
	SetCompletedAt(v string) *SubmitTodoResponseBody
	GetCompletedAt() *string
	SetCompletedByUserId(v string) *SubmitTodoResponseBody
	GetCompletedByUserId() *string
	SetMessage(v string) *SubmitTodoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitTodoResponseBody
	GetRequestId() *string
	SetStatus(v string) *SubmitTodoResponseBody
	GetStatus() *string
	SetTodoId(v string) *SubmitTodoResponseBody
	GetTodoId() *string
}

type SubmitTodoResponseBody struct {
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 完成时间（ISO8601 UTC；与其他北京时间格式化字段格式不同）
	//
	// example:
	//
	// 2026-10-12T02:00:00+00:00
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// 完成人用户ID
	//
	// example:
	//
	// 42
	CompletedByUserId *string `json:"completedByUserId,omitempty" xml:"completedByUserId,omitempty"`
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
	// 提交后状态（恒为 COMPLETED）
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 待办ID
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s SubmitTodoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTodoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTodoResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitTodoResponseBody) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *SubmitTodoResponseBody) GetCompletedByUserId() *string {
	return s.CompletedByUserId
}

func (s *SubmitTodoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitTodoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTodoResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SubmitTodoResponseBody) GetTodoId() *string {
	return s.TodoId
}

func (s *SubmitTodoResponseBody) SetCode(v string) *SubmitTodoResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTodoResponseBody) SetCompletedAt(v string) *SubmitTodoResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *SubmitTodoResponseBody) SetCompletedByUserId(v string) *SubmitTodoResponseBody {
	s.CompletedByUserId = &v
	return s
}

func (s *SubmitTodoResponseBody) SetMessage(v string) *SubmitTodoResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTodoResponseBody) SetRequestId(v string) *SubmitTodoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTodoResponseBody) SetStatus(v string) *SubmitTodoResponseBody {
	s.Status = &v
	return s
}

func (s *SubmitTodoResponseBody) SetTodoId(v string) *SubmitTodoResponseBody {
	s.TodoId = &v
	return s
}

func (s *SubmitTodoResponseBody) Validate() error {
	return dara.Validate(s)
}
