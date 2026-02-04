// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderUdPushConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityProviderUdPushConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityProviderUdPushConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityProviderUdPushConfigurationResponseBody) *GetIdentityProviderUdPushConfigurationResponse
	GetBody() *GetIdentityProviderUdPushConfigurationResponseBody
}

type GetIdentityProviderUdPushConfigurationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityProviderUdPushConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityProviderUdPushConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPushConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPushConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityProviderUdPushConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityProviderUdPushConfigurationResponse) GetBody() *GetIdentityProviderUdPushConfigurationResponseBody {
	return s.Body
}

func (s *GetIdentityProviderUdPushConfigurationResponse) SetHeaders(v map[string]*string) *GetIdentityProviderUdPushConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponse) SetStatusCode(v int32) *GetIdentityProviderUdPushConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponse) SetBody(v *GetIdentityProviderUdPushConfigurationResponseBody) *GetIdentityProviderUdPushConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
