// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlarmEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlarmEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlarmEventResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlarmEventResponseBody) *UpdateAlarmEventResponse
	GetBody() *UpdateAlarmEventResponseBody
}

type UpdateAlarmEventResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlarmEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlarmEventResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmEventResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlarmEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlarmEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlarmEventResponse) GetBody() *UpdateAlarmEventResponseBody {
	return s.Body
}

func (s *UpdateAlarmEventResponse) SetHeaders(v map[string]*string) *UpdateAlarmEventResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlarmEventResponse) SetStatusCode(v int32) *UpdateAlarmEventResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlarmEventResponse) SetBody(v *UpdateAlarmEventResponseBody) *UpdateAlarmEventResponse {
	s.Body = v
	return s
}

func (s *UpdateAlarmEventResponse) Validate() error {
	return dara.Validate(s)
}
