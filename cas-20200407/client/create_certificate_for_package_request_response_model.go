// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateForPackageRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCertificateForPackageRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCertificateForPackageRequestResponse
	GetStatusCode() *int32
	SetBody(v *CreateCertificateForPackageRequestResponseBody) *CreateCertificateForPackageRequestResponse
	GetBody() *CreateCertificateForPackageRequestResponseBody
}

type CreateCertificateForPackageRequestResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateForPackageRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCertificateForPackageRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateForPackageRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateForPackageRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCertificateForPackageRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCertificateForPackageRequestResponse) GetBody() *CreateCertificateForPackageRequestResponseBody {
	return s.Body
}

func (s *CreateCertificateForPackageRequestResponse) SetHeaders(v map[string]*string) *CreateCertificateForPackageRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateForPackageRequestResponse) SetStatusCode(v int32) *CreateCertificateForPackageRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateForPackageRequestResponse) SetBody(v *CreateCertificateForPackageRequestResponseBody) *CreateCertificateForPackageRequestResponse {
	s.Body = v
	return s
}

func (s *CreateCertificateForPackageRequestResponse) Validate() error {
	return dara.Validate(s)
}
