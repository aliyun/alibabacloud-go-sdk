// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMeetingRoomGroupResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateMeetingRoomGroupResponseBody
	GetResult() *bool
}

type UpdateMeetingRoomGroupResponseBody struct {
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

func (s UpdateMeetingRoomGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMeetingRoomGroupResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateMeetingRoomGroupResponseBody) SetRequestId(v string) *UpdateMeetingRoomGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMeetingRoomGroupResponseBody) SetResult(v bool) *UpdateMeetingRoomGroupResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateMeetingRoomGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
