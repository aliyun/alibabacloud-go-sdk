// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompanyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *CreateCompanyRequest
	GetCity() *string
	SetCompanyAddress(v string) *CreateCompanyRequest
	GetCompanyAddress() *string
	SetCompanyCode(v string) *CreateCompanyRequest
	GetCompanyCode() *string
	SetCompanyEmail(v string) *CreateCompanyRequest
	GetCompanyEmail() *string
	SetCompanyName(v string) *CreateCompanyRequest
	GetCompanyName() *string
	SetCompanyPhone(v string) *CreateCompanyRequest
	GetCompanyPhone() *string
	SetCompanyType(v int32) *CreateCompanyRequest
	GetCompanyType() *int32
	SetCountryCode(v string) *CreateCompanyRequest
	GetCountryCode() *string
	SetDepartment(v string) *CreateCompanyRequest
	GetDepartment() *string
	SetLang(v string) *CreateCompanyRequest
	GetLang() *string
	SetPostCode(v string) *CreateCompanyRequest
	GetPostCode() *string
	SetProvince(v string) *CreateCompanyRequest
	GetProvince() *string
}

type CreateCompanyRequest struct {
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
	// test@123.com
	CompanyEmail *string `json:"CompanyEmail,omitempty" xml:"CompanyEmail,omitempty"`
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
	// 1999
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
	// This parameter is required.
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

func (s CreateCompanyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCompanyRequest) GoString() string {
	return s.String()
}

func (s *CreateCompanyRequest) GetCity() *string {
	return s.City
}

func (s *CreateCompanyRequest) GetCompanyAddress() *string {
	return s.CompanyAddress
}

func (s *CreateCompanyRequest) GetCompanyCode() *string {
	return s.CompanyCode
}

func (s *CreateCompanyRequest) GetCompanyEmail() *string {
	return s.CompanyEmail
}

func (s *CreateCompanyRequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *CreateCompanyRequest) GetCompanyPhone() *string {
	return s.CompanyPhone
}

func (s *CreateCompanyRequest) GetCompanyType() *int32 {
	return s.CompanyType
}

func (s *CreateCompanyRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *CreateCompanyRequest) GetDepartment() *string {
	return s.Department
}

func (s *CreateCompanyRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateCompanyRequest) GetPostCode() *string {
	return s.PostCode
}

func (s *CreateCompanyRequest) GetProvince() *string {
	return s.Province
}

func (s *CreateCompanyRequest) SetCity(v string) *CreateCompanyRequest {
	s.City = &v
	return s
}

func (s *CreateCompanyRequest) SetCompanyAddress(v string) *CreateCompanyRequest {
	s.CompanyAddress = &v
	return s
}

func (s *CreateCompanyRequest) SetCompanyCode(v string) *CreateCompanyRequest {
	s.CompanyCode = &v
	return s
}

func (s *CreateCompanyRequest) SetCompanyEmail(v string) *CreateCompanyRequest {
	s.CompanyEmail = &v
	return s
}

func (s *CreateCompanyRequest) SetCompanyName(v string) *CreateCompanyRequest {
	s.CompanyName = &v
	return s
}

func (s *CreateCompanyRequest) SetCompanyPhone(v string) *CreateCompanyRequest {
	s.CompanyPhone = &v
	return s
}

func (s *CreateCompanyRequest) SetCompanyType(v int32) *CreateCompanyRequest {
	s.CompanyType = &v
	return s
}

func (s *CreateCompanyRequest) SetCountryCode(v string) *CreateCompanyRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateCompanyRequest) SetDepartment(v string) *CreateCompanyRequest {
	s.Department = &v
	return s
}

func (s *CreateCompanyRequest) SetLang(v string) *CreateCompanyRequest {
	s.Lang = &v
	return s
}

func (s *CreateCompanyRequest) SetPostCode(v string) *CreateCompanyRequest {
	s.PostCode = &v
	return s
}

func (s *CreateCompanyRequest) SetProvince(v string) *CreateCompanyRequest {
	s.Province = &v
	return s
}

func (s *CreateCompanyRequest) Validate() error {
	return dara.Validate(s)
}
