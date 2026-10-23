// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceManualMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceManualMemoryResponseBody
	GetCode() *string
	SetContent(v string) *ReplaceManualMemoryResponseBody
	GetContent() *string
	SetMessage(v string) *ReplaceManualMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReplaceManualMemoryResponseBody
	GetRequestId() *string
}

type ReplaceManualMemoryResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The complete content of the manual memory after it is saved. If the memory is successfully cleared, an empty string is returned.
	//
	// example:
	//
	// Prefers concise answers
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The error description.
	//
	// example:
	//
	// The current user is not an active member of the group cg_12345
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ReplaceManualMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceManualMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceManualMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceManualMemoryResponseBody) GetContent() *string {
	return s.Content
}

func (s *ReplaceManualMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceManualMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceManualMemoryResponseBody) SetCode(v string) *ReplaceManualMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceManualMemoryResponseBody) SetContent(v string) *ReplaceManualMemoryResponseBody {
	s.Content = &v
	return s
}

func (s *ReplaceManualMemoryResponseBody) SetMessage(v string) *ReplaceManualMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceManualMemoryResponseBody) SetRequestId(v string) *ReplaceManualMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceManualMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
