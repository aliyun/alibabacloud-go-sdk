// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCertificateOrderForPackageRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewCertificateOrderForPackageRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewCertificateOrderForPackageRequestResponse
	GetStatusCode() *int32
	SetBody(v *RenewCertificateOrderForPackageRequestResponseBody) *RenewCertificateOrderForPackageRequestResponse
	GetBody() *RenewCertificateOrderForPackageRequestResponseBody
}

type RenewCertificateOrderForPackageRequestResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewCertificateOrderForPackageRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewCertificateOrderForPackageRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewCertificateOrderForPackageRequestResponse) GoString() string {
	return s.String()
}

func (s *RenewCertificateOrderForPackageRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewCertificateOrderForPackageRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewCertificateOrderForPackageRequestResponse) GetBody() *RenewCertificateOrderForPackageRequestResponseBody {
	return s.Body
}

func (s *RenewCertificateOrderForPackageRequestResponse) SetHeaders(v map[string]*string) *RenewCertificateOrderForPackageRequestResponse {
	s.Headers = v
	return s
}

func (s *RenewCertificateOrderForPackageRequestResponse) SetStatusCode(v int32) *RenewCertificateOrderForPackageRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestResponse) SetBody(v *RenewCertificateOrderForPackageRequestResponseBody) *RenewCertificateOrderForPackageRequestResponse {
	s.Body = v
	return s
}

func (s *RenewCertificateOrderForPackageRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
