// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse
	GetBody() *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse struct {
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) GetBody() *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody {
	return s.Body
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) SetHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) SetStatusCode(v int32) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) SetBody(v *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponseBody) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse) Validate() error {
	return dara.Validate(s)
}
