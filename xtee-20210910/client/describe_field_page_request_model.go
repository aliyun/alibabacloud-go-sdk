// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeFieldPageRequest
	GetLang() *string
	SetClassify(v string) *DescribeFieldPageRequest
	GetClassify() *string
	SetCondition(v string) *DescribeFieldPageRequest
	GetCondition() *string
	SetCurrentPage(v string) *DescribeFieldPageRequest
	GetCurrentPage() *string
	SetName(v string) *DescribeFieldPageRequest
	GetName() *string
	SetPageSize(v string) *DescribeFieldPageRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeFieldPageRequest
	GetRegId() *string
	SetSource(v string) *DescribeFieldPageRequest
	GetSource() *string
	SetStatus(v string) *DescribeFieldPageRequest
	GetStatus() *string
	SetTitle(v string) *DescribeFieldPageRequest
	GetTitle() *string
	SetType(v string) *DescribeFieldPageRequest
	GetType() *string
}

type DescribeFieldPageRequest struct {
	// Set the language type for request and response messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Field classification
	//
	// example:
	//
	// REQUEST_PARAM
	Classify *string `json:"classify,omitempty" xml:"classify,omitempty"`
	// Query input parameter name or title
	//
	// example:
	//
	// age/年龄
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Field name
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Number of items per page, default value is 10
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
	// Field source
	//
	// example:
	//
	// DEFINE
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Title.
	//
	// example:
	//
	// 年龄
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Field type
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeFieldPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeFieldPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFieldPageRequest) GetClassify() *string {
	return s.Classify
}

func (s *DescribeFieldPageRequest) GetCondition() *string {
	return s.Condition
}

func (s *DescribeFieldPageRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeFieldPageRequest) GetName() *string {
	return s.Name
}

func (s *DescribeFieldPageRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeFieldPageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeFieldPageRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeFieldPageRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeFieldPageRequest) GetTitle() *string {
	return s.Title
}

func (s *DescribeFieldPageRequest) GetType() *string {
	return s.Type
}

func (s *DescribeFieldPageRequest) SetLang(v string) *DescribeFieldPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFieldPageRequest) SetClassify(v string) *DescribeFieldPageRequest {
	s.Classify = &v
	return s
}

func (s *DescribeFieldPageRequest) SetCondition(v string) *DescribeFieldPageRequest {
	s.Condition = &v
	return s
}

func (s *DescribeFieldPageRequest) SetCurrentPage(v string) *DescribeFieldPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeFieldPageRequest) SetName(v string) *DescribeFieldPageRequest {
	s.Name = &v
	return s
}

func (s *DescribeFieldPageRequest) SetPageSize(v string) *DescribeFieldPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFieldPageRequest) SetRegId(v string) *DescribeFieldPageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeFieldPageRequest) SetSource(v string) *DescribeFieldPageRequest {
	s.Source = &v
	return s
}

func (s *DescribeFieldPageRequest) SetStatus(v string) *DescribeFieldPageRequest {
	s.Status = &v
	return s
}

func (s *DescribeFieldPageRequest) SetTitle(v string) *DescribeFieldPageRequest {
	s.Title = &v
	return s
}

func (s *DescribeFieldPageRequest) SetType(v string) *DescribeFieldPageRequest {
	s.Type = &v
	return s
}

func (s *DescribeFieldPageRequest) Validate() error {
	return dara.Validate(s)
}
