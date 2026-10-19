// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTodosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCancelledCount(v int64) *BatchCancelTodosResponseBody
	GetCancelledCount() *int64
	SetCode(v string) *BatchCancelTodosResponseBody
	GetCode() *string
	SetMessage(v string) *BatchCancelTodosResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchCancelTodosResponseBody
	GetRequestId() *string
}

type BatchCancelTodosResponseBody struct {
	// 实际作废条数（幂等重复调用为 0）
	//
	// example:
	//
	// 2
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
}

func (s BatchCancelTodosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTodosResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCancelTodosResponseBody) GetCancelledCount() *int64 {
	return s.CancelledCount
}

func (s *BatchCancelTodosResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchCancelTodosResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchCancelTodosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCancelTodosResponseBody) SetCancelledCount(v int64) *BatchCancelTodosResponseBody {
	s.CancelledCount = &v
	return s
}

func (s *BatchCancelTodosResponseBody) SetCode(v string) *BatchCancelTodosResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCancelTodosResponseBody) SetMessage(v string) *BatchCancelTodosResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCancelTodosResponseBody) SetRequestId(v string) *BatchCancelTodosResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCancelTodosResponseBody) Validate() error {
	return dara.Validate(s)
}
