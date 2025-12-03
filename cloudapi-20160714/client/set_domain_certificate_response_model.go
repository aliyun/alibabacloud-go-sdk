// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDomainCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDomainCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetDomainCertificateResponseBody) *SetDomainCertificateResponse
	GetBody() *SetDomainCertificateResponseBody
}

type SetDomainCertificateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDomainCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetDomainCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDomainCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDomainCertificateResponse) GetBody() *SetDomainCertificateResponseBody {
	return s.Body
}

func (s *SetDomainCertificateResponse) SetHeaders(v map[string]*string) *SetDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetDomainCertificateResponse) SetStatusCode(v int32) *SetDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainCertificateResponse) SetBody(v *SetDomainCertificateResponseBody) *SetDomainCertificateResponse {
	s.Body = v
	return s
}

func (s *SetDomainCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
