// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvUserSaveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserList(v []*IsvUserSaveRequestUserList) *IsvUserSaveRequest
	GetUserList() []*IsvUserSaveRequestUserList
}

type IsvUserSaveRequest struct {
	UserList []*IsvUserSaveRequestUserList `json:"user_list,omitempty" xml:"user_list,omitempty" type:"Repeated"`
}

func (s IsvUserSaveRequest) String() string {
	return dara.Prettify(s)
}

func (s IsvUserSaveRequest) GoString() string {
	return s.String()
}

func (s *IsvUserSaveRequest) GetUserList() []*IsvUserSaveRequestUserList {
	return s.UserList
}

func (s *IsvUserSaveRequest) SetUserList(v []*IsvUserSaveRequestUserList) *IsvUserSaveRequest {
	s.UserList = v
	return s
}

func (s *IsvUserSaveRequest) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IsvUserSaveRequestUserList struct {
	// example:
	//
	// 330100,310100
	BaseCityCode *string `json:"base_city_code,omitempty" xml:"base_city_code,omitempty"`
	// example:
	//
	// 2000-01-01
	Birthday *string                               `json:"birthday,omitempty" xml:"birthday,omitempty"`
	CertList []*IsvUserSaveRequestUserListCertList `json:"cert_list,omitempty" xml:"cert_list,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	DepartId *int64 `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// example:
	//
	// 123@163.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// M
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// true
	IsAdmin *bool `json:"is_admin,omitempty" xml:"is_admin,omitempty"`
	// example:
	//
	// 8797
	JobNo *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// example:
	//
	// 0
	LeaveStatus *int32 `json:"leave_status,omitempty" xml:"leave_status,omitempty"`
	// example:
	//
	// 123456
	ManagerUserId *string `json:"manager_user_id,omitempty" xml:"manager_user_id,omitempty"`
	// example:
	//
	// 15364762829
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// 产品经理
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// 20
	PositionLevel *string `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// example:
	//
	// ceshi
	RealNameEn *string   `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
	RoleIdList []*string `json:"role_id_list,omitempty" xml:"role_id_list,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	ThirdDepartId     *string   `json:"third_depart_id,omitempty" xml:"third_depart_id,omitempty"`
	ThirdDepartIdList []*string `json:"third_depart_id_list,omitempty" xml:"third_depart_id_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 三儿
	UserNick *string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
}

func (s IsvUserSaveRequestUserList) String() string {
	return dara.Prettify(s)
}

func (s IsvUserSaveRequestUserList) GoString() string {
	return s.String()
}

func (s *IsvUserSaveRequestUserList) GetBaseCityCode() *string {
	return s.BaseCityCode
}

func (s *IsvUserSaveRequestUserList) GetBirthday() *string {
	return s.Birthday
}

func (s *IsvUserSaveRequestUserList) GetCertList() []*IsvUserSaveRequestUserListCertList {
	return s.CertList
}

func (s *IsvUserSaveRequestUserList) GetDepartId() *int64 {
	return s.DepartId
}

func (s *IsvUserSaveRequestUserList) GetEmail() *string {
	return s.Email
}

func (s *IsvUserSaveRequestUserList) GetGender() *string {
	return s.Gender
}

func (s *IsvUserSaveRequestUserList) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *IsvUserSaveRequestUserList) GetJobNo() *string {
	return s.JobNo
}

func (s *IsvUserSaveRequestUserList) GetLeaveStatus() *int32 {
	return s.LeaveStatus
}

func (s *IsvUserSaveRequestUserList) GetManagerUserId() *string {
	return s.ManagerUserId
}

func (s *IsvUserSaveRequestUserList) GetPhone() *string {
	return s.Phone
}

func (s *IsvUserSaveRequestUserList) GetPosition() *string {
	return s.Position
}

func (s *IsvUserSaveRequestUserList) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *IsvUserSaveRequestUserList) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *IsvUserSaveRequestUserList) GetRoleIdList() []*string {
	return s.RoleIdList
}

func (s *IsvUserSaveRequestUserList) GetThirdDepartId() *string {
	return s.ThirdDepartId
}

func (s *IsvUserSaveRequestUserList) GetThirdDepartIdList() []*string {
	return s.ThirdDepartIdList
}

func (s *IsvUserSaveRequestUserList) GetUserId() *string {
	return s.UserId
}

func (s *IsvUserSaveRequestUserList) GetUserName() *string {
	return s.UserName
}

func (s *IsvUserSaveRequestUserList) GetUserNick() *string {
	return s.UserNick
}

func (s *IsvUserSaveRequestUserList) SetBaseCityCode(v string) *IsvUserSaveRequestUserList {
	s.BaseCityCode = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetBirthday(v string) *IsvUserSaveRequestUserList {
	s.Birthday = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetCertList(v []*IsvUserSaveRequestUserListCertList) *IsvUserSaveRequestUserList {
	s.CertList = v
	return s
}

func (s *IsvUserSaveRequestUserList) SetDepartId(v int64) *IsvUserSaveRequestUserList {
	s.DepartId = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetEmail(v string) *IsvUserSaveRequestUserList {
	s.Email = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetGender(v string) *IsvUserSaveRequestUserList {
	s.Gender = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetIsAdmin(v bool) *IsvUserSaveRequestUserList {
	s.IsAdmin = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetJobNo(v string) *IsvUserSaveRequestUserList {
	s.JobNo = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetLeaveStatus(v int32) *IsvUserSaveRequestUserList {
	s.LeaveStatus = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetManagerUserId(v string) *IsvUserSaveRequestUserList {
	s.ManagerUserId = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetPhone(v string) *IsvUserSaveRequestUserList {
	s.Phone = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetPosition(v string) *IsvUserSaveRequestUserList {
	s.Position = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetPositionLevel(v string) *IsvUserSaveRequestUserList {
	s.PositionLevel = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetRealNameEn(v string) *IsvUserSaveRequestUserList {
	s.RealNameEn = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetRoleIdList(v []*string) *IsvUserSaveRequestUserList {
	s.RoleIdList = v
	return s
}

func (s *IsvUserSaveRequestUserList) SetThirdDepartId(v string) *IsvUserSaveRequestUserList {
	s.ThirdDepartId = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetThirdDepartIdList(v []*string) *IsvUserSaveRequestUserList {
	s.ThirdDepartIdList = v
	return s
}

func (s *IsvUserSaveRequestUserList) SetUserId(v string) *IsvUserSaveRequestUserList {
	s.UserId = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetUserName(v string) *IsvUserSaveRequestUserList {
	s.UserName = &v
	return s
}

func (s *IsvUserSaveRequestUserList) SetUserNick(v string) *IsvUserSaveRequestUserList {
	s.UserNick = &v
	return s
}

func (s *IsvUserSaveRequestUserList) Validate() error {
	if s.CertList != nil {
		for _, item := range s.CertList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IsvUserSaveRequestUserListCertList struct {
	// example:
	//
	// 2050-01-01
	CertExpiredTime *string `json:"cert_expired_time,omitempty" xml:"cert_expired_time,omitempty"`
	// example:
	//
	// CN
	CertNation *string `json:"cert_nation,omitempty" xml:"cert_nation,omitempty"`
	// example:
	//
	// 110101********1234
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// example:
	//
	// 0
	CertType *int32 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
}

func (s IsvUserSaveRequestUserListCertList) String() string {
	return dara.Prettify(s)
}

func (s IsvUserSaveRequestUserListCertList) GoString() string {
	return s.String()
}

func (s *IsvUserSaveRequestUserListCertList) GetCertExpiredTime() *string {
	return s.CertExpiredTime
}

func (s *IsvUserSaveRequestUserListCertList) GetCertNation() *string {
	return s.CertNation
}

func (s *IsvUserSaveRequestUserListCertList) GetCertNo() *string {
	return s.CertNo
}

func (s *IsvUserSaveRequestUserListCertList) GetCertType() *int32 {
	return s.CertType
}

func (s *IsvUserSaveRequestUserListCertList) GetNationality() *string {
	return s.Nationality
}

func (s *IsvUserSaveRequestUserListCertList) SetCertExpiredTime(v string) *IsvUserSaveRequestUserListCertList {
	s.CertExpiredTime = &v
	return s
}

func (s *IsvUserSaveRequestUserListCertList) SetCertNation(v string) *IsvUserSaveRequestUserListCertList {
	s.CertNation = &v
	return s
}

func (s *IsvUserSaveRequestUserListCertList) SetCertNo(v string) *IsvUserSaveRequestUserListCertList {
	s.CertNo = &v
	return s
}

func (s *IsvUserSaveRequestUserListCertList) SetCertType(v int32) *IsvUserSaveRequestUserListCertList {
	s.CertType = &v
	return s
}

func (s *IsvUserSaveRequestUserListCertList) SetNationality(v string) *IsvUserSaveRequestUserListCertList {
	s.Nationality = &v
	return s
}

func (s *IsvUserSaveRequestUserListCertList) Validate() error {
	return dara.Validate(s)
}
