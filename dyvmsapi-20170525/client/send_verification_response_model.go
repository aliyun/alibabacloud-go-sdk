// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendVerificationResponse
	GetStatusCode() *int32
	SetBody(v *SendVerificationResponseBody) *SendVerificationResponse
	GetBody() *SendVerificationResponseBody
}

type SendVerificationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendVerificationResponse) GetBody() *SendVerificationResponseBody {
	return s.Body
}

func (s *SendVerificationResponse) SetHeaders(v map[string]*string) *SendVerificationResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationResponse) SetStatusCode(v int32) *SendVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationResponse) SetBody(v *SendVerificationResponseBody) *SendVerificationResponse {
	s.Body = v
	return s
}

func (s *SendVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
