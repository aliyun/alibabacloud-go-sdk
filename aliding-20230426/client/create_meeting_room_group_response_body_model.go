// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMeetingRoomGroupResponseBody
	GetRequestId() *string
	SetResult(v int64) *CreateMeetingRoomGroupResponseBody
	GetResult() *int64
}

type CreateMeetingRoomGroupResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 172
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateMeetingRoomGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMeetingRoomGroupResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *CreateMeetingRoomGroupResponseBody) SetRequestId(v string) *CreateMeetingRoomGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMeetingRoomGroupResponseBody) SetResult(v int64) *CreateMeetingRoomGroupResponseBody {
	s.Result = &v
	return s
}

func (s *CreateMeetingRoomGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
