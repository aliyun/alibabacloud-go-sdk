// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebofficeTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateWebofficeTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateWebofficeTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenerateWebofficeTokenResponseBody) *GenerateWebofficeTokenResponse
	GetBody() *GenerateWebofficeTokenResponseBody
}

type GenerateWebofficeTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateWebofficeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateWebofficeTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebofficeTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateWebofficeTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateWebofficeTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateWebofficeTokenResponse) GetBody() *GenerateWebofficeTokenResponseBody {
	return s.Body
}

func (s *GenerateWebofficeTokenResponse) SetHeaders(v map[string]*string) *GenerateWebofficeTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateWebofficeTokenResponse) SetStatusCode(v int32) *GenerateWebofficeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateWebofficeTokenResponse) SetBody(v *GenerateWebofficeTokenResponseBody) *GenerateWebofficeTokenResponse {
	s.Body = v
	return s
}

func (s *GenerateWebofficeTokenResponse) Validate() error {
	return dara.Validate(s)
}
