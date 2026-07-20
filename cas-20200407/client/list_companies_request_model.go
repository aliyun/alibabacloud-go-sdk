// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompaniesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListCompaniesRequest
	GetCompanyId() *int64
	SetCurrentPage(v int32) *ListCompaniesRequest
	GetCurrentPage() *int32
	SetKeyword(v string) *ListCompaniesRequest
	GetKeyword() *string
	SetShowSize(v int32) *ListCompaniesRequest
	GetShowSize() *int32
}

type ListCompaniesRequest struct {
	// The company ID.
	//
	// example:
	//
	// 51001
	CompanyId *int64 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	// The page number of the current page. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The search keyword. For example, a keyword for the company name, province, country code, or city.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of contacts to display per page in a paged query.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCompaniesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCompaniesRequest) GoString() string {
	return s.String()
}

func (s *ListCompaniesRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListCompaniesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCompaniesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCompaniesRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCompaniesRequest) SetCompanyId(v int64) *ListCompaniesRequest {
	s.CompanyId = &v
	return s
}

func (s *ListCompaniesRequest) SetCurrentPage(v int32) *ListCompaniesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCompaniesRequest) SetKeyword(v string) *ListCompaniesRequest {
	s.Keyword = &v
	return s
}

func (s *ListCompaniesRequest) SetShowSize(v int32) *ListCompaniesRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCompaniesRequest) Validate() error {
	return dara.Validate(s)
}
