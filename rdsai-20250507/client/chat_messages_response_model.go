// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ChatMessagesResponseBody) *ChatMessagesResponse
	GetBody() *ChatMessagesResponseBody
}

type ChatMessagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatMessagesResponse) GoString() string {
	return s.String()
}

func (s *ChatMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatMessagesResponse) GetBody() *ChatMessagesResponseBody {
	return s.Body
}

func (s *ChatMessagesResponse) SetHeaders(v map[string]*string) *ChatMessagesResponse {
	s.Headers = v
	return s
}

func (s *ChatMessagesResponse) SetStatusCode(v int32) *ChatMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatMessagesResponse) SetBody(v *ChatMessagesResponseBody) *ChatMessagesResponse {
	s.Body = v
	return s
}

func (s *ChatMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
