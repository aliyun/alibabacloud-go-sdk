// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomControlDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRoomControlDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRoomControlDevicesResponse
	GetStatusCode() *int32
	SetBody(v *QueryRoomControlDevicesResponseBody) *QueryRoomControlDevicesResponse
	GetBody() *QueryRoomControlDevicesResponseBody
}

type QueryRoomControlDevicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRoomControlDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRoomControlDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesResponse) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRoomControlDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRoomControlDevicesResponse) GetBody() *QueryRoomControlDevicesResponseBody {
	return s.Body
}

func (s *QueryRoomControlDevicesResponse) SetHeaders(v map[string]*string) *QueryRoomControlDevicesResponse {
	s.Headers = v
	return s
}

func (s *QueryRoomControlDevicesResponse) SetStatusCode(v int32) *QueryRoomControlDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRoomControlDevicesResponse) SetBody(v *QueryRoomControlDevicesResponseBody) *QueryRoomControlDevicesResponse {
	s.Body = v
	return s
}

func (s *QueryRoomControlDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
