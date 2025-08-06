// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVodDomainSSLCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetVodDomainSSLCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetVodDomainSSLCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetVodDomainSSLCertificateResponseBody) *SetVodDomainSSLCertificateResponse
	GetBody() *SetVodDomainSSLCertificateResponseBody
}

type SetVodDomainSSLCertificateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetVodDomainSSLCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetVodDomainSSLCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainSSLCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetVodDomainSSLCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetVodDomainSSLCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetVodDomainSSLCertificateResponse) GetBody() *SetVodDomainSSLCertificateResponseBody {
	return s.Body
}

func (s *SetVodDomainSSLCertificateResponse) SetHeaders(v map[string]*string) *SetVodDomainSSLCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetVodDomainSSLCertificateResponse) SetStatusCode(v int32) *SetVodDomainSSLCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVodDomainSSLCertificateResponse) SetBody(v *SetVodDomainSSLCertificateResponseBody) *SetVodDomainSSLCertificateResponse {
	s.Body = v
	return s
}

func (s *SetVodDomainSSLCertificateResponse) Validate() error {
	return dara.Validate(s)
}
