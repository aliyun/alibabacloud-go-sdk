// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWaitingRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWaitingRoomResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWaitingRoomResponseBody) *UpdateWaitingRoomResponse
	GetBody() *UpdateWaitingRoomResponseBody
}

type UpdateWaitingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWaitingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWaitingRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWaitingRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWaitingRoomResponse) GetBody() *UpdateWaitingRoomResponseBody {
	return s.Body
}

func (s *UpdateWaitingRoomResponse) SetHeaders(v map[string]*string) *UpdateWaitingRoomResponse {
	s.Headers = v
	return s
}

func (s *UpdateWaitingRoomResponse) SetStatusCode(v int32) *UpdateWaitingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWaitingRoomResponse) SetBody(v *UpdateWaitingRoomResponseBody) *UpdateWaitingRoomResponse {
	s.Body = v
	return s
}

func (s *UpdateWaitingRoomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
