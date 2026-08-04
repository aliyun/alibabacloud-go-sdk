// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeptMemberDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedModels(v string) *DeptMemberDTO
	GetAllowedModels() *string
	SetAuthConfig(v string) *DeptMemberDTO
	GetAuthConfig() *string
	SetGmtCreate(v string) *DeptMemberDTO
	GetGmtCreate() *string
	SetId(v int64) *DeptMemberDTO
	GetId() *int64
	SetKeyCount(v int32) *DeptMemberDTO
	GetKeyCount() *int32
	SetLoginName(v string) *DeptMemberDTO
	GetLoginName() *string
	SetMonthlyBalance(v float64) *DeptMemberDTO
	GetMonthlyBalance() *float64
	SetName(v string) *DeptMemberDTO
	GetName() *string
	SetPermanentBalance(v float64) *DeptMemberDTO
	GetPermanentBalance() *float64
	SetPhone(v string) *DeptMemberDTO
	GetPhone() *string
	SetRoleCode(v string) *DeptMemberDTO
	GetRoleCode() *string
	SetRoleName(v string) *DeptMemberDTO
	GetRoleName() *string
}

type DeptMemberDTO struct {
	// example:
	//
	// 1,2
	AllowedModels *string `json:"allowedModels,omitempty" xml:"allowedModels,omitempty"`
	// example:
	//
	// inherit
	AuthConfig *string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	KeyCount *int32 `json:"keyCount,omitempty" xml:"keyCount,omitempty"`
	// example:
	//
	// zhangsan
	LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty"`
	// example:
	//
	// 10.00
	MonthlyBalance *float64 `json:"monthlyBalance,omitempty" xml:"monthlyBalance,omitempty"`
	// example:
	//
	// John Smith
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 20.00
	PermanentBalance *float64 `json:"permanentBalance,omitempty" xml:"permanentBalance,omitempty"`
	// example:
	//
	// 138****0000
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// member
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// example:
	//
	// Member
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s DeptMemberDTO) String() string {
	return dara.Prettify(s)
}

func (s DeptMemberDTO) GoString() string {
	return s.String()
}

func (s *DeptMemberDTO) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *DeptMemberDTO) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *DeptMemberDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeptMemberDTO) GetId() *int64 {
	return s.Id
}

func (s *DeptMemberDTO) GetKeyCount() *int32 {
	return s.KeyCount
}

func (s *DeptMemberDTO) GetLoginName() *string {
	return s.LoginName
}

func (s *DeptMemberDTO) GetMonthlyBalance() *float64 {
	return s.MonthlyBalance
}

func (s *DeptMemberDTO) GetName() *string {
	return s.Name
}

func (s *DeptMemberDTO) GetPermanentBalance() *float64 {
	return s.PermanentBalance
}

func (s *DeptMemberDTO) GetPhone() *string {
	return s.Phone
}

func (s *DeptMemberDTO) GetRoleCode() *string {
	return s.RoleCode
}

func (s *DeptMemberDTO) GetRoleName() *string {
	return s.RoleName
}

func (s *DeptMemberDTO) SetAllowedModels(v string) *DeptMemberDTO {
	s.AllowedModels = &v
	return s
}

func (s *DeptMemberDTO) SetAuthConfig(v string) *DeptMemberDTO {
	s.AuthConfig = &v
	return s
}

func (s *DeptMemberDTO) SetGmtCreate(v string) *DeptMemberDTO {
	s.GmtCreate = &v
	return s
}

func (s *DeptMemberDTO) SetId(v int64) *DeptMemberDTO {
	s.Id = &v
	return s
}

func (s *DeptMemberDTO) SetKeyCount(v int32) *DeptMemberDTO {
	s.KeyCount = &v
	return s
}

func (s *DeptMemberDTO) SetLoginName(v string) *DeptMemberDTO {
	s.LoginName = &v
	return s
}

func (s *DeptMemberDTO) SetMonthlyBalance(v float64) *DeptMemberDTO {
	s.MonthlyBalance = &v
	return s
}

func (s *DeptMemberDTO) SetName(v string) *DeptMemberDTO {
	s.Name = &v
	return s
}

func (s *DeptMemberDTO) SetPermanentBalance(v float64) *DeptMemberDTO {
	s.PermanentBalance = &v
	return s
}

func (s *DeptMemberDTO) SetPhone(v string) *DeptMemberDTO {
	s.Phone = &v
	return s
}

func (s *DeptMemberDTO) SetRoleCode(v string) *DeptMemberDTO {
	s.RoleCode = &v
	return s
}

func (s *DeptMemberDTO) SetRoleName(v string) *DeptMemberDTO {
	s.RoleName = &v
	return s
}

func (s *DeptMemberDTO) Validate() error {
	return dara.Validate(s)
}
