// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEciScalingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEciScalingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEciScalingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateEciScalingConfigurationResponseBody) *CreateEciScalingConfigurationResponse
	GetBody() *CreateEciScalingConfigurationResponseBody
}

type CreateEciScalingConfigurationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEciScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEciScalingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEciScalingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEciScalingConfigurationResponse) GetBody() *CreateEciScalingConfigurationResponseBody {
	return s.Body
}

func (s *CreateEciScalingConfigurationResponse) SetHeaders(v map[string]*string) *CreateEciScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateEciScalingConfigurationResponse) SetStatusCode(v int32) *CreateEciScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEciScalingConfigurationResponse) SetBody(v *CreateEciScalingConfigurationResponseBody) *CreateEciScalingConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateEciScalingConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
