// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPermissionPolicyToAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPermissionPolicyToAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPermissionPolicyToAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *AddPermissionPolicyToAccessConfigurationResponseBody) *AddPermissionPolicyToAccessConfigurationResponse
	GetBody() *AddPermissionPolicyToAccessConfigurationResponseBody
}

type AddPermissionPolicyToAccessConfigurationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPermissionPolicyToAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPermissionPolicyToAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionPolicyToAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) GetBody() *AddPermissionPolicyToAccessConfigurationResponseBody {
	return s.Body
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) SetHeaders(v map[string]*string) *AddPermissionPolicyToAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) SetStatusCode(v int32) *AddPermissionPolicyToAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) SetBody(v *AddPermissionPolicyToAccessConfigurationResponseBody) *AddPermissionPolicyToAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
