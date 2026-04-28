// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeResponse
	GetStatusCode() *int32
}

type AuthorizeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AuthorizeResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeResponse) SetHeaders(v map[string]*string) *AuthorizeResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeResponse) SetStatusCode(v int32) *AuthorizeResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeResponse) Validate() error {
	return dara.Validate(s)
}
