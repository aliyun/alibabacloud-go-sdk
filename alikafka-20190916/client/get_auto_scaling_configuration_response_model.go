// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoScalingConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoScalingConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoScalingConfigurationResponseBody) *GetAutoScalingConfigurationResponse
	GetBody() *GetAutoScalingConfigurationResponseBody
}

type GetAutoScalingConfigurationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoScalingConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoScalingConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoScalingConfigurationResponse) GetBody() *GetAutoScalingConfigurationResponseBody {
	return s.Body
}

func (s *GetAutoScalingConfigurationResponse) SetHeaders(v map[string]*string) *GetAutoScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScalingConfigurationResponse) SetStatusCode(v int32) *GetAutoScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScalingConfigurationResponse) SetBody(v *GetAutoScalingConfigurationResponseBody) *GetAutoScalingConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetAutoScalingConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
