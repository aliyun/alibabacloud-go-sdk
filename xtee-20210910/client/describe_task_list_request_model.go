// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeTaskListRequest
	GetCurrentPage() *string
	SetIsPage(v bool) *DescribeTaskListRequest
	GetIsPage() *bool
	SetLang(v string) *DescribeTaskListRequest
	GetLang() *string
	SetPageSize(v string) *DescribeTaskListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeTaskListRequest
	GetRegId() *string
}

type DescribeTaskListRequest struct {
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Whether to paginate.
	//
	// example:
	//
	// true
	IsPage *bool `json:"IsPage,omitempty" xml:"IsPage,omitempty"`
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Page size, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeTaskListRequest) GetIsPage() *bool {
	return s.IsPage
}

func (s *DescribeTaskListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTaskListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeTaskListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTaskListRequest) SetCurrentPage(v string) *DescribeTaskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTaskListRequest) SetIsPage(v bool) *DescribeTaskListRequest {
	s.IsPage = &v
	return s
}

func (s *DescribeTaskListRequest) SetLang(v string) *DescribeTaskListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTaskListRequest) SetPageSize(v string) *DescribeTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTaskListRequest) SetRegId(v string) *DescribeTaskListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTaskListRequest) Validate() error {
	return dara.Validate(s)
}
