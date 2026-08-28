// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ShareCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ShareCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ShareCertificateResponseBody) *ShareCertificateResponse
	GetBody() *ShareCertificateResponseBody
}

type ShareCertificateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShareCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ShareCertificateResponse) GoString() string {
	return s.String()
}

func (s *ShareCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ShareCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ShareCertificateResponse) GetBody() *ShareCertificateResponseBody {
	return s.Body
}

func (s *ShareCertificateResponse) SetHeaders(v map[string]*string) *ShareCertificateResponse {
	s.Headers = v
	return s
}

func (s *ShareCertificateResponse) SetStatusCode(v int32) *ShareCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ShareCertificateResponse) SetBody(v *ShareCertificateResponseBody) *ShareCertificateResponse {
	s.Body = v
	return s
}

func (s *ShareCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
