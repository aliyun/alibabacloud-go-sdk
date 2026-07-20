// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompaniesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyList(v []*ListCompaniesResponseBodyCompanyList) *ListCompaniesResponseBody
	GetCompanyList() []*ListCompaniesResponseBodyCompanyList
	SetCurrentPage(v int32) *ListCompaniesResponseBody
	GetCurrentPage() *int32
	SetRequestId(v string) *ListCompaniesResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListCompaniesResponseBody
	GetShowSize() *int32
	SetTotalCount(v int32) *ListCompaniesResponseBody
	GetTotalCount() *int32
}

type ListCompaniesResponseBody struct {
	// The list of companies.
	CompanyList []*ListCompaniesResponseBodyCompanyList `json:"CompanyList,omitempty" xml:"CompanyList,omitempty" type:"Repeated"`
	// Settings the page number of the current page in a paged query for paging. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 285BBE08-F12B-5A04-97BC-09EA7FF18646
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of certificates to display per page in a paged query. Default value: 10.
	//
	// example:
	//
	// 10
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of search results.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCompaniesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCompaniesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCompaniesResponseBody) GetCompanyList() []*ListCompaniesResponseBodyCompanyList {
	return s.CompanyList
}

func (s *ListCompaniesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCompaniesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCompaniesResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCompaniesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCompaniesResponseBody) SetCompanyList(v []*ListCompaniesResponseBodyCompanyList) *ListCompaniesResponseBody {
	s.CompanyList = v
	return s
}

func (s *ListCompaniesResponseBody) SetCurrentPage(v int32) *ListCompaniesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCompaniesResponseBody) SetRequestId(v string) *ListCompaniesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCompaniesResponseBody) SetShowSize(v int32) *ListCompaniesResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCompaniesResponseBody) SetTotalCount(v int32) *ListCompaniesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCompaniesResponseBody) Validate() error {
	if s.CompanyList != nil {
		for _, item := range s.CompanyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCompaniesResponseBodyCompanyList struct {
	// The city.
	//
	// example:
	//
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The company address.
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
	// The company email address.
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
	// The company phone number.
	//
	// example:
	//
	// 1511
	CompanyPhone *string `json:"CompanyPhone,omitempty" xml:"CompanyPhone,omitempty"`
	// The company code.
	//
	// example:
	//
	// xxx
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
	// test
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListCompaniesResponseBodyCompanyList) String() string {
	return dara.Prettify(s)
}

func (s ListCompaniesResponseBodyCompanyList) GoString() string {
	return s.String()
}

func (s *ListCompaniesResponseBodyCompanyList) GetCity() *string {
	return s.City
}

func (s *ListCompaniesResponseBodyCompanyList) GetCompanyAddress() *string {
	return s.CompanyAddress
}

func (s *ListCompaniesResponseBodyCompanyList) GetCompanyCode() *string {
	return s.CompanyCode
}

func (s *ListCompaniesResponseBodyCompanyList) GetCompanyEmail() *string {
	return s.CompanyEmail
}

func (s *ListCompaniesResponseBodyCompanyList) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListCompaniesResponseBodyCompanyList) GetCompanyName() *string {
	return s.CompanyName
}

func (s *ListCompaniesResponseBodyCompanyList) GetCompanyPhone() *string {
	return s.CompanyPhone
}

func (s *ListCompaniesResponseBodyCompanyList) GetCompanyType() *int32 {
	return s.CompanyType
}

func (s *ListCompaniesResponseBodyCompanyList) GetCountryCode() *string {
	return s.CountryCode
}

func (s *ListCompaniesResponseBodyCompanyList) GetDepartment() *string {
	return s.Department
}

func (s *ListCompaniesResponseBodyCompanyList) GetLang() *string {
	return s.Lang
}

func (s *ListCompaniesResponseBodyCompanyList) GetPostCode() *string {
	return s.PostCode
}

func (s *ListCompaniesResponseBodyCompanyList) GetProvince() *string {
	return s.Province
}

func (s *ListCompaniesResponseBodyCompanyList) SetCity(v string) *ListCompaniesResponseBodyCompanyList {
	s.City = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetCompanyAddress(v string) *ListCompaniesResponseBodyCompanyList {
	s.CompanyAddress = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetCompanyCode(v string) *ListCompaniesResponseBodyCompanyList {
	s.CompanyCode = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetCompanyEmail(v string) *ListCompaniesResponseBodyCompanyList {
	s.CompanyEmail = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetCompanyId(v int64) *ListCompaniesResponseBodyCompanyList {
	s.CompanyId = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetCompanyName(v string) *ListCompaniesResponseBodyCompanyList {
	s.CompanyName = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetCompanyPhone(v string) *ListCompaniesResponseBodyCompanyList {
	s.CompanyPhone = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetCompanyType(v int32) *ListCompaniesResponseBodyCompanyList {
	s.CompanyType = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetCountryCode(v string) *ListCompaniesResponseBodyCompanyList {
	s.CountryCode = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetDepartment(v string) *ListCompaniesResponseBodyCompanyList {
	s.Department = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetLang(v string) *ListCompaniesResponseBodyCompanyList {
	s.Lang = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetPostCode(v string) *ListCompaniesResponseBodyCompanyList {
	s.PostCode = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) SetProvince(v string) *ListCompaniesResponseBodyCompanyList {
	s.Province = &v
	return s
}

func (s *ListCompaniesResponseBodyCompanyList) Validate() error {
	return dara.Validate(s)
}
