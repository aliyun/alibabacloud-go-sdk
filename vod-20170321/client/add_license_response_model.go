// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLicenseResponse
	GetStatusCode() *int32
	SetBody(v *AddLicenseResponseBody) *AddLicenseResponse
	GetBody() *AddLicenseResponseBody
}

type AddLicenseResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLicenseResponse) GoString() string {
	return s.String()
}

func (s *AddLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLicenseResponse) GetBody() *AddLicenseResponseBody {
	return s.Body
}

func (s *AddLicenseResponse) SetHeaders(v map[string]*string) *AddLicenseResponse {
	s.Headers = v
	return s
}

func (s *AddLicenseResponse) SetStatusCode(v int32) *AddLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLicenseResponse) SetBody(v *AddLicenseResponseBody) *AddLicenseResponse {
	s.Body = v
	return s
}

func (s *AddLicenseResponse) Validate() error {
	return dara.Validate(s)
}
