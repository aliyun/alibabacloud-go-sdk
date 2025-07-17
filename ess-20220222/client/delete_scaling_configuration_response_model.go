// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScalingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScalingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScalingConfigurationResponseBody) *DeleteScalingConfigurationResponse
	GetBody() *DeleteScalingConfigurationResponseBody
}

type DeleteScalingConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScalingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScalingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScalingConfigurationResponse) GetBody() *DeleteScalingConfigurationResponseBody {
	return s.Body
}

func (s *DeleteScalingConfigurationResponse) SetHeaders(v map[string]*string) *DeleteScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingConfigurationResponse) SetStatusCode(v int32) *DeleteScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingConfigurationResponse) SetBody(v *DeleteScalingConfigurationResponseBody) *DeleteScalingConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteScalingConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
