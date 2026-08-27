// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteChatSessionResponseBody
	GetCode() *string
	SetDeleted(v bool) *DeleteChatSessionResponseBody
	GetDeleted() *bool
	SetHardDelete(v bool) *DeleteChatSessionResponseBody
	GetHardDelete() *bool
	SetMessage(v string) *DeleteChatSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteChatSessionResponseBody
	GetRequestId() *string
	SetSessionId(v string) *DeleteChatSessionResponseBody
	GetSessionId() *string
}

type DeleteChatSessionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether the session is deleted.
	//
	// example:
	//
	// true
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// Indicates whether the session is hard-deleted.
	//
	// example:
	//
	// true
	HardDelete *bool `json:"hardDelete,omitempty" xml:"hardDelete,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The unique identifier of the function session.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s DeleteChatSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChatSessionResponseBody) GetDeleted() *bool {
	return s.Deleted
}

func (s *DeleteChatSessionResponseBody) GetHardDelete() *bool {
	return s.HardDelete
}

func (s *DeleteChatSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteChatSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChatSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteChatSessionResponseBody) SetCode(v string) *DeleteChatSessionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChatSessionResponseBody) SetDeleted(v bool) *DeleteChatSessionResponseBody {
	s.Deleted = &v
	return s
}

func (s *DeleteChatSessionResponseBody) SetHardDelete(v bool) *DeleteChatSessionResponseBody {
	s.HardDelete = &v
	return s
}

func (s *DeleteChatSessionResponseBody) SetMessage(v string) *DeleteChatSessionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChatSessionResponseBody) SetRequestId(v string) *DeleteChatSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChatSessionResponseBody) SetSessionId(v string) *DeleteChatSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *DeleteChatSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
