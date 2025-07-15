// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisitorChatMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVisitorChatMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVisitorChatMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListVisitorChatMessagesResponseBody) *ListVisitorChatMessagesResponse
	GetBody() *ListVisitorChatMessagesResponseBody
}

type ListVisitorChatMessagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVisitorChatMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVisitorChatMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVisitorChatMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListVisitorChatMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVisitorChatMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVisitorChatMessagesResponse) GetBody() *ListVisitorChatMessagesResponseBody {
	return s.Body
}

func (s *ListVisitorChatMessagesResponse) SetHeaders(v map[string]*string) *ListVisitorChatMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListVisitorChatMessagesResponse) SetStatusCode(v int32) *ListVisitorChatMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVisitorChatMessagesResponse) SetBody(v *ListVisitorChatMessagesResponseBody) *ListVisitorChatMessagesResponse {
	s.Body = v
	return s
}

func (s *ListVisitorChatMessagesResponse) Validate() error {
	return dara.Validate(s)
}
