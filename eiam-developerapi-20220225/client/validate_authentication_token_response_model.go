// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateAuthenticationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateAuthenticationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateAuthenticationTokenResponse
	GetStatusCode() *int32
	SetBody(v *ValidateAuthenticationTokenResponseBody) *ValidateAuthenticationTokenResponse
	GetBody() *ValidateAuthenticationTokenResponseBody
}

type ValidateAuthenticationTokenResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateAuthenticationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateAuthenticationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateAuthenticationTokenResponse) GoString() string {
	return s.String()
}

func (s *ValidateAuthenticationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateAuthenticationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateAuthenticationTokenResponse) GetBody() *ValidateAuthenticationTokenResponseBody {
	return s.Body
}

func (s *ValidateAuthenticationTokenResponse) SetHeaders(v map[string]*string) *ValidateAuthenticationTokenResponse {
	s.Headers = v
	return s
}

func (s *ValidateAuthenticationTokenResponse) SetStatusCode(v int32) *ValidateAuthenticationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateAuthenticationTokenResponse) SetBody(v *ValidateAuthenticationTokenResponseBody) *ValidateAuthenticationTokenResponse {
	s.Body = v
	return s
}

func (s *ValidateAuthenticationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
