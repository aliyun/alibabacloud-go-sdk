// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceGroupMatchDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDeviceGroupMatchDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDeviceGroupMatchDevicesResponse
	GetStatusCode() *int32
	SetBody(v *AddDeviceGroupMatchDevicesResponseBody) *AddDeviceGroupMatchDevicesResponse
	GetBody() *AddDeviceGroupMatchDevicesResponseBody
}

type AddDeviceGroupMatchDevicesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDeviceGroupMatchDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDeviceGroupMatchDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceGroupMatchDevicesResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceGroupMatchDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDeviceGroupMatchDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDeviceGroupMatchDevicesResponse) GetBody() *AddDeviceGroupMatchDevicesResponseBody {
	return s.Body
}

func (s *AddDeviceGroupMatchDevicesResponse) SetHeaders(v map[string]*string) *AddDeviceGroupMatchDevicesResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceGroupMatchDevicesResponse) SetStatusCode(v int32) *AddDeviceGroupMatchDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDeviceGroupMatchDevicesResponse) SetBody(v *AddDeviceGroupMatchDevicesResponseBody) *AddDeviceGroupMatchDevicesResponse {
	s.Body = v
	return s
}

func (s *AddDeviceGroupMatchDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
