// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConversationMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppConversationMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppConversationMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListAppConversationMessagesResponseBody) *ListAppConversationMessagesResponse
	GetBody() *ListAppConversationMessagesResponseBody
}

type ListAppConversationMessagesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppConversationMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppConversationMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppConversationMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListAppConversationMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppConversationMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppConversationMessagesResponse) GetBody() *ListAppConversationMessagesResponseBody {
	return s.Body
}

func (s *ListAppConversationMessagesResponse) SetHeaders(v map[string]*string) *ListAppConversationMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListAppConversationMessagesResponse) SetStatusCode(v int32) *ListAppConversationMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppConversationMessagesResponse) SetBody(v *ListAppConversationMessagesResponseBody) *ListAppConversationMessagesResponse {
	s.Body = v
	return s
}

func (s *ListAppConversationMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
