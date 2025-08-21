// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDynamicSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDynamicSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDynamicSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDynamicSettingsResponseBody) *UpdateDynamicSettingsResponse
	GetBody() *UpdateDynamicSettingsResponseBody
}

type UpdateDynamicSettingsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDynamicSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDynamicSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDynamicSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateDynamicSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDynamicSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDynamicSettingsResponse) GetBody() *UpdateDynamicSettingsResponseBody {
	return s.Body
}

func (s *UpdateDynamicSettingsResponse) SetHeaders(v map[string]*string) *UpdateDynamicSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateDynamicSettingsResponse) SetStatusCode(v int32) *UpdateDynamicSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDynamicSettingsResponse) SetBody(v *UpdateDynamicSettingsResponseBody) *UpdateDynamicSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateDynamicSettingsResponse) Validate() error {
	return dara.Validate(s)
}
