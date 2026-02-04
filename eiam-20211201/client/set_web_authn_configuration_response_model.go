// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWebAuthnConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWebAuthnConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWebAuthnConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetWebAuthnConfigurationResponseBody) *SetWebAuthnConfigurationResponse
	GetBody() *SetWebAuthnConfigurationResponseBody
}

type SetWebAuthnConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWebAuthnConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWebAuthnConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWebAuthnConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetWebAuthnConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWebAuthnConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWebAuthnConfigurationResponse) GetBody() *SetWebAuthnConfigurationResponseBody {
	return s.Body
}

func (s *SetWebAuthnConfigurationResponse) SetHeaders(v map[string]*string) *SetWebAuthnConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetWebAuthnConfigurationResponse) SetStatusCode(v int32) *SetWebAuthnConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWebAuthnConfigurationResponse) SetBody(v *SetWebAuthnConfigurationResponseBody) *SetWebAuthnConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetWebAuthnConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
