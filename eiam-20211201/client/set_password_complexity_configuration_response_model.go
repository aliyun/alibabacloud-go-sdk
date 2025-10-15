// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordComplexityConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPasswordComplexityConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPasswordComplexityConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetPasswordComplexityConfigurationResponseBody) *SetPasswordComplexityConfigurationResponse
	GetBody() *SetPasswordComplexityConfigurationResponseBody
}

type SetPasswordComplexityConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordComplexityConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordComplexityConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordComplexityConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordComplexityConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPasswordComplexityConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPasswordComplexityConfigurationResponse) GetBody() *SetPasswordComplexityConfigurationResponseBody {
	return s.Body
}

func (s *SetPasswordComplexityConfigurationResponse) SetHeaders(v map[string]*string) *SetPasswordComplexityConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordComplexityConfigurationResponse) SetStatusCode(v int32) *SetPasswordComplexityConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordComplexityConfigurationResponse) SetBody(v *SetPasswordComplexityConfigurationResponseBody) *SetPasswordComplexityConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetPasswordComplexityConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
