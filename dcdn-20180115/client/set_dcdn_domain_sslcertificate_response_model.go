// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainSSLCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDcdnDomainSSLCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDcdnDomainSSLCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetDcdnDomainSSLCertificateResponseBody) *SetDcdnDomainSSLCertificateResponse
	GetBody() *SetDcdnDomainSSLCertificateResponseBody
}

type SetDcdnDomainSSLCertificateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDcdnDomainSSLCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDcdnDomainSSLCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainSSLCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainSSLCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDcdnDomainSSLCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDcdnDomainSSLCertificateResponse) GetBody() *SetDcdnDomainSSLCertificateResponseBody {
	return s.Body
}

func (s *SetDcdnDomainSSLCertificateResponse) SetHeaders(v map[string]*string) *SetDcdnDomainSSLCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnDomainSSLCertificateResponse) SetStatusCode(v int32) *SetDcdnDomainSSLCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateResponse) SetBody(v *SetDcdnDomainSSLCertificateResponseBody) *SetDcdnDomainSSLCertificateResponse {
	s.Body = v
	return s
}

func (s *SetDcdnDomainSSLCertificateResponse) Validate() error {
	return dara.Validate(s)
}
