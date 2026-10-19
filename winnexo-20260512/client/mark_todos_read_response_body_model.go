// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkTodosReadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MarkTodosReadResponseBody
	GetCode() *string
	SetMessage(v string) *MarkTodosReadResponseBody
	GetMessage() *string
	SetRequestId(v string) *MarkTodosReadResponseBody
	GetRequestId() *string
	SetUnreadCount(v int64) *MarkTodosReadResponseBody
	GetUnreadCount() *int64
	SetUpdatedCount(v int64) *MarkTodosReadResponseBody
	GetUpdatedCount() *int64
}

type MarkTodosReadResponseBody struct {
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
	// 标记后剩余未读数
	//
	// example:
	//
	// 0
	UnreadCount *int64 `json:"unreadCount,omitempty" xml:"unreadCount,omitempty"`
	// 本次标记条数
	//
	// example:
	//
	// 2
	UpdatedCount *int64 `json:"updatedCount,omitempty" xml:"updatedCount,omitempty"`
}

func (s MarkTodosReadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MarkTodosReadResponseBody) GoString() string {
	return s.String()
}

func (s *MarkTodosReadResponseBody) GetCode() *string {
	return s.Code
}

func (s *MarkTodosReadResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MarkTodosReadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MarkTodosReadResponseBody) GetUnreadCount() *int64 {
	return s.UnreadCount
}

func (s *MarkTodosReadResponseBody) GetUpdatedCount() *int64 {
	return s.UpdatedCount
}

func (s *MarkTodosReadResponseBody) SetCode(v string) *MarkTodosReadResponseBody {
	s.Code = &v
	return s
}

func (s *MarkTodosReadResponseBody) SetMessage(v string) *MarkTodosReadResponseBody {
	s.Message = &v
	return s
}

func (s *MarkTodosReadResponseBody) SetRequestId(v string) *MarkTodosReadResponseBody {
	s.RequestId = &v
	return s
}

func (s *MarkTodosReadResponseBody) SetUnreadCount(v int64) *MarkTodosReadResponseBody {
	s.UnreadCount = &v
	return s
}

func (s *MarkTodosReadResponseBody) SetUpdatedCount(v int64) *MarkTodosReadResponseBody {
	s.UpdatedCount = &v
	return s
}

func (s *MarkTodosReadResponseBody) Validate() error {
	return dara.Validate(s)
}
