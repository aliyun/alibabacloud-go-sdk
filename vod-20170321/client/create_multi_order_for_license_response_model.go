// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiOrderForLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultiOrderForLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultiOrderForLicenseResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultiOrderForLicenseResponseBody) *CreateMultiOrderForLicenseResponse
	GetBody() *CreateMultiOrderForLicenseResponseBody
}

type CreateMultiOrderForLicenseResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultiOrderForLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultiOrderForLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderForLicenseResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderForLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultiOrderForLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultiOrderForLicenseResponse) GetBody() *CreateMultiOrderForLicenseResponseBody {
	return s.Body
}

func (s *CreateMultiOrderForLicenseResponse) SetHeaders(v map[string]*string) *CreateMultiOrderForLicenseResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiOrderForLicenseResponse) SetStatusCode(v int32) *CreateMultiOrderForLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultiOrderForLicenseResponse) SetBody(v *CreateMultiOrderForLicenseResponseBody) *CreateMultiOrderForLicenseResponse {
	s.Body = v
	return s
}

func (s *CreateMultiOrderForLicenseResponse) Validate() error {
	return dara.Validate(s)
}
