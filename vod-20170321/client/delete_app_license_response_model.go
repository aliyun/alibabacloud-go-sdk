// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppLicenseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppLicenseResponseBody) *DeleteAppLicenseResponse
	GetBody() *DeleteAppLicenseResponseBody
}

type DeleteAppLicenseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppLicenseResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppLicenseResponse) GetBody() *DeleteAppLicenseResponseBody {
	return s.Body
}

func (s *DeleteAppLicenseResponse) SetHeaders(v map[string]*string) *DeleteAppLicenseResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppLicenseResponse) SetStatusCode(v int32) *DeleteAppLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppLicenseResponse) SetBody(v *DeleteAppLicenseResponseBody) *DeleteAppLicenseResponse {
	s.Body = v
	return s
}

func (s *DeleteAppLicenseResponse) Validate() error {
	return dara.Validate(s)
}
