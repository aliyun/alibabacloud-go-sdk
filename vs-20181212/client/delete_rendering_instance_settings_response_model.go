// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRenderingInstanceSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRenderingInstanceSettingsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRenderingInstanceSettingsResponseBody) *DeleteRenderingInstanceSettingsResponse
	GetBody() *DeleteRenderingInstanceSettingsResponseBody
}

type DeleteRenderingInstanceSettingsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRenderingInstanceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRenderingInstanceSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceSettingsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRenderingInstanceSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRenderingInstanceSettingsResponse) GetBody() *DeleteRenderingInstanceSettingsResponseBody {
	return s.Body
}

func (s *DeleteRenderingInstanceSettingsResponse) SetHeaders(v map[string]*string) *DeleteRenderingInstanceSettingsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRenderingInstanceSettingsResponse) SetStatusCode(v int32) *DeleteRenderingInstanceSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRenderingInstanceSettingsResponse) SetBody(v *DeleteRenderingInstanceSettingsResponseBody) *DeleteRenderingInstanceSettingsResponse {
	s.Body = v
	return s
}

func (s *DeleteRenderingInstanceSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
