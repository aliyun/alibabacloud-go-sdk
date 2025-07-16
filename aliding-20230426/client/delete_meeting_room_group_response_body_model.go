// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMeetingRoomGroupResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteMeetingRoomGroupResponseBody
	GetResult() *bool
}

type DeleteMeetingRoomGroupResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteMeetingRoomGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMeetingRoomGroupResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteMeetingRoomGroupResponseBody) SetRequestId(v string) *DeleteMeetingRoomGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMeetingRoomGroupResponseBody) SetResult(v bool) *DeleteMeetingRoomGroupResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteMeetingRoomGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
