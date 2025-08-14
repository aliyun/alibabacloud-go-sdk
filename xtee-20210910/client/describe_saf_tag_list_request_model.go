// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafTagListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSafTagListRequest
	GetLang() *string
	SetTagName(v string) *DescribeSafTagListRequest
	GetTagName() *string
	SetApiId(v string) *DescribeSafTagListRequest
	GetApiId() *string
	SetCurrentPage(v string) *DescribeSafTagListRequest
	GetCurrentPage() *string
	SetPageSize(v string) *DescribeSafTagListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeSafTagListRequest
	GetRegId() *string
}

type DescribeSafTagListRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Tag name. Fuzzy search.
	//
	// example:
	//
	// rn0301
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// API service ID.
	//
	// example:
	//
	// 34
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeSafTagListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafTagListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSafTagListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSafTagListRequest) GetTagName() *string {
	return s.TagName
}

func (s *DescribeSafTagListRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeSafTagListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSafTagListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSafTagListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSafTagListRequest) SetLang(v string) *DescribeSafTagListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSafTagListRequest) SetTagName(v string) *DescribeSafTagListRequest {
	s.TagName = &v
	return s
}

func (s *DescribeSafTagListRequest) SetApiId(v string) *DescribeSafTagListRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeSafTagListRequest) SetCurrentPage(v string) *DescribeSafTagListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSafTagListRequest) SetPageSize(v string) *DescribeSafTagListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSafTagListRequest) SetRegId(v string) *DescribeSafTagListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSafTagListRequest) Validate() error {
	return dara.Validate(s)
}
