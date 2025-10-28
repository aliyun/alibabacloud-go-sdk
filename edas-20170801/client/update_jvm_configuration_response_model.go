// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJvmConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJvmConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJvmConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJvmConfigurationResponseBody) *UpdateJvmConfigurationResponse
	GetBody() *UpdateJvmConfigurationResponseBody
}

type UpdateJvmConfigurationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJvmConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJvmConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJvmConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateJvmConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJvmConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJvmConfigurationResponse) GetBody() *UpdateJvmConfigurationResponseBody {
	return s.Body
}

func (s *UpdateJvmConfigurationResponse) SetHeaders(v map[string]*string) *UpdateJvmConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateJvmConfigurationResponse) SetStatusCode(v int32) *UpdateJvmConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJvmConfigurationResponse) SetBody(v *UpdateJvmConfigurationResponseBody) *UpdateJvmConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateJvmConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
