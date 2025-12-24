// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableIdentityProviderAdvancedAbilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableIdentityProviderAdvancedAbilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableIdentityProviderAdvancedAbilityResponse
	GetStatusCode() *int32
	SetBody(v *DisableIdentityProviderAdvancedAbilityResponseBody) *DisableIdentityProviderAdvancedAbilityResponse
	GetBody() *DisableIdentityProviderAdvancedAbilityResponseBody
}

type DisableIdentityProviderAdvancedAbilityResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableIdentityProviderAdvancedAbilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableIdentityProviderAdvancedAbilityResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableIdentityProviderAdvancedAbilityResponse) GoString() string {
	return s.String()
}

func (s *DisableIdentityProviderAdvancedAbilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableIdentityProviderAdvancedAbilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableIdentityProviderAdvancedAbilityResponse) GetBody() *DisableIdentityProviderAdvancedAbilityResponseBody {
	return s.Body
}

func (s *DisableIdentityProviderAdvancedAbilityResponse) SetHeaders(v map[string]*string) *DisableIdentityProviderAdvancedAbilityResponse {
	s.Headers = v
	return s
}

func (s *DisableIdentityProviderAdvancedAbilityResponse) SetStatusCode(v int32) *DisableIdentityProviderAdvancedAbilityResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableIdentityProviderAdvancedAbilityResponse) SetBody(v *DisableIdentityProviderAdvancedAbilityResponseBody) *DisableIdentityProviderAdvancedAbilityResponse {
	s.Body = v
	return s
}

func (s *DisableIdentityProviderAdvancedAbilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
