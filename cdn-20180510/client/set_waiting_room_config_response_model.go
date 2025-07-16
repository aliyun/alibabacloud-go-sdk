// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWaitingRoomConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWaitingRoomConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWaitingRoomConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetWaitingRoomConfigResponseBody) *SetWaitingRoomConfigResponse
	GetBody() *SetWaitingRoomConfigResponseBody
}

type SetWaitingRoomConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWaitingRoomConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWaitingRoomConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWaitingRoomConfigResponse) GoString() string {
	return s.String()
}

func (s *SetWaitingRoomConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWaitingRoomConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWaitingRoomConfigResponse) GetBody() *SetWaitingRoomConfigResponseBody {
	return s.Body
}

func (s *SetWaitingRoomConfigResponse) SetHeaders(v map[string]*string) *SetWaitingRoomConfigResponse {
	s.Headers = v
	return s
}

func (s *SetWaitingRoomConfigResponse) SetStatusCode(v int32) *SetWaitingRoomConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWaitingRoomConfigResponse) SetBody(v *SetWaitingRoomConfigResponseBody) *SetWaitingRoomConfigResponse {
	s.Body = v
	return s
}

func (s *SetWaitingRoomConfigResponse) Validate() error {
	return dara.Validate(s)
}
