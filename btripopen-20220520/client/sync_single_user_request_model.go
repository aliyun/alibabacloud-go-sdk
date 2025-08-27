// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncSingleUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *SyncSingleUserRequest
	GetEmail() *string
	SetJobNo(v string) *SyncSingleUserRequest
	GetJobNo() *string
	SetLeaveStatus(v int32) *SyncSingleUserRequest
	GetLeaveStatus() *int32
	SetManagerUserId(v string) *SyncSingleUserRequest
	GetManagerUserId() *string
	SetPhone(v string) *SyncSingleUserRequest
	GetPhone() *string
	SetPosition(v string) *SyncSingleUserRequest
	GetPosition() *string
	SetPositionLevel(v string) *SyncSingleUserRequest
	GetPositionLevel() *string
	SetRealNameEn(v string) *SyncSingleUserRequest
	GetRealNameEn() *string
	SetThirdDepartIdList(v []*string) *SyncSingleUserRequest
	GetThirdDepartIdList() []*string
	SetUserId(v string) *SyncSingleUserRequest
	GetUserId() *string
	SetUserName(v string) *SyncSingleUserRequest
	GetUserName() *string
}

type SyncSingleUserRequest struct {
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
	RealNameEn        *string   `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
	ThirdDepartIdList []*string `json:"third_depart_id_list,omitempty" xml:"third_depart_id_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 573263
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// This parameter is required.
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s SyncSingleUserRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncSingleUserRequest) GoString() string {
	return s.String()
}

func (s *SyncSingleUserRequest) GetEmail() *string {
	return s.Email
}

func (s *SyncSingleUserRequest) GetJobNo() *string {
	return s.JobNo
}

func (s *SyncSingleUserRequest) GetLeaveStatus() *int32 {
	return s.LeaveStatus
}

func (s *SyncSingleUserRequest) GetManagerUserId() *string {
	return s.ManagerUserId
}

func (s *SyncSingleUserRequest) GetPhone() *string {
	return s.Phone
}

func (s *SyncSingleUserRequest) GetPosition() *string {
	return s.Position
}

func (s *SyncSingleUserRequest) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *SyncSingleUserRequest) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *SyncSingleUserRequest) GetThirdDepartIdList() []*string {
	return s.ThirdDepartIdList
}

func (s *SyncSingleUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *SyncSingleUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *SyncSingleUserRequest) SetEmail(v string) *SyncSingleUserRequest {
	s.Email = &v
	return s
}

func (s *SyncSingleUserRequest) SetJobNo(v string) *SyncSingleUserRequest {
	s.JobNo = &v
	return s
}

func (s *SyncSingleUserRequest) SetLeaveStatus(v int32) *SyncSingleUserRequest {
	s.LeaveStatus = &v
	return s
}

func (s *SyncSingleUserRequest) SetManagerUserId(v string) *SyncSingleUserRequest {
	s.ManagerUserId = &v
	return s
}

func (s *SyncSingleUserRequest) SetPhone(v string) *SyncSingleUserRequest {
	s.Phone = &v
	return s
}

func (s *SyncSingleUserRequest) SetPosition(v string) *SyncSingleUserRequest {
	s.Position = &v
	return s
}

func (s *SyncSingleUserRequest) SetPositionLevel(v string) *SyncSingleUserRequest {
	s.PositionLevel = &v
	return s
}

func (s *SyncSingleUserRequest) SetRealNameEn(v string) *SyncSingleUserRequest {
	s.RealNameEn = &v
	return s
}

func (s *SyncSingleUserRequest) SetThirdDepartIdList(v []*string) *SyncSingleUserRequest {
	s.ThirdDepartIdList = v
	return s
}

func (s *SyncSingleUserRequest) SetUserId(v string) *SyncSingleUserRequest {
	s.UserId = &v
	return s
}

func (s *SyncSingleUserRequest) SetUserName(v string) *SyncSingleUserRequest {
	s.UserName = &v
	return s
}

func (s *SyncSingleUserRequest) Validate() error {
	return dara.Validate(s)
}
