// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendChatappMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendChatappMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendChatappMessageResponseBody) *SendChatappMessageResponse
	GetBody() *SendChatappMessageResponseBody
}

type SendChatappMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendChatappMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendChatappMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageResponse) GoString() string {
	return s.String()
}

func (s *SendChatappMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendChatappMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendChatappMessageResponse) GetBody() *SendChatappMessageResponseBody {
	return s.Body
}

func (s *SendChatappMessageResponse) SetHeaders(v map[string]*string) *SendChatappMessageResponse {
	s.Headers = v
	return s
}

func (s *SendChatappMessageResponse) SetStatusCode(v int32) *SendChatappMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendChatappMessageResponse) SetBody(v *SendChatappMessageResponseBody) *SendChatappMessageResponse {
	s.Body = v
	return s
}

func (s *SendChatappMessageResponse) Validate() error {
	return dara.Validate(s)
}
