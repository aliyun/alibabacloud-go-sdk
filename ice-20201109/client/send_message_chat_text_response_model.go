// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageChatTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendMessageChatTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendMessageChatTextResponse
	GetStatusCode() *int32
	SetBody(v *SendMessageChatTextResponseBody) *SendMessageChatTextResponse
	GetBody() *SendMessageChatTextResponseBody
}

type SendMessageChatTextResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageChatTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageChatTextResponse) String() string {
	return dara.Prettify(s)
}

func (s SendMessageChatTextResponse) GoString() string {
	return s.String()
}

func (s *SendMessageChatTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendMessageChatTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendMessageChatTextResponse) GetBody() *SendMessageChatTextResponseBody {
	return s.Body
}

func (s *SendMessageChatTextResponse) SetHeaders(v map[string]*string) *SendMessageChatTextResponse {
	s.Headers = v
	return s
}

func (s *SendMessageChatTextResponse) SetStatusCode(v int32) *SendMessageChatTextResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageChatTextResponse) SetBody(v *SendMessageChatTextResponseBody) *SendMessageChatTextResponse {
	s.Body = v
	return s
}

func (s *SendMessageChatTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
