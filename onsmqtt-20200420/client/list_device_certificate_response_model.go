// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceCertificateResponseBody) *ListDeviceCertificateResponse
	GetBody() *ListDeviceCertificateResponseBody
}

type ListDeviceCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceCertificateResponse) GetBody() *ListDeviceCertificateResponseBody {
	return s.Body
}

func (s *ListDeviceCertificateResponse) SetHeaders(v map[string]*string) *ListDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceCertificateResponse) SetStatusCode(v int32) *ListDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceCertificateResponse) SetBody(v *ListDeviceCertificateResponseBody) *ListDeviceCertificateResponse {
	s.Body = v
	return s
}

func (s *ListDeviceCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
