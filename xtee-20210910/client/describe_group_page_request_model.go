// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGroupPageRequest
	GetLang() *string
	SetCurrentPage(v string) *DescribeGroupPageRequest
	GetCurrentPage() *string
	SetDirection(v string) *DescribeGroupPageRequest
	GetDirection() *string
	SetOrder(v string) *DescribeGroupPageRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeGroupPageRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeGroupPageRequest
	GetRegId() *string
	SetTaskId(v string) *DescribeGroupPageRequest
	GetTaskId() *string
	SetTimeType(v string) *DescribeGroupPageRequest
	GetTimeType() *string
}

type DescribeGroupPageRequest struct {
	// Sets the language type for requests and responses, with a default value of **zh**. Values:
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
	// Order.
	//
	// example:
	//
	// 0
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// Sorting condition.
	//
	// example:
	//
	// asc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// Page size, with a default value of 10.
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
	// 6770764
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// Time type.
	//
	// example:
	//
	// 1
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
}

func (s DescribeGroupPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupPageRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeGroupPageRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeGroupPageRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeGroupPageRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeGroupPageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeGroupPageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeGroupPageRequest) GetTimeType() *string {
	return s.TimeType
}

func (s *DescribeGroupPageRequest) SetLang(v string) *DescribeGroupPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupPageRequest) SetCurrentPage(v string) *DescribeGroupPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupPageRequest) SetDirection(v string) *DescribeGroupPageRequest {
	s.Direction = &v
	return s
}

func (s *DescribeGroupPageRequest) SetOrder(v string) *DescribeGroupPageRequest {
	s.Order = &v
	return s
}

func (s *DescribeGroupPageRequest) SetPageSize(v string) *DescribeGroupPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupPageRequest) SetRegId(v string) *DescribeGroupPageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeGroupPageRequest) SetTaskId(v string) *DescribeGroupPageRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeGroupPageRequest) SetTimeType(v string) *DescribeGroupPageRequest {
	s.TimeType = &v
	return s
}

func (s *DescribeGroupPageRequest) Validate() error {
	return dara.Validate(s)
}
