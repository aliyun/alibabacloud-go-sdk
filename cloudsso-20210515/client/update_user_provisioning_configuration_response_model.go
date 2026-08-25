// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserProvisioningConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserProvisioningConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserProvisioningConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserProvisioningConfigurationResponseBody) *UpdateUserProvisioningConfigurationResponse
	GetBody() *UpdateUserProvisioningConfigurationResponseBody
}

type UpdateUserProvisioningConfigurationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserProvisioningConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserProvisioningConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserProvisioningConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserProvisioningConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserProvisioningConfigurationResponse) GetBody() *UpdateUserProvisioningConfigurationResponseBody {
	return s.Body
}

func (s *UpdateUserProvisioningConfigurationResponse) SetHeaders(v map[string]*string) *UpdateUserProvisioningConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponse) SetStatusCode(v int32) *UpdateUserProvisioningConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponse) SetBody(v *UpdateUserProvisioningConfigurationResponseBody) *UpdateUserProvisioningConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
