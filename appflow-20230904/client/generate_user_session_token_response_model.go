// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUserSessionTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateUserSessionTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateUserSessionTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenerateUserSessionTokenResponseBody) *GenerateUserSessionTokenResponse
	GetBody() *GenerateUserSessionTokenResponseBody
}

type GenerateUserSessionTokenResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUserSessionTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUserSessionTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateUserSessionTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateUserSessionTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateUserSessionTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateUserSessionTokenResponse) GetBody() *GenerateUserSessionTokenResponseBody {
	return s.Body
}

func (s *GenerateUserSessionTokenResponse) SetHeaders(v map[string]*string) *GenerateUserSessionTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateUserSessionTokenResponse) SetStatusCode(v int32) *GenerateUserSessionTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUserSessionTokenResponse) SetBody(v *GenerateUserSessionTokenResponseBody) *GenerateUserSessionTokenResponse {
	s.Body = v
	return s
}

func (s *GenerateUserSessionTokenResponse) Validate() error {
	return dara.Validate(s)
}
