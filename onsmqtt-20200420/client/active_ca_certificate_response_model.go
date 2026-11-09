// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ActiveCaCertificateResponseBody) *ActiveCaCertificateResponse
	GetBody() *ActiveCaCertificateResponseBody
}

type ActiveCaCertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *ActiveCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveCaCertificateResponse) GetBody() *ActiveCaCertificateResponseBody {
	return s.Body
}

func (s *ActiveCaCertificateResponse) SetHeaders(v map[string]*string) *ActiveCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *ActiveCaCertificateResponse) SetStatusCode(v int32) *ActiveCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveCaCertificateResponse) SetBody(v *ActiveCaCertificateResponseBody) *ActiveCaCertificateResponse {
	s.Body = v
	return s
}

func (s *ActiveCaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
