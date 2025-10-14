// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWaitingRoomEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWaitingRoomEventResponse
	GetStatusCode() *int32
	SetBody(v *CreateWaitingRoomEventResponseBody) *CreateWaitingRoomEventResponse
	GetBody() *CreateWaitingRoomEventResponseBody
}

type CreateWaitingRoomEventResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWaitingRoomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWaitingRoomEventResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomEventResponse) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWaitingRoomEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWaitingRoomEventResponse) GetBody() *CreateWaitingRoomEventResponseBody {
	return s.Body
}

func (s *CreateWaitingRoomEventResponse) SetHeaders(v map[string]*string) *CreateWaitingRoomEventResponse {
	s.Headers = v
	return s
}

func (s *CreateWaitingRoomEventResponse) SetStatusCode(v int32) *CreateWaitingRoomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWaitingRoomEventResponse) SetBody(v *CreateWaitingRoomEventResponseBody) *CreateWaitingRoomEventResponse {
	s.Body = v
	return s
}

func (s *CreateWaitingRoomEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
