// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaitingRoomEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWaitingRoomEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWaitingRoomEventResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWaitingRoomEventResponseBody) *DeleteWaitingRoomEventResponse
	GetBody() *DeleteWaitingRoomEventResponseBody
}

type DeleteWaitingRoomEventResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWaitingRoomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWaitingRoomEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaitingRoomEventResponse) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWaitingRoomEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWaitingRoomEventResponse) GetBody() *DeleteWaitingRoomEventResponseBody {
	return s.Body
}

func (s *DeleteWaitingRoomEventResponse) SetHeaders(v map[string]*string) *DeleteWaitingRoomEventResponse {
	s.Headers = v
	return s
}

func (s *DeleteWaitingRoomEventResponse) SetStatusCode(v int32) *DeleteWaitingRoomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWaitingRoomEventResponse) SetBody(v *DeleteWaitingRoomEventResponseBody) *DeleteWaitingRoomEventResponse {
	s.Body = v
	return s
}

func (s *DeleteWaitingRoomEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
