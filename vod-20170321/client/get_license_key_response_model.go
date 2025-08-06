// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicenseKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLicenseKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLicenseKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetLicenseKeyResponseBody) *GetLicenseKeyResponse
	GetBody() *GetLicenseKeyResponseBody
}

type GetLicenseKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLicenseKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLicenseKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseKeyResponse) GoString() string {
	return s.String()
}

func (s *GetLicenseKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLicenseKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLicenseKeyResponse) GetBody() *GetLicenseKeyResponseBody {
	return s.Body
}

func (s *GetLicenseKeyResponse) SetHeaders(v map[string]*string) *GetLicenseKeyResponse {
	s.Headers = v
	return s
}

func (s *GetLicenseKeyResponse) SetStatusCode(v int32) *GetLicenseKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLicenseKeyResponse) SetBody(v *GetLicenseKeyResponseBody) *GetLicenseKeyResponse {
	s.Body = v
	return s
}

func (s *GetLicenseKeyResponse) Validate() error {
	return dara.Validate(s)
}
