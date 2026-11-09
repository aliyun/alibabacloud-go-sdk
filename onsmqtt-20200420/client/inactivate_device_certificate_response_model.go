// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInactivateDeviceCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InactivateDeviceCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InactivateDeviceCertificateResponse
	GetStatusCode() *int32
	SetBody(v *InactivateDeviceCertificateResponseBody) *InactivateDeviceCertificateResponse
	GetBody() *InactivateDeviceCertificateResponseBody
}

type InactivateDeviceCertificateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InactivateDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InactivateDeviceCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s InactivateDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *InactivateDeviceCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InactivateDeviceCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InactivateDeviceCertificateResponse) GetBody() *InactivateDeviceCertificateResponseBody {
	return s.Body
}

func (s *InactivateDeviceCertificateResponse) SetHeaders(v map[string]*string) *InactivateDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *InactivateDeviceCertificateResponse) SetStatusCode(v int32) *InactivateDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *InactivateDeviceCertificateResponse) SetBody(v *InactivateDeviceCertificateResponseBody) *InactivateDeviceCertificateResponse {
	s.Body = v
	return s
}

func (s *InactivateDeviceCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
