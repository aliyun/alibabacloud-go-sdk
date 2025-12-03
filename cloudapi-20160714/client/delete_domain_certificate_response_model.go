// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDomainCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDomainCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDomainCertificateResponseBody) *DeleteDomainCertificateResponse
	GetBody() *DeleteDomainCertificateResponseBody
}

type DeleteDomainCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDomainCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDomainCertificateResponse) GetBody() *DeleteDomainCertificateResponseBody {
	return s.Body
}

func (s *DeleteDomainCertificateResponse) SetHeaders(v map[string]*string) *DeleteDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainCertificateResponse) SetStatusCode(v int32) *DeleteDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainCertificateResponse) SetBody(v *DeleteDomainCertificateResponseBody) *DeleteDomainCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteDomainCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
