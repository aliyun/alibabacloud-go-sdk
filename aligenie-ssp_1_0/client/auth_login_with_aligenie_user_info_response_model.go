// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithAligenieUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthLoginWithAligenieUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *AuthLoginWithAligenieUserInfoResponseBody) *AuthLoginWithAligenieUserInfoResponse
	GetBody() *AuthLoginWithAligenieUserInfoResponseBody
}

type AuthLoginWithAligenieUserInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthLoginWithAligenieUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoResponse) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthLoginWithAligenieUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthLoginWithAligenieUserInfoResponse) GetBody() *AuthLoginWithAligenieUserInfoResponseBody {
	return s.Body
}

func (s *AuthLoginWithAligenieUserInfoResponse) SetHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoResponse {
	s.Headers = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponse) SetStatusCode(v int32) *AuthLoginWithAligenieUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponse) SetBody(v *AuthLoginWithAligenieUserInfoResponseBody) *AuthLoginWithAligenieUserInfoResponse {
	s.Body = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoResponse) Validate() error {
	return dara.Validate(s)
}
