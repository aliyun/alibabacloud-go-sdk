// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ActivateClientCertificateResponseBody) *ActivateClientCertificateResponse
	GetBody() *ActivateClientCertificateResponseBody
}

type ActivateClientCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *ActivateClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateClientCertificateResponse) GetBody() *ActivateClientCertificateResponseBody {
	return s.Body
}

func (s *ActivateClientCertificateResponse) SetHeaders(v map[string]*string) *ActivateClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *ActivateClientCertificateResponse) SetStatusCode(v int32) *ActivateClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateClientCertificateResponse) SetBody(v *ActivateClientCertificateResponseBody) *ActivateClientCertificateResponse {
	s.Body = v
	return s
}

func (s *ActivateClientCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
