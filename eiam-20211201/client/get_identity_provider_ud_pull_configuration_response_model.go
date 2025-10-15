// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderUdPullConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityProviderUdPullConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityProviderUdPullConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityProviderUdPullConfigurationResponseBody) *GetIdentityProviderUdPullConfigurationResponse
	GetBody() *GetIdentityProviderUdPullConfigurationResponseBody
}

type GetIdentityProviderUdPullConfigurationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityProviderUdPullConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityProviderUdPullConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPullConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPullConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityProviderUdPullConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityProviderUdPullConfigurationResponse) GetBody() *GetIdentityProviderUdPullConfigurationResponseBody {
	return s.Body
}

func (s *GetIdentityProviderUdPullConfigurationResponse) SetHeaders(v map[string]*string) *GetIdentityProviderUdPullConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponse) SetStatusCode(v int32) *GetIdentityProviderUdPullConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponse) SetBody(v *GetIdentityProviderUdPullConfigurationResponseBody) *GetIdentityProviderUdPullConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
