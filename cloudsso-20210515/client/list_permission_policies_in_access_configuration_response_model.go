// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionPoliciesInAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPermissionPoliciesInAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPermissionPoliciesInAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ListPermissionPoliciesInAccessConfigurationResponseBody) *ListPermissionPoliciesInAccessConfigurationResponse
	GetBody() *ListPermissionPoliciesInAccessConfigurationResponseBody
}

type ListPermissionPoliciesInAccessConfigurationResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionPoliciesInAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionPoliciesInAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionPoliciesInAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) GetBody() *ListPermissionPoliciesInAccessConfigurationResponseBody {
	return s.Body
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) SetHeaders(v map[string]*string) *ListPermissionPoliciesInAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) SetStatusCode(v int32) *ListPermissionPoliciesInAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) SetBody(v *ListPermissionPoliciesInAccessConfigurationResponseBody) *ListPermissionPoliciesInAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
