// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyPhoneWithTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyPhoneWithTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyPhoneWithTokenResponse
	GetStatusCode() *int32
	SetBody(v *VerifyPhoneWithTokenResponseBody) *VerifyPhoneWithTokenResponse
	GetBody() *VerifyPhoneWithTokenResponseBody
}

type VerifyPhoneWithTokenResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyPhoneWithTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyPhoneWithTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyPhoneWithTokenResponse) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyPhoneWithTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyPhoneWithTokenResponse) GetBody() *VerifyPhoneWithTokenResponseBody {
	return s.Body
}

func (s *VerifyPhoneWithTokenResponse) SetHeaders(v map[string]*string) *VerifyPhoneWithTokenResponse {
	s.Headers = v
	return s
}

func (s *VerifyPhoneWithTokenResponse) SetStatusCode(v int32) *VerifyPhoneWithTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyPhoneWithTokenResponse) SetBody(v *VerifyPhoneWithTokenResponseBody) *VerifyPhoneWithTokenResponse {
	s.Body = v
	return s
}

func (s *VerifyPhoneWithTokenResponse) Validate() error {
	return dara.Validate(s)
}
