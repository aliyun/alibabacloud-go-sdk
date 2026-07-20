// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompanyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *GetCompanyResponseBody
	GetCity() *string
	SetCompanyAddress(v string) *GetCompanyResponseBody
	GetCompanyAddress() *string
	SetCompanyCode(v string) *GetCompanyResponseBody
	GetCompanyCode() *string
	SetCompanyEmail(v string) *GetCompanyResponseBody
	GetCompanyEmail() *string
	SetCompanyId(v int64) *GetCompanyResponseBody
	GetCompanyId() *int64
	SetCompanyName(v string) *GetCompanyResponseBody
	GetCompanyName() *string
	SetCompanyPhone(v string) *GetCompanyResponseBody
	GetCompanyPhone() *string
	SetCompanyType(v int32) *GetCompanyResponseBody
	GetCompanyType() *int32
	SetCountryCode(v string) *GetCompanyResponseBody
	GetCountryCode() *string
	SetDepartment(v string) *GetCompanyResponseBody
	GetDepartment() *string
	SetLang(v string) *GetCompanyResponseBody
	GetLang() *string
	SetPostCode(v string) *GetCompanyResponseBody
	GetPostCode() *string
	SetProvince(v string) *GetCompanyResponseBody
	GetProvince() *string
	SetRequestId(v string) *GetCompanyResponseBody
	GetRequestId() *string
}

type GetCompanyResponseBody struct {
	// The city.
	//
	// example:
	//
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The address of the company.
	//
	// example:
	//
	// test
	CompanyAddress *string `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty"`
	// The company code.
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
	// example:
	//
	// 51001
	CompanyId *int64 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	// The name of the company or organization.
	//
	// example:
	//
	// testYanwen045
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// The phone number of the company.
	//
	// example:
	//
	// 1511
	CompanyPhone *string `json:"CompanyPhone,omitempty" xml:"CompanyPhone,omitempty"`
	// The company type.
	//
	// example:
	//
	// 0
	CompanyType *int32 `json:"CompanyType,omitempty" xml:"CompanyType,omitempty"`
	// The country code.
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
	// example:
	//
	// 100000
	PostCode *string `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	// The province.
	//
	// example:
	//
	// Beijing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0068247C-A454-5FC9-93BF-C41CBB5CD19E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCompanyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCompanyResponseBody) GoString() string {
	return s.String()
}

func (s *GetCompanyResponseBody) GetCity() *string {
	return s.City
}

func (s *GetCompanyResponseBody) GetCompanyAddress() *string {
	return s.CompanyAddress
}

func (s *GetCompanyResponseBody) GetCompanyCode() *string {
	return s.CompanyCode
}

func (s *GetCompanyResponseBody) GetCompanyEmail() *string {
	return s.CompanyEmail
}

func (s *GetCompanyResponseBody) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *GetCompanyResponseBody) GetCompanyName() *string {
	return s.CompanyName
}

func (s *GetCompanyResponseBody) GetCompanyPhone() *string {
	return s.CompanyPhone
}

func (s *GetCompanyResponseBody) GetCompanyType() *int32 {
	return s.CompanyType
}

func (s *GetCompanyResponseBody) GetCountryCode() *string {
	return s.CountryCode
}

func (s *GetCompanyResponseBody) GetDepartment() *string {
	return s.Department
}

func (s *GetCompanyResponseBody) GetLang() *string {
	return s.Lang
}

func (s *GetCompanyResponseBody) GetPostCode() *string {
	return s.PostCode
}

func (s *GetCompanyResponseBody) GetProvince() *string {
	return s.Province
}

func (s *GetCompanyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCompanyResponseBody) SetCity(v string) *GetCompanyResponseBody {
	s.City = &v
	return s
}

func (s *GetCompanyResponseBody) SetCompanyAddress(v string) *GetCompanyResponseBody {
	s.CompanyAddress = &v
	return s
}

func (s *GetCompanyResponseBody) SetCompanyCode(v string) *GetCompanyResponseBody {
	s.CompanyCode = &v
	return s
}

func (s *GetCompanyResponseBody) SetCompanyEmail(v string) *GetCompanyResponseBody {
	s.CompanyEmail = &v
	return s
}

func (s *GetCompanyResponseBody) SetCompanyId(v int64) *GetCompanyResponseBody {
	s.CompanyId = &v
	return s
}

func (s *GetCompanyResponseBody) SetCompanyName(v string) *GetCompanyResponseBody {
	s.CompanyName = &v
	return s
}

func (s *GetCompanyResponseBody) SetCompanyPhone(v string) *GetCompanyResponseBody {
	s.CompanyPhone = &v
	return s
}

func (s *GetCompanyResponseBody) SetCompanyType(v int32) *GetCompanyResponseBody {
	s.CompanyType = &v
	return s
}

func (s *GetCompanyResponseBody) SetCountryCode(v string) *GetCompanyResponseBody {
	s.CountryCode = &v
	return s
}

func (s *GetCompanyResponseBody) SetDepartment(v string) *GetCompanyResponseBody {
	s.Department = &v
	return s
}

func (s *GetCompanyResponseBody) SetLang(v string) *GetCompanyResponseBody {
	s.Lang = &v
	return s
}

func (s *GetCompanyResponseBody) SetPostCode(v string) *GetCompanyResponseBody {
	s.PostCode = &v
	return s
}

func (s *GetCompanyResponseBody) SetProvince(v string) *GetCompanyResponseBody {
	s.Province = &v
	return s
}

func (s *GetCompanyResponseBody) SetRequestId(v string) *GetCompanyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCompanyResponseBody) Validate() error {
	return dara.Validate(s)
}
