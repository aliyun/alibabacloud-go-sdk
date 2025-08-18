// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *RevokeClientCertificateResponseBody) *RevokeClientCertificateResponse
	GetBody() *RevokeClientCertificateResponseBody
}

type RevokeClientCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *RevokeClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeClientCertificateResponse) GetBody() *RevokeClientCertificateResponseBody {
	return s.Body
}

func (s *RevokeClientCertificateResponse) SetHeaders(v map[string]*string) *RevokeClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *RevokeClientCertificateResponse) SetStatusCode(v int32) *RevokeClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeClientCertificateResponse) SetBody(v *RevokeClientCertificateResponseBody) *RevokeClientCertificateResponse {
	s.Body = v
	return s
}

func (s *RevokeClientCertificateResponse) Validate() error {
	return dara.Validate(s)
}
