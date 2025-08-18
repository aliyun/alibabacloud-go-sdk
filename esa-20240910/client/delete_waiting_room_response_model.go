// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaitingRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWaitingRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWaitingRoomResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWaitingRoomResponseBody) *DeleteWaitingRoomResponse
	GetBody() *DeleteWaitingRoomResponseBody
}

type DeleteWaitingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWaitingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWaitingRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaitingRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWaitingRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWaitingRoomResponse) GetBody() *DeleteWaitingRoomResponseBody {
	return s.Body
}

func (s *DeleteWaitingRoomResponse) SetHeaders(v map[string]*string) *DeleteWaitingRoomResponse {
	s.Headers = v
	return s
}

func (s *DeleteWaitingRoomResponse) SetStatusCode(v int32) *DeleteWaitingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWaitingRoomResponse) SetBody(v *DeleteWaitingRoomResponseBody) *DeleteWaitingRoomResponse {
	s.Body = v
	return s
}

func (s *DeleteWaitingRoomResponse) Validate() error {
	return dara.Validate(s)
}
