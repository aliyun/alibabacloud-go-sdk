// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCardSmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendCardSmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendCardSmsResponse
	GetStatusCode() *int32
	SetBody(v *SendCardSmsResponseBody) *SendCardSmsResponse
	GetBody() *SendCardSmsResponseBody
}

type SendCardSmsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCardSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendCardSmsResponse) String() string {
	return dara.Prettify(s)
}

func (s SendCardSmsResponse) GoString() string {
	return s.String()
}

func (s *SendCardSmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendCardSmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendCardSmsResponse) GetBody() *SendCardSmsResponseBody {
	return s.Body
}

func (s *SendCardSmsResponse) SetHeaders(v map[string]*string) *SendCardSmsResponse {
	s.Headers = v
	return s
}

func (s *SendCardSmsResponse) SetStatusCode(v int32) *SendCardSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCardSmsResponse) SetBody(v *SendCardSmsResponseBody) *SendCardSmsResponse {
	s.Body = v
	return s
}

func (s *SendCardSmsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
