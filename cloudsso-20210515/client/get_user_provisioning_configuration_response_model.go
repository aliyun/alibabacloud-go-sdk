// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserProvisioningConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserProvisioningConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetUserProvisioningConfigurationResponseBody) *GetUserProvisioningConfigurationResponse
	GetBody() *GetUserProvisioningConfigurationResponseBody
}

type GetUserProvisioningConfigurationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserProvisioningConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserProvisioningConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserProvisioningConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserProvisioningConfigurationResponse) GetBody() *GetUserProvisioningConfigurationResponseBody {
	return s.Body
}

func (s *GetUserProvisioningConfigurationResponse) SetHeaders(v map[string]*string) *GetUserProvisioningConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetUserProvisioningConfigurationResponse) SetStatusCode(v int32) *GetUserProvisioningConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponse) SetBody(v *GetUserProvisioningConfigurationResponseBody) *GetUserProvisioningConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetUserProvisioningConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
