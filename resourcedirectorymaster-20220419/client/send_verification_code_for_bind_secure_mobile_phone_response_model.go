// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationCodeForBindSecureMobilePhoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendVerificationCodeForBindSecureMobilePhoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendVerificationCodeForBindSecureMobilePhoneResponse
	GetStatusCode() *int32
	SetBody(v *SendVerificationCodeForBindSecureMobilePhoneResponseBody) *SendVerificationCodeForBindSecureMobilePhoneResponse
	GetBody() *SendVerificationCodeForBindSecureMobilePhoneResponseBody
}

type SendVerificationCodeForBindSecureMobilePhoneResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerificationCodeForBindSecureMobilePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponse) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) GetBody() *SendVerificationCodeForBindSecureMobilePhoneResponseBody {
	return s.Body
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) SetHeaders(v map[string]*string) *SendVerificationCodeForBindSecureMobilePhoneResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) SetStatusCode(v int32) *SendVerificationCodeForBindSecureMobilePhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) SetBody(v *SendVerificationCodeForBindSecureMobilePhoneResponseBody) *SendVerificationCodeForBindSecureMobilePhoneResponse {
	s.Body = v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
