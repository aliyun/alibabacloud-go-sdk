// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCertificateForPackageRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCertificateForPackageRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCertificateForPackageRequestResponse
	GetStatusCode() *int32
	SetBody(v *CancelCertificateForPackageRequestResponseBody) *CancelCertificateForPackageRequestResponse
	GetBody() *CancelCertificateForPackageRequestResponseBody
}

type CancelCertificateForPackageRequestResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCertificateForPackageRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCertificateForPackageRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCertificateForPackageRequestResponse) GoString() string {
	return s.String()
}

func (s *CancelCertificateForPackageRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCertificateForPackageRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCertificateForPackageRequestResponse) GetBody() *CancelCertificateForPackageRequestResponseBody {
	return s.Body
}

func (s *CancelCertificateForPackageRequestResponse) SetHeaders(v map[string]*string) *CancelCertificateForPackageRequestResponse {
	s.Headers = v
	return s
}

func (s *CancelCertificateForPackageRequestResponse) SetStatusCode(v int32) *CancelCertificateForPackageRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCertificateForPackageRequestResponse) SetBody(v *CancelCertificateForPackageRequestResponseBody) *CancelCertificateForPackageRequestResponse {
	s.Body = v
	return s
}

func (s *CancelCertificateForPackageRequestResponse) Validate() error {
	return dara.Validate(s)
}
