// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyEciScalingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyEciScalingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyEciScalingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ApplyEciScalingConfigurationResponseBody) *ApplyEciScalingConfigurationResponse
	GetBody() *ApplyEciScalingConfigurationResponseBody
}

type ApplyEciScalingConfigurationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyEciScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyEciScalingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyEciScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ApplyEciScalingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyEciScalingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyEciScalingConfigurationResponse) GetBody() *ApplyEciScalingConfigurationResponseBody {
	return s.Body
}

func (s *ApplyEciScalingConfigurationResponse) SetHeaders(v map[string]*string) *ApplyEciScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ApplyEciScalingConfigurationResponse) SetStatusCode(v int32) *ApplyEciScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyEciScalingConfigurationResponse) SetBody(v *ApplyEciScalingConfigurationResponseBody) *ApplyEciScalingConfigurationResponse {
	s.Body = v
	return s
}

func (s *ApplyEciScalingConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
