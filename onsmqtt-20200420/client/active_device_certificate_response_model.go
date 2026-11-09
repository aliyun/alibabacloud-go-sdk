// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveDeviceCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveDeviceCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveDeviceCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ActiveDeviceCertificateResponseBody) *ActiveDeviceCertificateResponse
	GetBody() *ActiveDeviceCertificateResponseBody
}

type ActiveDeviceCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveDeviceCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *ActiveDeviceCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveDeviceCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveDeviceCertificateResponse) GetBody() *ActiveDeviceCertificateResponseBody {
	return s.Body
}

func (s *ActiveDeviceCertificateResponse) SetHeaders(v map[string]*string) *ActiveDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *ActiveDeviceCertificateResponse) SetStatusCode(v int32) *ActiveDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveDeviceCertificateResponse) SetBody(v *ActiveDeviceCertificateResponseBody) *ActiveDeviceCertificateResponse {
	s.Body = v
	return s
}

func (s *ActiveDeviceCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
