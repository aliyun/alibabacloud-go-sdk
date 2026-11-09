// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *RegisterCaCertificateResponseBody) *RegisterCaCertificateResponse
	GetBody() *RegisterCaCertificateResponseBody
}

type RegisterCaCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *RegisterCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterCaCertificateResponse) GetBody() *RegisterCaCertificateResponseBody {
	return s.Body
}

func (s *RegisterCaCertificateResponse) SetHeaders(v map[string]*string) *RegisterCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *RegisterCaCertificateResponse) SetStatusCode(v int32) *RegisterCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterCaCertificateResponse) SetBody(v *RegisterCaCertificateResponseBody) *RegisterCaCertificateResponse {
	s.Body = v
	return s
}

func (s *RegisterCaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
