// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIdentityProviderResponseBody) *UpdateIdentityProviderResponse
	GetBody() *UpdateIdentityProviderResponseBody
}

type UpdateIdentityProviderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIdentityProviderResponse) GetBody() *UpdateIdentityProviderResponseBody {
	return s.Body
}

func (s *UpdateIdentityProviderResponse) SetHeaders(v map[string]*string) *UpdateIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateIdentityProviderResponse) SetStatusCode(v int32) *UpdateIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIdentityProviderResponse) SetBody(v *UpdateIdentityProviderResponseBody) *UpdateIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
