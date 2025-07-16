// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *QueryMeetingRoomGroupResponseBody
	GetGroupId() *int64
	SetGroupName(v string) *QueryMeetingRoomGroupResponseBody
	GetGroupName() *string
	SetParentId(v int64) *QueryMeetingRoomGroupResponseBody
	GetParentId() *int64
	SetRequestId(v string) *QueryMeetingRoomGroupResponseBody
	GetRequestId() *string
}

type QueryMeetingRoomGroupResponseBody struct {
	// example:
	//
	// 172
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryMeetingRoomGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupResponseBody) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QueryMeetingRoomGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryMeetingRoomGroupResponseBody) GetParentId() *int64 {
	return s.ParentId
}

func (s *QueryMeetingRoomGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMeetingRoomGroupResponseBody) SetGroupId(v int64) *QueryMeetingRoomGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomGroupResponseBody) SetGroupName(v string) *QueryMeetingRoomGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *QueryMeetingRoomGroupResponseBody) SetParentId(v int64) *QueryMeetingRoomGroupResponseBody {
	s.ParentId = &v
	return s
}

func (s *QueryMeetingRoomGroupResponseBody) SetRequestId(v string) *QueryMeetingRoomGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMeetingRoomGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
