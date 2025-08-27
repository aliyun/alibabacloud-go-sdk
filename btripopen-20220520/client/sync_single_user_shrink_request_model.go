// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncSingleUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *SyncSingleUserShrinkRequest
	GetEmail() *string
	SetJobNo(v string) *SyncSingleUserShrinkRequest
	GetJobNo() *string
	SetLeaveStatus(v int32) *SyncSingleUserShrinkRequest
	GetLeaveStatus() *int32
	SetManagerUserId(v string) *SyncSingleUserShrinkRequest
	GetManagerUserId() *string
	SetPhone(v string) *SyncSingleUserShrinkRequest
	GetPhone() *string
	SetPosition(v string) *SyncSingleUserShrinkRequest
	GetPosition() *string
	SetPositionLevel(v string) *SyncSingleUserShrinkRequest
	GetPositionLevel() *string
	SetRealNameEn(v string) *SyncSingleUserShrinkRequest
	GetRealNameEn() *string
	SetThirdDepartIdListShrink(v string) *SyncSingleUserShrinkRequest
	GetThirdDepartIdListShrink() *string
	SetUserId(v string) *SyncSingleUserShrinkRequest
	GetUserId() *string
	SetUserName(v string) *SyncSingleUserShrinkRequest
	GetUserName() *string
}

type SyncSingleUserShrinkRequest struct {
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 1001
	JobNo *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// example:
	//
	// 0
	LeaveStatus *int32 `json:"leave_status,omitempty" xml:"leave_status,omitempty"`
	// example:
	//
	// 72369
	ManagerUserId *string `json:"manager_user_id,omitempty" xml:"manager_user_id,omitempty"`
	// example:
	//
	// 16392740204
	Phone    *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// 1
	PositionLevel *string `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// example:
	//
	// ce/shi
	RealNameEn              *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
	ThirdDepartIdListShrink *string `json:"third_depart_id_list,omitempty" xml:"third_depart_id_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 573263
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// This parameter is required.
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s SyncSingleUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncSingleUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncSingleUserShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *SyncSingleUserShrinkRequest) GetJobNo() *string {
	return s.JobNo
}

func (s *SyncSingleUserShrinkRequest) GetLeaveStatus() *int32 {
	return s.LeaveStatus
}

func (s *SyncSingleUserShrinkRequest) GetManagerUserId() *string {
	return s.ManagerUserId
}

func (s *SyncSingleUserShrinkRequest) GetPhone() *string {
	return s.Phone
}

func (s *SyncSingleUserShrinkRequest) GetPosition() *string {
	return s.Position
}

func (s *SyncSingleUserShrinkRequest) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *SyncSingleUserShrinkRequest) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *SyncSingleUserShrinkRequest) GetThirdDepartIdListShrink() *string {
	return s.ThirdDepartIdListShrink
}

func (s *SyncSingleUserShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *SyncSingleUserShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *SyncSingleUserShrinkRequest) SetEmail(v string) *SyncSingleUserShrinkRequest {
	s.Email = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetJobNo(v string) *SyncSingleUserShrinkRequest {
	s.JobNo = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetLeaveStatus(v int32) *SyncSingleUserShrinkRequest {
	s.LeaveStatus = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetManagerUserId(v string) *SyncSingleUserShrinkRequest {
	s.ManagerUserId = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetPhone(v string) *SyncSingleUserShrinkRequest {
	s.Phone = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetPosition(v string) *SyncSingleUserShrinkRequest {
	s.Position = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetPositionLevel(v string) *SyncSingleUserShrinkRequest {
	s.PositionLevel = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetRealNameEn(v string) *SyncSingleUserShrinkRequest {
	s.RealNameEn = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetThirdDepartIdListShrink(v string) *SyncSingleUserShrinkRequest {
	s.ThirdDepartIdListShrink = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetUserId(v string) *SyncSingleUserShrinkRequest {
	s.UserId = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) SetUserName(v string) *SyncSingleUserShrinkRequest {
	s.UserName = &v
	return s
}

func (s *SyncSingleUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
