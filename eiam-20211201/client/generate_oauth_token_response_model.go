// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOauthTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateOauthTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateOauthTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenerateOauthTokenResponseBody) *GenerateOauthTokenResponse
	GetBody() *GenerateOauthTokenResponseBody
}

type GenerateOauthTokenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateOauthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateOauthTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateOauthTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateOauthTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateOauthTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateOauthTokenResponse) GetBody() *GenerateOauthTokenResponseBody {
	return s.Body
}

func (s *GenerateOauthTokenResponse) SetHeaders(v map[string]*string) *GenerateOauthTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateOauthTokenResponse) SetStatusCode(v int32) *GenerateOauthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateOauthTokenResponse) SetBody(v *GenerateOauthTokenResponseBody) *GenerateOauthTokenResponse {
	s.Body = v
	return s
}

func (s *GenerateOauthTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
