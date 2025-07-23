// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateWithCsrRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCertificateWithCsrRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCertificateWithCsrRequestResponse
	GetStatusCode() *int32
	SetBody(v *CreateCertificateWithCsrRequestResponseBody) *CreateCertificateWithCsrRequestResponse
	GetBody() *CreateCertificateWithCsrRequestResponseBody
}

type CreateCertificateWithCsrRequestResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateWithCsrRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCertificateWithCsrRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateWithCsrRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithCsrRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCertificateWithCsrRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCertificateWithCsrRequestResponse) GetBody() *CreateCertificateWithCsrRequestResponseBody {
	return s.Body
}

func (s *CreateCertificateWithCsrRequestResponse) SetHeaders(v map[string]*string) *CreateCertificateWithCsrRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateWithCsrRequestResponse) SetStatusCode(v int32) *CreateCertificateWithCsrRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateWithCsrRequestResponse) SetBody(v *CreateCertificateWithCsrRequestResponseBody) *CreateCertificateWithCsrRequestResponse {
	s.Body = v
	return s
}

func (s *CreateCertificateWithCsrRequestResponse) Validate() error {
	return dara.Validate(s)
}
