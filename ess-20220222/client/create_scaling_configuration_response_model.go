// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScalingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScalingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateScalingConfigurationResponseBody) *CreateScalingConfigurationResponse
	GetBody() *CreateScalingConfigurationResponseBody
}

type CreateScalingConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScalingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScalingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScalingConfigurationResponse) GetBody() *CreateScalingConfigurationResponseBody {
	return s.Body
}

func (s *CreateScalingConfigurationResponse) SetHeaders(v map[string]*string) *CreateScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingConfigurationResponse) SetStatusCode(v int32) *CreateScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScalingConfigurationResponse) SetBody(v *CreateScalingConfigurationResponseBody) *CreateScalingConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateScalingConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
