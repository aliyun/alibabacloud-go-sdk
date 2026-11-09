// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceCertificateResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceCertificateResponseBody) *GetDeviceCertificateResponse
	GetBody() *GetDeviceCertificateResponseBody
}

type GetDeviceCertificateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceCertificateResponse) GetBody() *GetDeviceCertificateResponseBody {
	return s.Body
}

func (s *GetDeviceCertificateResponse) SetHeaders(v map[string]*string) *GetDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceCertificateResponse) SetStatusCode(v int32) *GetDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceCertificateResponse) SetBody(v *GetDeviceCertificateResponseBody) *GetDeviceCertificateResponse {
	s.Body = v
	return s
}

func (s *GetDeviceCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
