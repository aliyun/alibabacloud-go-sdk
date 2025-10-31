// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigInstanceIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigInstanceIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *ConfigInstanceIpAddressResponseBody) *ConfigInstanceIpAddressResponse
	GetBody() *ConfigInstanceIpAddressResponseBody
}

type ConfigInstanceIpAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigInstanceIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceIpAddressResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigInstanceIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigInstanceIpAddressResponse) GetBody() *ConfigInstanceIpAddressResponseBody {
	return s.Body
}

func (s *ConfigInstanceIpAddressResponse) SetHeaders(v map[string]*string) *ConfigInstanceIpAddressResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceIpAddressResponse) SetStatusCode(v int32) *ConfigInstanceIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceIpAddressResponse) SetBody(v *ConfigInstanceIpAddressResponseBody) *ConfigInstanceIpAddressResponse {
	s.Body = v
	return s
}

func (s *ConfigInstanceIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
