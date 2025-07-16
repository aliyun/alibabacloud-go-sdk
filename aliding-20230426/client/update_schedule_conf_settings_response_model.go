// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConfSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScheduleConfSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScheduleConfSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScheduleConfSettingsResponseBody) *UpdateScheduleConfSettingsResponse
	GetBody() *UpdateScheduleConfSettingsResponseBody
}

type UpdateScheduleConfSettingsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduleConfSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduleConfSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScheduleConfSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScheduleConfSettingsResponse) GetBody() *UpdateScheduleConfSettingsResponseBody {
	return s.Body
}

func (s *UpdateScheduleConfSettingsResponse) SetHeaders(v map[string]*string) *UpdateScheduleConfSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleConfSettingsResponse) SetStatusCode(v int32) *UpdateScheduleConfSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduleConfSettingsResponse) SetBody(v *UpdateScheduleConfSettingsResponseBody) *UpdateScheduleConfSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateScheduleConfSettingsResponse) Validate() error {
	return dara.Validate(s)
}
