// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TokenResponse
	GetStatusCode() *int32
	SetBody(v *Token) *TokenResponse
	GetBody() *Token
}

type TokenResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Token             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TokenResponse) String() string {
	return dara.Prettify(s)
}

func (s TokenResponse) GoString() string {
	return s.String()
}

func (s *TokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TokenResponse) GetBody() *Token {
	return s.Body
}

func (s *TokenResponse) SetHeaders(v map[string]*string) *TokenResponse {
	s.Headers = v
	return s
}

func (s *TokenResponse) SetStatusCode(v int32) *TokenResponse {
	s.StatusCode = &v
	return s
}

func (s *TokenResponse) SetBody(v *Token) *TokenResponse {
	s.Body = v
	return s
}

func (s *TokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
