// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingInstanceSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRenderingInstanceSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRenderingInstanceSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRenderingInstanceSettingsResponseBody) *UpdateRenderingInstanceSettingsResponse
	GetBody() *UpdateRenderingInstanceSettingsResponseBody
}

type UpdateRenderingInstanceSettingsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRenderingInstanceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRenderingInstanceSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRenderingInstanceSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRenderingInstanceSettingsResponse) GetBody() *UpdateRenderingInstanceSettingsResponseBody {
	return s.Body
}

func (s *UpdateRenderingInstanceSettingsResponse) SetHeaders(v map[string]*string) *UpdateRenderingInstanceSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateRenderingInstanceSettingsResponse) SetStatusCode(v int32) *UpdateRenderingInstanceSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRenderingInstanceSettingsResponse) SetBody(v *UpdateRenderingInstanceSettingsResponseBody) *UpdateRenderingInstanceSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateRenderingInstanceSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
