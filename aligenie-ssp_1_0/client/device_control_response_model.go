// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeviceControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeviceControlResponse
	GetStatusCode() *int32
	SetBody(v *DeviceControlResponseBody) *DeviceControlResponse
	GetBody() *DeviceControlResponseBody
}

type DeviceControlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeviceControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeviceControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlResponse) GoString() string {
	return s.String()
}

func (s *DeviceControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeviceControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeviceControlResponse) GetBody() *DeviceControlResponseBody {
	return s.Body
}

func (s *DeviceControlResponse) SetHeaders(v map[string]*string) *DeviceControlResponse {
	s.Headers = v
	return s
}

func (s *DeviceControlResponse) SetStatusCode(v int32) *DeviceControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceControlResponse) SetBody(v *DeviceControlResponseBody) *DeviceControlResponse {
	s.Body = v
	return s
}

func (s *DeviceControlResponse) Validate() error {
	return dara.Validate(s)
}
