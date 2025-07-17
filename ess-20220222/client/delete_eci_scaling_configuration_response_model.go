// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEciScalingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEciScalingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEciScalingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEciScalingConfigurationResponseBody) *DeleteEciScalingConfigurationResponse
	GetBody() *DeleteEciScalingConfigurationResponseBody
}

type DeleteEciScalingConfigurationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEciScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEciScalingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEciScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteEciScalingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEciScalingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEciScalingConfigurationResponse) GetBody() *DeleteEciScalingConfigurationResponseBody {
	return s.Body
}

func (s *DeleteEciScalingConfigurationResponse) SetHeaders(v map[string]*string) *DeleteEciScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteEciScalingConfigurationResponse) SetStatusCode(v int32) *DeleteEciScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEciScalingConfigurationResponse) SetBody(v *DeleteEciScalingConfigurationResponseBody) *DeleteEciScalingConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteEciScalingConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
