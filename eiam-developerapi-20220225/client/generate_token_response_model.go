// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenerateTokenResponseBody) *GenerateTokenResponse
	GetBody() *GenerateTokenResponseBody
}

type GenerateTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateTokenResponse) GetBody() *GenerateTokenResponseBody {
	return s.Body
}

func (s *GenerateTokenResponse) SetHeaders(v map[string]*string) *GenerateTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateTokenResponse) SetStatusCode(v int32) *GenerateTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTokenResponse) SetBody(v *GenerateTokenResponseBody) *GenerateTokenResponse {
	s.Body = v
	return s
}

func (s *GenerateTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
