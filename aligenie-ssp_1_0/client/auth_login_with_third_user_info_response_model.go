// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithThirdUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthLoginWithThirdUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthLoginWithThirdUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *AuthLoginWithThirdUserInfoResponseBody) *AuthLoginWithThirdUserInfoResponse
	GetBody() *AuthLoginWithThirdUserInfoResponseBody
}

type AuthLoginWithThirdUserInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthLoginWithThirdUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthLoginWithThirdUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoResponse) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthLoginWithThirdUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthLoginWithThirdUserInfoResponse) GetBody() *AuthLoginWithThirdUserInfoResponseBody {
	return s.Body
}

func (s *AuthLoginWithThirdUserInfoResponse) SetHeaders(v map[string]*string) *AuthLoginWithThirdUserInfoResponse {
	s.Headers = v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponse) SetStatusCode(v int32) *AuthLoginWithThirdUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponse) SetBody(v *AuthLoginWithThirdUserInfoResponseBody) *AuthLoginWithThirdUserInfoResponse {
	s.Body = v
	return s
}

func (s *AuthLoginWithThirdUserInfoResponse) Validate() error {
	return dara.Validate(s)
}
