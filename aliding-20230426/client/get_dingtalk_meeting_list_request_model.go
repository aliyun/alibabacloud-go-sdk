// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetDingtalkMeetingListRequestTenantContext) *GetDingtalkMeetingListRequest
	GetTenantContext() *GetDingtalkMeetingListRequestTenantContext
	SetCurrentPage(v int32) *GetDingtalkMeetingListRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *GetDingtalkMeetingListRequest
	GetEndTime() *int64
	SetOrgId(v string) *GetDingtalkMeetingListRequest
	GetOrgId() *string
	SetPageSize(v int32) *GetDingtalkMeetingListRequest
	GetPageSize() *int32
	SetRoomCode(v string) *GetDingtalkMeetingListRequest
	GetRoomCode() *string
	SetRoomName(v string) *GetDingtalkMeetingListRequest
	GetRoomName() *string
	SetStartTime(v int64) *GetDingtalkMeetingListRequest
	GetStartTime() *int64
	SetWorkNo(v string) *GetDingtalkMeetingListRequest
	GetWorkNo() *string
}

type GetDingtalkMeetingListRequest struct {
	TenantContext *GetDingtalkMeetingListRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s GetDingtalkMeetingListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingListRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingListRequest) GetTenantContext() *GetDingtalkMeetingListRequestTenantContext {
	return s.TenantContext
}

func (s *GetDingtalkMeetingListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDingtalkMeetingListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDingtalkMeetingListRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDingtalkMeetingListRequest) GetRoomCode() *string {
	return s.RoomCode
}

func (s *GetDingtalkMeetingListRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *GetDingtalkMeetingListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDingtalkMeetingListRequest) GetWorkNo() *string {
	return s.WorkNo
}

func (s *GetDingtalkMeetingListRequest) SetTenantContext(v *GetDingtalkMeetingListRequestTenantContext) *GetDingtalkMeetingListRequest {
	s.TenantContext = v
	return s
}

func (s *GetDingtalkMeetingListRequest) SetCurrentPage(v int32) *GetDingtalkMeetingListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDingtalkMeetingListRequest) SetEndTime(v int64) *GetDingtalkMeetingListRequest {
	s.EndTime = &v
	return s
}

func (s *GetDingtalkMeetingListRequest) SetOrgId(v string) *GetDingtalkMeetingListRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingListRequest) SetPageSize(v int32) *GetDingtalkMeetingListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDingtalkMeetingListRequest) SetRoomCode(v string) *GetDingtalkMeetingListRequest {
	s.RoomCode = &v
	return s
}

func (s *GetDingtalkMeetingListRequest) SetRoomName(v string) *GetDingtalkMeetingListRequest {
	s.RoomName = &v
	return s
}

func (s *GetDingtalkMeetingListRequest) SetStartTime(v int64) *GetDingtalkMeetingListRequest {
	s.StartTime = &v
	return s
}

func (s *GetDingtalkMeetingListRequest) SetWorkNo(v string) *GetDingtalkMeetingListRequest {
	s.WorkNo = &v
	return s
}

func (s *GetDingtalkMeetingListRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkMeetingListRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDingtalkMeetingListRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingListRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingListRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDingtalkMeetingListRequestTenantContext) SetTenantId(v string) *GetDingtalkMeetingListRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDingtalkMeetingListRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
