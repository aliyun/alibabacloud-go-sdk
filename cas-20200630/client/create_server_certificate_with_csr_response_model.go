// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerCertificateWithCsrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServerCertificateWithCsrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServerCertificateWithCsrResponse
	GetStatusCode() *int32
	SetBody(v *CreateServerCertificateWithCsrResponseBody) *CreateServerCertificateWithCsrResponse
	GetBody() *CreateServerCertificateWithCsrResponseBody
}

type CreateServerCertificateWithCsrResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServerCertificateWithCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServerCertificateWithCsrResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServerCertificateWithCsrResponse) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateWithCsrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServerCertificateWithCsrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServerCertificateWithCsrResponse) GetBody() *CreateServerCertificateWithCsrResponseBody {
	return s.Body
}

func (s *CreateServerCertificateWithCsrResponse) SetHeaders(v map[string]*string) *CreateServerCertificateWithCsrResponse {
	s.Headers = v
	return s
}

func (s *CreateServerCertificateWithCsrResponse) SetStatusCode(v int32) *CreateServerCertificateWithCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponse) SetBody(v *CreateServerCertificateWithCsrResponseBody) *CreateServerCertificateWithCsrResponse {
	s.Body = v
	return s
}

func (s *CreateServerCertificateWithCsrResponse) Validate() error {
	return dara.Validate(s)
}
