// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMeetingRoomListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMeetingRoomListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMeetingRoomListResponseBody) *QueryMeetingRoomListResponse
	GetBody() *QueryMeetingRoomListResponseBody
}

type QueryMeetingRoomListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMeetingRoomListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMeetingRoomListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMeetingRoomListResponse) GetBody() *QueryMeetingRoomListResponseBody {
	return s.Body
}

func (s *QueryMeetingRoomListResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomListResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomListResponse) SetStatusCode(v int32) *QueryMeetingRoomListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomListResponse) SetBody(v *QueryMeetingRoomListResponseBody) *QueryMeetingRoomListResponse {
	s.Body = v
	return s
}

func (s *QueryMeetingRoomListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
