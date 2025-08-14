// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneEventPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSceneEventPageListRequest
	GetLang() *string
	SetCurrentPage(v string) *DescribeSceneEventPageListRequest
	GetCurrentPage() *string
	SetNameOrCode(v string) *DescribeSceneEventPageListRequest
	GetNameOrCode() *string
	SetPageSize(v string) *DescribeSceneEventPageListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeSceneEventPageListRequest
	GetRegId() *string
}

type DescribeSceneEventPageListRequest struct {
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
	// Search name or service code
	//
	// example:
	//
	// servicer_code
	NameOrCode *string `json:"nameOrCode,omitempty" xml:"nameOrCode,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeSceneEventPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneEventPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSceneEventPageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSceneEventPageListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSceneEventPageListRequest) GetNameOrCode() *string {
	return s.NameOrCode
}

func (s *DescribeSceneEventPageListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSceneEventPageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSceneEventPageListRequest) SetLang(v string) *DescribeSceneEventPageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSceneEventPageListRequest) SetCurrentPage(v string) *DescribeSceneEventPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSceneEventPageListRequest) SetNameOrCode(v string) *DescribeSceneEventPageListRequest {
	s.NameOrCode = &v
	return s
}

func (s *DescribeSceneEventPageListRequest) SetPageSize(v string) *DescribeSceneEventPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSceneEventPageListRequest) SetRegId(v string) *DescribeSceneEventPageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSceneEventPageListRequest) Validate() error {
	return dara.Validate(s)
}
