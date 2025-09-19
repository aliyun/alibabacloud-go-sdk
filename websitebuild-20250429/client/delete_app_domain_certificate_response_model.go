// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppDomainCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppDomainCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppDomainCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppDomainCertificateResponseBody) *DeleteAppDomainCertificateResponse
	GetBody() *DeleteAppDomainCertificateResponseBody
}

type DeleteAppDomainCertificateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppDomainCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppDomainCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppDomainCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppDomainCertificateResponse) GetBody() *DeleteAppDomainCertificateResponseBody {
	return s.Body
}

func (s *DeleteAppDomainCertificateResponse) SetHeaders(v map[string]*string) *DeleteAppDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppDomainCertificateResponse) SetStatusCode(v int32) *DeleteAppDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppDomainCertificateResponse) SetBody(v *DeleteAppDomainCertificateResponseBody) *DeleteAppDomainCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteAppDomainCertificateResponse) Validate() error {
	return dara.Validate(s)
}
