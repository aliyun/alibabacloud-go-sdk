// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendChatMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendChatMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendChatMessageResponseBody) *SendChatMessageResponse
	GetBody() *SendChatMessageResponseBody
}

type SendChatMessageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendChatMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendChatMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageResponse) GoString() string {
	return s.String()
}

func (s *SendChatMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendChatMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendChatMessageResponse) GetBody() *SendChatMessageResponseBody {
	return s.Body
}

func (s *SendChatMessageResponse) SetHeaders(v map[string]*string) *SendChatMessageResponse {
	s.Headers = v
	return s
}

func (s *SendChatMessageResponse) SetStatusCode(v int32) *SendChatMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendChatMessageResponse) SetBody(v *SendChatMessageResponseBody) *SendChatMessageResponse {
	s.Body = v
	return s
}

func (s *SendChatMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
