// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientCertificateWithCsrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClientCertificateWithCsrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClientCertificateWithCsrResponse
	GetStatusCode() *int32
	SetBody(v *CreateClientCertificateWithCsrResponseBody) *CreateClientCertificateWithCsrResponse
	GetBody() *CreateClientCertificateWithCsrResponseBody
}

type CreateClientCertificateWithCsrResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientCertificateWithCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientCertificateWithCsrResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateWithCsrResponse) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateWithCsrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClientCertificateWithCsrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClientCertificateWithCsrResponse) GetBody() *CreateClientCertificateWithCsrResponseBody {
	return s.Body
}

func (s *CreateClientCertificateWithCsrResponse) SetHeaders(v map[string]*string) *CreateClientCertificateWithCsrResponse {
	s.Headers = v
	return s
}

func (s *CreateClientCertificateWithCsrResponse) SetStatusCode(v int32) *CreateClientCertificateWithCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponse) SetBody(v *CreateClientCertificateWithCsrResponseBody) *CreateClientCertificateWithCsrResponse {
	s.Body = v
	return s
}

func (s *CreateClientCertificateWithCsrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
