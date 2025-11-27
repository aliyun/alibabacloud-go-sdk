// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingInstanceConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRenderingInstanceConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRenderingInstanceConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRenderingInstanceConfigurationResponseBody) *UpdateRenderingInstanceConfigurationResponse
	GetBody() *UpdateRenderingInstanceConfigurationResponseBody
}

type UpdateRenderingInstanceConfigurationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRenderingInstanceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRenderingInstanceConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRenderingInstanceConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRenderingInstanceConfigurationResponse) GetBody() *UpdateRenderingInstanceConfigurationResponseBody {
	return s.Body
}

func (s *UpdateRenderingInstanceConfigurationResponse) SetHeaders(v map[string]*string) *UpdateRenderingInstanceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateRenderingInstanceConfigurationResponse) SetStatusCode(v int32) *UpdateRenderingInstanceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRenderingInstanceConfigurationResponse) SetBody(v *UpdateRenderingInstanceConfigurationResponseBody) *UpdateRenderingInstanceConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateRenderingInstanceConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
