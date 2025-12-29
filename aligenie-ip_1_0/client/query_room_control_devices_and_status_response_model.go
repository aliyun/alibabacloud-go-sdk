// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomControlDevicesAndStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRoomControlDevicesAndStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRoomControlDevicesAndStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryRoomControlDevicesAndStatusResponseBody) *QueryRoomControlDevicesAndStatusResponse
	GetBody() *QueryRoomControlDevicesAndStatusResponseBody
}

type QueryRoomControlDevicesAndStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRoomControlDevicesAndStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRoomControlDevicesAndStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRoomControlDevicesAndStatusResponse) GetBody() *QueryRoomControlDevicesAndStatusResponseBody {
	return s.Body
}

func (s *QueryRoomControlDevicesAndStatusResponse) SetHeaders(v map[string]*string) *QueryRoomControlDevicesAndStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponse) SetStatusCode(v int32) *QueryRoomControlDevicesAndStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponse) SetBody(v *QueryRoomControlDevicesAndStatusResponseBody) *QueryRoomControlDevicesAndStatusResponse {
	s.Body = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
