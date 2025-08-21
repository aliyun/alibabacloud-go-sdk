// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendMessageResponseBody) *SendMessageResponse
	GetBody() *SendMessageResponseBody
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendMessageResponse) GetBody() *SendMessageResponseBody {
	return s.Body
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

func (s *SendMessageResponse) Validate() error {
	return dara.Validate(s)
}
