// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlarmResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlarmResponseBody) *UpdateAlarmResponse
	GetBody() *UpdateAlarmResponseBody
}

type UpdateAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlarmResponse) GetBody() *UpdateAlarmResponseBody {
	return s.Body
}

func (s *UpdateAlarmResponse) SetHeaders(v map[string]*string) *UpdateAlarmResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlarmResponse) SetStatusCode(v int32) *UpdateAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlarmResponse) SetBody(v *UpdateAlarmResponseBody) *UpdateAlarmResponse {
	s.Body = v
	return s
}

func (s *UpdateAlarmResponse) Validate() error {
	return dara.Validate(s)
}
