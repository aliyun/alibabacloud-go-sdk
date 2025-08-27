// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountEmail(v string) *AddEmployeeShrinkRequest
	GetAccountEmail() *string
	SetAccountPhone(v string) *AddEmployeeShrinkRequest
	GetAccountPhone() *string
	SetAttribute(v string) *AddEmployeeShrinkRequest
	GetAttribute() *string
	SetAvatar(v string) *AddEmployeeShrinkRequest
	GetAvatar() *string
	SetBaseCityCodeListShrink(v string) *AddEmployeeShrinkRequest
	GetBaseCityCodeListShrink() *string
	SetBaseLocationListShrink(v string) *AddEmployeeShrinkRequest
	GetBaseLocationListShrink() *string
	SetBirthday(v string) *AddEmployeeShrinkRequest
	GetBirthday() *string
	SetCertListShrink(v string) *AddEmployeeShrinkRequest
	GetCertListShrink() *string
	SetCustomRoleCodeListShrink(v string) *AddEmployeeShrinkRequest
	GetCustomRoleCodeListShrink() *string
	SetEmail(v string) *AddEmployeeShrinkRequest
	GetEmail() *string
	SetGender(v string) *AddEmployeeShrinkRequest
	GetGender() *string
	SetIsAdmin(v bool) *AddEmployeeShrinkRequest
	GetIsAdmin() *bool
	SetIsBoss(v bool) *AddEmployeeShrinkRequest
	GetIsBoss() *bool
	SetIsDeptLeader(v bool) *AddEmployeeShrinkRequest
	GetIsDeptLeader() *bool
	SetJobNo(v string) *AddEmployeeShrinkRequest
	GetJobNo() *string
	SetManagerUserId(v string) *AddEmployeeShrinkRequest
	GetManagerUserId() *string
	SetOutDeptIdListShrink(v string) *AddEmployeeShrinkRequest
	GetOutDeptIdListShrink() *string
	SetPhone(v string) *AddEmployeeShrinkRequest
	GetPhone() *string
	SetPositionLevel(v string) *AddEmployeeShrinkRequest
	GetPositionLevel() *string
	SetRealName(v string) *AddEmployeeShrinkRequest
	GetRealName() *string
	SetRealNameEn(v string) *AddEmployeeShrinkRequest
	GetRealNameEn() *string
	SetUnionId(v string) *AddEmployeeShrinkRequest
	GetUnionId() *string
	SetUserId(v string) *AddEmployeeShrinkRequest
	GetUserId() *string
	SetUserNick(v string) *AddEmployeeShrinkRequest
	GetUserNick() *string
}

type AddEmployeeShrinkRequest struct {
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
	// example:
	//
	// union123
	UnionId *string `json:"union_id,omitempty" xml:"union_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// This parameter is required.
	UserNick *string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
}

func (s AddEmployeeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeeShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddEmployeeShrinkRequest) GetAccountEmail() *string {
	return s.AccountEmail
}

func (s *AddEmployeeShrinkRequest) GetAccountPhone() *string {
	return s.AccountPhone
}

func (s *AddEmployeeShrinkRequest) GetAttribute() *string {
	return s.Attribute
}

func (s *AddEmployeeShrinkRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *AddEmployeeShrinkRequest) GetBaseCityCodeListShrink() *string {
	return s.BaseCityCodeListShrink
}

func (s *AddEmployeeShrinkRequest) GetBaseLocationListShrink() *string {
	return s.BaseLocationListShrink
}

func (s *AddEmployeeShrinkRequest) GetBirthday() *string {
	return s.Birthday
}

func (s *AddEmployeeShrinkRequest) GetCertListShrink() *string {
	return s.CertListShrink
}

func (s *AddEmployeeShrinkRequest) GetCustomRoleCodeListShrink() *string {
	return s.CustomRoleCodeListShrink
}

func (s *AddEmployeeShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *AddEmployeeShrinkRequest) GetGender() *string {
	return s.Gender
}

func (s *AddEmployeeShrinkRequest) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *AddEmployeeShrinkRequest) GetIsBoss() *bool {
	return s.IsBoss
}

func (s *AddEmployeeShrinkRequest) GetIsDeptLeader() *bool {
	return s.IsDeptLeader
}

func (s *AddEmployeeShrinkRequest) GetJobNo() *string {
	return s.JobNo
}

func (s *AddEmployeeShrinkRequest) GetManagerUserId() *string {
	return s.ManagerUserId
}

func (s *AddEmployeeShrinkRequest) GetOutDeptIdListShrink() *string {
	return s.OutDeptIdListShrink
}

func (s *AddEmployeeShrinkRequest) GetPhone() *string {
	return s.Phone
}

func (s *AddEmployeeShrinkRequest) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *AddEmployeeShrinkRequest) GetRealName() *string {
	return s.RealName
}

func (s *AddEmployeeShrinkRequest) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *AddEmployeeShrinkRequest) GetUnionId() *string {
	return s.UnionId
}

func (s *AddEmployeeShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddEmployeeShrinkRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *AddEmployeeShrinkRequest) SetAccountEmail(v string) *AddEmployeeShrinkRequest {
	s.AccountEmail = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetAccountPhone(v string) *AddEmployeeShrinkRequest {
	s.AccountPhone = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetAttribute(v string) *AddEmployeeShrinkRequest {
	s.Attribute = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetAvatar(v string) *AddEmployeeShrinkRequest {
	s.Avatar = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetBaseCityCodeListShrink(v string) *AddEmployeeShrinkRequest {
	s.BaseCityCodeListShrink = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetBaseLocationListShrink(v string) *AddEmployeeShrinkRequest {
	s.BaseLocationListShrink = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetBirthday(v string) *AddEmployeeShrinkRequest {
	s.Birthday = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetCertListShrink(v string) *AddEmployeeShrinkRequest {
	s.CertListShrink = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetCustomRoleCodeListShrink(v string) *AddEmployeeShrinkRequest {
	s.CustomRoleCodeListShrink = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetEmail(v string) *AddEmployeeShrinkRequest {
	s.Email = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetGender(v string) *AddEmployeeShrinkRequest {
	s.Gender = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetIsAdmin(v bool) *AddEmployeeShrinkRequest {
	s.IsAdmin = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetIsBoss(v bool) *AddEmployeeShrinkRequest {
	s.IsBoss = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetIsDeptLeader(v bool) *AddEmployeeShrinkRequest {
	s.IsDeptLeader = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetJobNo(v string) *AddEmployeeShrinkRequest {
	s.JobNo = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetManagerUserId(v string) *AddEmployeeShrinkRequest {
	s.ManagerUserId = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetOutDeptIdListShrink(v string) *AddEmployeeShrinkRequest {
	s.OutDeptIdListShrink = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetPhone(v string) *AddEmployeeShrinkRequest {
	s.Phone = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetPositionLevel(v string) *AddEmployeeShrinkRequest {
	s.PositionLevel = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetRealName(v string) *AddEmployeeShrinkRequest {
	s.RealName = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetRealNameEn(v string) *AddEmployeeShrinkRequest {
	s.RealNameEn = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetUnionId(v string) *AddEmployeeShrinkRequest {
	s.UnionId = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetUserId(v string) *AddEmployeeShrinkRequest {
	s.UserId = &v
	return s
}

func (s *AddEmployeeShrinkRequest) SetUserNick(v string) *AddEmployeeShrinkRequest {
	s.UserNick = &v
	return s
}

func (s *AddEmployeeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
