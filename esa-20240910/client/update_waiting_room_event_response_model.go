// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWaitingRoomEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWaitingRoomEventResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWaitingRoomEventResponseBody) *UpdateWaitingRoomEventResponse
	GetBody() *UpdateWaitingRoomEventResponseBody
}

type UpdateWaitingRoomEventResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWaitingRoomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWaitingRoomEventResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomEventResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWaitingRoomEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWaitingRoomEventResponse) GetBody() *UpdateWaitingRoomEventResponseBody {
	return s.Body
}

func (s *UpdateWaitingRoomEventResponse) SetHeaders(v map[string]*string) *UpdateWaitingRoomEventResponse {
	s.Headers = v
	return s
}

func (s *UpdateWaitingRoomEventResponse) SetStatusCode(v int32) *UpdateWaitingRoomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWaitingRoomEventResponse) SetBody(v *UpdateWaitingRoomEventResponseBody) *UpdateWaitingRoomEventResponse {
	s.Body = v
	return s
}

func (s *UpdateWaitingRoomEventResponse) Validate() error {
	return dara.Validate(s)
}
