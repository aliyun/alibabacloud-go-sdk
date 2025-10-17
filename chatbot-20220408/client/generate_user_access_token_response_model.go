// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUserAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateUserAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateUserAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenerateUserAccessTokenResponseBody) *GenerateUserAccessTokenResponse
	GetBody() *GenerateUserAccessTokenResponseBody
}

type GenerateUserAccessTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUserAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUserAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateUserAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateUserAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateUserAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateUserAccessTokenResponse) GetBody() *GenerateUserAccessTokenResponseBody {
	return s.Body
}

func (s *GenerateUserAccessTokenResponse) SetHeaders(v map[string]*string) *GenerateUserAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateUserAccessTokenResponse) SetStatusCode(v int32) *GenerateUserAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUserAccessTokenResponse) SetBody(v *GenerateUserAccessTokenResponseBody) *GenerateUserAccessTokenResponse {
	s.Body = v
	return s
}

func (s *GenerateUserAccessTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
