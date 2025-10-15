// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableIdentityProviderAuthnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableIdentityProviderAuthnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableIdentityProviderAuthnResponse
	GetStatusCode() *int32
	SetBody(v *DisableIdentityProviderAuthnResponseBody) *DisableIdentityProviderAuthnResponse
	GetBody() *DisableIdentityProviderAuthnResponseBody
}

type DisableIdentityProviderAuthnResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableIdentityProviderAuthnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableIdentityProviderAuthnResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableIdentityProviderAuthnResponse) GoString() string {
	return s.String()
}

func (s *DisableIdentityProviderAuthnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableIdentityProviderAuthnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableIdentityProviderAuthnResponse) GetBody() *DisableIdentityProviderAuthnResponseBody {
	return s.Body
}

func (s *DisableIdentityProviderAuthnResponse) SetHeaders(v map[string]*string) *DisableIdentityProviderAuthnResponse {
	s.Headers = v
	return s
}

func (s *DisableIdentityProviderAuthnResponse) SetStatusCode(v int32) *DisableIdentityProviderAuthnResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableIdentityProviderAuthnResponse) SetBody(v *DisableIdentityProviderAuthnResponseBody) *DisableIdentityProviderAuthnResponse {
	s.Body = v
	return s
}

func (s *DisableIdentityProviderAuthnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
