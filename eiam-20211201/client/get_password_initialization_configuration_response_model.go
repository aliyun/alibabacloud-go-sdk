// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordInitializationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPasswordInitializationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPasswordInitializationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetPasswordInitializationConfigurationResponseBody) *GetPasswordInitializationConfigurationResponse
	GetBody() *GetPasswordInitializationConfigurationResponseBody
}

type GetPasswordInitializationConfigurationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordInitializationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordInitializationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordInitializationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordInitializationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPasswordInitializationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPasswordInitializationConfigurationResponse) GetBody() *GetPasswordInitializationConfigurationResponseBody {
	return s.Body
}

func (s *GetPasswordInitializationConfigurationResponse) SetHeaders(v map[string]*string) *GetPasswordInitializationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordInitializationConfigurationResponse) SetStatusCode(v int32) *GetPasswordInitializationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordInitializationConfigurationResponse) SetBody(v *GetPasswordInitializationConfigurationResponseBody) *GetPasswordInitializationConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetPasswordInitializationConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
