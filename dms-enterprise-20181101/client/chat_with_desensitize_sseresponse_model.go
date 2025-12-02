// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeSSEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatWithDesensitizeSSEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatWithDesensitizeSSEResponse
	GetStatusCode() *int32
	SetBody(v *ChatWithDesensitizeSSEResponseBody) *ChatWithDesensitizeSSEResponse
	GetBody() *ChatWithDesensitizeSSEResponseBody
}

type ChatWithDesensitizeSSEResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatWithDesensitizeSSEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatWithDesensitizeSSEResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeSSEResponse) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeSSEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatWithDesensitizeSSEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatWithDesensitizeSSEResponse) GetBody() *ChatWithDesensitizeSSEResponseBody {
	return s.Body
}

func (s *ChatWithDesensitizeSSEResponse) SetHeaders(v map[string]*string) *ChatWithDesensitizeSSEResponse {
	s.Headers = v
	return s
}

func (s *ChatWithDesensitizeSSEResponse) SetStatusCode(v int32) *ChatWithDesensitizeSSEResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatWithDesensitizeSSEResponse) SetBody(v *ChatWithDesensitizeSSEResponseBody) *ChatWithDesensitizeSSEResponse {
	s.Body = v
	return s
}

func (s *ChatWithDesensitizeSSEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
