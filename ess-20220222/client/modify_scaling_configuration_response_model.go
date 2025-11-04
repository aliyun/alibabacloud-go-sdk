// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyScalingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyScalingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyScalingConfigurationResponseBody) *ModifyScalingConfigurationResponse
	GetBody() *ModifyScalingConfigurationResponseBody
}

type ModifyScalingConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScalingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyScalingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyScalingConfigurationResponse) GetBody() *ModifyScalingConfigurationResponseBody {
	return s.Body
}

func (s *ModifyScalingConfigurationResponse) SetHeaders(v map[string]*string) *ModifyScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingConfigurationResponse) SetStatusCode(v int32) *ModifyScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScalingConfigurationResponse) SetBody(v *ModifyScalingConfigurationResponseBody) *ModifyScalingConfigurationResponse {
	s.Body = v
	return s
}

func (s *ModifyScalingConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
