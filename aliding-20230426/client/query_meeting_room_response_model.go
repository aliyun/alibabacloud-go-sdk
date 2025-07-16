// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMeetingRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMeetingRoomResponse
	GetStatusCode() *int32
	SetBody(v *QueryMeetingRoomResponseBody) *QueryMeetingRoomResponse
	GetBody() *QueryMeetingRoomResponseBody
}

type QueryMeetingRoomResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMeetingRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMeetingRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMeetingRoomResponse) GetBody() *QueryMeetingRoomResponseBody {
	return s.Body
}

func (s *QueryMeetingRoomResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomResponse) SetStatusCode(v int32) *QueryMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomResponse) SetBody(v *QueryMeetingRoomResponseBody) *QueryMeetingRoomResponse {
	s.Body = v
	return s
}

func (s *QueryMeetingRoomResponse) Validate() error {
	return dara.Validate(s)
}
