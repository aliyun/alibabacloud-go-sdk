// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainSMCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCdnDomainSMCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCdnDomainSMCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetCdnDomainSMCertificateResponseBody) *SetCdnDomainSMCertificateResponse
	GetBody() *SetCdnDomainSMCertificateResponseBody
}

type SetCdnDomainSMCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCdnDomainSMCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCdnDomainSMCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainSMCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetCdnDomainSMCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCdnDomainSMCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCdnDomainSMCertificateResponse) GetBody() *SetCdnDomainSMCertificateResponseBody {
	return s.Body
}

func (s *SetCdnDomainSMCertificateResponse) SetHeaders(v map[string]*string) *SetCdnDomainSMCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetCdnDomainSMCertificateResponse) SetStatusCode(v int32) *SetCdnDomainSMCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCdnDomainSMCertificateResponse) SetBody(v *SetCdnDomainSMCertificateResponseBody) *SetCdnDomainSMCertificateResponse {
	s.Body = v
	return s
}

func (s *SetCdnDomainSMCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
