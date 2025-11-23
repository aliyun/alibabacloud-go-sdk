// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatWithDesensitizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatWithDesensitizeResponse
	GetStatusCode() *int32
	SetBody(v *ChatWithDesensitizeResponseBody) *ChatWithDesensitizeResponse
	GetBody() *ChatWithDesensitizeResponseBody
}

type ChatWithDesensitizeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatWithDesensitizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatWithDesensitizeResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeResponse) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatWithDesensitizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatWithDesensitizeResponse) GetBody() *ChatWithDesensitizeResponseBody {
	return s.Body
}

func (s *ChatWithDesensitizeResponse) SetHeaders(v map[string]*string) *ChatWithDesensitizeResponse {
	s.Headers = v
	return s
}

func (s *ChatWithDesensitizeResponse) SetStatusCode(v int32) *ChatWithDesensitizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatWithDesensitizeResponse) SetBody(v *ChatWithDesensitizeResponseBody) *ChatWithDesensitizeResponse {
	s.Body = v
	return s
}

func (s *ChatWithDesensitizeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
