// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordExpirationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPasswordExpirationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPasswordExpirationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetPasswordExpirationConfigurationResponseBody) *SetPasswordExpirationConfigurationResponse
	GetBody() *SetPasswordExpirationConfigurationResponseBody
}

type SetPasswordExpirationConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordExpirationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordExpirationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordExpirationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordExpirationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPasswordExpirationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPasswordExpirationConfigurationResponse) GetBody() *SetPasswordExpirationConfigurationResponseBody {
	return s.Body
}

func (s *SetPasswordExpirationConfigurationResponse) SetHeaders(v map[string]*string) *SetPasswordExpirationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordExpirationConfigurationResponse) SetStatusCode(v int32) *SetPasswordExpirationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordExpirationConfigurationResponse) SetBody(v *SetPasswordExpirationConfigurationResponseBody) *SetPasswordExpirationConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetPasswordExpirationConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
