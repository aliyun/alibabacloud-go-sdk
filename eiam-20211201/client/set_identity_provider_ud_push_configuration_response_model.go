// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentityProviderUdPushConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetIdentityProviderUdPushConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetIdentityProviderUdPushConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetIdentityProviderUdPushConfigurationResponseBody) *SetIdentityProviderUdPushConfigurationResponse
	GetBody() *SetIdentityProviderUdPushConfigurationResponseBody
}

type SetIdentityProviderUdPushConfigurationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetIdentityProviderUdPushConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetIdentityProviderUdPushConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPushConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPushConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetIdentityProviderUdPushConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetIdentityProviderUdPushConfigurationResponse) GetBody() *SetIdentityProviderUdPushConfigurationResponseBody {
	return s.Body
}

func (s *SetIdentityProviderUdPushConfigurationResponse) SetHeaders(v map[string]*string) *SetIdentityProviderUdPushConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationResponse) SetStatusCode(v int32) *SetIdentityProviderUdPushConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationResponse) SetBody(v *SetIdentityProviderUdPushConfigurationResponseBody) *SetIdentityProviderUdPushConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
