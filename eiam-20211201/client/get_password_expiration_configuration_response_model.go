// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordExpirationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPasswordExpirationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPasswordExpirationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetPasswordExpirationConfigurationResponseBody) *GetPasswordExpirationConfigurationResponse
	GetBody() *GetPasswordExpirationConfigurationResponseBody
}

type GetPasswordExpirationConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordExpirationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordExpirationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordExpirationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordExpirationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPasswordExpirationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPasswordExpirationConfigurationResponse) GetBody() *GetPasswordExpirationConfigurationResponseBody {
	return s.Body
}

func (s *GetPasswordExpirationConfigurationResponse) SetHeaders(v map[string]*string) *GetPasswordExpirationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordExpirationConfigurationResponse) SetStatusCode(v int32) *GetPasswordExpirationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponse) SetBody(v *GetPasswordExpirationConfigurationResponseBody) *GetPasswordExpirationConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetPasswordExpirationConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
