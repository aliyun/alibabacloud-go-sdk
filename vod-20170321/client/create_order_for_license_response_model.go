// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrderForLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrderForLicenseResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrderForLicenseResponseBody) *CreateOrderForLicenseResponse
	GetBody() *CreateOrderForLicenseResponseBody
}

type CreateOrderForLicenseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrderForLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrderForLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForLicenseResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderForLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrderForLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrderForLicenseResponse) GetBody() *CreateOrderForLicenseResponseBody {
	return s.Body
}

func (s *CreateOrderForLicenseResponse) SetHeaders(v map[string]*string) *CreateOrderForLicenseResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderForLicenseResponse) SetStatusCode(v int32) *CreateOrderForLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderForLicenseResponse) SetBody(v *CreateOrderForLicenseResponseBody) *CreateOrderForLicenseResponse {
	s.Body = v
	return s
}

func (s *CreateOrderForLicenseResponse) Validate() error {
	return dara.Validate(s)
}
