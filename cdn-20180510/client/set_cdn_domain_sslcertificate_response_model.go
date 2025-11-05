// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainSSLCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCdnDomainSSLCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCdnDomainSSLCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetCdnDomainSSLCertificateResponseBody) *SetCdnDomainSSLCertificateResponse
	GetBody() *SetCdnDomainSSLCertificateResponseBody
}

type SetCdnDomainSSLCertificateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCdnDomainSSLCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCdnDomainSSLCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainSSLCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetCdnDomainSSLCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCdnDomainSSLCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCdnDomainSSLCertificateResponse) GetBody() *SetCdnDomainSSLCertificateResponseBody {
	return s.Body
}

func (s *SetCdnDomainSSLCertificateResponse) SetHeaders(v map[string]*string) *SetCdnDomainSSLCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetCdnDomainSSLCertificateResponse) SetStatusCode(v int32) *SetCdnDomainSSLCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCdnDomainSSLCertificateResponse) SetBody(v *SetCdnDomainSSLCertificateResponseBody) *SetCdnDomainSSLCertificateResponse {
	s.Body = v
	return s
}

func (s *SetCdnDomainSSLCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
