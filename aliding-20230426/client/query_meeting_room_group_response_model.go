// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMeetingRoomGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMeetingRoomGroupResponse
	GetStatusCode() *int32
	SetBody(v *QueryMeetingRoomGroupResponseBody) *QueryMeetingRoomGroupResponse
	GetBody() *QueryMeetingRoomGroupResponseBody
}

type QueryMeetingRoomGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMeetingRoomGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMeetingRoomGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMeetingRoomGroupResponse) GetBody() *QueryMeetingRoomGroupResponseBody {
	return s.Body
}

func (s *QueryMeetingRoomGroupResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomGroupResponse) SetStatusCode(v int32) *QueryMeetingRoomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomGroupResponse) SetBody(v *QueryMeetingRoomGroupResponseBody) *QueryMeetingRoomGroupResponse {
	s.Body = v
	return s
}

func (s *QueryMeetingRoomGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
