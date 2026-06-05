// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppChatMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppChatMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppChatMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListAppChatMessagesResponseBody) *ListAppChatMessagesResponse
	GetBody() *ListAppChatMessagesResponseBody
}

type ListAppChatMessagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppChatMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppChatMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppChatMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListAppChatMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppChatMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppChatMessagesResponse) GetBody() *ListAppChatMessagesResponseBody {
	return s.Body
}

func (s *ListAppChatMessagesResponse) SetHeaders(v map[string]*string) *ListAppChatMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListAppChatMessagesResponse) SetStatusCode(v int32) *ListAppChatMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppChatMessagesResponse) SetBody(v *ListAppChatMessagesResponseBody) *ListAppChatMessagesResponse {
	s.Body = v
	return s
}

func (s *ListAppChatMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
