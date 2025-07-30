// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ListClientCertificateResponseBody) *ListClientCertificateResponse
	GetBody() *ListClientCertificateResponseBody
}

type ListClientCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientCertificateResponse) GetBody() *ListClientCertificateResponseBody {
	return s.Body
}

func (s *ListClientCertificateResponse) SetHeaders(v map[string]*string) *ListClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListClientCertificateResponse) SetStatusCode(v int32) *ListClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientCertificateResponse) SetBody(v *ListClientCertificateResponseBody) *ListClientCertificateResponse {
	s.Body = v
	return s
}

func (s *ListClientCertificateResponse) Validate() error {
	return dara.Validate(s)
}
