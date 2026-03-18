// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstateAuthenticationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReinstateAuthenticationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReinstateAuthenticationTokenResponse
	GetStatusCode() *int32
}

type ReinstateAuthenticationTokenResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ReinstateAuthenticationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ReinstateAuthenticationTokenResponse) GoString() string {
	return s.String()
}

func (s *ReinstateAuthenticationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReinstateAuthenticationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReinstateAuthenticationTokenResponse) SetHeaders(v map[string]*string) *ReinstateAuthenticationTokenResponse {
	s.Headers = v
	return s
}

func (s *ReinstateAuthenticationTokenResponse) SetStatusCode(v int32) *ReinstateAuthenticationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ReinstateAuthenticationTokenResponse) Validate() error {
	return dara.Validate(s)
}
