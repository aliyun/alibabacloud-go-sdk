// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendTaskPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRecommendTaskPageListRequest
	GetLang() *string
	SetCurrentPage(v string) *DescribeRecommendTaskPageListRequest
	GetCurrentPage() *string
	SetPageSize(v string) *DescribeRecommendTaskPageListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeRecommendTaskPageListRequest
	GetRegId() *string
	SetTaskName(v string) *DescribeRecommendTaskPageListRequest
	GetTaskName() *string
}

type DescribeRecommendTaskPageListRequest struct {
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
	// Current page number
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10
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
	// Task name
	//
	// example:
	//
	// 策略推荐任务
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s DescribeRecommendTaskPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskPageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRecommendTaskPageListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeRecommendTaskPageListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeRecommendTaskPageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRecommendTaskPageListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeRecommendTaskPageListRequest) SetLang(v string) *DescribeRecommendTaskPageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecommendTaskPageListRequest) SetCurrentPage(v string) *DescribeRecommendTaskPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRecommendTaskPageListRequest) SetPageSize(v string) *DescribeRecommendTaskPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecommendTaskPageListRequest) SetRegId(v string) *DescribeRecommendTaskPageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRecommendTaskPageListRequest) SetTaskName(v string) *DescribeRecommendTaskPageListRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeRecommendTaskPageListRequest) Validate() error {
	return dara.Validate(s)
}
