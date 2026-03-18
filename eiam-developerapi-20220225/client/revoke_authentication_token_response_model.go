// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAuthenticationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeAuthenticationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeAuthenticationTokenResponse
	GetStatusCode() *int32
}

type RevokeAuthenticationTokenResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RevokeAuthenticationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeAuthenticationTokenResponse) GoString() string {
	return s.String()
}

func (s *RevokeAuthenticationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeAuthenticationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeAuthenticationTokenResponse) SetHeaders(v map[string]*string) *RevokeAuthenticationTokenResponse {
	s.Headers = v
	return s
}

func (s *RevokeAuthenticationTokenResponse) SetStatusCode(v int32) *RevokeAuthenticationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeAuthenticationTokenResponse) Validate() error {
	return dara.Validate(s)
}
