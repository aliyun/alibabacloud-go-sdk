// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ListCertificateResponseBody) *ListCertificateResponse
	GetBody() *ListCertificateResponseBody
}

type ListCertificateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCertificateResponse) GetBody() *ListCertificateResponseBody {
	return s.Body
}

func (s *ListCertificateResponse) SetHeaders(v map[string]*string) *ListCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListCertificateResponse) SetStatusCode(v int32) *ListCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCertificateResponse) SetBody(v *ListCertificateResponseBody) *ListCertificateResponse {
	s.Body = v
	return s
}

func (s *ListCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
