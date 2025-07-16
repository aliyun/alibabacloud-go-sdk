// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMeetingRoomsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddMeetingRoomsResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddMeetingRoomsResponseBody
	GetResult() *bool
}

type AddMeetingRoomsResponseBody struct {
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

func (s AddMeetingRoomsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMeetingRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *AddMeetingRoomsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMeetingRoomsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddMeetingRoomsResponseBody) SetRequestId(v string) *AddMeetingRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMeetingRoomsResponseBody) SetResult(v bool) *AddMeetingRoomsResponseBody {
	s.Result = &v
	return s
}

func (s *AddMeetingRoomsResponseBody) Validate() error {
	return dara.Validate(s)
}
