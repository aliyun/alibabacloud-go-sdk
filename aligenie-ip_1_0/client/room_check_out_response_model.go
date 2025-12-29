// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoomCheckOutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RoomCheckOutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RoomCheckOutResponse
	GetStatusCode() *int32
	SetBody(v *RoomCheckOutResponseBody) *RoomCheckOutResponse
	GetBody() *RoomCheckOutResponseBody
}

type RoomCheckOutResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RoomCheckOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RoomCheckOutResponse) String() string {
	return dara.Prettify(s)
}

func (s RoomCheckOutResponse) GoString() string {
	return s.String()
}

func (s *RoomCheckOutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RoomCheckOutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RoomCheckOutResponse) GetBody() *RoomCheckOutResponseBody {
	return s.Body
}

func (s *RoomCheckOutResponse) SetHeaders(v map[string]*string) *RoomCheckOutResponse {
	s.Headers = v
	return s
}

func (s *RoomCheckOutResponse) SetStatusCode(v int32) *RoomCheckOutResponse {
	s.StatusCode = &v
	return s
}

func (s *RoomCheckOutResponse) SetBody(v *RoomCheckOutResponseBody) *RoomCheckOutResponse {
	s.Body = v
	return s
}

func (s *RoomCheckOutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
