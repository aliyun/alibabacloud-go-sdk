// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificatePackageCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCertificatePackageCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCertificatePackageCountResponse
	GetStatusCode() *int32
	SetBody(v *GetCertificatePackageCountResponseBody) *GetCertificatePackageCountResponse
	GetBody() *GetCertificatePackageCountResponseBody
}

type GetCertificatePackageCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCertificatePackageCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCertificatePackageCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCertificatePackageCountResponse) GoString() string {
	return s.String()
}

func (s *GetCertificatePackageCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCertificatePackageCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCertificatePackageCountResponse) GetBody() *GetCertificatePackageCountResponseBody {
	return s.Body
}

func (s *GetCertificatePackageCountResponse) SetHeaders(v map[string]*string) *GetCertificatePackageCountResponse {
	s.Headers = v
	return s
}

func (s *GetCertificatePackageCountResponse) SetStatusCode(v int32) *GetCertificatePackageCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCertificatePackageCountResponse) SetBody(v *GetCertificatePackageCountResponseBody) *GetCertificatePackageCountResponse {
	s.Body = v
	return s
}

func (s *GetCertificatePackageCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
