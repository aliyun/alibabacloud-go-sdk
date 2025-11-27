// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRenderingInstanceConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRenderingInstanceConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRenderingInstanceConfigurationResponseBody) *DeleteRenderingInstanceConfigurationResponse
	GetBody() *DeleteRenderingInstanceConfigurationResponseBody
}

type DeleteRenderingInstanceConfigurationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRenderingInstanceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRenderingInstanceConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRenderingInstanceConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRenderingInstanceConfigurationResponse) GetBody() *DeleteRenderingInstanceConfigurationResponseBody {
	return s.Body
}

func (s *DeleteRenderingInstanceConfigurationResponse) SetHeaders(v map[string]*string) *DeleteRenderingInstanceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteRenderingInstanceConfigurationResponse) SetStatusCode(v int32) *DeleteRenderingInstanceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRenderingInstanceConfigurationResponse) SetBody(v *DeleteRenderingInstanceConfigurationResponseBody) *DeleteRenderingInstanceConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteRenderingInstanceConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
