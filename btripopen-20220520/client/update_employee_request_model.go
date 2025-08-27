// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmployeeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountEmail(v string) *UpdateEmployeeRequest
	GetAccountEmail() *string
	SetAccountPhone(v string) *UpdateEmployeeRequest
	GetAccountPhone() *string
	SetAttribute(v string) *UpdateEmployeeRequest
	GetAttribute() *string
	SetAvatar(v string) *UpdateEmployeeRequest
	GetAvatar() *string
	SetBaseCityCodeList(v []*string) *UpdateEmployeeRequest
	GetBaseCityCodeList() []*string
	SetBaseLocationList(v []*UpdateEmployeeRequestBaseLocationList) *UpdateEmployeeRequest
	GetBaseLocationList() []*UpdateEmployeeRequestBaseLocationList
	SetBirthday(v string) *UpdateEmployeeRequest
	GetBirthday() *string
	SetCertList(v []*UpdateEmployeeRequestCertList) *UpdateEmployeeRequest
	GetCertList() []*UpdateEmployeeRequestCertList
	SetCustomRoleCodeList(v []*string) *UpdateEmployeeRequest
	GetCustomRoleCodeList() []*string
	SetEmail(v string) *UpdateEmployeeRequest
	GetEmail() *string
	SetGender(v string) *UpdateEmployeeRequest
	GetGender() *string
	SetIsAdmin(v bool) *UpdateEmployeeRequest
	GetIsAdmin() *bool
	SetIsBoss(v bool) *UpdateEmployeeRequest
	GetIsBoss() *bool
	SetIsDeptLeader(v bool) *UpdateEmployeeRequest
	GetIsDeptLeader() *bool
	SetJobNo(v string) *UpdateEmployeeRequest
	GetJobNo() *string
	SetManagerUserId(v string) *UpdateEmployeeRequest
	GetManagerUserId() *string
	SetOutDeptIdList(v []*string) *UpdateEmployeeRequest
	GetOutDeptIdList() []*string
	SetPhone(v string) *UpdateEmployeeRequest
	GetPhone() *string
	SetPositionLevel(v string) *UpdateEmployeeRequest
	GetPositionLevel() *string
	SetRealName(v string) *UpdateEmployeeRequest
	GetRealName() *string
	SetRealNameEn(v string) *UpdateEmployeeRequest
	GetRealNameEn() *string
	SetUserId(v string) *UpdateEmployeeRequest
	GetUserId() *string
	SetUserNick(v string) *UpdateEmployeeRequest
	GetUserNick() *string
}

type UpdateEmployeeRequest struct {
	AccountEmail *string `json:"account_email,omitempty" xml:"account_email,omitempty"`
	AccountPhone *string `json:"account_phone,omitempty" xml:"account_phone,omitempty"`
	Attribute    *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// https://static-legacy.dingtalk.com/media/lADPF8XMoxJeUkbNA2LNA5s_923_866.jpg
	Avatar           *string                                  `json:"avatar,omitempty" xml:"avatar,omitempty"`
	BaseCityCodeList []*string                                `json:"base_city_code_list,omitempty" xml:"base_city_code_list,omitempty" type:"Repeated"`
	BaseLocationList []*UpdateEmployeeRequestBaseLocationList `json:"base_location_list,omitempty" xml:"base_location_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2000-01-02
	Birthday           *string                          `json:"birthday,omitempty" xml:"birthday,omitempty"`
	CertList           []*UpdateEmployeeRequestCertList `json:"cert_list,omitempty" xml:"cert_list,omitempty" type:"Repeated"`
	CustomRoleCodeList []*string                        `json:"custom_role_code_list,omitempty" xml:"custom_role_code_list,omitempty" type:"Repeated"`
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
	ManagerUserId *string   `json:"manager_user_id,omitempty" xml:"manager_user_id,omitempty"`
	OutDeptIdList []*string `json:"out_dept_id_list,omitempty" xml:"out_dept_id_list,omitempty" type:"Repeated"`
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

func (s UpdateEmployeeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeRequest) GetAccountEmail() *string {
	return s.AccountEmail
}

func (s *UpdateEmployeeRequest) GetAccountPhone() *string {
	return s.AccountPhone
}

func (s *UpdateEmployeeRequest) GetAttribute() *string {
	return s.Attribute
}

func (s *UpdateEmployeeRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateEmployeeRequest) GetBaseCityCodeList() []*string {
	return s.BaseCityCodeList
}

func (s *UpdateEmployeeRequest) GetBaseLocationList() []*UpdateEmployeeRequestBaseLocationList {
	return s.BaseLocationList
}

func (s *UpdateEmployeeRequest) GetBirthday() *string {
	return s.Birthday
}

func (s *UpdateEmployeeRequest) GetCertList() []*UpdateEmployeeRequestCertList {
	return s.CertList
}

func (s *UpdateEmployeeRequest) GetCustomRoleCodeList() []*string {
	return s.CustomRoleCodeList
}

func (s *UpdateEmployeeRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateEmployeeRequest) GetGender() *string {
	return s.Gender
}

func (s *UpdateEmployeeRequest) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *UpdateEmployeeRequest) GetIsBoss() *bool {
	return s.IsBoss
}

func (s *UpdateEmployeeRequest) GetIsDeptLeader() *bool {
	return s.IsDeptLeader
}

func (s *UpdateEmployeeRequest) GetJobNo() *string {
	return s.JobNo
}

func (s *UpdateEmployeeRequest) GetManagerUserId() *string {
	return s.ManagerUserId
}

func (s *UpdateEmployeeRequest) GetOutDeptIdList() []*string {
	return s.OutDeptIdList
}

func (s *UpdateEmployeeRequest) GetPhone() *string {
	return s.Phone
}

func (s *UpdateEmployeeRequest) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *UpdateEmployeeRequest) GetRealName() *string {
	return s.RealName
}

func (s *UpdateEmployeeRequest) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *UpdateEmployeeRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateEmployeeRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *UpdateEmployeeRequest) SetAccountEmail(v string) *UpdateEmployeeRequest {
	s.AccountEmail = &v
	return s
}

func (s *UpdateEmployeeRequest) SetAccountPhone(v string) *UpdateEmployeeRequest {
	s.AccountPhone = &v
	return s
}

func (s *UpdateEmployeeRequest) SetAttribute(v string) *UpdateEmployeeRequest {
	s.Attribute = &v
	return s
}

func (s *UpdateEmployeeRequest) SetAvatar(v string) *UpdateEmployeeRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateEmployeeRequest) SetBaseCityCodeList(v []*string) *UpdateEmployeeRequest {
	s.BaseCityCodeList = v
	return s
}

func (s *UpdateEmployeeRequest) SetBaseLocationList(v []*UpdateEmployeeRequestBaseLocationList) *UpdateEmployeeRequest {
	s.BaseLocationList = v
	return s
}

func (s *UpdateEmployeeRequest) SetBirthday(v string) *UpdateEmployeeRequest {
	s.Birthday = &v
	return s
}

func (s *UpdateEmployeeRequest) SetCertList(v []*UpdateEmployeeRequestCertList) *UpdateEmployeeRequest {
	s.CertList = v
	return s
}

func (s *UpdateEmployeeRequest) SetCustomRoleCodeList(v []*string) *UpdateEmployeeRequest {
	s.CustomRoleCodeList = v
	return s
}

func (s *UpdateEmployeeRequest) SetEmail(v string) *UpdateEmployeeRequest {
	s.Email = &v
	return s
}

func (s *UpdateEmployeeRequest) SetGender(v string) *UpdateEmployeeRequest {
	s.Gender = &v
	return s
}

func (s *UpdateEmployeeRequest) SetIsAdmin(v bool) *UpdateEmployeeRequest {
	s.IsAdmin = &v
	return s
}

func (s *UpdateEmployeeRequest) SetIsBoss(v bool) *UpdateEmployeeRequest {
	s.IsBoss = &v
	return s
}

func (s *UpdateEmployeeRequest) SetIsDeptLeader(v bool) *UpdateEmployeeRequest {
	s.IsDeptLeader = &v
	return s
}

func (s *UpdateEmployeeRequest) SetJobNo(v string) *UpdateEmployeeRequest {
	s.JobNo = &v
	return s
}

func (s *UpdateEmployeeRequest) SetManagerUserId(v string) *UpdateEmployeeRequest {
	s.ManagerUserId = &v
	return s
}

func (s *UpdateEmployeeRequest) SetOutDeptIdList(v []*string) *UpdateEmployeeRequest {
	s.OutDeptIdList = v
	return s
}

func (s *UpdateEmployeeRequest) SetPhone(v string) *UpdateEmployeeRequest {
	s.Phone = &v
	return s
}

func (s *UpdateEmployeeRequest) SetPositionLevel(v string) *UpdateEmployeeRequest {
	s.PositionLevel = &v
	return s
}

func (s *UpdateEmployeeRequest) SetRealName(v string) *UpdateEmployeeRequest {
	s.RealName = &v
	return s
}

func (s *UpdateEmployeeRequest) SetRealNameEn(v string) *UpdateEmployeeRequest {
	s.RealNameEn = &v
	return s
}

func (s *UpdateEmployeeRequest) SetUserId(v string) *UpdateEmployeeRequest {
	s.UserId = &v
	return s
}

func (s *UpdateEmployeeRequest) SetUserNick(v string) *UpdateEmployeeRequest {
	s.UserNick = &v
	return s
}

func (s *UpdateEmployeeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateEmployeeRequestBaseLocationList struct {
	Code  *string `json:"code,omitempty" xml:"code,omitempty"`
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
}

func (s UpdateEmployeeRequestBaseLocationList) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeRequestBaseLocationList) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeRequestBaseLocationList) GetCode() *string {
	return s.Code
}

func (s *UpdateEmployeeRequestBaseLocationList) GetLevel() *string {
	return s.Level
}

func (s *UpdateEmployeeRequestBaseLocationList) SetCode(v string) *UpdateEmployeeRequestBaseLocationList {
	s.Code = &v
	return s
}

func (s *UpdateEmployeeRequestBaseLocationList) SetLevel(v string) *UpdateEmployeeRequestBaseLocationList {
	s.Level = &v
	return s
}

func (s *UpdateEmployeeRequestBaseLocationList) Validate() error {
	return dara.Validate(s)
}

type UpdateEmployeeRequestCertList struct {
	// example:
	//
	// 2000-01-02
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// example:
	//
	// 2099-03-12
	CertExpiredTime *string `json:"cert_expired_time,omitempty" xml:"cert_expired_time,omitempty"`
	// example:
	//
	// CN
	CertNation *string `json:"cert_nation,omitempty" xml:"cert_nation,omitempty"`
	// example:
	//
	// 123
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// example:
	//
	// 0
	CertType *int32 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// example:
	//
	// F
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// example:
	//
	// 13111111111
	Phone    *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RealName *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// example:
	//
	// John/Wilson
	RealNameEn *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
}

func (s UpdateEmployeeRequestCertList) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeRequestCertList) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeRequestCertList) GetBirthday() *string {
	return s.Birthday
}

func (s *UpdateEmployeeRequestCertList) GetCertExpiredTime() *string {
	return s.CertExpiredTime
}

func (s *UpdateEmployeeRequestCertList) GetCertNation() *string {
	return s.CertNation
}

func (s *UpdateEmployeeRequestCertList) GetCertNo() *string {
	return s.CertNo
}

func (s *UpdateEmployeeRequestCertList) GetCertType() *int32 {
	return s.CertType
}

func (s *UpdateEmployeeRequestCertList) GetGender() *string {
	return s.Gender
}

func (s *UpdateEmployeeRequestCertList) GetNationality() *string {
	return s.Nationality
}

func (s *UpdateEmployeeRequestCertList) GetPhone() *string {
	return s.Phone
}

func (s *UpdateEmployeeRequestCertList) GetRealName() *string {
	return s.RealName
}

func (s *UpdateEmployeeRequestCertList) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *UpdateEmployeeRequestCertList) SetBirthday(v string) *UpdateEmployeeRequestCertList {
	s.Birthday = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) SetCertExpiredTime(v string) *UpdateEmployeeRequestCertList {
	s.CertExpiredTime = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) SetCertNation(v string) *UpdateEmployeeRequestCertList {
	s.CertNation = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) SetCertNo(v string) *UpdateEmployeeRequestCertList {
	s.CertNo = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) SetCertType(v int32) *UpdateEmployeeRequestCertList {
	s.CertType = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) SetGender(v string) *UpdateEmployeeRequestCertList {
	s.Gender = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) SetNationality(v string) *UpdateEmployeeRequestCertList {
	s.Nationality = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) SetPhone(v string) *UpdateEmployeeRequestCertList {
	s.Phone = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) SetRealName(v string) *UpdateEmployeeRequestCertList {
	s.RealName = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) SetRealNameEn(v string) *UpdateEmployeeRequestCertList {
	s.RealNameEn = &v
	return s
}

func (s *UpdateEmployeeRequestCertList) Validate() error {
	return dara.Validate(s)
}
