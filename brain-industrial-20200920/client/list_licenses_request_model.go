// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLicensesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListLicensesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListLicensesRequest
	GetPageSize() *int32
	SetQueryStr(v string) *ListLicensesRequest
	GetQueryStr() *string
}

type ListLicensesRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 12
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryStr *string `json:"QueryStr,omitempty" xml:"QueryStr,omitempty"`
}

func (s ListLicensesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLicensesRequest) GoString() string {
	return s.String()
}

func (s *ListLicensesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListLicensesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLicensesRequest) GetQueryStr() *string {
	return s.QueryStr
}

func (s *ListLicensesRequest) SetCurrentPage(v int32) *ListLicensesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListLicensesRequest) SetPageSize(v int32) *ListLicensesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLicensesRequest) SetQueryStr(v string) *ListLicensesRequest {
	s.QueryStr = &v
	return s
}

func (s *ListLicensesRequest) Validate() error {
	return dara.Validate(s)
}
