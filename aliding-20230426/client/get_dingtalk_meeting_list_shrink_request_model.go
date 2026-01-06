// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetDingtalkMeetingListShrinkRequest
	GetTenantContextShrink() *string
	SetCurrentPage(v int32) *GetDingtalkMeetingListShrinkRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *GetDingtalkMeetingListShrinkRequest
	GetEndTime() *int64
	SetOrgId(v string) *GetDingtalkMeetingListShrinkRequest
	GetOrgId() *string
	SetPageSize(v int32) *GetDingtalkMeetingListShrinkRequest
	GetPageSize() *int32
	SetRoomCode(v string) *GetDingtalkMeetingListShrinkRequest
	GetRoomCode() *string
	SetRoomName(v string) *GetDingtalkMeetingListShrinkRequest
	GetRoomName() *string
	SetStartTime(v int64) *GetDingtalkMeetingListShrinkRequest
	GetStartTime() *int64
	SetWorkNo(v string) *GetDingtalkMeetingListShrinkRequest
	GetWorkNo() *string
}

type GetDingtalkMeetingListShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 1732953600000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 21001
	OrgId *string `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 123456789
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	// example:
	//
	// 会议室A
	RoomName *string `json:"roomName,omitempty" xml:"roomName,omitempty"`
	// example:
	//
	// 1732867200000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 34343
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetDingtalkMeetingListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingListShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDingtalkMeetingListShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDingtalkMeetingListShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDingtalkMeetingListShrinkRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDingtalkMeetingListShrinkRequest) GetRoomCode() *string {
	return s.RoomCode
}

func (s *GetDingtalkMeetingListShrinkRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *GetDingtalkMeetingListShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDingtalkMeetingListShrinkRequest) GetWorkNo() *string {
	return s.WorkNo
}

func (s *GetDingtalkMeetingListShrinkRequest) SetTenantContextShrink(v string) *GetDingtalkMeetingListShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkRequest) SetCurrentPage(v int32) *GetDingtalkMeetingListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkRequest) SetEndTime(v int64) *GetDingtalkMeetingListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkRequest) SetOrgId(v string) *GetDingtalkMeetingListShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkRequest) SetPageSize(v int32) *GetDingtalkMeetingListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkRequest) SetRoomCode(v string) *GetDingtalkMeetingListShrinkRequest {
	s.RoomCode = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkRequest) SetRoomName(v string) *GetDingtalkMeetingListShrinkRequest {
	s.RoomName = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkRequest) SetStartTime(v int64) *GetDingtalkMeetingListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkRequest) SetWorkNo(v string) *GetDingtalkMeetingListShrinkRequest {
	s.WorkNo = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
