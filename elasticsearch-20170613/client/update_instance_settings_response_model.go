// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceSettingsResponseBody) *UpdateInstanceSettingsResponse
	GetBody() *UpdateInstanceSettingsResponseBody
}

type UpdateInstanceSettingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceSettingsResponse) GetBody() *UpdateInstanceSettingsResponseBody {
	return s.Body
}

func (s *UpdateInstanceSettingsResponse) SetHeaders(v map[string]*string) *UpdateInstanceSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceSettingsResponse) SetStatusCode(v int32) *UpdateInstanceSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceSettingsResponse) SetBody(v *UpdateInstanceSettingsResponseBody) *UpdateInstanceSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
