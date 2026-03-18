// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateJwtAuthenticationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateJwtAuthenticationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateJwtAuthenticationTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenerateJwtAuthenticationTokenResponseBody) *GenerateJwtAuthenticationTokenResponse
	GetBody() *GenerateJwtAuthenticationTokenResponseBody
}

type GenerateJwtAuthenticationTokenResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateJwtAuthenticationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateJwtAuthenticationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateJwtAuthenticationTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateJwtAuthenticationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateJwtAuthenticationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateJwtAuthenticationTokenResponse) GetBody() *GenerateJwtAuthenticationTokenResponseBody {
	return s.Body
}

func (s *GenerateJwtAuthenticationTokenResponse) SetHeaders(v map[string]*string) *GenerateJwtAuthenticationTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponse) SetStatusCode(v int32) *GenerateJwtAuthenticationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponse) SetBody(v *GenerateJwtAuthenticationTokenResponseBody) *GenerateJwtAuthenticationTokenResponse {
	s.Body = v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
