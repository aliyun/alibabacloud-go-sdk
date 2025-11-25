// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigDomainSecurityProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigDomainSecurityProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigDomainSecurityProfileResponse
	GetStatusCode() *int32
	SetBody(v *ConfigDomainSecurityProfileResponseBody) *ConfigDomainSecurityProfileResponse
	GetBody() *ConfigDomainSecurityProfileResponseBody
}

type ConfigDomainSecurityProfileResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigDomainSecurityProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigDomainSecurityProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigDomainSecurityProfileResponse) GoString() string {
	return s.String()
}

func (s *ConfigDomainSecurityProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigDomainSecurityProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigDomainSecurityProfileResponse) GetBody() *ConfigDomainSecurityProfileResponseBody {
	return s.Body
}

func (s *ConfigDomainSecurityProfileResponse) SetHeaders(v map[string]*string) *ConfigDomainSecurityProfileResponse {
	s.Headers = v
	return s
}

func (s *ConfigDomainSecurityProfileResponse) SetStatusCode(v int32) *ConfigDomainSecurityProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigDomainSecurityProfileResponse) SetBody(v *ConfigDomainSecurityProfileResponseBody) *ConfigDomainSecurityProfileResponse {
	s.Body = v
	return s
}

func (s *ConfigDomainSecurityProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
