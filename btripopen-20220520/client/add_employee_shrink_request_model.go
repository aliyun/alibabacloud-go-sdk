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
	// The account email address. This parameter can be specified when the enterprise activation method is email activation.
	//
	// example:
	//
	// j*********@example.com
	AccountEmail *string `json:"account_email,omitempty" xml:"account_email,omitempty"`
	// The account phone number. For enterprises with international phone numbers enabled, specify international numbers, Hong Kong (China), Macao (China), and Taiwan (China) numbers in the format +xx-xxxxxx.
	//
	// example:
	//
	// +86-18812345678
	AccountPhone *string `json:"account_phone,omitempty" xml:"account_phone,omitempty"`
	// The custom extension field for the employee, which supports key-value pairs.
	//
	// - Format: JSON string.
	//
	// example:
	//
	// {
	//
	//     "sio": "123",
	//
	//     "location": "22222",
	//
	//     "isForeigner": "Y"
	//
	// }
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The avatar of the employee. Specify the URL of the image.
	//
	// example:
	//
	// https://example.com/example.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The 6-digit administrative code of the work location.
	//
	// - A maximum of two different administrative codes can be specified.
	BaseCityCodeListShrink *string `json:"base_city_code_list,omitempty" xml:"base_city_code_list,omitempty"`
	// The base location information of the employee.
	BaseLocationListShrink *string `json:"base_location_list,omitempty" xml:"base_location_list,omitempty"`
	// The birthday of the employee.
	//
	// - Format: yy-MM-dd.
	//
	// example:
	//
	// 2000-01-01
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// The certificate information of the employee.
	CertListShrink *string `json:"cert_list,omitempty" xml:"cert_list,omitempty"`
	// The collection of role IDs associated with the employee. The number of roles associated with a single employee must be less than or equal to 200. Otherwise, the employee synchronization fails.
	CustomRoleCodeListShrink *string `json:"custom_role_code_list,omitempty" xml:"custom_role_code_list,omitempty"`
	// The email address of the employee.
	//
	// example:
	//
	// j*********@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The gender of the employee.
	//
	// example:
	//
	// F
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// Specifies whether the employee is an Alibaba Business Travel enterprise administrator.
	//
	// example:
	//
	// false
	IsAdmin *bool `json:"is_admin,omitempty" xml:"is_admin,omitempty"`
	// Specifies whether the employee is the boss.
	//
	// example:
	//
	// false
	IsBoss *bool `json:"is_boss,omitempty" xml:"is_boss,omitempty"`
	// Specifies whether the employee is a department manager.
	//
	// example:
	//
	// false
	IsDeptLeader *bool `json:"is_dept_leader,omitempty" xml:"is_dept_leader,omitempty"`
	// The employee number, which serves as a unique identifier for the employee along with `user_id`. Ensure that the value is unique.
	//
	// example:
	//
	// job_1234
	JobNo *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// The ID of the direct manager of the employee.
	//
	// example:
	//
	// user_001
	ManagerUserId *string `json:"manager_user_id,omitempty" xml:"manager_user_id,omitempty"`
	// The list of departments to which the employee belongs.
	OutDeptIdListShrink *string `json:"out_dept_id_list,omitempty" xml:"out_dept_id_list,omitempty"`
	// The mobile phone number of the employee.
	//
	// - This field is commonly used for booking business travel services across various categories. **In this case, it is required.**
	//
	// - If your enterprise is a government agency or other special enterprise, call 400-800-5890 to contact an Alibaba Business Travel customer service representative.
	//
	// example:
	//
	// 131****8888
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The position level of the employee, which is commonly used to match different travel standards.
	//
	// example:
	//
	// 高级
	PositionLevel *string `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// The name of the employee.
	//
	// example:
	//
	// 张三
	RealName *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// The English name of the employee. Follow these format requirements:
	//
	// - Separate the last name and first name with "/", for example: LastName/FirstName.
	//
	// - Do not include spaces between the last name and first name.
	//
	// example:
	//
	// John/Wilson
	RealNameEn *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
	// The unique identifier of the same employee across multiple enterprises (parent and subsidiary enterprises).
	//
	// example:
	//
	// union_0123
	UnionId *string `json:"union_id,omitempty" xml:"union_id,omitempty"`
	// The employee ID, which is the unique identifier of the employee within the enterprise. Ensure that this value is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_1234
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// The nickname of the employee.
	//
	// This parameter is required.
	//
	// example:
	//
	// 弓长
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
