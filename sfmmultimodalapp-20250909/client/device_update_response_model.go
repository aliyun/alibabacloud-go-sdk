// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeviceUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeviceUpdateResponse
	GetStatusCode() *int32
	SetBody(v *DeviceUpdateResponseBody) *DeviceUpdateResponse
	GetBody() *DeviceUpdateResponseBody
}

type DeviceUpdateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeviceUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeviceUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeviceUpdateResponse) GoString() string {
	return s.String()
}

func (s *DeviceUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeviceUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeviceUpdateResponse) GetBody() *DeviceUpdateResponseBody {
	return s.Body
}

func (s *DeviceUpdateResponse) SetHeaders(v map[string]*string) *DeviceUpdateResponse {
	s.Headers = v
	return s
}

func (s *DeviceUpdateResponse) SetStatusCode(v int32) *DeviceUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceUpdateResponse) SetBody(v *DeviceUpdateResponseBody) *DeviceUpdateResponse {
	s.Body = v
	return s
}

func (s *DeviceUpdateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
