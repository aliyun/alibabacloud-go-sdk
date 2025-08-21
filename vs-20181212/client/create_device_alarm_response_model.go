// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeviceAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeviceAlarmResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeviceAlarmResponseBody) *CreateDeviceAlarmResponse
	GetBody() *CreateDeviceAlarmResponseBody
}

type CreateDeviceAlarmResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeviceAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeviceAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeviceAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeviceAlarmResponse) GetBody() *CreateDeviceAlarmResponseBody {
	return s.Body
}

func (s *CreateDeviceAlarmResponse) SetHeaders(v map[string]*string) *CreateDeviceAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceAlarmResponse) SetStatusCode(v int32) *CreateDeviceAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeviceAlarmResponse) SetBody(v *CreateDeviceAlarmResponseBody) *CreateDeviceAlarmResponse {
	s.Body = v
	return s
}

func (s *CreateDeviceAlarmResponse) Validate() error {
	return dara.Validate(s)
}
