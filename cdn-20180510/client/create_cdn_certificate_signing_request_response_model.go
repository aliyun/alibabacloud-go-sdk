// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdnCertificateSigningRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCdnCertificateSigningRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCdnCertificateSigningRequestResponse
	GetStatusCode() *int32
	SetBody(v *CreateCdnCertificateSigningRequestResponseBody) *CreateCdnCertificateSigningRequestResponse
	GetBody() *CreateCdnCertificateSigningRequestResponseBody
}

type CreateCdnCertificateSigningRequestResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCdnCertificateSigningRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCdnCertificateSigningRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCdnCertificateSigningRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCdnCertificateSigningRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCdnCertificateSigningRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCdnCertificateSigningRequestResponse) GetBody() *CreateCdnCertificateSigningRequestResponseBody {
	return s.Body
}

func (s *CreateCdnCertificateSigningRequestResponse) SetHeaders(v map[string]*string) *CreateCdnCertificateSigningRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponse) SetStatusCode(v int32) *CreateCdnCertificateSigningRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponse) SetBody(v *CreateCdnCertificateSigningRequestResponseBody) *CreateCdnCertificateSigningRequestResponse {
	s.Body = v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
