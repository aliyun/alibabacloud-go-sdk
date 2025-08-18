// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCertificatesResponse
	GetStatusCode() *int32
	SetBody(v *ListCertificatesResponseBody) *ListCertificatesResponse
	GetBody() *ListCertificatesResponseBody
}

type ListCertificatesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCertificatesResponse) GetBody() *ListCertificatesResponseBody {
	return s.Body
}

func (s *ListCertificatesResponse) SetHeaders(v map[string]*string) *ListCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListCertificatesResponse) SetStatusCode(v int32) *ListCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCertificatesResponse) SetBody(v *ListCertificatesResponseBody) *ListCertificatesResponse {
	s.Body = v
	return s
}

func (s *ListCertificatesResponse) Validate() error {
	return dara.Validate(s)
}
