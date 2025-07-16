// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainCSRCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCdnDomainCSRCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCdnDomainCSRCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetCdnDomainCSRCertificateResponseBody) *SetCdnDomainCSRCertificateResponse
	GetBody() *SetCdnDomainCSRCertificateResponseBody
}

type SetCdnDomainCSRCertificateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCdnDomainCSRCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCdnDomainCSRCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainCSRCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetCdnDomainCSRCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCdnDomainCSRCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCdnDomainCSRCertificateResponse) GetBody() *SetCdnDomainCSRCertificateResponseBody {
	return s.Body
}

func (s *SetCdnDomainCSRCertificateResponse) SetHeaders(v map[string]*string) *SetCdnDomainCSRCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetCdnDomainCSRCertificateResponse) SetStatusCode(v int32) *SetCdnDomainCSRCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCdnDomainCSRCertificateResponse) SetBody(v *SetCdnDomainCSRCertificateResponseBody) *SetCdnDomainCSRCertificateResponse {
	s.Body = v
	return s
}

func (s *SetCdnDomainCSRCertificateResponse) Validate() error {
	return dara.Validate(s)
}
