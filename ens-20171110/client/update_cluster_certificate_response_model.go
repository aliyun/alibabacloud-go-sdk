// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClusterCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClusterCertificateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClusterCertificateResponseBody) *UpdateClusterCertificateResponse
	GetBody() *UpdateClusterCertificateResponseBody
}

type UpdateClusterCertificateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterCertificateResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClusterCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClusterCertificateResponse) GetBody() *UpdateClusterCertificateResponseBody {
	return s.Body
}

func (s *UpdateClusterCertificateResponse) SetHeaders(v map[string]*string) *UpdateClusterCertificateResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterCertificateResponse) SetStatusCode(v int32) *UpdateClusterCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterCertificateResponse) SetBody(v *UpdateClusterCertificateResponseBody) *UpdateClusterCertificateResponse {
	s.Body = v
	return s
}

func (s *UpdateClusterCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
