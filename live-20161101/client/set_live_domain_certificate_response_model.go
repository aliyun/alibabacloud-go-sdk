// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveDomainCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveDomainCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveDomainCertificateResponseBody) *SetLiveDomainCertificateResponse
	GetBody() *SetLiveDomainCertificateResponseBody
}

type SetLiveDomainCertificateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveDomainCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetLiveDomainCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveDomainCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveDomainCertificateResponse) GetBody() *SetLiveDomainCertificateResponseBody {
	return s.Body
}

func (s *SetLiveDomainCertificateResponse) SetHeaders(v map[string]*string) *SetLiveDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetLiveDomainCertificateResponse) SetStatusCode(v int32) *SetLiveDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveDomainCertificateResponse) SetBody(v *SetLiveDomainCertificateResponseBody) *SetLiveDomainCertificateResponse {
	s.Body = v
	return s
}

func (s *SetLiveDomainCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
