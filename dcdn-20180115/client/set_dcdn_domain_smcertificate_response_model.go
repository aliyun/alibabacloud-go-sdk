// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainSMCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDcdnDomainSMCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDcdnDomainSMCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetDcdnDomainSMCertificateResponseBody) *SetDcdnDomainSMCertificateResponse
	GetBody() *SetDcdnDomainSMCertificateResponseBody
}

type SetDcdnDomainSMCertificateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDcdnDomainSMCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDcdnDomainSMCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainSMCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainSMCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDcdnDomainSMCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDcdnDomainSMCertificateResponse) GetBody() *SetDcdnDomainSMCertificateResponseBody {
	return s.Body
}

func (s *SetDcdnDomainSMCertificateResponse) SetHeaders(v map[string]*string) *SetDcdnDomainSMCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnDomainSMCertificateResponse) SetStatusCode(v int32) *SetDcdnDomainSMCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDcdnDomainSMCertificateResponse) SetBody(v *SetDcdnDomainSMCertificateResponseBody) *SetDcdnDomainSMCertificateResponse {
	s.Body = v
	return s
}

func (s *SetDcdnDomainSMCertificateResponse) Validate() error {
	return dara.Validate(s)
}
