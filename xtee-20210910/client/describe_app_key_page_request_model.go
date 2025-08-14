// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppKeyPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAppKeyPageRequest
	GetLang() *string
	SetCurrentPage(v string) *DescribeAppKeyPageRequest
	GetCurrentPage() *string
	SetPageSize(v string) *DescribeAppKeyPageRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeAppKeyPageRequest
	GetRegId() *string
}

type DescribeAppKeyPageRequest struct {
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
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Number of items per page, default value is 10.
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

func (s DescribeAppKeyPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppKeyPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAppKeyPageRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeAppKeyPageRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAppKeyPageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAppKeyPageRequest) SetLang(v string) *DescribeAppKeyPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAppKeyPageRequest) SetCurrentPage(v string) *DescribeAppKeyPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAppKeyPageRequest) SetPageSize(v string) *DescribeAppKeyPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppKeyPageRequest) SetRegId(v string) *DescribeAppKeyPageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAppKeyPageRequest) Validate() error {
	return dara.Validate(s)
}
