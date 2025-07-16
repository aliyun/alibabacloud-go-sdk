// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMeetingRoomResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteMeetingRoomResponseBody
	GetResult() *bool
}

type DeleteMeetingRoomResponseBody struct {
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

func (s DeleteMeetingRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMeetingRoomResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteMeetingRoomResponseBody) SetRequestId(v string) *DeleteMeetingRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMeetingRoomResponseBody) SetResult(v bool) *DeleteMeetingRoomResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteMeetingRoomResponseBody) Validate() error {
	return dara.Validate(s)
}
