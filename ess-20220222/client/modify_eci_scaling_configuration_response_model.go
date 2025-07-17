// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEciScalingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEciScalingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEciScalingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEciScalingConfigurationResponseBody) *ModifyEciScalingConfigurationResponse
	GetBody() *ModifyEciScalingConfigurationResponseBody
}

type ModifyEciScalingConfigurationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEciScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEciScalingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEciScalingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEciScalingConfigurationResponse) GetBody() *ModifyEciScalingConfigurationResponseBody {
	return s.Body
}

func (s *ModifyEciScalingConfigurationResponse) SetHeaders(v map[string]*string) *ModifyEciScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyEciScalingConfigurationResponse) SetStatusCode(v int32) *ModifyEciScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEciScalingConfigurationResponse) SetBody(v *ModifyEciScalingConfigurationResponseBody) *ModifyEciScalingConfigurationResponse {
	s.Body = v
	return s
}

func (s *ModifyEciScalingConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
