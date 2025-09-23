// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7CertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer7CertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer7CertResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer7CertResponseBody) *ConfigLayer7CertResponse
	GetBody() *ConfigLayer7CertResponseBody
}

type ConfigLayer7CertResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7CertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer7CertResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7CertResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer7CertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer7CertResponse) GetBody() *ConfigLayer7CertResponseBody {
	return s.Body
}

func (s *ConfigLayer7CertResponse) SetHeaders(v map[string]*string) *ConfigLayer7CertResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7CertResponse) SetStatusCode(v int32) *ConfigLayer7CertResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7CertResponse) SetBody(v *ConfigLayer7CertResponseBody) *ConfigLayer7CertResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer7CertResponse) Validate() error {
	return dara.Validate(s)
}
