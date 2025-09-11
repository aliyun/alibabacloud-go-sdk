// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendSmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendSmsResponse
	GetStatusCode() *int32
	SetBody(v *SendSmsResponseBody) *SendSmsResponse
	GetBody() *SendSmsResponseBody
}

type SendSmsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSmsResponse) String() string {
	return dara.Prettify(s)
}

func (s SendSmsResponse) GoString() string {
	return s.String()
}

func (s *SendSmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendSmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendSmsResponse) GetBody() *SendSmsResponseBody {
	return s.Body
}

func (s *SendSmsResponse) SetHeaders(v map[string]*string) *SendSmsResponse {
	s.Headers = v
	return s
}

func (s *SendSmsResponse) SetStatusCode(v int32) *SendSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSmsResponse) SetBody(v *SendSmsResponseBody) *SendSmsResponse {
	s.Body = v
	return s
}

func (s *SendSmsResponse) Validate() error {
	return dara.Validate(s)
}
