// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVodDomainCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetVodDomainCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetVodDomainCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SetVodDomainCertificateResponseBody) *SetVodDomainCertificateResponse
	GetBody() *SetVodDomainCertificateResponseBody
}

type SetVodDomainCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetVodDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetVodDomainCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetVodDomainCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetVodDomainCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetVodDomainCertificateResponse) GetBody() *SetVodDomainCertificateResponseBody {
	return s.Body
}

func (s *SetVodDomainCertificateResponse) SetHeaders(v map[string]*string) *SetVodDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetVodDomainCertificateResponse) SetStatusCode(v int32) *SetVodDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVodDomainCertificateResponse) SetBody(v *SetVodDomainCertificateResponseBody) *SetVodDomainCertificateResponse {
	s.Body = v
	return s
}

func (s *SetVodDomainCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
