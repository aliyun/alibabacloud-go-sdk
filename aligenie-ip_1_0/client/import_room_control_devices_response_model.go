// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomControlDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportRoomControlDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportRoomControlDevicesResponse
	GetStatusCode() *int32
	SetBody(v *ImportRoomControlDevicesResponseBody) *ImportRoomControlDevicesResponse
	GetBody() *ImportRoomControlDevicesResponseBody
}

type ImportRoomControlDevicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportRoomControlDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportRoomControlDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomControlDevicesResponse) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportRoomControlDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportRoomControlDevicesResponse) GetBody() *ImportRoomControlDevicesResponseBody {
	return s.Body
}

func (s *ImportRoomControlDevicesResponse) SetHeaders(v map[string]*string) *ImportRoomControlDevicesResponse {
	s.Headers = v
	return s
}

func (s *ImportRoomControlDevicesResponse) SetStatusCode(v int32) *ImportRoomControlDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportRoomControlDevicesResponse) SetBody(v *ImportRoomControlDevicesResponseBody) *ImportRoomControlDevicesResponse {
	s.Body = v
	return s
}

func (s *ImportRoomControlDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
