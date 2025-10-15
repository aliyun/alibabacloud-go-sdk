// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentityProviderUdPullConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetIdentityProviderUdPullConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetIdentityProviderUdPullConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetIdentityProviderUdPullConfigurationResponseBody) *SetIdentityProviderUdPullConfigurationResponse
	GetBody() *SetIdentityProviderUdPullConfigurationResponseBody
}

type SetIdentityProviderUdPullConfigurationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetIdentityProviderUdPullConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetIdentityProviderUdPullConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPullConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPullConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetIdentityProviderUdPullConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetIdentityProviderUdPullConfigurationResponse) GetBody() *SetIdentityProviderUdPullConfigurationResponseBody {
	return s.Body
}

func (s *SetIdentityProviderUdPullConfigurationResponse) SetHeaders(v map[string]*string) *SetIdentityProviderUdPullConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationResponse) SetStatusCode(v int32) *SetIdentityProviderUdPullConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationResponse) SetBody(v *SetIdentityProviderUdPullConfigurationResponseBody) *SetIdentityProviderUdPullConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
