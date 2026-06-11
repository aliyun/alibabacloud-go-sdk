// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDevicePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DevicePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DevicePageResponse
	GetStatusCode() *int32
	SetBody(v *DevicePageResponseBody) *DevicePageResponse
	GetBody() *DevicePageResponseBody
}

type DevicePageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DevicePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DevicePageResponse) String() string {
	return dara.Prettify(s)
}

func (s DevicePageResponse) GoString() string {
	return s.String()
}

func (s *DevicePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DevicePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DevicePageResponse) GetBody() *DevicePageResponseBody {
	return s.Body
}

func (s *DevicePageResponse) SetHeaders(v map[string]*string) *DevicePageResponse {
	s.Headers = v
	return s
}

func (s *DevicePageResponse) SetStatusCode(v int32) *DevicePageResponse {
	s.StatusCode = &v
	return s
}

func (s *DevicePageResponse) SetBody(v *DevicePageResponseBody) *DevicePageResponse {
	s.Body = v
	return s
}

func (s *DevicePageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
