// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePermissionPolicyFromAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePermissionPolicyFromAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePermissionPolicyFromAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *RemovePermissionPolicyFromAccessConfigurationResponseBody) *RemovePermissionPolicyFromAccessConfigurationResponse
	GetBody() *RemovePermissionPolicyFromAccessConfigurationResponseBody
}

type RemovePermissionPolicyFromAccessConfigurationResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePermissionPolicyFromAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePermissionPolicyFromAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePermissionPolicyFromAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) GetBody() *RemovePermissionPolicyFromAccessConfigurationResponseBody {
	return s.Body
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) SetHeaders(v map[string]*string) *RemovePermissionPolicyFromAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) SetStatusCode(v int32) *RemovePermissionPolicyFromAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) SetBody(v *RemovePermissionPolicyFromAccessConfigurationResponseBody) *RemovePermissionPolicyFromAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
