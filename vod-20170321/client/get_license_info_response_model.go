// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicenseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLicenseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLicenseInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetLicenseInfoResponseBody) *GetLicenseInfoResponse
	GetBody() *GetLicenseInfoResponseBody
}

type GetLicenseInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLicenseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLicenseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLicenseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLicenseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLicenseInfoResponse) GetBody() *GetLicenseInfoResponseBody {
	return s.Body
}

func (s *GetLicenseInfoResponse) SetHeaders(v map[string]*string) *GetLicenseInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLicenseInfoResponse) SetStatusCode(v int32) *GetLicenseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLicenseInfoResponse) SetBody(v *GetLicenseInfoResponseBody) *GetLicenseInfoResponse {
	s.Body = v
	return s
}

func (s *GetLicenseInfoResponse) Validate() error {
	return dara.Validate(s)
}
