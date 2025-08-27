// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmployeeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountEmail(v string) *UpdateEmployeeShrinkRequest
	GetAccountEmail() *string
	SetAccountPhone(v string) *UpdateEmployeeShrinkRequest
	GetAccountPhone() *string
	SetAttribute(v string) *UpdateEmployeeShrinkRequest
	GetAttribute() *string
	SetAvatar(v string) *UpdateEmployeeShrinkRequest
	GetAvatar() *string
	SetBaseCityCodeListShrink(v string) *UpdateEmployeeShrinkRequest
	GetBaseCityCodeListShrink() *string
	SetBaseLocationListShrink(v string) *UpdateEmployeeShrinkRequest
	GetBaseLocationListShrink() *string
	SetBirthday(v string) *UpdateEmployeeShrinkRequest
	GetBirthday() *string
	SetCertListShrink(v string) *UpdateEmployeeShrinkRequest
	GetCertListShrink() *string
	SetCustomRoleCodeListShrink(v string) *UpdateEmployeeShrinkRequest
	GetCustomRoleCodeListShrink() *string
	SetEmail(v string) *UpdateEmployeeShrinkRequest
	GetEmail() *string
	SetGender(v string) *UpdateEmployeeShrinkRequest
	GetGender() *string
	SetIsAdmin(v bool) *UpdateEmployeeShrinkRequest
	GetIsAdmin() *bool
	SetIsBoss(v bool) *UpdateEmployeeShrinkRequest
	GetIsBoss() *bool
	SetIsDeptLeader(v bool) *UpdateEmployeeShrinkRequest
	GetIsDeptLeader() *bool
	SetJobNo(v string) *UpdateEmployeeShrinkRequest
	GetJobNo() *string
	SetManagerUserId(v string) *UpdateEmployeeShrinkRequest
	GetManagerUserId() *string
	SetOutDeptIdListShrink(v string) *UpdateEmployeeShrinkRequest
	GetOutDeptIdListShrink() *string
	SetPhone(v string) *UpdateEmployeeShrinkRequest
	GetPhone() *string
	SetPositionLevel(v string) *UpdateEmployeeShrinkRequest
	GetPositionLevel() *string
	SetRealName(v string) *UpdateEmployeeShrinkRequest
	GetRealName() *string
	SetRealNameEn(v string) *UpdateEmployeeShrinkRequest
	GetRealNameEn() *string
	SetUserId(v string) *UpdateEmployeeShrinkRequest
	GetUserId() *string
	SetUserNick(v string) *UpdateEmployeeShrinkRequest
	GetUserNick() *string
}

type UpdateEmployeeShrinkRequest struct {
	AccountEmail *string `json:"account_email,omitempty" xml:"account_email,omitempty"`
	AccountPhone *string `json:"account_phone,omitempty" xml:"account_phone,omitempty"`
	Attribute    *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// https://static-legacy.dingtalk.com/media/lADPF8XMoxJeUkbNA2LNA5s_923_866.jpg
	Avatar                 *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	BaseCityCodeListShrink *string `json:"base_city_code_list,omitempty" xml:"base_city_code_list,omitempty"`
	BaseLocationListShrink *string `json:"base_location_list,omitempty" xml:"base_location_list,omitempty"`
	// example:
	//
	// 2000-01-02
	Birthday                 *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	CertListShrink           *string `json:"cert_list,omitempty" xml:"cert_list,omitempty"`
	CustomRoleCodeListShrink *string `json:"custom_role_code_list,omitempty" xml:"custom_role_code_list,omitempty"`
	// example:
	//
	// 123@163.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// F
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// false
	IsAdmin *bool `json:"is_admin,omitempty" xml:"is_admin,omitempty"`
	// example:
	//
	// false
	IsBoss *bool `json:"is_boss,omitempty" xml:"is_boss,omitempty"`
	// example:
	//
	// false
	IsDeptLeader *bool `json:"is_dept_leader,omitempty" xml:"is_dept_leader,omitempty"`
	// example:
	//
	// 1001
	JobNo *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// example:
	//
	// user456
	ManagerUserId       *string `json:"manager_user_id,omitempty" xml:"manager_user_id,omitempty"`
	OutDeptIdListShrink *string `json:"out_dept_id_list,omitempty" xml:"out_dept_id_list,omitempty"`
	// example:
	//
	// 13111111111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// M4
	PositionLevel *string `json:"position_level,omitempty" xml:"position_level,omitempty"`
	RealName      *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// example:
	//
	// John/Wilson
	RealNameEn *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserNick *string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
}

func (s UpdateEmployeeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeShrinkRequest) GetAccountEmail() *string {
	return s.AccountEmail
}

func (s *UpdateEmployeeShrinkRequest) GetAccountPhone() *string {
	return s.AccountPhone
}

func (s *UpdateEmployeeShrinkRequest) GetAttribute() *string {
	return s.Attribute
}

func (s *UpdateEmployeeShrinkRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateEmployeeShrinkRequest) GetBaseCityCodeListShrink() *string {
	return s.BaseCityCodeListShrink
}

func (s *UpdateEmployeeShrinkRequest) GetBaseLocationListShrink() *string {
	return s.BaseLocationListShrink
}

func (s *UpdateEmployeeShrinkRequest) GetBirthday() *string {
	return s.Birthday
}

func (s *UpdateEmployeeShrinkRequest) GetCertListShrink() *string {
	return s.CertListShrink
}

func (s *UpdateEmployeeShrinkRequest) GetCustomRoleCodeListShrink() *string {
	return s.CustomRoleCodeListShrink
}

func (s *UpdateEmployeeShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateEmployeeShrinkRequest) GetGender() *string {
	return s.Gender
}

func (s *UpdateEmployeeShrinkRequest) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *UpdateEmployeeShrinkRequest) GetIsBoss() *bool {
	return s.IsBoss
}

func (s *UpdateEmployeeShrinkRequest) GetIsDeptLeader() *bool {
	return s.IsDeptLeader
}

func (s *UpdateEmployeeShrinkRequest) GetJobNo() *string {
	return s.JobNo
}

func (s *UpdateEmployeeShrinkRequest) GetManagerUserId() *string {
	return s.ManagerUserId
}

func (s *UpdateEmployeeShrinkRequest) GetOutDeptIdListShrink() *string {
	return s.OutDeptIdListShrink
}

func (s *UpdateEmployeeShrinkRequest) GetPhone() *string {
	return s.Phone
}

func (s *UpdateEmployeeShrinkRequest) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *UpdateEmployeeShrinkRequest) GetRealName() *string {
	return s.RealName
}

func (s *UpdateEmployeeShrinkRequest) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *UpdateEmployeeShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateEmployeeShrinkRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *UpdateEmployeeShrinkRequest) SetAccountEmail(v string) *UpdateEmployeeShrinkRequest {
	s.AccountEmail = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetAccountPhone(v string) *UpdateEmployeeShrinkRequest {
	s.AccountPhone = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetAttribute(v string) *UpdateEmployeeShrinkRequest {
	s.Attribute = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetAvatar(v string) *UpdateEmployeeShrinkRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetBaseCityCodeListShrink(v string) *UpdateEmployeeShrinkRequest {
	s.BaseCityCodeListShrink = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetBaseLocationListShrink(v string) *UpdateEmployeeShrinkRequest {
	s.BaseLocationListShrink = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetBirthday(v string) *UpdateEmployeeShrinkRequest {
	s.Birthday = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetCertListShrink(v string) *UpdateEmployeeShrinkRequest {
	s.CertListShrink = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetCustomRoleCodeListShrink(v string) *UpdateEmployeeShrinkRequest {
	s.CustomRoleCodeListShrink = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetEmail(v string) *UpdateEmployeeShrinkRequest {
	s.Email = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetGender(v string) *UpdateEmployeeShrinkRequest {
	s.Gender = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetIsAdmin(v bool) *UpdateEmployeeShrinkRequest {
	s.IsAdmin = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetIsBoss(v bool) *UpdateEmployeeShrinkRequest {
	s.IsBoss = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetIsDeptLeader(v bool) *UpdateEmployeeShrinkRequest {
	s.IsDeptLeader = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetJobNo(v string) *UpdateEmployeeShrinkRequest {
	s.JobNo = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetManagerUserId(v string) *UpdateEmployeeShrinkRequest {
	s.ManagerUserId = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetOutDeptIdListShrink(v string) *UpdateEmployeeShrinkRequest {
	s.OutDeptIdListShrink = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetPhone(v string) *UpdateEmployeeShrinkRequest {
	s.Phone = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetPositionLevel(v string) *UpdateEmployeeShrinkRequest {
	s.PositionLevel = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetRealName(v string) *UpdateEmployeeShrinkRequest {
	s.RealName = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetRealNameEn(v string) *UpdateEmployeeShrinkRequest {
	s.RealNameEn = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetUserId(v string) *UpdateEmployeeShrinkRequest {
	s.UserId = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) SetUserNick(v string) *UpdateEmployeeShrinkRequest {
	s.UserNick = &v
	return s
}

func (s *UpdateEmployeeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
