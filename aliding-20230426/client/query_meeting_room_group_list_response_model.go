// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMeetingRoomGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMeetingRoomGroupListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMeetingRoomGroupListResponseBody) *QueryMeetingRoomGroupListResponse
	GetBody() *QueryMeetingRoomGroupListResponseBody
}

type QueryMeetingRoomGroupListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMeetingRoomGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupListResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMeetingRoomGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMeetingRoomGroupListResponse) GetBody() *QueryMeetingRoomGroupListResponseBody {
	return s.Body
}

func (s *QueryMeetingRoomGroupListResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomGroupListResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomGroupListResponse) SetStatusCode(v int32) *QueryMeetingRoomGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomGroupListResponse) SetBody(v *QueryMeetingRoomGroupListResponseBody) *QueryMeetingRoomGroupListResponse {
	s.Body = v
	return s
}

func (s *QueryMeetingRoomGroupListResponse) Validate() error {
	return dara.Validate(s)
}
