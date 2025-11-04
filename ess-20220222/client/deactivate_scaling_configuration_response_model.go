// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateScalingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactivateScalingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactivateScalingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeactivateScalingConfigurationResponseBody) *DeactivateScalingConfigurationResponse
	GetBody() *DeactivateScalingConfigurationResponseBody
}

type DeactivateScalingConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactivateScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactivateScalingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactivateScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeactivateScalingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactivateScalingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactivateScalingConfigurationResponse) GetBody() *DeactivateScalingConfigurationResponseBody {
	return s.Body
}

func (s *DeactivateScalingConfigurationResponse) SetHeaders(v map[string]*string) *DeactivateScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeactivateScalingConfigurationResponse) SetStatusCode(v int32) *DeactivateScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateScalingConfigurationResponse) SetBody(v *DeactivateScalingConfigurationResponseBody) *DeactivateScalingConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeactivateScalingConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
