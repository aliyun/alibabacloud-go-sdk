// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAtiAlertSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAtiAlertSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAtiAlertSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAtiAlertSettingsResponseBody) *UpdateAtiAlertSettingsResponse
	GetBody() *UpdateAtiAlertSettingsResponseBody
}

type UpdateAtiAlertSettingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAtiAlertSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAtiAlertSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiAlertSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateAtiAlertSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAtiAlertSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAtiAlertSettingsResponse) GetBody() *UpdateAtiAlertSettingsResponseBody {
	return s.Body
}

func (s *UpdateAtiAlertSettingsResponse) SetHeaders(v map[string]*string) *UpdateAtiAlertSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateAtiAlertSettingsResponse) SetStatusCode(v int32) *UpdateAtiAlertSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponse) SetBody(v *UpdateAtiAlertSettingsResponseBody) *UpdateAtiAlertSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateAtiAlertSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
