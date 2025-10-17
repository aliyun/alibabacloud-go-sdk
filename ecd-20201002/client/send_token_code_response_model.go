// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTokenCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendTokenCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendTokenCodeResponse
	GetStatusCode() *int32
	SetBody(v *SendTokenCodeResponseBody) *SendTokenCodeResponse
	GetBody() *SendTokenCodeResponseBody
}

type SendTokenCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTokenCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTokenCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s SendTokenCodeResponse) GoString() string {
	return s.String()
}

func (s *SendTokenCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendTokenCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendTokenCodeResponse) GetBody() *SendTokenCodeResponseBody {
	return s.Body
}

func (s *SendTokenCodeResponse) SetHeaders(v map[string]*string) *SendTokenCodeResponse {
	s.Headers = v
	return s
}

func (s *SendTokenCodeResponse) SetStatusCode(v int32) *SendTokenCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTokenCodeResponse) SetBody(v *SendTokenCodeResponseBody) *SendTokenCodeResponse {
	s.Body = v
	return s
}

func (s *SendTokenCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
