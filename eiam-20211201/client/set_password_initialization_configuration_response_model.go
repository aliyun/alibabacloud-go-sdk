// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordInitializationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPasswordInitializationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPasswordInitializationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetPasswordInitializationConfigurationResponseBody) *SetPasswordInitializationConfigurationResponse
	GetBody() *SetPasswordInitializationConfigurationResponseBody
}

type SetPasswordInitializationConfigurationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordInitializationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordInitializationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordInitializationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordInitializationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPasswordInitializationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPasswordInitializationConfigurationResponse) GetBody() *SetPasswordInitializationConfigurationResponseBody {
	return s.Body
}

func (s *SetPasswordInitializationConfigurationResponse) SetHeaders(v map[string]*string) *SetPasswordInitializationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordInitializationConfigurationResponse) SetStatusCode(v int32) *SetPasswordInitializationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordInitializationConfigurationResponse) SetBody(v *SetPasswordInitializationConfigurationResponseBody) *SetPasswordInitializationConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetPasswordInitializationConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
