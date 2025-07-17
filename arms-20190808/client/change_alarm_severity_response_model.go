// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAlarmSeverityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeAlarmSeverityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeAlarmSeverityResponse
	GetStatusCode() *int32
	SetBody(v *ChangeAlarmSeverityResponseBody) *ChangeAlarmSeverityResponse
	GetBody() *ChangeAlarmSeverityResponseBody
}

type ChangeAlarmSeverityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAlarmSeverityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAlarmSeverityResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeAlarmSeverityResponse) GoString() string {
	return s.String()
}

func (s *ChangeAlarmSeverityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeAlarmSeverityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeAlarmSeverityResponse) GetBody() *ChangeAlarmSeverityResponseBody {
	return s.Body
}

func (s *ChangeAlarmSeverityResponse) SetHeaders(v map[string]*string) *ChangeAlarmSeverityResponse {
	s.Headers = v
	return s
}

func (s *ChangeAlarmSeverityResponse) SetStatusCode(v int32) *ChangeAlarmSeverityResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAlarmSeverityResponse) SetBody(v *ChangeAlarmSeverityResponseBody) *ChangeAlarmSeverityResponse {
	s.Body = v
	return s
}

func (s *ChangeAlarmSeverityResponse) Validate() error {
	return dara.Validate(s)
}
