// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWaitingRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWaitingRoomResponse
	GetStatusCode() *int32
	SetBody(v *CreateWaitingRoomResponseBody) *CreateWaitingRoomResponse
	GetBody() *CreateWaitingRoomResponseBody
}

type CreateWaitingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWaitingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWaitingRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWaitingRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWaitingRoomResponse) GetBody() *CreateWaitingRoomResponseBody {
	return s.Body
}

func (s *CreateWaitingRoomResponse) SetHeaders(v map[string]*string) *CreateWaitingRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateWaitingRoomResponse) SetStatusCode(v int32) *CreateWaitingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWaitingRoomResponse) SetBody(v *CreateWaitingRoomResponseBody) *CreateWaitingRoomResponse {
	s.Body = v
	return s
}

func (s *CreateWaitingRoomResponse) Validate() error {
	return dara.Validate(s)
}
