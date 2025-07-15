// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupChatMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupChatMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupChatMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupChatMessagesResponseBody) *ListGroupChatMessagesResponse
	GetBody() *ListGroupChatMessagesResponseBody
}

type ListGroupChatMessagesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupChatMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupChatMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupChatMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListGroupChatMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupChatMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupChatMessagesResponse) GetBody() *ListGroupChatMessagesResponseBody {
	return s.Body
}

func (s *ListGroupChatMessagesResponse) SetHeaders(v map[string]*string) *ListGroupChatMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListGroupChatMessagesResponse) SetStatusCode(v int32) *ListGroupChatMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupChatMessagesResponse) SetBody(v *ListGroupChatMessagesResponseBody) *ListGroupChatMessagesResponse {
	s.Body = v
	return s
}

func (s *ListGroupChatMessagesResponse) Validate() error {
	return dara.Validate(s)
}
