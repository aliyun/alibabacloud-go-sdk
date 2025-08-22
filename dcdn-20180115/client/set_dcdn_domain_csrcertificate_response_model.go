// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainCSRCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDcdnDomainCSRCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDcdnDomainCSRCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetDcdnDomainCSRCertificateResponseBody) *SetDcdnDomainCSRCertificateResponse
	GetBody() *SetDcdnDomainCSRCertificateResponseBody
}

type SetDcdnDomainCSRCertificateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDcdnDomainCSRCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDcdnDomainCSRCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainCSRCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainCSRCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDcdnDomainCSRCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDcdnDomainCSRCertificateResponse) GetBody() *SetDcdnDomainCSRCertificateResponseBody {
	return s.Body
}

func (s *SetDcdnDomainCSRCertificateResponse) SetHeaders(v map[string]*string) *SetDcdnDomainCSRCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnDomainCSRCertificateResponse) SetStatusCode(v int32) *SetDcdnDomainCSRCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDcdnDomainCSRCertificateResponse) SetBody(v *SetDcdnDomainCSRCertificateResponseBody) *SetDcdnDomainCSRCertificateResponse {
	s.Body = v
	return s
}

func (s *SetDcdnDomainCSRCertificateResponse) Validate() error {
	return dara.Validate(s)
}
