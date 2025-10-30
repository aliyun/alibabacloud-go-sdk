// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRevokeClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRevokeClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRevokeClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *CreateRevokeClientCertificateResponseBody) *CreateRevokeClientCertificateResponse
	GetBody() *CreateRevokeClientCertificateResponseBody
}

type CreateRevokeClientCertificateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRevokeClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRevokeClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRevokeClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateRevokeClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRevokeClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRevokeClientCertificateResponse) GetBody() *CreateRevokeClientCertificateResponseBody {
	return s.Body
}

func (s *CreateRevokeClientCertificateResponse) SetHeaders(v map[string]*string) *CreateRevokeClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateRevokeClientCertificateResponse) SetStatusCode(v int32) *CreateRevokeClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRevokeClientCertificateResponse) SetBody(v *CreateRevokeClientCertificateResponseBody) *CreateRevokeClientCertificateResponse {
	s.Body = v
	return s
}

func (s *CreateRevokeClientCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
