// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewAppLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewAppLicenseResponse
	GetStatusCode() *int32
	SetBody(v *RenewAppLicenseResponseBody) *RenewAppLicenseResponse
	GetBody() *RenewAppLicenseResponseBody
}

type RenewAppLicenseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAppLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAppLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewAppLicenseResponse) GoString() string {
	return s.String()
}

func (s *RenewAppLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewAppLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewAppLicenseResponse) GetBody() *RenewAppLicenseResponseBody {
	return s.Body
}

func (s *RenewAppLicenseResponse) SetHeaders(v map[string]*string) *RenewAppLicenseResponse {
	s.Headers = v
	return s
}

func (s *RenewAppLicenseResponse) SetStatusCode(v int32) *RenewAppLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAppLicenseResponse) SetBody(v *RenewAppLicenseResponseBody) *RenewAppLicenseResponse {
	s.Body = v
	return s
}

func (s *RenewAppLicenseResponse) Validate() error {
	return dara.Validate(s)
}
