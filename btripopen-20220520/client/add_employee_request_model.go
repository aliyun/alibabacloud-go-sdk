// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountEmail(v string) *AddEmployeeRequest
	GetAccountEmail() *string
	SetAccountPhone(v string) *AddEmployeeRequest
	GetAccountPhone() *string
	SetAttribute(v string) *AddEmployeeRequest
	GetAttribute() *string
	SetAvatar(v string) *AddEmployeeRequest
	GetAvatar() *string
	SetBaseCityCodeList(v []*string) *AddEmployeeRequest
	GetBaseCityCodeList() []*string
	SetBaseLocationList(v []*AddEmployeeRequestBaseLocationList) *AddEmployeeRequest
	GetBaseLocationList() []*AddEmployeeRequestBaseLocationList
	SetBirthday(v string) *AddEmployeeRequest
	GetBirthday() *string
	SetCertList(v []*AddEmployeeRequestCertList) *AddEmployeeRequest
	GetCertList() []*AddEmployeeRequestCertList
	SetCustomRoleCodeList(v []*string) *AddEmployeeRequest
	GetCustomRoleCodeList() []*string
	SetEmail(v string) *AddEmployeeRequest
	GetEmail() *string
	SetGender(v string) *AddEmployeeRequest
	GetGender() *string
	SetIsAdmin(v bool) *AddEmployeeRequest
	GetIsAdmin() *bool
	SetIsBoss(v bool) *AddEmployeeRequest
	GetIsBoss() *bool
	SetIsDeptLeader(v bool) *AddEmployeeRequest
	GetIsDeptLeader() *bool
	SetJobNo(v string) *AddEmployeeRequest
	GetJobNo() *string
	SetManagerUserId(v string) *AddEmployeeRequest
	GetManagerUserId() *string
	SetOutDeptIdList(v []*string) *AddEmployeeRequest
	GetOutDeptIdList() []*string
	SetPhone(v string) *AddEmployeeRequest
	GetPhone() *string
	SetPositionLevel(v string) *AddEmployeeRequest
	GetPositionLevel() *string
	SetRealName(v string) *AddEmployeeRequest
	GetRealName() *string
	SetRealNameEn(v string) *AddEmployeeRequest
	GetRealNameEn() *string
	SetUnionId(v string) *AddEmployeeRequest
	GetUnionId() *string
	SetUserId(v string) *AddEmployeeRequest
	GetUserId() *string
	SetUserNick(v string) *AddEmployeeRequest
	GetUserNick() *string
}

type AddEmployeeRequest struct {
	AccountEmail *string `json:"account_email,omitempty" xml:"account_email,omitempty"`
	AccountPhone *string `json:"account_phone,omitempty" xml:"account_phone,omitempty"`
	Attribute    *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// https://static-legacy.dingtalk.com/media/lADPF8XMoxJeUkbNA2LNA5s_923_866.jpg
	Avatar           *string                               `json:"avatar,omitempty" xml:"avatar,omitempty"`
	BaseCityCodeList []*string                             `json:"base_city_code_list,omitempty" xml:"base_city_code_list,omitempty" type:"Repeated"`
	BaseLocationList []*AddEmployeeRequestBaseLocationList `json:"base_location_list,omitempty" xml:"base_location_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2000-01-02
	Birthday           *string                       `json:"birthday,omitempty" xml:"birthday,omitempty"`
	CertList           []*AddEmployeeRequestCertList `json:"cert_list,omitempty" xml:"cert_list,omitempty" type:"Repeated"`
	CustomRoleCodeList []*string                     `json:"custom_role_code_list,omitempty" xml:"custom_role_code_list,omitempty" type:"Repeated"`
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

func (s AddEmployeeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeeRequest) GoString() string {
	return s.String()
}

func (s *AddEmployeeRequest) GetAccountEmail() *string {
	return s.AccountEmail
}

func (s *AddEmployeeRequest) GetAccountPhone() *string {
	return s.AccountPhone
}

func (s *AddEmployeeRequest) GetAttribute() *string {
	return s.Attribute
}

func (s *AddEmployeeRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *AddEmployeeRequest) GetBaseCityCodeList() []*string {
	return s.BaseCityCodeList
}

func (s *AddEmployeeRequest) GetBaseLocationList() []*AddEmployeeRequestBaseLocationList {
	return s.BaseLocationList
}

func (s *AddEmployeeRequest) GetBirthday() *string {
	return s.Birthday
}

func (s *AddEmployeeRequest) GetCertList() []*AddEmployeeRequestCertList {
	return s.CertList
}

func (s *AddEmployeeRequest) GetCustomRoleCodeList() []*string {
	return s.CustomRoleCodeList
}

func (s *AddEmployeeRequest) GetEmail() *string {
	return s.Email
}

func (s *AddEmployeeRequest) GetGender() *string {
	return s.Gender
}

func (s *AddEmployeeRequest) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *AddEmployeeRequest) GetIsBoss() *bool {
	return s.IsBoss
}

func (s *AddEmployeeRequest) GetIsDeptLeader() *bool {
	return s.IsDeptLeader
}

func (s *AddEmployeeRequest) GetJobNo() *string {
	return s.JobNo
}

func (s *AddEmployeeRequest) GetManagerUserId() *string {
	return s.ManagerUserId
}

func (s *AddEmployeeRequest) GetOutDeptIdList() []*string {
	return s.OutDeptIdList
}

func (s *AddEmployeeRequest) GetPhone() *string {
	return s.Phone
}

func (s *AddEmployeeRequest) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *AddEmployeeRequest) GetRealName() *string {
	return s.RealName
}

func (s *AddEmployeeRequest) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *AddEmployeeRequest) GetUnionId() *string {
	return s.UnionId
}

func (s *AddEmployeeRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddEmployeeRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *AddEmployeeRequest) SetAccountEmail(v string) *AddEmployeeRequest {
	s.AccountEmail = &v
	return s
}

func (s *AddEmployeeRequest) SetAccountPhone(v string) *AddEmployeeRequest {
	s.AccountPhone = &v
	return s
}

func (s *AddEmployeeRequest) SetAttribute(v string) *AddEmployeeRequest {
	s.Attribute = &v
	return s
}

func (s *AddEmployeeRequest) SetAvatar(v string) *AddEmployeeRequest {
	s.Avatar = &v
	return s
}

func (s *AddEmployeeRequest) SetBaseCityCodeList(v []*string) *AddEmployeeRequest {
	s.BaseCityCodeList = v
	return s
}

func (s *AddEmployeeRequest) SetBaseLocationList(v []*AddEmployeeRequestBaseLocationList) *AddEmployeeRequest {
	s.BaseLocationList = v
	return s
}

func (s *AddEmployeeRequest) SetBirthday(v string) *AddEmployeeRequest {
	s.Birthday = &v
	return s
}

func (s *AddEmployeeRequest) SetCertList(v []*AddEmployeeRequestCertList) *AddEmployeeRequest {
	s.CertList = v
	return s
}

func (s *AddEmployeeRequest) SetCustomRoleCodeList(v []*string) *AddEmployeeRequest {
	s.CustomRoleCodeList = v
	return s
}

func (s *AddEmployeeRequest) SetEmail(v string) *AddEmployeeRequest {
	s.Email = &v
	return s
}

func (s *AddEmployeeRequest) SetGender(v string) *AddEmployeeRequest {
	s.Gender = &v
	return s
}

func (s *AddEmployeeRequest) SetIsAdmin(v bool) *AddEmployeeRequest {
	s.IsAdmin = &v
	return s
}

func (s *AddEmployeeRequest) SetIsBoss(v bool) *AddEmployeeRequest {
	s.IsBoss = &v
	return s
}

func (s *AddEmployeeRequest) SetIsDeptLeader(v bool) *AddEmployeeRequest {
	s.IsDeptLeader = &v
	return s
}

func (s *AddEmployeeRequest) SetJobNo(v string) *AddEmployeeRequest {
	s.JobNo = &v
	return s
}

func (s *AddEmployeeRequest) SetManagerUserId(v string) *AddEmployeeRequest {
	s.ManagerUserId = &v
	return s
}

func (s *AddEmployeeRequest) SetOutDeptIdList(v []*string) *AddEmployeeRequest {
	s.OutDeptIdList = v
	return s
}

func (s *AddEmployeeRequest) SetPhone(v string) *AddEmployeeRequest {
	s.Phone = &v
	return s
}

func (s *AddEmployeeRequest) SetPositionLevel(v string) *AddEmployeeRequest {
	s.PositionLevel = &v
	return s
}

func (s *AddEmployeeRequest) SetRealName(v string) *AddEmployeeRequest {
	s.RealName = &v
	return s
}

func (s *AddEmployeeRequest) SetRealNameEn(v string) *AddEmployeeRequest {
	s.RealNameEn = &v
	return s
}

func (s *AddEmployeeRequest) SetUnionId(v string) *AddEmployeeRequest {
	s.UnionId = &v
	return s
}

func (s *AddEmployeeRequest) SetUserId(v string) *AddEmployeeRequest {
	s.UserId = &v
	return s
}

func (s *AddEmployeeRequest) SetUserNick(v string) *AddEmployeeRequest {
	s.UserNick = &v
	return s
}

func (s *AddEmployeeRequest) Validate() error {
	return dara.Validate(s)
}

type AddEmployeeRequestBaseLocationList struct {
	Code  *string `json:"code,omitempty" xml:"code,omitempty"`
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
}

func (s AddEmployeeRequestBaseLocationList) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeeRequestBaseLocationList) GoString() string {
	return s.String()
}

func (s *AddEmployeeRequestBaseLocationList) GetCode() *string {
	return s.Code
}

func (s *AddEmployeeRequestBaseLocationList) GetLevel() *string {
	return s.Level
}

func (s *AddEmployeeRequestBaseLocationList) SetCode(v string) *AddEmployeeRequestBaseLocationList {
	s.Code = &v
	return s
}

func (s *AddEmployeeRequestBaseLocationList) SetLevel(v string) *AddEmployeeRequestBaseLocationList {
	s.Level = &v
	return s
}

func (s *AddEmployeeRequestBaseLocationList) Validate() error {
	return dara.Validate(s)
}

type AddEmployeeRequestCertList struct {
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

func (s AddEmployeeRequestCertList) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeeRequestCertList) GoString() string {
	return s.String()
}

func (s *AddEmployeeRequestCertList) GetBirthday() *string {
	return s.Birthday
}

func (s *AddEmployeeRequestCertList) GetCertExpiredTime() *string {
	return s.CertExpiredTime
}

func (s *AddEmployeeRequestCertList) GetCertNation() *string {
	return s.CertNation
}

func (s *AddEmployeeRequestCertList) GetCertNo() *string {
	return s.CertNo
}

func (s *AddEmployeeRequestCertList) GetCertType() *int32 {
	return s.CertType
}

func (s *AddEmployeeRequestCertList) GetGender() *string {
	return s.Gender
}

func (s *AddEmployeeRequestCertList) GetNationality() *string {
	return s.Nationality
}

func (s *AddEmployeeRequestCertList) GetPhone() *string {
	return s.Phone
}

func (s *AddEmployeeRequestCertList) GetRealName() *string {
	return s.RealName
}

func (s *AddEmployeeRequestCertList) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *AddEmployeeRequestCertList) SetBirthday(v string) *AddEmployeeRequestCertList {
	s.Birthday = &v
	return s
}

func (s *AddEmployeeRequestCertList) SetCertExpiredTime(v string) *AddEmployeeRequestCertList {
	s.CertExpiredTime = &v
	return s
}

func (s *AddEmployeeRequestCertList) SetCertNation(v string) *AddEmployeeRequestCertList {
	s.CertNation = &v
	return s
}

func (s *AddEmployeeRequestCertList) SetCertNo(v string) *AddEmployeeRequestCertList {
	s.CertNo = &v
	return s
}

func (s *AddEmployeeRequestCertList) SetCertType(v int32) *AddEmployeeRequestCertList {
	s.CertType = &v
	return s
}

func (s *AddEmployeeRequestCertList) SetGender(v string) *AddEmployeeRequestCertList {
	s.Gender = &v
	return s
}

func (s *AddEmployeeRequestCertList) SetNationality(v string) *AddEmployeeRequestCertList {
	s.Nationality = &v
	return s
}

func (s *AddEmployeeRequestCertList) SetPhone(v string) *AddEmployeeRequestCertList {
	s.Phone = &v
	return s
}

func (s *AddEmployeeRequestCertList) SetRealName(v string) *AddEmployeeRequestCertList {
	s.RealName = &v
	return s
}

func (s *AddEmployeeRequestCertList) SetRealNameEn(v string) *AddEmployeeRequestCertList {
	s.RealNameEn = &v
	return s
}

func (s *AddEmployeeRequestCertList) Validate() error {
	return dara.Validate(s)
}
