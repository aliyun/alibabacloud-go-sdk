// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderSCIMResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIdentityProviderSCIMResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIdentityProviderSCIMResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIdentityProviderSCIMResponseBody) *UpdateIdentityProviderSCIMResponse
	GetBody() *UpdateIdentityProviderSCIMResponseBody
}

type UpdateIdentityProviderSCIMResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIdentityProviderSCIMResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIdentityProviderSCIMResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMResponse) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIdentityProviderSCIMResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIdentityProviderSCIMResponse) GetBody() *UpdateIdentityProviderSCIMResponseBody {
	return s.Body
}

func (s *UpdateIdentityProviderSCIMResponse) SetHeaders(v map[string]*string) *UpdateIdentityProviderSCIMResponse {
	s.Headers = v
	return s
}

func (s *UpdateIdentityProviderSCIMResponse) SetStatusCode(v int32) *UpdateIdentityProviderSCIMResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponse) SetBody(v *UpdateIdentityProviderSCIMResponseBody) *UpdateIdentityProviderSCIMResponse {
	s.Body = v
	return s
}

func (s *UpdateIdentityProviderSCIMResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
