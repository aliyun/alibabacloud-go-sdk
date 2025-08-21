// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVsDomainCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetVsDomainCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetVsDomainCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetVsDomainCertificateResponseBody) *SetVsDomainCertificateResponse
	GetBody() *SetVsDomainCertificateResponseBody
}

type SetVsDomainCertificateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetVsDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetVsDomainCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetVsDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetVsDomainCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetVsDomainCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetVsDomainCertificateResponse) GetBody() *SetVsDomainCertificateResponseBody {
	return s.Body
}

func (s *SetVsDomainCertificateResponse) SetHeaders(v map[string]*string) *SetVsDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetVsDomainCertificateResponse) SetStatusCode(v int32) *SetVsDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVsDomainCertificateResponse) SetBody(v *SetVsDomainCertificateResponseBody) *SetVsDomainCertificateResponse {
	s.Body = v
	return s
}

func (s *SetVsDomainCertificateResponse) Validate() error {
	return dara.Validate(s)
}
