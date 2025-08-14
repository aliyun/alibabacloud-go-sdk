// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarkPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeMarkPageRequest
	GetLang() *string
	SetCurrentPage(v string) *DescribeMarkPageRequest
	GetCurrentPage() *string
	SetDirection(v string) *DescribeMarkPageRequest
	GetDirection() *string
	SetIsPage(v bool) *DescribeMarkPageRequest
	GetIsPage() *bool
	SetOrder(v string) *DescribeMarkPageRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeMarkPageRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeMarkPageRequest
	GetRegId() *string
	SetTaskLogId(v string) *DescribeMarkPageRequest
	GetTaskLogId() *string
}

type DescribeMarkPageRequest struct {
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
	// Order direction.
	//
	// example:
	//
	// 0
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// Whether to paginate.
	//
	// example:
	//
	// true
	IsPage *bool `json:"isPage,omitempty" xml:"isPage,omitempty"`
	// Sorting condition.
	//
	// example:
	//
	// asc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
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
	// Task ID.
	//
	// example:
	//
	// 443
	TaskLogId *string `json:"taskLogId,omitempty" xml:"taskLogId,omitempty"`
}

func (s DescribeMarkPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarkPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMarkPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeMarkPageRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeMarkPageRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeMarkPageRequest) GetIsPage() *bool {
	return s.IsPage
}

func (s *DescribeMarkPageRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeMarkPageRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeMarkPageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeMarkPageRequest) GetTaskLogId() *string {
	return s.TaskLogId
}

func (s *DescribeMarkPageRequest) SetLang(v string) *DescribeMarkPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeMarkPageRequest) SetCurrentPage(v string) *DescribeMarkPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeMarkPageRequest) SetDirection(v string) *DescribeMarkPageRequest {
	s.Direction = &v
	return s
}

func (s *DescribeMarkPageRequest) SetIsPage(v bool) *DescribeMarkPageRequest {
	s.IsPage = &v
	return s
}

func (s *DescribeMarkPageRequest) SetOrder(v string) *DescribeMarkPageRequest {
	s.Order = &v
	return s
}

func (s *DescribeMarkPageRequest) SetPageSize(v string) *DescribeMarkPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMarkPageRequest) SetRegId(v string) *DescribeMarkPageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeMarkPageRequest) SetTaskLogId(v string) *DescribeMarkPageRequest {
	s.TaskLogId = &v
	return s
}

func (s *DescribeMarkPageRequest) Validate() error {
	return dara.Validate(s)
}
