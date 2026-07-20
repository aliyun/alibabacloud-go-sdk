// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompanyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *UpdateCompanyRequest
	GetCity() *string
	SetCompanyAddress(v string) *UpdateCompanyRequest
	GetCompanyAddress() *string
	SetCompanyCode(v string) *UpdateCompanyRequest
	GetCompanyCode() *string
	SetCompanyEmail(v string) *UpdateCompanyRequest
	GetCompanyEmail() *string
	SetCompanyId(v int64) *UpdateCompanyRequest
	GetCompanyId() *int64
	SetCompanyName(v string) *UpdateCompanyRequest
	GetCompanyName() *string
	SetCompanyPhone(v string) *UpdateCompanyRequest
	GetCompanyPhone() *string
	SetCompanyType(v int32) *UpdateCompanyRequest
	GetCompanyType() *int32
	SetCountryCode(v string) *UpdateCompanyRequest
	GetCountryCode() *string
	SetDepartment(v string) *UpdateCompanyRequest
	GetDepartment() *string
	SetLang(v string) *UpdateCompanyRequest
	GetLang() *string
	SetPostCode(v string) *UpdateCompanyRequest
	GetPostCode() *string
	SetProvince(v string) *UpdateCompanyRequest
	GetProvince() *string
}

type UpdateCompanyRequest struct {
	// The city.
	//
	// This parameter is required.
	//
	// example:
	//
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The address of the company.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	CompanyAddress *string `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty"`
	// The company code.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	CompanyCode *string `json:"CompanyCode,omitempty" xml:"CompanyCode,omitempty"`
	// The email address of the company.
	//
	// example:
	//
	// test@163.com
	CompanyEmail *string `json:"CompanyEmail,omitempty" xml:"CompanyEmail,omitempty"`
	// The company ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51001
	CompanyId *int64 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	// The name of the company or organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// testYanwen045
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// The phone number of the company.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1511
	CompanyPhone *string `json:"CompanyPhone,omitempty" xml:"CompanyPhone,omitempty"`
	// The company type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	CompanyType *int32 `json:"CompanyType,omitempty" xml:"CompanyType,omitempty"`
	// The country code.
	//
	// This parameter is required.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The department.
	//
	// example:
	//
	// test
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The language.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The postal code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000
	PostCode *string `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	// The province.
	//
	// This parameter is required.
	//
	// example:
	//
	// Beijing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s UpdateCompanyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompanyRequest) GoString() string {
	return s.String()
}

func (s *UpdateCompanyRequest) GetCity() *string {
	return s.City
}

func (s *UpdateCompanyRequest) GetCompanyAddress() *string {
	return s.CompanyAddress
}

func (s *UpdateCompanyRequest) GetCompanyCode() *string {
	return s.CompanyCode
}

func (s *UpdateCompanyRequest) GetCompanyEmail() *string {
	return s.CompanyEmail
}

func (s *UpdateCompanyRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *UpdateCompanyRequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *UpdateCompanyRequest) GetCompanyPhone() *string {
	return s.CompanyPhone
}

func (s *UpdateCompanyRequest) GetCompanyType() *int32 {
	return s.CompanyType
}

func (s *UpdateCompanyRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *UpdateCompanyRequest) GetDepartment() *string {
	return s.Department
}

func (s *UpdateCompanyRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateCompanyRequest) GetPostCode() *string {
	return s.PostCode
}

func (s *UpdateCompanyRequest) GetProvince() *string {
	return s.Province
}

func (s *UpdateCompanyRequest) SetCity(v string) *UpdateCompanyRequest {
	s.City = &v
	return s
}

func (s *UpdateCompanyRequest) SetCompanyAddress(v string) *UpdateCompanyRequest {
	s.CompanyAddress = &v
	return s
}

func (s *UpdateCompanyRequest) SetCompanyCode(v string) *UpdateCompanyRequest {
	s.CompanyCode = &v
	return s
}

func (s *UpdateCompanyRequest) SetCompanyEmail(v string) *UpdateCompanyRequest {
	s.CompanyEmail = &v
	return s
}

func (s *UpdateCompanyRequest) SetCompanyId(v int64) *UpdateCompanyRequest {
	s.CompanyId = &v
	return s
}

func (s *UpdateCompanyRequest) SetCompanyName(v string) *UpdateCompanyRequest {
	s.CompanyName = &v
	return s
}

func (s *UpdateCompanyRequest) SetCompanyPhone(v string) *UpdateCompanyRequest {
	s.CompanyPhone = &v
	return s
}

func (s *UpdateCompanyRequest) SetCompanyType(v int32) *UpdateCompanyRequest {
	s.CompanyType = &v
	return s
}

func (s *UpdateCompanyRequest) SetCountryCode(v string) *UpdateCompanyRequest {
	s.CountryCode = &v
	return s
}

func (s *UpdateCompanyRequest) SetDepartment(v string) *UpdateCompanyRequest {
	s.Department = &v
	return s
}

func (s *UpdateCompanyRequest) SetLang(v string) *UpdateCompanyRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCompanyRequest) SetPostCode(v string) *UpdateCompanyRequest {
	s.PostCode = &v
	return s
}

func (s *UpdateCompanyRequest) SetProvince(v string) *UpdateCompanyRequest {
	s.Province = &v
	return s
}

func (s *UpdateCompanyRequest) Validate() error {
	return dara.Validate(s)
}
