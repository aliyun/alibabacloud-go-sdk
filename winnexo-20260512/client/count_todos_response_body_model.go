// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountTodosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CountTodosResponseBody
	GetCode() *string
	SetMessage(v string) *CountTodosResponseBody
	GetMessage() *string
	SetPendingCount(v int64) *CountTodosResponseBody
	GetPendingCount() *int64
	SetRequestId(v string) *CountTodosResponseBody
	GetRequestId() *string
	SetUnreadCount(v int64) *CountTodosResponseBody
	GetUnreadCount() *int64
}

type CountTodosResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The number of pending items with the primary status of PENDING.
	//
	// example:
	//
	// 3
	PendingCount *int64 `json:"pendingCount,omitempty" xml:"pendingCount,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of unread items where the readAt field of the handler is empty.
	//
	// example:
	//
	// 1
	UnreadCount *int64 `json:"unreadCount,omitempty" xml:"unreadCount,omitempty"`
}

func (s CountTodosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CountTodosResponseBody) GoString() string {
	return s.String()
}

func (s *CountTodosResponseBody) GetCode() *string {
	return s.Code
}

func (s *CountTodosResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CountTodosResponseBody) GetPendingCount() *int64 {
	return s.PendingCount
}

func (s *CountTodosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CountTodosResponseBody) GetUnreadCount() *int64 {
	return s.UnreadCount
}

func (s *CountTodosResponseBody) SetCode(v string) *CountTodosResponseBody {
	s.Code = &v
	return s
}

func (s *CountTodosResponseBody) SetMessage(v string) *CountTodosResponseBody {
	s.Message = &v
	return s
}

func (s *CountTodosResponseBody) SetPendingCount(v int64) *CountTodosResponseBody {
	s.PendingCount = &v
	return s
}

func (s *CountTodosResponseBody) SetRequestId(v string) *CountTodosResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountTodosResponseBody) SetUnreadCount(v int64) *CountTodosResponseBody {
	s.UnreadCount = &v
	return s
}

func (s *CountTodosResponseBody) Validate() error {
	return dara.Validate(s)
}
