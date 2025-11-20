// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *QueryMeetingRoomListResponseBody
	GetHasMore() *bool
	SetNextToken(v int64) *QueryMeetingRoomListResponseBody
	GetNextToken() *int64
	SetRequestId(v string) *QueryMeetingRoomListResponseBody
	GetRequestId() *string
	SetResult(v []*QueryMeetingRoomListResponseBodyResult) *QueryMeetingRoomListResponseBody
	GetResult() []*QueryMeetingRoomListResponseBodyResult
}

type QueryMeetingRoomListResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 123
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*QueryMeetingRoomListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryMeetingRoomListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *QueryMeetingRoomListResponseBody) GetNextToken() *int64 {
	return s.NextToken
}

func (s *QueryMeetingRoomListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMeetingRoomListResponseBody) GetResult() []*QueryMeetingRoomListResponseBodyResult {
	return s.Result
}

func (s *QueryMeetingRoomListResponseBody) SetHasMore(v bool) *QueryMeetingRoomListResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryMeetingRoomListResponseBody) SetNextToken(v int64) *QueryMeetingRoomListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryMeetingRoomListResponseBody) SetRequestId(v string) *QueryMeetingRoomListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBody) SetResult(v []*QueryMeetingRoomListResponseBodyResult) *QueryMeetingRoomListResponseBody {
	s.Result = v
	return s
}

func (s *QueryMeetingRoomListResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMeetingRoomListResponseBodyResult struct {
	// example:
	//
	// ding994axxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId *string `json:"IsvRoomId,omitempty" xml:"IsvRoomId,omitempty"`
	// example:
	//
	// 10
	RoomCapacity *int32                                           `json:"RoomCapacity,omitempty" xml:"RoomCapacity,omitempty"`
	RoomGroup    *QueryMeetingRoomListResponseBodyResultRoomGroup `json:"RoomGroup,omitempty" xml:"RoomGroup,omitempty" type:"Struct"`
	// example:
	//
	// 0ffb7xxxxx
	RoomId       *string                                             `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	RoomLabels   []*QueryMeetingRoomListResponseBodyResultRoomLabels `json:"RoomLabels,omitempty" xml:"RoomLabels,omitempty" type:"Repeated"`
	RoomLocation *QueryMeetingRoomListResponseBodyResultRoomLocation `json:"RoomLocation,omitempty" xml:"RoomLocation,omitempty" type:"Struct"`
	// example:
	//
	// 测试会议室
	RoomName *string `json:"RoomName,omitempty" xml:"RoomName,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/lADPxxxxx.jpg
	RoomPicture *string `json:"RoomPicture,omitempty" xml:"RoomPicture,omitempty"`
	// example:
	//
	// 012241xxxxx
	RoomStaffId *string `json:"RoomStaffId,omitempty" xml:"RoomStaffId,omitempty"`
	// example:
	//
	// 0
	RoomStatus *int32 `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty"`
}

func (s QueryMeetingRoomListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBodyResult) GetCorpId() *string {
	return s.CorpId
}

func (s *QueryMeetingRoomListResponseBodyResult) GetIsvRoomId() *string {
	return s.IsvRoomId
}

func (s *QueryMeetingRoomListResponseBodyResult) GetRoomCapacity() *int32 {
	return s.RoomCapacity
}

func (s *QueryMeetingRoomListResponseBodyResult) GetRoomGroup() *QueryMeetingRoomListResponseBodyResultRoomGroup {
	return s.RoomGroup
}

func (s *QueryMeetingRoomListResponseBodyResult) GetRoomId() *string {
	return s.RoomId
}

func (s *QueryMeetingRoomListResponseBodyResult) GetRoomLabels() []*QueryMeetingRoomListResponseBodyResultRoomLabels {
	return s.RoomLabels
}

func (s *QueryMeetingRoomListResponseBodyResult) GetRoomLocation() *QueryMeetingRoomListResponseBodyResultRoomLocation {
	return s.RoomLocation
}

func (s *QueryMeetingRoomListResponseBodyResult) GetRoomName() *string {
	return s.RoomName
}

func (s *QueryMeetingRoomListResponseBodyResult) GetRoomPicture() *string {
	return s.RoomPicture
}

func (s *QueryMeetingRoomListResponseBodyResult) GetRoomStaffId() *string {
	return s.RoomStaffId
}

func (s *QueryMeetingRoomListResponseBodyResult) GetRoomStatus() *int32 {
	return s.RoomStatus
}

func (s *QueryMeetingRoomListResponseBodyResult) SetCorpId(v string) *QueryMeetingRoomListResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetIsvRoomId(v string) *QueryMeetingRoomListResponseBodyResult {
	s.IsvRoomId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomCapacity(v int32) *QueryMeetingRoomListResponseBodyResult {
	s.RoomCapacity = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomGroup(v *QueryMeetingRoomListResponseBodyResultRoomGroup) *QueryMeetingRoomListResponseBodyResult {
	s.RoomGroup = v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomId(v string) *QueryMeetingRoomListResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomLabels(v []*QueryMeetingRoomListResponseBodyResultRoomLabels) *QueryMeetingRoomListResponseBodyResult {
	s.RoomLabels = v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomLocation(v *QueryMeetingRoomListResponseBodyResultRoomLocation) *QueryMeetingRoomListResponseBodyResult {
	s.RoomLocation = v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomName(v string) *QueryMeetingRoomListResponseBodyResult {
	s.RoomName = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomPicture(v string) *QueryMeetingRoomListResponseBodyResult {
	s.RoomPicture = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomStaffId(v string) *QueryMeetingRoomListResponseBodyResult {
	s.RoomStaffId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomStatus(v int32) *QueryMeetingRoomListResponseBodyResult {
	s.RoomStatus = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) Validate() error {
	if s.RoomGroup != nil {
		if err := s.RoomGroup.Validate(); err != nil {
			return err
		}
	}
	if s.RoomLabels != nil {
		for _, item := range s.RoomLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RoomLocation != nil {
		if err := s.RoomLocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMeetingRoomListResponseBodyResultRoomGroup struct {
	// example:
	//
	// 1
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

func (s QueryMeetingRoomListResponseBodyResultRoomGroup) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListResponseBodyResultRoomGroup) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) GetParentId() *int64 {
	return s.ParentId
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) SetGroupId(v int64) *QueryMeetingRoomListResponseBodyResultRoomGroup {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) SetGroupName(v string) *QueryMeetingRoomListResponseBodyResultRoomGroup {
	s.GroupName = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) SetParentId(v int64) *QueryMeetingRoomListResponseBodyResultRoomGroup {
	s.ParentId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomListResponseBodyResultRoomLabels struct {
	// example:
	//
	// 1
	LabelId *int64 `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	// example:
	//
	// 电视
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
}

func (s QueryMeetingRoomListResponseBodyResultRoomLabels) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListResponseBodyResultRoomLabels) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLabels) GetLabelId() *int64 {
	return s.LabelId
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLabels) GetLabelName() *string {
	return s.LabelName
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLabels) SetLabelId(v int64) *QueryMeetingRoomListResponseBodyResultRoomLabels {
	s.LabelId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLabels) SetLabelName(v string) *QueryMeetingRoomListResponseBodyResultRoomLabels {
	s.LabelName = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLabels) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomListResponseBodyResultRoomLocation struct {
	// example:
	//
	// xx市xx区xx街道xx号
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// xxx公司
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryMeetingRoomListResponseBodyResultRoomLocation) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListResponseBodyResultRoomLocation) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLocation) GetDesc() *string {
	return s.Desc
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLocation) GetTitle() *string {
	return s.Title
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLocation) SetDesc(v string) *QueryMeetingRoomListResponseBodyResultRoomLocation {
	s.Desc = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLocation) SetTitle(v string) *QueryMeetingRoomListResponseBodyResultRoomLocation {
	s.Title = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLocation) Validate() error {
	return dara.Validate(s)
}
