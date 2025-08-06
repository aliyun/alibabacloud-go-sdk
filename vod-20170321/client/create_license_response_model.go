// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLicenseResponse
	GetStatusCode() *int32
	SetBody(v *CreateLicenseResponseBody) *CreateLicenseResponse
	GetBody() *CreateLicenseResponseBody
}

type CreateLicenseResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLicenseResponse) GoString() string {
	return s.String()
}

func (s *CreateLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLicenseResponse) GetBody() *CreateLicenseResponseBody {
	return s.Body
}

func (s *CreateLicenseResponse) SetHeaders(v map[string]*string) *CreateLicenseResponse {
	s.Headers = v
	return s
}

func (s *CreateLicenseResponse) SetStatusCode(v int32) *CreateLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLicenseResponse) SetBody(v *CreateLicenseResponseBody) *CreateLicenseResponse {
	s.Body = v
	return s
}

func (s *CreateLicenseResponse) Validate() error {
	return dara.Validate(s)
}
