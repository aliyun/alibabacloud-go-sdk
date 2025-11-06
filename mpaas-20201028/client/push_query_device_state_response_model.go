// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushQueryDeviceStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushQueryDeviceStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushQueryDeviceStateResponse
	GetStatusCode() *int32
	SetBody(v *PushQueryDeviceStateResponseBody) *PushQueryDeviceStateResponse
	GetBody() *PushQueryDeviceStateResponseBody
}

type PushQueryDeviceStateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushQueryDeviceStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushQueryDeviceStateResponse) String() string {
	return dara.Prettify(s)
}

func (s PushQueryDeviceStateResponse) GoString() string {
	return s.String()
}

func (s *PushQueryDeviceStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushQueryDeviceStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushQueryDeviceStateResponse) GetBody() *PushQueryDeviceStateResponseBody {
	return s.Body
}

func (s *PushQueryDeviceStateResponse) SetHeaders(v map[string]*string) *PushQueryDeviceStateResponse {
	s.Headers = v
	return s
}

func (s *PushQueryDeviceStateResponse) SetStatusCode(v int32) *PushQueryDeviceStateResponse {
	s.StatusCode = &v
	return s
}

func (s *PushQueryDeviceStateResponse) SetBody(v *PushQueryDeviceStateResponseBody) *PushQueryDeviceStateResponse {
	s.Body = v
	return s
}

func (s *PushQueryDeviceStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
