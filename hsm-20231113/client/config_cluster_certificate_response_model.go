// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigClusterCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigClusterCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ConfigClusterCertificateResponseBody) *ConfigClusterCertificateResponse
	GetBody() *ConfigClusterCertificateResponseBody
}

type ConfigClusterCertificateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigClusterCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigClusterCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterCertificateResponse) GoString() string {
	return s.String()
}

func (s *ConfigClusterCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigClusterCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigClusterCertificateResponse) GetBody() *ConfigClusterCertificateResponseBody {
	return s.Body
}

func (s *ConfigClusterCertificateResponse) SetHeaders(v map[string]*string) *ConfigClusterCertificateResponse {
	s.Headers = v
	return s
}

func (s *ConfigClusterCertificateResponse) SetStatusCode(v int32) *ConfigClusterCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigClusterCertificateResponse) SetBody(v *ConfigClusterCertificateResponseBody) *ConfigClusterCertificateResponse {
	s.Body = v
	return s
}

func (s *ConfigClusterCertificateResponse) Validate() error {
	return dara.Validate(s)
}
