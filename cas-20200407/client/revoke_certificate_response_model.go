// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeCertificateResponse
	GetStatusCode() *int32
	SetBody(v *RevokeCertificateResponseBody) *RevokeCertificateResponse
	GetBody() *RevokeCertificateResponseBody
}

type RevokeCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeCertificateResponse) GoString() string {
	return s.String()
}

func (s *RevokeCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeCertificateResponse) GetBody() *RevokeCertificateResponseBody {
	return s.Body
}

func (s *RevokeCertificateResponse) SetHeaders(v map[string]*string) *RevokeCertificateResponse {
	s.Headers = v
	return s
}

func (s *RevokeCertificateResponse) SetStatusCode(v int32) *RevokeCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeCertificateResponse) SetBody(v *RevokeCertificateResponseBody) *RevokeCertificateResponse {
	s.Body = v
	return s
}

func (s *RevokeCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
