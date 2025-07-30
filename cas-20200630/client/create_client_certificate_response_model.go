// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *CreateClientCertificateResponseBody) *CreateClientCertificateResponse
	GetBody() *CreateClientCertificateResponseBody
}

type CreateClientCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClientCertificateResponse) GetBody() *CreateClientCertificateResponseBody {
	return s.Body
}

func (s *CreateClientCertificateResponse) SetHeaders(v map[string]*string) *CreateClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateClientCertificateResponse) SetStatusCode(v int32) *CreateClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientCertificateResponse) SetBody(v *CreateClientCertificateResponseBody) *CreateClientCertificateResponse {
	s.Body = v
	return s
}

func (s *CreateClientCertificateResponse) Validate() error {
	return dara.Validate(s)
}
