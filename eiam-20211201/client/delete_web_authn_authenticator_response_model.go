// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebAuthnAuthenticatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebAuthnAuthenticatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebAuthnAuthenticatorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWebAuthnAuthenticatorResponseBody) *DeleteWebAuthnAuthenticatorResponse
	GetBody() *DeleteWebAuthnAuthenticatorResponseBody
}

type DeleteWebAuthnAuthenticatorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebAuthnAuthenticatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebAuthnAuthenticatorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebAuthnAuthenticatorResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebAuthnAuthenticatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebAuthnAuthenticatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebAuthnAuthenticatorResponse) GetBody() *DeleteWebAuthnAuthenticatorResponseBody {
	return s.Body
}

func (s *DeleteWebAuthnAuthenticatorResponse) SetHeaders(v map[string]*string) *DeleteWebAuthnAuthenticatorResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebAuthnAuthenticatorResponse) SetStatusCode(v int32) *DeleteWebAuthnAuthenticatorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebAuthnAuthenticatorResponse) SetBody(v *DeleteWebAuthnAuthenticatorResponseBody) *DeleteWebAuthnAuthenticatorResponse {
	s.Body = v
	return s
}

func (s *DeleteWebAuthnAuthenticatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
