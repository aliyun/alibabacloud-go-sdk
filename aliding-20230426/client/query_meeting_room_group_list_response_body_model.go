// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMeetingRoomGroupListResponseBody
	GetRequestId() *string
	SetResult(v []*QueryMeetingRoomGroupListResponseBodyResult) *QueryMeetingRoomGroupListResponseBody
	GetResult() []*QueryMeetingRoomGroupListResponseBodyResult
}

type QueryMeetingRoomGroupListResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*QueryMeetingRoomGroupListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryMeetingRoomGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMeetingRoomGroupListResponseBody) GetResult() []*QueryMeetingRoomGroupListResponseBodyResult {
	return s.Result
}

func (s *QueryMeetingRoomGroupListResponseBody) SetRequestId(v string) *QueryMeetingRoomGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMeetingRoomGroupListResponseBody) SetResult(v []*QueryMeetingRoomGroupListResponseBodyResult) *QueryMeetingRoomGroupListResponseBody {
	s.Result = v
	return s
}

func (s *QueryMeetingRoomGroupListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomGroupListResponseBodyResult struct {
	// example:
	//
	// 172
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s QueryMeetingRoomGroupListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) GetParentId() *int64 {
	return s.ParentId
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) SetGroupId(v int64) *QueryMeetingRoomGroupListResponseBodyResult {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) SetGroupName(v string) *QueryMeetingRoomGroupListResponseBodyResult {
	s.GroupName = &v
	return s
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) SetParentId(v int64) *QueryMeetingRoomGroupListResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
