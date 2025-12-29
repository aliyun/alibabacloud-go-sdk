// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iControlRoomDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ControlRoomDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ControlRoomDeviceResponse
	GetStatusCode() *int32
	SetBody(v *ControlRoomDeviceResponseBody) *ControlRoomDeviceResponse
	GetBody() *ControlRoomDeviceResponseBody
}

type ControlRoomDeviceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ControlRoomDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ControlRoomDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s ControlRoomDeviceResponse) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ControlRoomDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ControlRoomDeviceResponse) GetBody() *ControlRoomDeviceResponseBody {
	return s.Body
}

func (s *ControlRoomDeviceResponse) SetHeaders(v map[string]*string) *ControlRoomDeviceResponse {
	s.Headers = v
	return s
}

func (s *ControlRoomDeviceResponse) SetStatusCode(v int32) *ControlRoomDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ControlRoomDeviceResponse) SetBody(v *ControlRoomDeviceResponseBody) *ControlRoomDeviceResponse {
	s.Body = v
	return s
}

func (s *ControlRoomDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
