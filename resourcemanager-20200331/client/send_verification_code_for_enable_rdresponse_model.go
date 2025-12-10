// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationCodeForEnableRDResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendVerificationCodeForEnableRDResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendVerificationCodeForEnableRDResponse
	GetStatusCode() *int32
	SetBody(v *SendVerificationCodeForEnableRDResponseBody) *SendVerificationCodeForEnableRDResponse
	GetBody() *SendVerificationCodeForEnableRDResponseBody
}

type SendVerificationCodeForEnableRDResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerificationCodeForEnableRDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerificationCodeForEnableRDResponse) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationCodeForEnableRDResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForEnableRDResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendVerificationCodeForEnableRDResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendVerificationCodeForEnableRDResponse) GetBody() *SendVerificationCodeForEnableRDResponseBody {
	return s.Body
}

func (s *SendVerificationCodeForEnableRDResponse) SetHeaders(v map[string]*string) *SendVerificationCodeForEnableRDResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationCodeForEnableRDResponse) SetStatusCode(v int32) *SendVerificationCodeForEnableRDResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationCodeForEnableRDResponse) SetBody(v *SendVerificationCodeForEnableRDResponseBody) *SendVerificationCodeForEnableRDResponse {
	s.Body = v
	return s
}

func (s *SendVerificationCodeForEnableRDResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
