// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithTaobaoUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthLoginWithTaobaoUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthLoginWithTaobaoUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *AuthLoginWithTaobaoUserInfoResponseBody) *AuthLoginWithTaobaoUserInfoResponse
	GetBody() *AuthLoginWithTaobaoUserInfoResponseBody
}

type AuthLoginWithTaobaoUserInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthLoginWithTaobaoUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthLoginWithTaobaoUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoResponse) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthLoginWithTaobaoUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthLoginWithTaobaoUserInfoResponse) GetBody() *AuthLoginWithTaobaoUserInfoResponseBody {
	return s.Body
}

func (s *AuthLoginWithTaobaoUserInfoResponse) SetHeaders(v map[string]*string) *AuthLoginWithTaobaoUserInfoResponse {
	s.Headers = v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponse) SetStatusCode(v int32) *AuthLoginWithTaobaoUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponse) SetBody(v *AuthLoginWithTaobaoUserInfoResponseBody) *AuthLoginWithTaobaoUserInfoResponse {
	s.Body = v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoResponse) Validate() error {
	return dara.Validate(s)
}
