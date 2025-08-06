// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckLicenseResponse
	GetStatusCode() *int32
	SetBody(v *CheckLicenseResponseBody) *CheckLicenseResponse
	GetBody() *CheckLicenseResponseBody
}

type CheckLicenseResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckLicenseResponse) GoString() string {
	return s.String()
}

func (s *CheckLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckLicenseResponse) GetBody() *CheckLicenseResponseBody {
	return s.Body
}

func (s *CheckLicenseResponse) SetHeaders(v map[string]*string) *CheckLicenseResponse {
	s.Headers = v
	return s
}

func (s *CheckLicenseResponse) SetStatusCode(v int32) *CheckLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckLicenseResponse) SetBody(v *CheckLicenseResponseBody) *CheckLicenseResponse {
	s.Body = v
	return s
}

func (s *CheckLicenseResponse) Validate() error {
	return dara.Validate(s)
}
