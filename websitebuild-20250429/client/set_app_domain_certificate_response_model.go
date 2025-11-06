// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppDomainCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAppDomainCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAppDomainCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetAppDomainCertificateResponseBody) *SetAppDomainCertificateResponse
	GetBody() *SetAppDomainCertificateResponseBody
}

type SetAppDomainCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAppDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAppDomainCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAppDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetAppDomainCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAppDomainCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAppDomainCertificateResponse) GetBody() *SetAppDomainCertificateResponseBody {
	return s.Body
}

func (s *SetAppDomainCertificateResponse) SetHeaders(v map[string]*string) *SetAppDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetAppDomainCertificateResponse) SetStatusCode(v int32) *SetAppDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAppDomainCertificateResponse) SetBody(v *SetAppDomainCertificateResponseBody) *SetAppDomainCertificateResponse {
	s.Body = v
	return s
}

func (s *SetAppDomainCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
