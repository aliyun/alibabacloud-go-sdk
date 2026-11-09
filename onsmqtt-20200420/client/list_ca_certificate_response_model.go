// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ListCaCertificateResponseBody) *ListCaCertificateResponse
	GetBody() *ListCaCertificateResponseBody
}

type ListCaCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCaCertificateResponse) GetBody() *ListCaCertificateResponseBody {
	return s.Body
}

func (s *ListCaCertificateResponse) SetHeaders(v map[string]*string) *ListCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListCaCertificateResponse) SetStatusCode(v int32) *ListCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCaCertificateResponse) SetBody(v *ListCaCertificateResponseBody) *ListCaCertificateResponse {
	s.Body = v
	return s
}

func (s *ListCaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
