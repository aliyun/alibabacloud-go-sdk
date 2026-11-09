// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeviceCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeviceCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeviceCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeviceCertificateResponseBody) *DeleteDeviceCertificateResponse
	GetBody() *DeleteDeviceCertificateResponseBody
}

type DeleteDeviceCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeviceCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeviceCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeviceCertificateResponse) GetBody() *DeleteDeviceCertificateResponseBody {
	return s.Body
}

func (s *DeleteDeviceCertificateResponse) SetHeaders(v map[string]*string) *DeleteDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceCertificateResponse) SetStatusCode(v int32) *DeleteDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeviceCertificateResponse) SetBody(v *DeleteDeviceCertificateResponseBody) *DeleteDeviceCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteDeviceCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
