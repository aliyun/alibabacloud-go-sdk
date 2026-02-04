// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderAdvancedConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityProviderAdvancedConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityProviderAdvancedConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityProviderAdvancedConfigurationResponseBody) *GetIdentityProviderAdvancedConfigurationResponse
	GetBody() *GetIdentityProviderAdvancedConfigurationResponseBody
}

type GetIdentityProviderAdvancedConfigurationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityProviderAdvancedConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityProviderAdvancedConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderAdvancedConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderAdvancedConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityProviderAdvancedConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityProviderAdvancedConfigurationResponse) GetBody() *GetIdentityProviderAdvancedConfigurationResponseBody {
	return s.Body
}

func (s *GetIdentityProviderAdvancedConfigurationResponse) SetHeaders(v map[string]*string) *GetIdentityProviderAdvancedConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponse) SetStatusCode(v int32) *GetIdentityProviderAdvancedConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponse) SetBody(v *GetIdentityProviderAdvancedConfigurationResponseBody) *GetIdentityProviderAdvancedConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
