// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordComplexityConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPasswordComplexityConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPasswordComplexityConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetPasswordComplexityConfigurationResponseBody) *GetPasswordComplexityConfigurationResponse
	GetBody() *GetPasswordComplexityConfigurationResponseBody
}

type GetPasswordComplexityConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordComplexityConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordComplexityConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordComplexityConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPasswordComplexityConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPasswordComplexityConfigurationResponse) GetBody() *GetPasswordComplexityConfigurationResponseBody {
	return s.Body
}

func (s *GetPasswordComplexityConfigurationResponse) SetHeaders(v map[string]*string) *GetPasswordComplexityConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordComplexityConfigurationResponse) SetStatusCode(v int32) *GetPasswordComplexityConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordComplexityConfigurationResponse) SetBody(v *GetPasswordComplexityConfigurationResponseBody) *GetPasswordComplexityConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetPasswordComplexityConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
