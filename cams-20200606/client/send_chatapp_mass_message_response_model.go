// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMassMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendChatappMassMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendChatappMassMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendChatappMassMessageResponseBody) *SendChatappMassMessageResponse
	GetBody() *SendChatappMassMessageResponseBody
}

type SendChatappMassMessageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendChatappMassMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendChatappMassMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageResponse) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendChatappMassMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendChatappMassMessageResponse) GetBody() *SendChatappMassMessageResponseBody {
	return s.Body
}

func (s *SendChatappMassMessageResponse) SetHeaders(v map[string]*string) *SendChatappMassMessageResponse {
	s.Headers = v
	return s
}

func (s *SendChatappMassMessageResponse) SetStatusCode(v int32) *SendChatappMassMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendChatappMassMessageResponse) SetBody(v *SendChatappMassMessageResponseBody) *SendChatappMassMessageResponse {
	s.Body = v
	return s
}

func (s *SendChatappMassMessageResponse) Validate() error {
	return dara.Validate(s)
}
